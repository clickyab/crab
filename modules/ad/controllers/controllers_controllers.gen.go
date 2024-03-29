// Code generated build with restful DO NOT EDIT.

package controllers

import (
	"context"
	"net/http"

	"github.com/clickyab/services/framework"
)

// acceptCampaignCreativeStatus to campaign
// @Route {
// 		url = /accept-campaign-creative/:id
//		method = patch
//		payload = changeStatus
//		middleware = authz.Authenticate
//		resource = change_creative_status:superGlobal
//		200 = CreativeStatusChangeResult
//		400 = controller.ErrorResponseSimple
//		401 = controller.ErrorResponseSimple
//		403 = controller.ErrorResponseSimple
// }
func (c Controller) acceptCampaignCreativeStatusPatch(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	pl := c.MustGetPayload(ctx).(*changeStatus)
	res, err := c.acceptCampaignCreativeStatus(ctx, r, pl)
	if err != nil {
		framework.Write(w, err, http.StatusBadRequest)
		return
	}
	framework.Write(w, res, http.StatusOK)
}

// addBannerCreative to campaign
// @Route {
// 		url = /banner
//		method = post
//		payload = createBannerPayload
//		middleware = authz.Authenticate
//		resource = create_creative:self
//		200 = createBannerResponse
//		400 = controller.ErrorResponseSimple
//		401 = controller.ErrorResponseSimple
//		403 = controller.ErrorResponseSimple
// }
func (c Controller) addBannerCreativePost(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	pl := c.MustGetPayload(ctx).(*createBannerPayload)
	res, err := c.addBannerCreative(ctx, r, pl)
	if err != nil {
		framework.Write(w, err, http.StatusBadRequest)
		return
	}
	framework.Write(w, res, http.StatusOK)
}

// addNativeCreative to campaign
// @Route {
// 		url = /native
//		method = post
//		payload = createNativePayload
//		middleware = authz.Authenticate
//		resource = create_creative:self
//		200 = orm.CreativeSaveResult
//		400 = controller.ErrorResponseSimple
//		401 = controller.ErrorResponseSimple
//		403 = controller.ErrorResponseSimple
// }
func (c Controller) addNativeCreativePost(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	pl := c.MustGetPayload(ctx).(*createNativePayload)
	res, err := c.addNativeCreative(ctx, r, pl)
	if err != nil {
		framework.Write(w, err, http.StatusBadRequest)
		return
	}
	framework.Write(w, res, http.StatusOK)
}

// addVastCreative to campaign
// @Route {
// 		url = /vast
//		method = post
//		payload = createVastPayLoad
//		middleware = authz.Authenticate
//		resource = create_creative:self
//		200 = orm.CreativeSaveResult
//		400 = controller.ErrorResponseSimple
//		401 = controller.ErrorResponseSimple
//		403 = controller.ErrorResponseSimple
// }
func (c Controller) addVastCreativePost(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	pl := c.MustGetPayload(ctx).(*createVastPayLoad)
	res, err := c.addVastCreative(ctx, r, pl)
	if err != nil {
		framework.Write(w, err, http.StatusBadRequest)
		return
	}
	framework.Write(w, res, http.StatusOK)
}

// archiveCreative archive creative with creative id
// @Route {
// 		url = /archive/:id
//		method = patch
//		middleware = authz.Authenticate
//		resource = archive_creative:self
//		200 = adOrm.Creative
//		400 = controller.ErrorResponseSimple
//		401 = controller.ErrorResponseSimple
//		403 = controller.ErrorResponseSimple
// }
func (c *Controller) archiveCreativePatch(ctx context.Context, w http.ResponseWriter, r *http.Request) {

	res, err := c.archiveCreative(ctx, r)
	if err != nil {
		framework.Write(w, err, http.StatusBadRequest)
		return
	}
	framework.Write(w, res, http.StatusOK)
}

// changeCreativesStatus bulk approve reject creatives status of a campaign
// @Route {
// 		url = /change-creatives-status/:id
//		method = put
//		payload = creativesStatusPayload
//		middleware = authz.Authenticate
//		resource = change_creatives_status:superGlobal
//		200 = ChangeStatusResult
//		400 = controller.ErrorResponseSimple
//		401 = controller.ErrorResponseSimple
//		403 = controller.ErrorResponseSimple
// }
func (c *Controller) changeCreativesStatusPut(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	pl := c.MustGetPayload(ctx).(*creativesStatusPayload)
	res, err := c.changeCreativesStatus(ctx, r, pl)
	if err != nil {
		framework.Write(w, err, http.StatusBadRequest)
		return
	}
	framework.Write(w, res, http.StatusOK)
}

// addNativeCreative to campaign
// @Route {
// 		url = /native/:id
//		method = put
//		payload = editNativePayload
//		middleware = authz.Authenticate
//		resource = edit_creative:self
//		200 = orm.CreativeSaveResult
//		400 = controller.ErrorResponseSimple
//		401 = controller.ErrorResponseSimple
//		403 = controller.ErrorResponseSimple
// }
func (c Controller) editNativeCreativePut(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	pl := c.MustGetPayload(ctx).(*editNativePayload)
	res, err := c.editNativeCreative(ctx, r, pl)
	if err != nil {
		framework.Write(w, err, http.StatusBadRequest)
		return
	}
	framework.Write(w, res, http.StatusOK)
}

// getCreative to get creative by id
// @Route {
// 		url = /creative/:id
//		method = get
//		middleware = authz.Authenticate
//		resource = get_creative:self
//		200 = orm.CreativeSaveResult
//		400 = controller.ErrorResponseSimple
//		401 = controller.ErrorResponseSimple
//		403 = controller.ErrorResponseSimple
// }
func (c Controller) getCreativeGet(ctx context.Context, w http.ResponseWriter, r *http.Request) {

	res, err := c.getCreative(ctx, r)
	if err != nil {
		framework.Write(w, err, http.StatusBadRequest)
		return
	}
	framework.Write(w, res, http.StatusOK)
}

// getCreativeRejectReasons to get list of creative reject reasons
// @Route {
// 		url = /creative-reject-reasons
//		method = get
//		middleware = authz.Authenticate
//		resource = get_creative_reject_reason:self
//		200 = rejectReasons
//		400 = controller.ErrorResponseSimple
//		401 = controller.ErrorResponseSimple
//		403 = controller.ErrorResponseSimple
// }
func (c Controller) getCreativeRejectReasonsGet(ctx context.Context, w http.ResponseWriter, r *http.Request) {

	res, err := c.getCreativeRejectReasons(ctx, r)
	if err != nil {
		framework.Write(w, err, http.StatusBadRequest)
		return
	}
	framework.Write(w, res, http.StatusOK)
}
