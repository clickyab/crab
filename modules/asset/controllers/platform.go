package controllers

import (
	"context"
	"net/http"

	"clickyab.com/crab/modules/asset/orm"
	orm2 "clickyab.com/crab/modules/campaign/orm"
	"github.com/clickyab/services/gettext/t9e"
	"github.com/rs/xmux"
)

type platformResponse []orm.Platform

// platform return list all is (e.g. desktop,mobile, ...)
// @Rest {
// 		url = /platform/:kind
//		method = get
//		protected = true
// }
func (c *Controller) platform(ctx context.Context, r *http.Request) (platformResponse, error) {
	kind := xmux.Param(ctx, "kind")
	campaignKind := orm2.CampaignKind(kind)
	if !campaignKind.IsValid() {
		return nil, t9e.G("campaign kind not valid")
	}
	m := orm.NewOrmManager()
	if campaignKind == orm2.AppCampaign {
		return platformResponse(m.ListPlatformsWithFilter("status=? AND name!=?", orm.EnableAssetStatus, "desktop")), nil
	}

	return platformResponse(m.ListPlatformsWithFilter("status=?", orm.EnableAssetStatus)), nil
}

// platform return list all is (e.g. desktop,mobile, ...)
// @Rest {
// 		url = /platform
//		method = get
//		protected = true
// }
func (c *Controller) allPlatform(ctx context.Context, r *http.Request) (platformResponse, error) {
	m := orm.NewOrmManager()
	return platformResponse(m.ListPlatformsWithFilter("status=?", orm.EnableAssetStatus)), nil
}
