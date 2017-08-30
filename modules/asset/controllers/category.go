package controllers

import (
	"context"
	"net/http"

	"clickyab.com/crab/modules/asset/asst"
)

type catResponse []asst.Category

// category return list iab categories
// @Route {
// 		url = /category
//		method = get
//		200 = catResponse
//		middleware = authz.Authenticate
// }
func (c *Controller) category(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	m := asst.NewAsstManager()
	c.OKResponse(w, m.ListCategories())
}
