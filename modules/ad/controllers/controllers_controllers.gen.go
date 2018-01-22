// Code generated build with restful DO NOT EDIT.

package controllers

import (
	"context"
	"net/http"

	"github.com/clickyab/services/framework"
)

// assignNormalBanner assignNormalBanner module is banner type (banner/native)
// @Route {
// 		url = /:banner_type/:id
//		method = post
//		payload = assignBannerPayload
//		middleware = authz.Authenticate
//		resource = assign_banner:self
//		200 = adResponse
//		400 = controller.ErrorResponseSimple
//		401 = controller.ErrorResponseSimple
//		403 = controller.ErrorResponseSimple
// }
func (c *Controller) assignNormalBannerPost(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	pl := c.MustGetPayload(ctx).(*assignBannerPayload)
	res, err := c.assignNormalBanner(ctx, r, pl)
	if err != nil {
		framework.Write(w, err, http.StatusBadRequest)
		return
	}
	framework.Write(w, res, http.StatusOK)
}
