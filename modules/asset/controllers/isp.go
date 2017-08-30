package controllers

import (
	"context"
	"net/http"

	"clickyab.com/crab/modules/asset/asst"
)

type ispResponse []asst.ISP

// isp return list all is (e.g. irancell, ...)
// @Route {
// 		url = /isp
//		method = get
//		200 = ispResponse
//		middleware = authz.Authenticate
// }
func (c *Controller) isp(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	m := asst.NewAsstManager()
	c.OKResponse(w, m.ListISPS())
}
