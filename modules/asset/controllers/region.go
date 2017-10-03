package controllers

import (
	"context"
	"net/http"

	"clickyab.com/crab/modules/asset/orm"
)

type regionList []orm.Region

// region return list iab categories
// @Route {
// 		url = /region
//		method = get
//		200 = regionList
//		middleware = authz.Authenticate
// }
func (c *Controller) region(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	m := orm.NewOrmManager()
	c.OKResponse(w, regionList(m.ListRegions()))
}