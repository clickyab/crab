// Code generated build with restful DO NOT EDIT.

package controllers

import (
	"context"
	"net/http"

	"github.com/clickyab/services/framework"
)

// upload route for upload banner,native,avatar,...
// @Route {
// 		url = /module/:module
//		method = post
//		middleware = authz.Authenticate
//		200 = uploadResponse
//		400 = controller.ErrorResponseSimple
//		401 = controller.ErrorResponseSimple
// }
func (c *Controller) uploadPost(ctx context.Context, w http.ResponseWriter, r *http.Request) {

	res, err := c.upload(ctx, r)
	if err != nil {
		framework.Write(w, err, http.StatusBadRequest)
		return
	}
	framework.Write(w, res, http.StatusOK)
}

// getVideoReady find video into the system
// @Route {
// 		url = /video/:id
//		method = get
//		middleware = authz.Authenticate
//		200 = getVideoResponse
//		400 = controller.ErrorResponseSimple
//		401 = controller.ErrorResponseSimple
// }
func (c *Controller) getVideoReadyGet(ctx context.Context, w http.ResponseWriter, r *http.Request) {

	res, err := c.getVideoReady(ctx, r)
	if err != nil {
		framework.Write(w, err, http.StatusBadRequest)
		return
	}
	framework.Write(w, res, http.StatusOK)
}
