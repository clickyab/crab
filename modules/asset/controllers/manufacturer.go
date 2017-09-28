package controllers

import (
	"context"
	"net/http"

	"clickyab.com/crab/modules/asset/models"
	"github.com/clickyab/services/assert"
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

// Manufacturers return list all mobile manufacturers (e.g. Apple, Samsung)
// @Route {
// 		url = /manufacturers
//		method = get
//		200 = models.Manufacturers
//		middleware = authz.Authenticate
// }
func (c *Controller) manufacturer(ctx context.Context, w http.ResponseWriter, r *http.Request) {

	res, err := models.NewModelsManager().ListActiveManufacturers()
	assert.Nil(err)
	c.OKResponse(w, res)
}
