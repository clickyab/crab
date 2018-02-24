// Code generated build with restful DO NOT EDIT.

package controllers

import (
	"context"
	"net/http"

	"github.com/clickyab/services/framework"
)

// whiteBlackLists return all user inventories
// @Route {
// 		url = /presets
//		method = get
//		middleware = authz.Authenticate
//		200 = whiteBlackLists
//		400 = controller.ErrorResponseSimple
//		401 = controller.ErrorResponseSimple
// }
func (ctrl *Controller) whiteBlackListsGet(ctx context.Context, w http.ResponseWriter, r *http.Request) {

	res, err := ctrl.whiteBlackLists(ctx, r)
	if err != nil {
		framework.Write(w, err, http.StatusBadRequest)
		return
	}
	framework.Write(w, res, http.StatusOK)
}

// whiteBlackList return a user inventory
// @Route {
// 		url = /preset/:id
//		method = get
//		middleware = authz.Authenticate
//		200 = orm.WhiteBlackList
//		400 = controller.ErrorResponseSimple
//		401 = controller.ErrorResponseSimple
// }
func (ctrl *Controller) whiteBlackListGet(ctx context.Context, w http.ResponseWriter, r *http.Request) {

	res, err := ctrl.whiteBlackList(ctx, r)
	if err != nil {
		framework.Write(w, err, http.StatusBadRequest)
		return
	}
	framework.Write(w, res, http.StatusOK)
}

// addPreset get a new whitelist blacklist for user
// @Route {
// 		url = /preset
//		method = post
//		payload = whiteBlackList
//		middleware = authz.Authenticate
//		200 = orm.WhiteBlackList
//		400 = controller.ErrorResponseSimple
//		401 = controller.ErrorResponseSimple
// }
func (ctrl *Controller) addPresetPost(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	pl := ctrl.MustGetPayload(ctx).(*whiteBlackList)
	res, err := ctrl.addPreset(ctx, r, pl)
	if err != nil {
		framework.Write(w, err, http.StatusBadRequest)
		return
	}
	framework.Write(w, res, http.StatusOK)
}
