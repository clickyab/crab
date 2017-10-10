package controllers

import (
	"context"
	"net/http"

	"clickyab.com/crab/modules/asset/orm"
)

type osResponse []orm.OS

// os return list all is (e.g. linux, ...)
// @Route {
// 		url = /os
//		method = get
//		200 = osResponse
//		middleware = authz.Authenticate
// }
func (c *Controller) os(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	m := orm.NewOrmManager()
	c.OKResponse(w, m.ListOSWithFilter("active=?", true))
}
