package controllers

import (
	"context"
	"net/http"

	"clickyab.com/crab/modules/asset/orm"
	"github.com/clickyab/services/framework/controller"
)

// Controller is the controller for the asset package
// @Route {
//		group = /asset
//		middleware = domain.Access
// }
type Controller struct {
	controller.Base
}

type manufacturers []orm.Manufacturer

// Manufacturers return list all mobile manufacturers (e.g. Apple, Samsung)
// @Route {
// 		url = /manufacturers
//		method = get
//		200 = manufacturers
//		middleware = authz.Authenticate
// }
func (c *Controller) manufacturer(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	m := orm.NewOrmManager()
	c.OKResponse(w, m.ListManufacturers())
}
