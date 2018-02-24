package controllers

import (
	"context"
	"net/http"

	"clickyab.com/crab/modules/asset/orm"
	"github.com/clickyab/services/gettext/t9e"
	"github.com/rs/xmux"
)

type ispResponse []orm.ISP

// isp return list all is (e.g. irancell, ...)
// @Rest {
// 		url = /isp/:kind
//		method = get
//		protected = true
// }
func (c *Controller) isp(ctx context.Context, r *http.Request) (ispResponse, error) {
	m := orm.NewOrmManager()
	kind := xmux.Param(ctx, "kind")
	if !orm.ISPKind(kind).IsValid() {
		return nil, t9e.G("not valid isp kind")

	}
	return ispResponse(m.ListISPSWithFilter("active=? AND kind=? OR kind=?", true, kind, orm.BothISPKind)), nil
}
