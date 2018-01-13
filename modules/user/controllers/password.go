package user

import (
	"context"
	"net/http"

	"github.com/clickyab/services/assert"
	"github.com/rs/xmux"

	"errors"

	"fmt"
	"net/url"

	"strings"

	"clickyab.com/crab/modules/user/aaa"
	"clickyab.com/crab/modules/user/mailer"
	"github.com/clickyab/services/trans"
)

// @Validate {
// }
type forgetPayload struct {
	Email string `json:"email" validate:"required,email"`
}

// forgetPassword
// @Route {
//		url = /password/forget
//		method = post
//		payload = forgetPayload
//		200 = controller.NormalResponse
//		400 = controller.ErrorResponseSimple
// }
func (c Controller) forgetPassword(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	p, _ := c.MustGetPayload(ctx).(*forgetPayload)
	u, err := aaa.NewAaaManager().FindUserByEmail(p.Email)
	if err != nil {
		c.BadResponse(w, errors.New("email not found"))
		return
	}

	ur, co, e := genVerifyCode(u, passwordVerifyPath)
	if e == errTooSoon {
		c.OKResponse(w, nil)
		return
	}
	assert.Nil(e)

	ul := &url.URL{
		Scheme: func() string {
			if r.TLS != nil {
				return "https"
			}
			return "http"
		}(),
		Host: r.Host,
		Path: fmt.Sprintf("/user/recover/verification/%s", ur),
	}
	temp := fmt.Sprintf(`
	%s
	code: %s
	`, ul.String(), co)

	// TODO: Change email template
	mailer.SendMail(u, "Password recovery", temp)

	c.OKResponse(w, nil)
}

// forgetCallBack is the url coming from sent email
// 		@Route {
// 		url = /password/verify/:token
// 		method = get
//		200 = ResponseLoginOK
// 		403 = controller.ErrorResponseSimple
// }
func (c Controller) checkForgetHash(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	t := xmux.Param(ctx, "token")
	u, e := verifyCode(t)
	if e != nil {
		c.ForbiddenResponse(w, nil)
		return
	}
	s, _, e := genVerifyCode(u, "change password")
	assert.Nil(e)
	c.createLoginResponseWithToken(w, u, s)
}

// @Validate {
// }
type forgetCodePayload struct {
	Email string `json:"email" validate:"required,email"`
	Code  string `json:"code" validate:"required"`
}

// forgetCallBack is the url coming from sent email
// 		@Route {
// 		url = /password/verify/
// 		method = post
//		payload = forgetCodePayload
//		200 = ResponseLoginOK
// 		403 = controller.ErrorResponseSimple
// }
func (c Controller) checkForgetCode(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	p := c.MustGetPayload(ctx).(*forgetCodePayload)
	u, e := verifyCode(fmt.Sprintf("%s%s%s", hasher(p.Email+passwordVerifyPath), delimiter, p.Code))
	if e != nil || strings.ToLower(p.Email) != strings.ToLower(u.Email) {
		c.ForbiddenResponse(w, nil)
		return
	}

	s, _, e := genVerifyCode(u, "change password")
	assert.Nil(e)
	c.createLoginResponseWithToken(w, u, s)
}

// @Validate {
// }
type callBackPayload struct {
	NewPassword string `json:"new_password" validate:"gt=5" error:"password is too short"`
}

// 		@Route {
// 		url = /password/change/:token
// 		method = put
// 		payload = callBackPayload
//		200 = ResponseLoginOK
// 		400 = controller.ErrorResponseSimple
// }
func (c Controller) changeForgetPassword(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	t := xmux.Param(ctx, "token")
	p := c.MustGetPayload(ctx).(*callBackPayload)

	u, e := verifyCode(t)
	if e != nil {
		c.ForbiddenResponse(w, e)
		return
	}

	err := u.ChangePassword(p.NewPassword)
	if err != nil {
		if err == aaa.ErrorOldPass {
			c.BadResponse(w, trans.EE(err))
			return
		}
		c.BadResponse(w, trans.E("Can't change password!"))
		return
	}

	c.createLoginResponse(w, u)
}

// @Validate {
// }
type changePassword struct {
	CurrentPassword string `json:"current_password" validate:"gt=5"`
	NewPassword     string `json:"new_password" validate:"gt=5"`
}

// changePassword
// @Route {
//		url = /password/change
//		method = put
//		payload = changePassword
//		middleware = authz.Authenticate
//		200 = controller.NormalResponse
//		400 = controller.ErrorResponseSimple
// }
func (c Controller) changePassword(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	u := c.MustGetUser(ctx)
	payload := c.MustGetPayload(ctx).(*changePassword)
	e := u.UpdatePassword(payload.CurrentPassword, payload.NewPassword)
	if e != nil {
		switch e {
		case aaa.ErrorWrongPassword:
			c.BadResponse(w, trans.EE(e))
			return
		case aaa.ErrorOldPass:
			c.BadResponse(w, trans.EE(e))
			return
		default:
			assert.Nil(e)
		}
	}

	c.OKResponse(w, nil)
}
