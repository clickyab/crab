package controllers

import (
	"context"
	"net/http"

	"clickyab.com/crab/modules/asset/orm"
	"github.com/clickyab/services/assert"
)

type browserResponse []orm.Browser

// browser return list all browsers
// @Route {
// 		url = /browser
//		method = get
//		200 = browserResponse
//		middleware = authz.Authenticate
// }
func (c *Controller) browser(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	browsers, err := orm.NewOrmManager().ListActiveBrowsers()
	assert.Nil(err)
	c.OKResponse(w, browsers)
}
