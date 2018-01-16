package user

import (
	"context"
	"net/http"

	"clickyab.com/crab/modules/user/aaa"
	"github.com/clickyab/services/assert"
	"github.com/clickyab/services/framework/controller"
)

// @Validate {
// }
type changePassword struct {
	CurrentPassword string `json:"current_password" validate:"gt=5"`
	NewPassword     string `json:"new_password" validate:"gt=5"`
}

// changePassword change password
// @Rest {
//		url = /password/change
//		protected = true
// 		method = put
// }
func (c *Controller) changePassword(ctx context.Context, r *http.Request, p *changePassword) (*controller.NormalResponse, error) {
	u := c.MustGetUser(ctx)
	e := u.UpdatePassword(p.CurrentPassword, p.NewPassword)
	if e != nil {
		switch e {
		case aaa.ErrorWrongPassword:
			return nil, e
		case aaa.ErrorOldPass:
			return nil, e
		default:
			assert.Nil(e)
		}
	}
	return nil, nil
}
