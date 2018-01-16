package user

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"

	"clickyab.com/crab/modules/user/aaa"
	"clickyab.com/crab/modules/user/mailer"
	"github.com/clickyab/services/assert"
	"github.com/clickyab/services/framework/controller"
)

// @Validate {
// }
type forgetPayload struct {
	Email string `json:"email" validate:"required,email"`
}

// forgetPassword
// @Rest {
//		url = /password/forget
// 		method = post
// }
func (c *Controller) forgetPassword(ctx context.Context, r *http.Request, p *forgetPayload) (*controller.NormalResponse, error) {
	u, err := aaa.NewAaaManager().FindUserByEmail(p.Email)
	if err != nil {
		return nil, errors.New("email not found")
	}

	ur, co, e := genVerifyCode(u, passwordVerifyPath.String())
	if e == errTooSoon {
		return nil, nil
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

	return nil, nil
}
