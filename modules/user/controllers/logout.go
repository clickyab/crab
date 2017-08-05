package user

import (
	"context"
	"net/http"

	"clickyab.com/crab/modules/user/middleware"
	"github.com/clickyab/services/assert"
	"github.com/clickyab/services/eav"
)

// closeSession closes current session by deleting its key in redis
// @Route {
// 		url = /session/close
//		method = get
//		middleware = authz.Authenticate
// }
func (c Controller) closeSession(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	token := authz.MustGetToken(ctx)
	err := eav.NewEavStore(token).Drop()
	assert.Nil(err)

	c.OKResponse(w, struct {
		Status string `json:"status"`
	}{
		Status: "session deleted successfully",
	})
}
