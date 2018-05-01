// Code generated build with restful DO NOT EDIT.

package controllers

import (
	"context"
	"net/http"

	"github.com/clickyab/services/framework"
)

// archive will archive campaign
// @Route {
// 		url = /archive/:id
//		method = patch
//		middleware = authz.Authenticate
//		resource = archive_campaign:self
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

// assignInventory in campaign
// @Route {
// 		url = /inventory/:id
//		method = put
//		payload = assignInventoryPayload
//		middleware = authz.Authenticate
//		resource = edit_campaign:self
//		200 = orm.Campaign
//		400 = controller.ErrorResponseSimple
//		401 = controller.ErrorResponseSimple
//		403 = controller.ErrorResponseSimple
// }
func (c Controller) assignInventoryPut(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	pl := c.MustGetPayload(ctx).(*assignInventoryPayload)
	res, err := c.assignInventory(ctx, r, pl)
	if err != nil {
		framework.Write(w, err, http.StatusBadRequest)
		return
	}
	framework.Write(w, res, http.StatusOK)
}

// attributes will update campaign attribute
// @Route {
// 		url = /attributes/:id
//		method = put
//		payload = attributesPayload
//		middleware = authz.Authenticate
//		resource = edit_attributes:self
//		200 = attributesResult
//		400 = controller.ErrorResponseSimple
//		401 = controller.ErrorResponseSimple
//		403 = controller.ErrorResponseSimple
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

// updateBase campaign
// @Route {
// 		url = /base/:id
//		method = put
//		payload = campaignBase
//		middleware = authz.Authenticate
//		resource = edit_campaign:self
//		200 = updateResult
//		400 = controller.ErrorResponseSimple
//		401 = controller.ErrorResponseSimple
//		403 = controller.ErrorResponseSimple
// }
func (c Controller) updateBasePut(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	pl := c.MustGetPayload(ctx).(*campaignBase)
	res, err := c.updateBase(ctx, r, pl)
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
//		resource = edit_budget:self
//		200 = orm.Campaign
//		400 = controller.ErrorResponseSimple
//		401 = controller.ErrorResponseSimple
//		403 = controller.ErrorResponseSimple
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

// changeStatus will update campaign finance status=start,pause
// @Route {
// 		url = /status/:id
//		method = patch
//		payload = changeCampaignStatus
//		middleware = authz.Authenticate
//		resource = change_campaign_status:self
//		200 = controller.NormalResponse
//		400 = controller.ErrorResponseSimple
//		401 = controller.ErrorResponseSimple
//		403 = controller.ErrorResponseSimple
// }
func (c *Controller) changeStatusPatch(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	pl := c.MustGetPayload(ctx).(*changeCampaignStatus)
	res, err := c.changeStatus(ctx, r, pl)
	if err != nil {
		framework.Write(w, err, http.StatusBadRequest)
		return
	}
	framework.Write(w, res, http.StatusOK)
}

// copy a campaign by id
// @Route {
// 		url = /copy/:id
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

// createBase campaign
// @Route {
// 		url = /create
//		method = post
//		payload = createCampaignPayload
//		middleware = authz.Authenticate
//		resource = edit_campaign:self
//		200 = baseResult
//		400 = controller.ErrorResponseSimple
//		401 = controller.ErrorResponseSimple
//		403 = controller.ErrorResponseSimple
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

// finalize
// @Route {
// 		url = /finalize/:id
//		method = put
//		middleware = authz.Authenticate
//		resource = edit_campaign:self
//		200 = finalizeResult
//		400 = controller.ErrorResponseSimple
//		401 = controller.ErrorResponseSimple
//		403 = controller.ErrorResponseSimple
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
//		200 = campaignGetResponse
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

// getCreativeByCampaign to get creative by id
// @Route {
// 		url = /creative/:id
//		method = get
//		middleware = authz.Authenticate
//		resource = get_creative:self
//		200 = getCreativeResp
//		400 = controller.ErrorResponseSimple
//		401 = controller.ErrorResponseSimple
//		403 = controller.ErrorResponseSimple
// }
func (c Controller) getCreativeByCampaignGet(ctx context.Context, w http.ResponseWriter, r *http.Request) {

	res, err := c.getCreativeByCampaign(ctx, r)
	if err != nil {
		framework.Write(w, err, http.StatusBadRequest)
		return
	}
	framework.Write(w, res, http.StatusOK)
}

// getCampaignProgress getCampaignProgress
// @Route {
// 		url = /progress/:id
//		method = GET
//		middleware = authz.Authenticate
//		resource = get_campaign:self
//		200 = orm.CampaignProgress
//		400 = controller.ErrorResponseSimple
//		401 = controller.ErrorResponseSimple
//		403 = controller.ErrorResponseSimple
// }
func (c Controller) getCampaignProgressGet(ctx context.Context, w http.ResponseWriter, r *http.Request) {

	res, err := c.getCampaignProgress(ctx, r)
	if err != nil {
		framework.Write(w, err, http.StatusBadRequest)
		return
	}
	framework.Write(w, res, http.StatusOK)
}
