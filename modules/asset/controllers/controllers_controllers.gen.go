// Code generated build with restful DO NOT EDIT.

package controllers

import (
	"context"
	"net/http"

	"github.com/clickyab/services/framework"
)

// browser return list all browsers
// @Route {
// 		url = /browser
//		method = get
//		middleware = authz.Authenticate
//		200 = browserResponse
//		400 = controller.ErrorResponseSimple
//		401 = controller.ErrorResponseSimple
// }
func (c *Controller) browserGet(ctx context.Context, w http.ResponseWriter, r *http.Request) {

	res, err := c.browser(ctx, r)
	if err != nil {
		framework.Write(w, err, http.StatusBadRequest)
		return
	}
	framework.Write(w, res, http.StatusOK)
}

// category return list iab categories
// @Route {
// 		url = /category
//		method = get
//		middleware = authz.Authenticate
//		200 = catResponse
//		400 = controller.ErrorResponseSimple
//		401 = controller.ErrorResponseSimple
// }
func (c *Controller) categoryGet(ctx context.Context, w http.ResponseWriter, r *http.Request) {

	res, err := c.category(ctx, r)
	if err != nil {
		framework.Write(w, err, http.StatusBadRequest)
		return
	}
	framework.Write(w, res, http.StatusOK)
}

// isp return list all is (e.g. irancell, ...)
// @Route {
// 		url = /isp/:kind
//		method = get
//		middleware = authz.Authenticate
//		200 = ispResponse
//		400 = controller.ErrorResponseSimple
//		401 = controller.ErrorResponseSimple
// }
func (c *Controller) ispGet(ctx context.Context, w http.ResponseWriter, r *http.Request) {

	res, err := c.isp(ctx, r)
	if err != nil {
		framework.Write(w, err, http.StatusBadRequest)
		return
	}
	framework.Write(w, res, http.StatusOK)
}

// Manufacturers return list all mobile manufacturers (e.g. Apple, Samsung)
// @Route {
// 		url = /manufacturers
//		method = get
//		middleware = authz.Authenticate
//		200 = manufacturers
//		400 = controller.ErrorResponseSimple
//		401 = controller.ErrorResponseSimple
// }
func (c *Controller) manufacturerGet(ctx context.Context, w http.ResponseWriter, r *http.Request) {

	res, err := c.manufacturer(ctx, r)
	if err != nil {
		framework.Write(w, err, http.StatusBadRequest)
		return
	}
	framework.Write(w, res, http.StatusOK)
}

// os return list all is (e.g. linux, ...)
// @Route {
// 		url = /os/:kind
//		method = get
//		middleware = authz.Authenticate
//		200 = osResponse
//		400 = controller.ErrorResponseSimple
//		401 = controller.ErrorResponseSimple
// }
func (c *Controller) osGet(ctx context.Context, w http.ResponseWriter, r *http.Request) {

	res, err := c.os(ctx, r)
	if err != nil {
		framework.Write(w, err, http.StatusBadRequest)
		return
	}
	framework.Write(w, res, http.StatusOK)
}

// os return list all is (e.g. linux, ...)
// @Route {
// 		url = /os
//		method = get
//		middleware = authz.Authenticate
//		200 = osResponse
//		400 = controller.ErrorResponseSimple
//		401 = controller.ErrorResponseSimple
// }
func (c *Controller) allOsGet(ctx context.Context, w http.ResponseWriter, r *http.Request) {

	res, err := c.allOs(ctx, r)
	if err != nil {
		framework.Write(w, err, http.StatusBadRequest)
		return
	}
	framework.Write(w, res, http.StatusOK)
}

// platform return list all is (e.g. desktop,mobile, ...)
// @Route {
// 		url = /platform/:kind
//		method = get
//		middleware = authz.Authenticate
//		200 = platformResponse
//		400 = controller.ErrorResponseSimple
//		401 = controller.ErrorResponseSimple
// }
func (c *Controller) platformGet(ctx context.Context, w http.ResponseWriter, r *http.Request) {

	res, err := c.platform(ctx, r)
	if err != nil {
		framework.Write(w, err, http.StatusBadRequest)
		return
	}
	framework.Write(w, res, http.StatusOK)
}

// platform return list all is (e.g. desktop,mobile, ...)
// @Route {
// 		url = /platform
//		method = get
//		middleware = authz.Authenticate
//		200 = platformResponse
//		400 = controller.ErrorResponseSimple
//		401 = controller.ErrorResponseSimple
// }
func (c *Controller) allPlatformGet(ctx context.Context, w http.ResponseWriter, r *http.Request) {

	res, err := c.allPlatform(ctx, r)
	if err != nil {
		framework.Write(w, err, http.StatusBadRequest)
		return
	}
	framework.Write(w, res, http.StatusOK)
}
