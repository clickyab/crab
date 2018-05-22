package controllers

import (
	"context"
	"net/http"
	"strconv"

	"clickyab.com/crab/modules/domain/middleware/domain"
	"clickyab.com/crab/modules/financial/errors"
	"clickyab.com/crab/modules/financial/orm"
	"clickyab.com/crab/modules/user/aaa"
	"clickyab.com/crab/modules/user/middleware/authz"
	"github.com/clickyab/services/permission"
	"github.com/rs/xmux"
)

// getGateway gets a gateway by id
// @Rest {
// 		url = /gateways/:id
//		protected = true
// 		method = get
//		resource = god:global
// }
func (c *Controller) getGateway(ctx context.Context, r *http.Request) (*orm.Gateway, error) {
	userDomain := domain.MustGetDomain(ctx)
	currentUser := authz.MustGetUser(ctx)
	// check access
	_, ok := aaa.CheckPermOn(currentUser, currentUser, "god", userDomain.ID, permission.ScopeGlobal)
	if !ok {
		return nil, errors.AccessDenied
	}
	// get gateway id from url params
	gatewayID, err := strconv.ParseInt(xmux.Param(ctx, "id"), 10, 64)
	if err != nil {
		return nil, errors.InvalidIDErr
	}
	// find gateway
	m := orm.NewOrmManager()
	gatewayObj, err := m.FindGatewayByID(gatewayID)
	if err != nil {
		return nil, errors.NotFoundGateway
	}
	return gatewayObj, nil
}
