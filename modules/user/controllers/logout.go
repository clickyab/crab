package user

import (
	"context"
	"net/http"

	"clickyab.com/crab/modules/domain/middleware/domain"
	"clickyab.com/crab/modules/user/aaa"
	"clickyab.com/crab/modules/user/errors"
	"clickyab.com/crab/modules/user/middleware/authz"
	"github.com/clickyab/services/assert"
	"github.com/clickyab/services/kv"
	"github.com/clickyab/services/random"
	"github.com/clickyab/services/xlog"
)

// closeSession closes current session
// @Rest {
// 		url = /logout
//		protected = true
// 		method = get
// }
func (c *Controller) closeSession(ctx context.Context, r *http.Request) (*ResponseLoginOK, error) {
	token := authz.MustGetToken(ctx)
	// check if user is impersonated
	impersonatorToken := aaa.ImpersonatorToken(token)
	if impersonatorToken != "" {
		// user is impersonated so drop user session
		err := kv.NewEavStore(token).Drop()
		assert.Nil(err)
		// return impersonator user detail
		impAccessToken := kv.NewEavStore(impersonatorToken).SubKey("token")
		m := aaa.NewAaaManager()
		user, err := m.FindUserByAccessToken(impAccessToken)
		if err != nil {
			xlog.Get(ctx).Error("find user by access token error:", err)
			return nil, errors.InvalidTokenErr
		}
		// get impersonator permissions
		currentDomain := domain.MustGetDomain(ctx)
		userPerms, err := user.GetAllUserPerms(currentDomain.ID)
		if err != nil {
			xlog.GetWithError(ctx, err).Debug("database error when get user permissions:", err)
			return nil, errors.GetUserPermsDbErr
		}
		assert.Nil(err)
		res := &ResponseLoginOK{
			Token:   impersonatorToken,
			Account: c.createUserResponse(user, userPerms, nil),
		}
		return res, nil
	}
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
