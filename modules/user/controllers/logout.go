package user

import (
	"context"
	"net/http"

	"clickyab.com/crab/modules/user/middleware/authz"
	"clickyab.com/crab/modules/user/models"
	"github.com/clickyab/services/assert"
	"github.com/clickyab/services/kv"
	"github.com/clickyab/services/random"
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

// closeAllOtherSession closes all of clients sessions but current one
// @Route {
// 		url = /logout/closeother
//		method = get
//		middleware = authz.Authenticate
//		200 = ResponseLoginOK
// }
func (c Controller) closeAllOtherSession(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	token := authz.MustGetToken(ctx)
	store := kv.NewEavStore(token)
	dbToken := store.SubKey("token")

	m := models.NewModelsManager()

	user, err := m.FindUserByAccessToken(dbToken)
	assert.Nil(err)

	newToken := <-random.ID
	user.AccessToken = newToken

	err = m.UpdateUser(user)
	assert.Nil(err)

	c.createLoginResponse(w, user)
}
