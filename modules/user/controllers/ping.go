package user

import (
	"context"
	"net/http"

	"clickyab.com/crab/modules/user/middleware/authz"
)

// ping is for ping
// @Rest {
// 		url = /ping
//		protected = true
// 		method = get
// }
func (c *Controller) ping(ctx context.Context, r *http.Request) (*ResponseLoginOK, error) {
	user := authz.MustGetUser(ctx)
	return c.createLoginResponseWithToken(user, authz.MustGetToken(ctx)), nil
}
