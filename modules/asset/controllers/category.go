package controllers

import (
	"context"
	"net/http"

	"clickyab.com/crab/modules/asset/models"
	"github.com/clickyab/services/assert"
)

// category return list iab categories
// @Route {
// 		url = /category
//		method = get
//		200 = models.Categories
//		middleware = authz.Authenticate
// }
func (c *Controller) category(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	res, err := models.NewModelsManager().ListActiveCategories()
	assert.Nil(err)
	c.OKResponse(w, res)
}
