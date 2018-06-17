package controllers

import (
	"context"
	"net/http"

	"clickyab.com/crab/modules/asset/orm"
)

type osResponse []orm.OSList

// os return list all is (e.g. linux, ...)
// @Rest {
// 		url = /os
//		method = get
//		protected = true
// }
func (c *Controller) os(ctx context.Context, r *http.Request) (osResponse, error) {
	m := orm.NewOrmManager()
	return osResponse(m.ListOSWithBrowser()), nil
}
