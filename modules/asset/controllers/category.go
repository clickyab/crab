package controllers

import (
	"context"
	"net/http"

	"clickyab.com/crab/modules/asset/orm"
)

type catResponse []orm.Category

// category return list iab categories
// @Route {
// 		url = /category
//		method = get
//		200 = catResponse
//		middleware = authz.Authenticate
// }
func (c *Controller) category(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	m := orm.NewOrmManager()
	c.OKResponse(w, m.ListCategories())
}
