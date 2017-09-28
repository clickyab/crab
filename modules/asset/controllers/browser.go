package controllers

import (
	"context"
	"net/http"

	"clickyab.com/crab/modules/asset/models"
	"github.com/clickyab/services/assert"
)

// browser return list all browsers
// @Route {
// 		url = /browser
//		method = get
//		200 = models.Browsers
//		middleware = authz.Authenticate
// }
func (c *Controller) browser(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	res, err := models.NewModelsManager().ListActiveBrowsers()
	assert.Nil(err)
	c.OKResponse(w, res)
}
