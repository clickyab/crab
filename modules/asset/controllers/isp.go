package controllers

import (
	"context"
	"net/http"

	"errors"

	"clickyab.com/crab/modules/asset/orm"
	"github.com/rs/xmux"
)

type ispResponse []orm.ISP

// isp return list all is (e.g. irancell, ...)
// @Route {
// 		url = /isp/:kind
//		method = get
//		200 = ispResponse
//		400 = controller.ErrorResponseSimple
//		middleware = authz.Authenticate
// }
func (c *Controller) isp(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	m := orm.NewOrmManager()
	kind := xmux.Param(ctx, "kind")
	if !orm.ISPKind(kind).IsValid() {
		c.BadResponse(w, errors.New("not valid isp kind"))
		return
	}
	c.OKResponse(w, m.ListISPSWithFilter("active=? AND kind=? OR kind=?", true, kind, orm.BothISPKind))
}
