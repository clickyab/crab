package user

import (
	"context"
	"net/http"

	"errors"

	"clickyab.com/crab/modules/user/aaa"
	"github.com/clickyab/services/assert"
	"github.com/clickyab/services/framework/controller"
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
		return nil, errors.New("cant find user")
	}
	e = verifyEmail(u, r)
	if e == errTooSoon {
		return nil, errors.New("verify too soon error")
	}
	assert.Nil(e)
	return nil, nil
}
