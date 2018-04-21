package controllers

import (
	"context"
	"net/http"

	"clickyab.com/crab/modules/financial/orm"
)

type getGateResp []orm.Gateway

// getGateways get payment data
// @Rest {
// 		url = /gateways
//		protected = true
// 		method = get
// }
func (c *Controller) getGateways(ctx context.Context, r *http.Request) (getGateResp, error) {
	gateways := orm.NewOrmManager().ListGateways()
	return getGateResp(gateways), nil
}
