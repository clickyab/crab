package controllers

import (
	"context"
	"net/http"

	"clickyab.com/crab/modules/asset/models"
	"github.com/clickyab/services/assert"
)

// region return list iab categories
// @Route {
// 		url = /region
//		method = get
//		200 = models.Regions
//		middleware = authz.Authenticate
// }
func (c *Controller) region(ctx context.Context, w http.ResponseWriter, r *http.Request) {

	res, err := models.NewModelsManager().ListActiveRegions()
	assert.Nil(err)
	c.OKResponse(w, res)
}
