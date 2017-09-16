package user

import (
	"context"
	"fmt"
	"net/http"
	"strconv"

	"errors"

	"time"

	"crypto/sha1"

	"clickyab.com/crab/modules/user/aaa"
	"github.com/clickyab/services/assert"
	"github.com/clickyab/services/kv"

	"strings"

	"net/url"

	"github.com/clickyab/services/notification"
	"github.com/clickyab/services/random"
	"github.com/clickyab/services/trans"
	"github.com/rs/xmux"
)

const (
	verifyKey          = "VERIFY"
	subID              = "ID"
	dump               = "d"
	emailVerifyPath    = "user/email/verify"
	passwordVerifyPath = "user/password/verify"
)

var (
	tooSoonError = errors.New("code has been sent")
)

type verifyIdResponse responseLoginOK

// verifyId is verify code
// @Route {
// 		url = /email/verify/:token
//		method = get
//		200 = verifyIdResponse
//		401 = controller.ErrorResponseSimple
//		403 = controller.ErrorResponseSimple
// }
func (ctrl *Controller) verifyEmail(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	u, e := verifyCode(xmux.Param(ctx, "token"))

	if e != nil || u.Status != aaa.RegisteredUserStatus {
		ctrl.ForbiddenResponse(w, nil)
		return
	}
	u.Status = aaa.ActiveUserStatus
	assert.Nil(aaa.NewAaaManager().UpdateUser(u))
	ctrl.createLoginResponse(w, u)
}

// @Validate{
// }
type verifyResendPayload struct {
	Email string `json:"email_string" validate:"required,email"`
}

// verifyResend will send an email again
// @Route {
// 		url = /email/verify/resend
//		method = post
//		payload = verifyResendPayload
//      200 = controller.NormalResponse
//		404 = controller.ErrorResponseSimple
// }
func (ctrl *Controller) verifyResend(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	pl := ctrl.MustGetPayload(ctx).(*verifyResendPayload)

	u, e := aaa.NewAaaManager().FindUserByEmail(pl.Email)
	if e != nil {
		ctrl.NotFoundResponse(w, nil)
		return
	}
	e = verifyEmail(u, r)
	if e == tooSoonError {
		ctrl.OKResponse(w, nil)
		return
	}
	assert.Nil(e)
}

func verifyEmail(u *aaa.User, r *http.Request) error {
	ur, e := genVerificationURL(u, emailVerifyPath, r)
	if e != nil {
		return e
	}

	// TODO: Change email template
	notification.Send(trans.T("Welcome to Clickyab!").String(), ur.String(), notification.Packet{
		Platform: notification.MailType,
		To:       []string{u.Email},
	})
	return nil
}

func genVerificationURL(u *aaa.User, p string, r *http.Request) (*url.URL, error) {
	k, e := genVerifyCode(u, p)
	if e != nil {
		return nil, e
	}

	return &url.URL{
		Scheme: func() string {
			if r.TLS != nil {
				return "https"
			}
			return "http"
		}(),
		Host:     r.Host,
		Path:     p,
		RawQuery: fmt.Sprintf("key=%s", k),
	}, nil

}

const delimiter = "-"

func verifyCode(c string) (*aaa.User, error) {
	s := strings.Split(c, delimiter)
	if len(s) != 2 {
		return nil, errors.New("code is not valid")
	}

	hash, key := s[0], s[1]

	kw := kv.NewEavStore(fmt.Sprintf("%s_%s", verifyKey, hash))

	if kw.SubKey(key) != dump {
		return nil, errors.New("code is not valid")
	}
	defer kw.Drop()
	id, e := strconv.ParseInt(kw.SubKey(subID), 10, 64)
	assert.Nil(e)
	m := aaa.NewAaaManager()
	cu, e := m.FindUserByID(id)
	assert.Nil(e)
	return cu, nil
}

var exp = 5 * time.Hour
var saltError = errors.New("salt should not be empty")

func genVerifyCode(u *aaa.User, salt string) (string, error) {
	if salt == "" {
		return "", saltError
	}
	hash := func() string {
		h := sha1.New()
		h.Write([]byte(u.Email + salt))
		return fmt.Sprintf("%x", h.Sum(nil))
	}()

	kw := kv.NewEavStore(fmt.Sprintf("%s_%s", verifyKey, hash))
	// TODO: get it from config

	if len(kw.AllKeys()) != 0 {
		t := time.Now().Add(exp).Add(-2 * time.Minute).Sub(time.Now())
		if t < kw.TTL() {
			return "", tooSoonError
		}
	}

	key := fmt.Sprintf("%s%s", <-random.ID, <-random.ID)
	kw.SetSubKey(subID, fmt.Sprintf("%d", u.ID))
	kw.SetSubKey(key, dump)
	assert.Nil(kw.Save(exp))

	return fmt.Sprintf("%s%s%s", hash, delimiter, key), nil

}
