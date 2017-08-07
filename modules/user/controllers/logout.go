package user

import (
	"context"
	"net/http"

	"clickyab.com/crab/modules/user/aaa"
	"clickyab.com/crab/modules/user/middleware"
	"github.com/clickyab/services/assert"
	"github.com/clickyab/services/config"
	"github.com/clickyab/services/eav"
	"github.com/clickyab/services/random"
)

var sessionTTL = config.GetDuration("crab.session.ttl")

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

// closeAllOtherSession closes all of clients sessions but current one
// @Route {
// 		url = /session/closeother
//		method = get
//		middleware = authz.Authenticate
// }
func (c Controller) closeAllOtherSession(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	token := authz.MustGetToken(ctx)
	store := eav.NewEavStore(token)
	dbToken := store.SubKey("token")

	m := aaa.NewAaaManager()
	user, err := m.FindUserByAccessToken(dbToken)
	assert.Nil(err)

	newToken := <-random.ID
	user.AccessToken = newToken
	err = m.UpdateUser(user)
	assert.Nil(err)

	store.SetSubKey("token", newToken)
	assert.Nil(store.Save(sessionTTL))

	c.OKResponse(w, struct {
		Status string `json:"status"`
	}{
		Status: "all other sessions were deleted successfully",
	})
}
