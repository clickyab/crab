package controllers

import (
	"context"
	"net/http"

	"clickyab.com/crab/modules/asset/orm"
)

type platformResponse []orm.Platform

// platform return list all is (e.g. desktop,mobile, ...)
// @Route {
// 		url = /platform
//		method = get
//		200 = platformResponse
//		middleware = authz.Authenticate
// }
func (c *Controller) platform(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	m := orm.NewOrmManager()
	c.OKResponse(w, m.ListPlatformsWithFilter("active=?", true))
}
