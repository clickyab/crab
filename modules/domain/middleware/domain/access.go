package domain

import (
	"context"
	"net/http"

	"clickyab.com/crab/modules/domain/dmn"
	"github.com/clickyab/services/assert"
	"github.com/clickyab/services/framework"
	"github.com/clickyab/services/framework/controller"
	"github.com/clickyab/services/trans"
	"github.com/sirupsen/logrus"
)

type contextKey string

const (
	// ContextDomain is the context key for the body unmarshalled object
	ContextDomain contextKey = "_domain"
)

// Access is a middleware used for domain access
func Access(next framework.Handler) framework.Handler {
	return func(ctx context.Context, w http.ResponseWriter, r *http.Request) {
		// check if d is valid
		d, err := dmn.NewDmnManager().FindActiveDomainByName(r.Host)
		if err != nil {
			logrus.WithError(err).WithField("domain", r.Host).Debug("domain not found")
			framework.JSON(w, http.StatusNotFound, controller.ErrorResponseSimple{Error: trans.E(http.StatusText(http.StatusNotFound))})
			return
		}
		ctx = context.WithValue(ctx, ContextDomain, d)
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
