package user

import (
	"context"
	"net/http"

	"clickyab.com/crab/modules/user/aaa"
	"clickyab.com/crab/modules/user/middleware/authz"
	"github.com/clickyab/services/framework/controller"
	"github.com/clickyab/services/gettext/t9e"
	"github.com/clickyab/services/mysql"
)

type storePayload struct {
	Data mysql.GenericJSONField `json:"data"`
}

// register is for register user
// @Rest {
// 		url = /store
// 		method = post
// 		protected = true
// }
func (c *Controller) store(ctx context.Context, r *http.Request, p *storePayload) (*controller.NormalResponse, error) {
	currentUser := authz.MustGetUser(ctx)
	currentUser.Attributes = p.Data
	err := aaa.NewAaaManager().UpdateUser(currentUser)
	if err != nil {
		return nil, t9e.G("error while updating user")
	}
	return nil, nil
}
