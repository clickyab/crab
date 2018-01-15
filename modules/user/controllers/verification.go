package user

import (
	"context"
	"fmt"
	"net/http"
	"strconv"

	"github.com/sirupsen/logrus"

	"errors"

	"time"

	"crypto/sha1"

	"clickyab.com/crab/modules/user/aaa"
	"github.com/clickyab/services/assert"
	"github.com/clickyab/services/kv"

	"strings"

	"net/url"

	"clickyab.com/crab/modules/user/mailer"
	"github.com/clickyab/services/config"
	"github.com/clickyab/services/random"
	"github.com/clickyab/services/trans"
	"github.com/rs/xmux"
)

const (
	verifyKeyPrefix         = "VERIFY"
	verifyTokenRedisKey     = "verifyToken"
	verifyShortCodeRedisKey = "verifyShortCode"
	userIDRedisKey          = "userId"
	emailVerifyPath         = "user/email/verify"
	passwordVerifyPath      = "user/password/verify"
)

var (
	errTooSoon = errors.New("code has been sent")
	exp        = config.RegisterDuration("crab.modules.user.verification.ttl", 5*time.Hour, "how long the token should be saved")
	resend     = config.RegisterDuration("crab.modules.user.verification.resend", 1*time.Minute, "Duration between resend")
)

// verifyId is verify code
// @Route {
// 		url = /email/verify/:token
//		method = get
//		200 = ResponseLoginOK
//		401 = controller.ErrorResponseSimple
//		403 = controller.ErrorResponseSimple
// }
func (ctrl *Controller) verifyEmail(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	u, e := verifyCode(xmux.Param(ctx, "token"))

	if e != nil {
		ctrl.BadResponse(w, e)
		return
	}

	if u.Status != aaa.RegisteredUserStatus {
		ctrl.BadResponse(w, errors.New("User status is not registered"))
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
//		200 = ResponseLoginOK
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
	Email string `json:"email" validate:"required,email"`
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
	if e == errTooSoon {
		ctrl.OKResponse(w, nil)
		return
	}
	assert.Nil(e)
	ctrl.OKResponse(w, nil)
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
		Path: fmt.Sprintf("/api/user/email/verify/%s", h),
	}
	temp := fmt.Sprintf(`
	%s
	code: %s
	`, ul.String(), c)
	// TODO: Change email template
	mailer.SendMail(u, trans.T("Welcome to Clickyab!").String(), temp)
	return nil
}

const delimiter = "-"

func verifyCode(c string) (*aaa.User, error) {
	data := strings.Split(c, delimiter)
	if len(data) != 2 {
		return nil, errors.New("code is not valid")
	}

	userEmailHash, verifyToken := data[0], data[1]

	kw := kv.NewEavStore(fmt.Sprintf("%s_%s", verifyKeyPrefix, userEmailHash))

	if kw.SubKey(verifyTokenRedisKey) != verifyToken {
		return nil, errors.New("code is not valid")
	}

	userID := kw.SubKey(userIDRedisKey)
	if userID == "" {
		return nil, errors.New("Can't find user")
	}

	id, err := strconv.ParseInt(userID, 10, 64)
	if err != nil {
		return nil, errors.New("invalid user id")
	}

	m := aaa.NewAaaManager()
	user, err := m.FindUserByID(id)

	if err == nil {
		_ = kw.Drop()
	} else {
		logrus.Debug(err)

		err = errors.New("Can't find user")
	}

	return user, err
}

func genVerifyCode(user *aaa.User, salt string) (string, string, error) {
	assert.True(salt != "")

	emailHash := hasher(user.Email + salt)
	redisKey := fmt.Sprintf("%s_%s", verifyKeyPrefix, emailHash)

	kw := kv.NewEavStore(redisKey)

	if len(kw.AllKeys()) != 0 {
		if exp.Duration()-resend.Duration() < kw.TTL() {
			return "", "", errTooSoon
		}
	}

	verifyToken := fmt.Sprintf("%s%s", <-random.ID, <-random.ID)
	intToken, err := strconv.ParseInt(verifyToken[:10], 16, 64)
	if err != nil {
		return "", "", errors.New("Can't generate verify short code")
	}

	verifyShortCode := fmt.Sprintf("%d", intToken)[:8]

	kw.SetSubKey(verifyTokenRedisKey, verifyToken)
	kw.SetSubKey(verifyShortCodeRedisKey, verifyShortCode)
	kw.SetSubKey(userIDRedisKey, fmt.Sprintf("%d", user.ID))

	return fmt.Sprintf("%s%s%s", emailHash, delimiter, verifyToken), verifyShortCode, kw.Save(exp.Duration())
}

func hasher(s string) string {
	h := sha1.New()

	_, err := h.Write([]byte(s))
	assert.Nil(err)

	return fmt.Sprintf("%x", h.Sum(nil))
}
