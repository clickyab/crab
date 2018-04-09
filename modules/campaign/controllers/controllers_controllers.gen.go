// Code generated build with restful DO NOT EDIT.

package controllers

import (
	"context"
	"net/http"

	"github.com/clickyab/services/framework"
)

// attributes will update campaign attribute
// @Route {
// 		url = /attributes/:id
//		method = put
//		payload = attributesPayload
//		middleware = authz.Authenticate
//		200 = orm.Campaign
//		400 = controller.ErrorResponseSimple
//		401 = controller.ErrorResponseSimple
// }
func (c *Controller) attributesPut(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	pl := c.MustGetPayload(ctx).(*attributesPayload)
	res, err := c.attributes(ctx, r, pl)
	if err != nil {
		framework.Write(w, err, http.StatusBadRequest)
		return
	}
	framework.Write(w, res, http.StatusOK)
}

// budget will update campaign finance
// @Route {
// 		url = /budget/:id
//		method = put
//		payload = budgetPayload
//		middleware = authz.Authenticate
//		200 = orm.Campaign
//		400 = controller.ErrorResponseSimple
//		401 = controller.ErrorResponseSimple
// }
func (c *Controller) budgetPut(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	pl := c.MustGetPayload(ctx).(*budgetPayload)
	res, err := c.budget(ctx, r, pl)
	if err != nil {
		framework.Write(w, err, http.StatusBadRequest)
		return
	}
	framework.Write(w, res, http.StatusOK)
}

// createBase campaign
// @Route {
// 		url = /create
//		method = post
//		payload = createCampaignPayload
//		middleware = authz.Authenticate
//		200 = orm.Campaign
//		400 = controller.ErrorResponseSimple
//		401 = controller.ErrorResponseSimple
// }
func (c Controller) createBasePost(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	pl := c.MustGetPayload(ctx).(*createCampaignPayload)
	res, err := c.createBase(ctx, r, pl)
	if err != nil {
		framework.Write(w, err, http.StatusBadRequest)
		return
	}
	framework.Write(w, res, http.StatusOK)
}

// updateBase campaign
// @Route {
// 		url = /base/:id
//		method = put
//		payload = campaignStatus
//		middleware = authz.Authenticate
//		resource = edit_campaign:self
//		200 = orm.Campaign
//		400 = controller.ErrorResponseSimple
//		401 = controller.ErrorResponseSimple
//		403 = controller.ErrorResponseSimple
// }
func (c Controller) updateBasePut(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	pl := c.MustGetPayload(ctx).(*campaignStatus)
	res, err := c.updateBase(ctx, r, pl)
	if err != nil {
		framework.Write(w, err, http.StatusBadRequest)
		return
	}
	framework.Write(w, res, http.StatusOK)
}

// finalize
// @Route {
// 		url = /finalize/:id
//		method = put
//		middleware = authz.Authenticate
//		200 = orm.Campaign
//		400 = controller.ErrorResponseSimple
//		401 = controller.ErrorResponseSimple
// }
func (c *Controller) finalizePut(ctx context.Context, w http.ResponseWriter, r *http.Request) {

	res, err := c.finalize(ctx, r)
	if err != nil {
		framework.Write(w, err, http.StatusBadRequest)
		return
	}
	framework.Write(w, res, http.StatusOK)
}

// get gets a campaign by id
// @Route {
// 		url = /get/:id
//		method = get
//		middleware = authz.Authenticate
//		resource = get_campaign:self
//		200 = orm.Campaign
//		400 = controller.ErrorResponseSimple
//		401 = controller.ErrorResponseSimple
//		403 = controller.ErrorResponseSimple
// }
func (c *Controller) getGet(ctx context.Context, w http.ResponseWriter, r *http.Request) {

	res, err := c.get(ctx, r)
	if err != nil {
		framework.Write(w, err, http.StatusBadRequest)
		return
	}
	framework.Write(w, res, http.StatusOK)
}

// budget will update campaign finance stat=archive,start,pause
// @Route {
// 		url = /:id/:stat
//		method = patch
//		middleware = authz.Authenticate
//		resource = change_campaign:self
//		200 = controller.NormalResponse
//		400 = controller.ErrorResponseSimple
//		401 = controller.ErrorResponseSimple
//		403 = controller.ErrorResponseSimple
// }
func (c *Controller) archivePatch(ctx context.Context, w http.ResponseWriter, r *http.Request) {

	res, err := c.archive(ctx, r)
	if err != nil {
		framework.Write(w, err, http.StatusBadRequest)
		return
	}
	framework.Write(w, res, http.StatusOK)
}

// copy a campaign by id
// @Route {
// 		url = /:id
//		method = patch
//		payload = copyCampaignPayload
//		middleware = authz.Authenticate
//		resource = copy_campaign:self
//		200 = orm.Campaign
//		400 = controller.ErrorResponseSimple
//		401 = controller.ErrorResponseSimple
//		403 = controller.ErrorResponseSimple
// }
func (c Controller) copyCampaignPatch(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	pl := c.MustGetPayload(ctx).(*copyCampaignPayload)
	res, err := c.copyCampaign(ctx, r, pl)
	if err != nil {
		framework.Write(w, err, http.StatusBadRequest)
		return
	}
	framework.Write(w, res, http.StatusOK)
}

// getCampaignAds get all campaign ads
// @Route {
// 		url = /get/:id/ad
//		method = get
//		middleware = authz.Authenticate
//		resource = get_banner:self
//		200 = sliceAds
//		400 = controller.ErrorResponseSimple
//		401 = controller.ErrorResponseSimple
//		403 = controller.ErrorResponseSimple
// }
func (c Controller) getCampaignAdsGet(ctx context.Context, w http.ResponseWriter, r *http.Request) {

	res, err := c.getCampaignAds(ctx, r)
	if err != nil {
		framework.Write(w, err, http.StatusBadRequest)
		return
	}
	framework.Write(w, res, http.StatusOK)
}

// getNativeData getNativeData
// @Route {
// 		url = /native/fetch
//		method = post
//		payload = getNativeDataPayload
//		middleware = authz.Authenticate
//		200 = getNativeDataResp
//		400 = controller.ErrorResponseSimple
//		401 = controller.ErrorResponseSimple
// }
func (c Controller) getNativeDataPost(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	pl := c.MustGetPayload(ctx).(*getNativeDataPayload)
	res, err := c.getNativeData(ctx, r, pl)
	if err != nil {
		framework.Write(w, err, http.StatusBadRequest)
		return
	}
	framework.Write(w, res, http.StatusOK)
}
