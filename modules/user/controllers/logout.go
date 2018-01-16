package user

import (
	"context"
	"net/http"

	"clickyab.com/crab/modules/user/aaa"
	"clickyab.com/crab/modules/user/middleware/authz"
	"github.com/clickyab/services/assert"
	"github.com/clickyab/services/framework/controller"
	"github.com/clickyab/services/kv"
	"github.com/clickyab/services/random"
)

// closeSession closes current session
// @Rest {
// 		url = /logout
//		protected = true
// 		method = get
// }
func (c *Controller) closeSession(ctx context.Context, r *http.Request) (*controller.NormalResponse, error) {
	token := authz.MustGetToken(ctx)
	err := kv.NewEavStore(token).Drop()
	assert.Nil(err)
	return nil, nil
}

// closeSession closes current session
// @Rest {
// 		url = /logout/closeother
//		protected = true
// 		method = get
// }
func (c *Controller) closeAllOtherSession(ctx context.Context, r *http.Request) (*ResponseLoginOK, error) {
	token := authz.MustGetToken(ctx)
	store := kv.NewEavStore(token)
	dbToken := store.SubKey("token")

	m := aaa.NewAaaManager()

	user, err := m.FindUserByAccessToken(dbToken)
	assert.Nil(err)

	newToken := <-random.ID
	user.AccessToken = newToken

	err = m.UpdateUser(user)
	assert.Nil(err)

	return c.createLoginResponse(user), nil
}
