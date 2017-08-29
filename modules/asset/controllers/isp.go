package controllers

import (
	"context"
	"net/http"

	"clickyab.com/crab/modules/asset/orm"
)

type ispResponse []orm.ISP

// isp return list all is (e.g. irancell, ...)
// @Route {
// 		url = /isp
//		method = get
//		200 = ispResponse
//		middleware = authz.Authenticate
// }
func (c *Controller) isp(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	m := orm.NewOrmManager()
	c.OKResponse(w, m.ListISPS())
}
