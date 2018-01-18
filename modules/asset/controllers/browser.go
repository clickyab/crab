package controllers

import (
	"context"
	"net/http"

	"clickyab.com/crab/modules/asset/orm"
	"github.com/clickyab/services/assert"
)

type browserResponse []orm.Browser

// browser return list all browsers
// @Rest {
// 		url = /browser
//		method = get
//		protected = true
// }
func (c *Controller) browser(ctx context.Context, r *http.Request) (browserResponse, error) {
	browsers, err := orm.NewOrmManager().ListActiveBrowsers()
	assert.Nil(err)
	return browserResponse(browsers), nil
}
