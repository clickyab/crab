package user

import (
	"context"
	"net/http"

	"clickyab.com/crab/modules/user/middleware/authz"
	"github.com/clickyab/services/assert"
	"github.com/clickyab/services/kv"
)

// closeSession closes current session
// @Route {
// 		url = /logout
// 		method = get
//		middleware = authz.Authenticate
// 		200 = controller.NormalResponse
// }
func (c Controller) closeSession(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	token := authz.MustGetToken(ctx)
	err := kv.NewEavStore(token).Drop()
	assert.Nil(err)

	c.OKResponse(w, nil)
}
