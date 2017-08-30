package authz

import (
	"context"
	"net/http"

	"github.com/clickyab/services/framework"
	"github.com/clickyab/services/permission"
	"github.com/clickyab/services/trans"
)

const (
	scopeGranted = "__granted_scope__"
	permGranted  = "__granted_perm__"
)

// Authorize is a middleware used for authorization in exchange console
func AuthorizeGenerator(resource permission.Token, scope permission.UserScope) framework.Middleware {
	return func(next framework.Handler) framework.Handler {
		return func(ctx context.Context, w http.ResponseWriter, r *http.Request) {
			u := MustGetUser(ctx)
			grantedScope, ok := u.Has(scope, resource)
			if !ok {
				grantedScope, ok = u.Has(permission.ScopeGlobal, permission.God)
				resource = permission.God
			}

			if !ok {
				framework.JSON(w, http.StatusForbidden, struct {
					Error trans.T9String `json:"error"`
				}{
					Error: trans.T("unauthorised user"),
				})
				return
			}

			ctx = context.WithValue(ctx, scopeGranted, grantedScope)
			ctx = context.WithValue(ctx, permGranted, resource)

			next(ctx, w, r)
		}
	}
}
