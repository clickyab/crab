package controllers

import (
	"context"
	"net/http"

	"clickyab.com/crab/modules/asset/orm"
)

type regionList []orm.Region

// region return list iab categories
// @Rest {
// 		url = /region
//		method = get
//		protected = true
// }
func (c *Controller) region(ctx context.Context, r *http.Request) (regionList, error) {
	m := orm.NewOrmManager()
	return regionList(m.ListRegions()), nil
}
