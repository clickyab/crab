package user

import (
	"context"
	"net/http"

	"clickyab.com/crab/modules/user/middleware/authz"
)

// logout is for the logout from the system
// @Route {
// 		url = /ping
//		method = get
//		middleware = authz.Authenticate
//		200 = responseLoginOK
// }
func (c Controller) ping(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	user := authz.MustGetUser(ctx)
	c.createLoginResponseWithToken(w, user, authz.MustGetToken(ctx))
}
