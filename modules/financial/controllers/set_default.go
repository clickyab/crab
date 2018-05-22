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
	"github.com/clickyab/services/xlog"
	"github.com/rs/xmux"
)

// setGatewayDefault to financial, to set a gateway to default
// @Rest {
// 		url = /gateways/:id
//		protected = true
// 		method = patch
// 		resource = set_default_gateway:global
// }
func (c *Controller) setGatewayDefault(ctx context.Context, r *http.Request) (*orm.Gateway, error) {
	currentDomain := domain.MustGetDomain(ctx)
	currentUser := authz.MustGetUser(ctx)
	//check if gateway id is valid
	idInt, err := strconv.ParseInt(xmux.Param(ctx, "id"), 10, 0)
	if err != nil {
		return nil, errors.InvalidIDErr
	}
	m := orm.NewOrmManager()
	targetGateway, err := m.FindGatewayByID(idInt)
	if err != nil {
		return nil, errors.InvalidIDErr
	}
	// check permission
	_, ok := aaa.CheckPermOn(currentUser, currentUser, "set_default_gateway", currentDomain.ID, permission.ScopeGlobal)
	if !ok {
		return nil, errors.AccessDenied
	}
	// set gateway status to default and all others to not default
	err = m.JustDefault(targetGateway)
	if err != nil {
		xlog.GetWithError(ctx, err).Debug("change gateway is_default error")
		return nil, errors.EditGatewayErr
	}
	return targetGateway, nil
}
