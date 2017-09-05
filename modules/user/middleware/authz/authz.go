package authz

import (
	"context"
	"net/http"

	"clickyab.com/crab/modules/domain/dmn"
	"clickyab.com/crab/modules/domain/middleware/domain"
	"clickyab.com/crab/modules/user/aaa"
	"github.com/clickyab/services/assert"
	"github.com/clickyab/services/framework"
	"github.com/clickyab/services/framework/controller"
	"github.com/clickyab/services/kv"
	"github.com/clickyab/services/trans"
	"github.com/sirupsen/logrus"
)

type dataType string

const dataKey dataType = "__user_data__"
const tokenKey dataType = "__token_data__"

// Authenticate is a middleware used for authorization in exchange console
func Authenticate(next framework.Handler) framework.Handler {
	return func(ctx context.Context, w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("token")
		logrus.WithField("token", token).Debug("user token in authz")
		if token != "" {
			val := kv.NewEavStore(token).SubKey("token")
			if val != "" {
				// TODO : Write me
				userDomain, ok := ctx.Value(domain.ContextDomain).(*dmn.Domain)
				assert.True(ok, "[BUG] no domain in context")
				usr, err := aaa.NewAaaManager().FindUserByAccessTokenDomain(val, userDomain.ID)
				if err == nil {
					ctx = context.WithValue(ctx, dataKey, usr)
					ctx = context.WithValue(ctx, tokenKey, token)
					next(ctx, w, r)
					return
				}
			}
		}
		framework.JSON(w, http.StatusUnauthorized, controller.ErrorResponseSimple{Error: trans.E("Unauthorized")})
	}
}

// GetUser is the helper function to extract user data from context
func GetUser(ctx context.Context) (*aaa.User, bool) {
	rd, ok := ctx.Value(dataKey).(*aaa.User)
	if !ok {
		return nil, false
	}

	return rd, true
}

// MustGetUser try to get user data, or panic if there is no user data
func MustGetUser(ctx context.Context) *aaa.User {
	rd, ok := GetUser(ctx)
	assert.True(ok, "[BUG] no user in context")
	return rd
}

// GetToken is the helper function to extract user data from context
func GetToken(ctx context.Context) (string, bool) {
	rd, ok := ctx.Value(tokenKey).(string)
	if !ok {
		return "", false
	}

	return rd, true
}

// MustGetToken try to get user data, or panic if there is no user data
func MustGetToken(ctx context.Context) string {
	rd, ok := GetToken(ctx)
	assert.True(ok, "[BUG] no user in context")
	return rd
}
