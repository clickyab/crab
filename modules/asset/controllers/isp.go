package controllers

import (
	"context"
	"net/http"

	"clickyab.com/crab/modules/asset/models"
	"github.com/clickyab/services/assert"
)

type ispResponse []models.ISP

// isp return list all is (e.g. irancell, ...)
// @Route {
// 		url = /isp
//		method = get
//		200 = models.ISPs
//		middleware = authz.Authenticate
// }
func (c *Controller) isp(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	res, err := models.NewModelsManager().ListActiveISPs()
	assert.Nil(err)
	c.OKResponse(w, ispResponse(res))
}
