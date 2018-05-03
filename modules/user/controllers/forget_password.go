package user

import (
	"context"
	"fmt"
	"net/http"
	"net/url"

	"clickyab.com/crab/modules/user/aaa"
	"clickyab.com/crab/modules/user/mailer"
	"github.com/clickyab/services/assert"
	"github.com/clickyab/services/config"
	"github.com/clickyab/services/framework"
	"github.com/clickyab/services/framework/controller"
	"github.com/clickyab/services/gettext/t9e"
)

var frontRecoverPath = config.RegisterString("crab.modules.user.recover.front.path", "/user/recover/verification", "front redirect forget pass route")

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
		return nil, t9e.G("user with email %s not found", p.Email)
	}

	ur, co, e := genVerifyCode(u, passwordVerifyPath.String())
	if e == errTooSoon {
		return nil, nil
	}
	assert.Nil(e)

	ul := &url.URL{
		Scheme: framework.Scheme(r),
		Host:   r.Host,
		Path:   fmt.Sprintf(frontRecoverPath.String()+"/%s", ur),
	}
	temp := fmt.Sprintf(`
	%s
	code: %s
	`, ul.String(), co)

	// TODO: Change email template
	mailer.SendMail(u, "Password recovery", temp)

	return nil, nil
}
