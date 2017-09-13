package user

import (
	"context"
	"net/http"

	"github.com/clickyab/services/assert"
	"github.com/clickyab/services/notification"
	"github.com/rs/xmux"

	"errors"

	"clickyab.com/crab/modules/user/aaa"
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

	ur, e := genVerificationURL(u, passwordVerifyPath, r)
	if e == tooSoonError {
		c.OKResponse(w, nil)
		return
	}
	assert.Nil(e)

	// TODO: Change email template
	notification.Send(trans.T("Password recovery").String(), ur.String(), notification.Packet{
		Platform: notification.MailType,
		To:       []string{u.Email},
	})

	c.OKResponse(w, nil)
}

// forgetCallBack is the url coming from sent email
// 		@Route {
// 		url = /password/verify/:token
// 		method = get
//		200 = responseLoginOK
// 		403 = controller.ErrorResponseSimple
// }
func (c Controller) checkForgetCode(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	t := xmux.Param(ctx, "token")
	u, e := verifyCode(t)
	if e != nil {
		c.ForbiddenResponse(w, nil)
		return
	}
	s, e := genVerifyCode(u, "change password")
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
//		200 = responseLoginOK
// 		400 = controller.ErrorResponseSimple
// }
func (c Controller) changeForgetPassword(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	t := xmux.Param(ctx, "token")
	p := c.MustGetPayload(ctx).(*callBackPayload)

	u, e := verifyCode(t)
	if e != nil {
		c.ForbiddenResponse(w, nil)
		return
	}
	u.ChangePassword(p.NewPassword)
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
