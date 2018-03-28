package controllers

import (
	"context"
	"net/http"

	"clickyab.com/crab/modules/asset/orm"
	orm2 "clickyab.com/crab/modules/campaign/orm"
	"github.com/clickyab/services/gettext/t9e"
	"github.com/rs/xmux"
)

type osResponse []orm.OS

// os return list all is (e.g. linux, ...)
// @Rest {
// 		url = /os/:kind
//		method = get
//		protected = true
// }
func (c *Controller) os(ctx context.Context, r *http.Request) (osResponse, error) {
	kind := xmux.Param(ctx, "kind")
	campaignKind := orm2.CampaignKind(kind)
	if !campaignKind.IsValid() {
		return nil, t9e.G("campaign kind not valid")
	}
	m := orm.NewOrmManager()
	if campaignKind == orm2.AppCampaign {
		return osResponse(m.ListOSWithFilter("status=? AND name IN (?,?)", orm.EnableAssetStatus, "android", "ios")), nil
	}

	return osResponse(m.ListOSWithFilter("status=?", orm.EnableAssetStatus)), nil
}

// os return list all is (e.g. linux, ...)
// @Rest {
// 		url = /os
//		method = get
//		protected = true
// }
func (c *Controller) allOs(ctx context.Context, r *http.Request) (osResponse, error) {
	m := orm.NewOrmManager()
	return osResponse(m.ListOSWithFilter("status=?", orm.EnableAssetStatus)), nil
}
