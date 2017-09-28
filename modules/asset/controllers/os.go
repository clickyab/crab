package controllers

import (
	"context"
	"net/http"

	"clickyab.com/crab/modules/asset/models"
	"github.com/clickyab/services/assert"
)

// os return list all is (e.g. linux, ...)
// @Route {
// 		url = /os
//		method = get
//		200 = models.OSes
//		middleware = authz.Authenticate
// }
func (c *Controller) os(ctx context.Context, w http.ResponseWriter, r *http.Request) {

	res, err := models.NewModelsManager().ListActiveOSes()
	assert.Nil(err)
	c.OKResponse(w, res)
}
