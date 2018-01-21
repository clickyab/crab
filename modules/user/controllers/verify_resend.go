package user

import (
	"context"
	"net/http"

	"clickyab.com/crab/modules/user/aaa"
	"github.com/clickyab/services/assert"
	"github.com/clickyab/services/framework/controller"
	"github.com/clickyab/services/gettext/t9e"
)

// @Validate{
// }
type verifyResendPayload struct {
	Email string `json:"email" validate:"required,email"`
}

// verifyResend will send an email again
// @Rest {
// 		url = /email/verify/resend
// 		method = post
// }
func (c *Controller) verifyResend(ctx context.Context, r *http.Request, p *verifyResendPayload) (*controller.NormalResponse, error) {
	u, e := aaa.NewAaaManager().FindUserByEmail(p.Email)
	if e != nil {
		return nil, t9e.G("cant find user")
	}
	e = verifyEmail(u, r)
	if e == errTooSoon {
		return nil, t9e.G("verify too soon error")
	}
	assert.Nil(e)
	return nil, nil
}
