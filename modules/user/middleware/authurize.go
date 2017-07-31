package authz

import (
	"context"
	"net/http"

	"github.com/clickyab/services/framework"
	"github.com/clickyab/services/permission"
)

// Authorize is a middleware used for authorization in exchange console
func AuthorizeGenerator(resource permission.Token, scope permission.UserScope) framework.Middleware {
	return func(next framework.Handler) framework.Handler {
		return func(ctx context.Context, w http.ResponseWriter, r *http.Request) {
			// TODO : place holder for real authorization
			next(ctx, w, r)
		}
	}
}
