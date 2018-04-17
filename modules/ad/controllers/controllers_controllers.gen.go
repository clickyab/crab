// Code generated build with restful DO NOT EDIT.

package controllers

import (
	"context"
	"net/http"

	"github.com/clickyab/services/framework"
)

// addNativeCreative to campaign
// @Route {
// 		url = /native
//		method = post
//		payload = createNativePayload
//		middleware = authz.Authenticate
//		resource = add_creative:self
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
