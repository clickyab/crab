package middleware

import (
	"context"
	"net/http"

	"clickyab.com/crab/modules/domain/dmn"
	"github.com/clickyab/services/assert"
	"github.com/clickyab/services/framework"
	"github.com/clickyab/services/framework/controller"
	"github.com/clickyab/services/trans"
)

type contextKey string

const (
	// ContextBody is the context key for the body unmarshalled object
	ContextDomain contextKey = "_domain"
)

// Access is a middleware used for domain access
func Access(next framework.Handler) framework.Handler {
	return func(ctx context.Context, w http.ResponseWriter, r *http.Request) {
		rUrl := r.URL
		// check if domain is valid
		domain, err := dmn.NewDmnManager().FindActiveDomainByName(rUrl.Host)
		if err != nil {
			framework.JSON(w, http.StatusUnauthorized, controller.ErrorResponseSimple{Error: trans.E("Unauthorized")})
			return
		}
		ctx = context.WithValue(ctx, ContextDomain, domain)
		next(ctx, w, r)
	}
}

// MustGetDomain is the helper function to extract user domain from context
func MustGetDomain(ctx context.Context) *dmn.Domain {
	rd, ok := GetDomain(ctx)
	assert.True(ok, "[BUG] no domain in context")
	return rd
}

// GetDomain is the helper function to extract user domain from context
func GetDomain(ctx context.Context) (*dmn.Domain, bool) {
	rd, ok := ctx.Value(ContextDomain).(*dmn.Domain)
	if !ok {
		return nil, false
	}

	return rd, true
}
