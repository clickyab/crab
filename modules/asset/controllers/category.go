package controllers

import (
	"context"
	"net/http"

	"clickyab.com/crab/modules/asset/orm"
)

type catResponse []orm.Category

// category return list iab categories
// @Rest {
// 		url = /category
//		method = get
//		protected = true
// }
func (c *Controller) category(ctx context.Context, r *http.Request) (catResponse, error) {
	m := orm.NewOrmManager()
	return catResponse(m.ListCategoriesWithFilter("active=?", true)), nil
}
