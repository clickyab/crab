package controllers

import (
	"context"
	"net/http"

	"clickyab.com/crab/modules/asset/orm"
)

type platformResponse []orm.Platform

// platform return list all is (e.g. desktop,mobile, ...)
// @Rest {
// 		url = /platform
//		method = get
//		protected = true
// }
func (c *Controller) platform(ctx context.Context, r *http.Request) (platformResponse, error) {
	m := orm.NewOrmManager()
	return platformResponse(m.ListPlatformsWithFilter("active=?", true)), nil
}
