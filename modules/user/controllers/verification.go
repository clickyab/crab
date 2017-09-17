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

	"github.com/clickyab/services/config"
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
	exp          = config.RegisterDuration("crab.modules.user.verification.ttl", 5*time.Hour, "how long the token should be saved")
	resend       = config.RegisterDuration("crab.modules.user.verification.resend", 2*time.Minute, "Duration between resend")
)

type verifyIdResponse responseLoginOK

// verifyId is verify code
// @Route {
// 		url = /email/verify/:token
//		method = get
//		200 = responseLoginOK
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

// @Validate {
// }
type verifyEmailCodePayload struct {
	Email string `json:"email" validate:"required,email"`
	Code  string `json:"code" validate:"required"`
}

// @Route {
// 		url = /email/verify
//		method = post
//		payload = verifyEmailCodePayload
//		200 = responseLoginOK
//		401 = controller.ErrorResponseSimple
//		403 = controller.ErrorResponseSimple
// }
func (ctrl *Controller) verifyEmailCode(ctx context.Context, w http.ResponseWriter, r *http.Request) {

	p := ctrl.MustGetPayload(ctx).(*verifyEmailCodePayload)
	u, e := verifyCode(fmt.Sprintf("%s%s%s", hasher(p.Email+emailVerifyPath), delimiter, p.Code))

	if e != nil || u.Status != aaa.RegisteredUserStatus || strings.ToLower(p.Email) != strings.ToLower(u.Email) {
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
	h, c, e := genVerifyCode(u, emailVerifyPath)
	if e != nil {
		return e
	}
	ul := &url.URL{
		Scheme: func() string {
			if r.TLS != nil {
				return "https"
			}
			return "http"
		}(),
		Host: r.Host,
		Path: fmt.Sprintf("/user/register/verification/%s", h),
	}
	temp := fmt.Sprintf(`
	%s
	code: %s
	`, ul.String(), c)
	// TODO: Change email template
	notification.Send(trans.T("Welcome to Clickyab!").String(), temp, notification.Packet{
		Platform: notification.MailType,
		To:       []string{u.Email},
	})
	return nil
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

var saltError = errors.New("salt should not be empty")

func genVerifyCode(u *aaa.User, salt string) (string, string, error) {
	assert.True(salt != "")

	hash := hasher(u.Email + salt)
	kw := kv.NewEavStore(fmt.Sprintf("%s_%s", verifyKey, hash))

	if len(kw.AllKeys()) != 0 {
		if exp.Duration()-resend.Duration() < kw.TTL() {
			return "", "", tooSoonError
		}
	}

	key := fmt.Sprintf("%s%s", <-random.ID, <-random.ID)
	kw.SetSubKey(subID, fmt.Sprintf("%d", u.ID))
	kw.SetSubKey(key, dump)

	cc, ee := strconv.ParseInt(key[:10], 16, 64)
	assert.Nil(ee)
	code := fmt.Sprintf("%d", cc)[:8]
	kw.SetSubKey(code, dump)

	return fmt.Sprintf("%s%s%s", hash, delimiter, key), code, kw.Save(exp.Duration())

}

func hasher(s string) string {
	h := sha1.New()
	h.Write([]byte(s))
	return fmt.Sprintf("%x", h.Sum(nil))
}
