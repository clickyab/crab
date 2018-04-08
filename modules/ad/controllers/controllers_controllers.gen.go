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
//		payload = nativeCreativePayload
//		middleware = authz.Authenticate
//		200 = orm.CreativeSaveResult
//		400 = controller.ErrorResponseSimple
//		401 = controller.ErrorResponseSimple
// }
func (c Controller) addNativeCreativePost(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	pl := c.MustGetPayload(ctx).(*nativeCreativePayload)
	res, err := c.addNativeCreative(ctx, r, pl)
	if err != nil {
		framework.Write(w, err, http.StatusBadRequest)
		return
	}
	framework.Write(w, res, http.StatusOK)
}

// editNativeCreative to campaign
// @Route {
// 		url = /native/:creative_id
//		method = put
//		payload = nativeCreativePayload
//		middleware = authz.Authenticate
//		200 = orm.CreativeSaveResult
//		400 = controller.ErrorResponseSimple
//		401 = controller.ErrorResponseSimple
// }
func (c Controller) editNativeCreativePut(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	pl := c.MustGetPayload(ctx).(*nativeCreativePayload)
	res, err := c.editNativeCreative(ctx, r, pl)
	if err != nil {
		framework.Write(w, err, http.StatusBadRequest)
		return
	}
	framework.Write(w, res, http.StatusOK)
}
