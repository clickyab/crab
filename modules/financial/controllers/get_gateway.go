package controllers

import (
	"context"
	"net/http"
	"strconv"

	"clickyab.com/crab/modules/financial/errors"
	"clickyab.com/crab/modules/financial/orm"
	"github.com/rs/xmux"
)

// getGateway gets a gateway by id
// @Rest {
// 		url = /gateways/:id
//		protected = true
// 		method = get
// }
func (c *Controller) getGateway(ctx context.Context, r *http.Request) (*orm.Gateway, error) {
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
