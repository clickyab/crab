package user

import (
	"context"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

	"crypto/sha1"
	"fmt"

	"clickyab.com/crab/modules/user/aaa"
	"clickyab.com/crab/modules/user/mailer"
	"github.com/clickyab/services/assert"
	"github.com/clickyab/services/config"
	"github.com/clickyab/services/gettext/t9e"
	"github.com/clickyab/services/kv"
	"github.com/clickyab/services/random"
	"github.com/clickyab/services/xlog"
	"github.com/rs/xmux"
)

const (
	verifyKeyPrefix = "VERIFY"
	checkToken      = "vk"
	userIDRedisKey  = "uid"
)

var (
	errTooSoon         error = t9e.G("code has been sent")
	exp                      = config.RegisterDuration("crab.modules.user.verification.ttl", 5*time.Hour, "how long the token should be saved")
	resend                   = config.RegisterDuration("crab.modules.user.verification.resend", 1*time.Minute, "Duration between resend")
	emailVerifyPath          = config.RegisterString("crab.modules.user.verification.email.path", "user/email/verify", "email verify client url")
	passwordVerifyPath       = config.RegisterString("crab.modules.user.verification.password.path", "user/password/verify", "password verify client url")
)

func verifyEmail(u *aaa.User, r *http.Request) error {
	h, c, e := genVerifyCode(u, emailVerifyPath.String())
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
	mailer.SendMail(u, t9e.G("Welcome to Clickyab!").String(), temp)
	return nil
}

const delimiter = "-"

func verifyCode(ctx context.Context, c string) (*aaa.User, error) {
	data := strings.Split(c, delimiter)
	if len(data) != 2 {
		return nil, t9e.G("code is not valid")
	}

	userEmailHash, verifyToken := data[0], data[1]

	kw := kv.NewEavStore(fmt.Sprintf("%s_%s", verifyKeyPrefix, userEmailHash))

	if kw.SubKey(verifyToken) != checkToken {
		return nil, t9e.G("code is not valid")
	}

	userID := kw.SubKey(userIDRedisKey)
	if userID == "" {
		return nil, t9e.G("can't find user")
	}

	id, err := strconv.ParseInt(userID, 10, 64)
	if err != nil {
		return nil, t9e.G("invalid user id")
	}

	m := aaa.NewAaaManager()
	user, err := m.FindUserByID(id)

	if err == nil {
		_ = kw.Drop()
	} else {
		xlog.GetWithError(ctx, err).Debug("can't find user in check verify code")

		err = t9e.G("can't find user")
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
		return "", "", t9e.G("can't generate verify short code")
	}

	verifyShortCode := fmt.Sprintf("%d", intToken)[:8]

	kw.SetSubKey(verifyToken, checkToken)
	kw.SetSubKey(verifyShortCode, checkToken)
	kw.SetSubKey(userIDRedisKey, fmt.Sprintf("%d", user.ID))

	return fmt.Sprintf("%s%s%s", emailHash, delimiter, verifyToken), verifyShortCode, kw.Save(exp.Duration())
}

func hasher(s string) string {
	h := sha1.New()

	_, err := h.Write([]byte(s))
	assert.Nil(err)

	return fmt.Sprintf("%x", h.Sum(nil))
}

// verifyEmail is verify code
// @Rest {
// 		url = /email/verify/:token
// 		method = get
// }
func (c *Controller) verifyEmail(ctx context.Context, r *http.Request) (*ResponseLoginOK, error) {
	u, e := verifyCode(ctx, xmux.Param(ctx, "token"))

	if e != nil {
		return nil, e
	}

	if u.Status != aaa.RegisteredUserStatus {
		return nil, t9e.G("user status is not registered")
	}

	u.Status = aaa.ActiveUserStatus
	assert.Nil(aaa.NewAaaManager().UpdateUser(u))
	return c.createLoginResponse(u), nil
}
