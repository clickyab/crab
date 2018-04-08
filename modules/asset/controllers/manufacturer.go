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
// @Rest {
// 		url = /manufacturers
//		method = get
//		protected = true
// }
func (c *Controller) manufacturer(ctx context.Context, r *http.Request) (manufacturers, error) {
	m := orm.NewOrmManager()
	return manufacturers(m.ListManufacturersWithFilter("status=?", orm.EnableAssetStatus)), nil
}
