// Code generated build with restful DO NOT EDIT.

package controllers

import (
	"context"
	"net/http"

	"github.com/clickyab/services/framework"
)

// addPreset edit inventory
// @Route {
// 		url = /addpub/:id
//		method = patch
//		payload = addInventoryPayload
//		middleware = authz.Authenticate
//		resource = edit_inventory:self
//		200 = orm.Inventory
//		400 = controller.ErrorResponseSimple
//		401 = controller.ErrorResponseSimple
//		403 = controller.ErrorResponseSimple
// }
func (ctrl *Controller) addPresetPatch(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	pl := ctrl.MustGetPayload(ctx).(*addInventoryPayload)
	res, err := ctrl.addPreset(ctx, r, pl)
	if err != nil {
		framework.Write(w, err, http.StatusBadRequest)
		return
	}
	framework.Write(w, res, http.StatusOK)
}

// removePreset edit inventory
// @Route {
// 		url = /:id
//		method = put
//		payload = changeLabelPayload
//		middleware = authz.Authenticate
//		resource = edit_inventory:self
//		200 = orm.Inventory
//		400 = controller.ErrorResponseSimple
//		401 = controller.ErrorResponseSimple
//		403 = controller.ErrorResponseSimple
// }
func (ctrl *Controller) changeLabelPut(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	pl := ctrl.MustGetPayload(ctx).(*changeLabelPayload)
	res, err := ctrl.changeLabel(ctx, r, pl)
	if err != nil {
		framework.Write(w, err, http.StatusBadRequest)
		return
	}
	framework.Write(w, res, http.StatusOK)
}

// changeStatus change inventory status
// @Route {
// 		url = /inventory/:id
//		method = patch
//		payload = changeStatusPayload
//		middleware = authz.Authenticate
//		resource = edit_inventory:self
//		200 = orm.Inventory
//		400 = controller.ErrorResponseSimple
//		401 = controller.ErrorResponseSimple
//		403 = controller.ErrorResponseSimple
// }
func (ctrl *Controller) changeStatusPatch(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	pl := ctrl.MustGetPayload(ctx).(*changeStatusPayload)
	res, err := ctrl.changeStatus(ctx, r, pl)
	if err != nil {
		framework.Write(w, err, http.StatusBadRequest)
		return
	}
	framework.Write(w, res, http.StatusOK)
}

// createPreset get a new whitelist blacklist for user
// @Route {
// 		url = /create
//		method = post
//		payload = createInventoryPayload
//		middleware = authz.Authenticate
//		resource = add_inventory:self
//		200 = orm.Inventory
//		400 = controller.ErrorResponseSimple
//		401 = controller.ErrorResponseSimple
//		403 = controller.ErrorResponseSimple
// }
func (ctrl *Controller) createPresetPost(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	pl := ctrl.MustGetPayload(ctx).(*createInventoryPayload)
	res, err := ctrl.createPreset(ctx, r, pl)
	if err != nil {
		framework.Write(w, err, http.StatusBadRequest)
		return
	}
	framework.Write(w, res, http.StatusOK)
}

// removePreset edit inventory
// @Route {
// 		url = /removepub/:id
//		method = patch
//		payload = removeInventoryPayload
//		middleware = authz.Authenticate
//		resource = edit_inventory:self
//		200 = orm.Inventory
//		400 = controller.ErrorResponseSimple
//		401 = controller.ErrorResponseSimple
//		403 = controller.ErrorResponseSimple
// }
func (ctrl *Controller) removePresetPatch(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	pl := ctrl.MustGetPayload(ctx).(*removeInventoryPayload)
	res, err := ctrl.removePreset(ctx, r, pl)
	if err != nil {
		framework.Write(w, err, http.StatusBadRequest)
		return
	}
	framework.Write(w, res, http.StatusOK)
}
