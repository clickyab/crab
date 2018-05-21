// Code generated build with restful DO NOT EDIT.

package controllers

import (
	"context"
	"net/http"

	"github.com/clickyab/services/framework"
)

// createDomain to domain
// @Route {
// 		url = /create
//		method = post
//		payload = createDomainPayload
//		middleware = authz.Authenticate
//		resource = god:global
//		200 = orm.Domain
//		400 = controller.ErrorResponseSimple
//		401 = controller.ErrorResponseSimple
//		403 = controller.ErrorResponseSimple
// }
func (c *Controller) createDomainPost(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	pl := c.MustGetPayload(ctx).(*createDomainPayload)
	res, err := c.createDomain(ctx, r, pl)
	if err != nil {
		framework.Write(w, err, http.StatusBadRequest)
		return
	}
	framework.Write(w, res, http.StatusOK)
}

// editDomain to domain
// @Route {
// 		url = /edit/:id
//		method = put
//		payload = editDomainPayload
//		middleware = authz.Authenticate
//		resource = god:global
//		200 = orm.Domain
//		400 = controller.ErrorResponseSimple
//		401 = controller.ErrorResponseSimple
//		403 = controller.ErrorResponseSimple
// }
func (c *Controller) editDomainPut(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	pl := c.MustGetPayload(ctx).(*editDomainPayload)
	res, err := c.editDomain(ctx, r, pl)
	if err != nil {
		framework.Write(w, err, http.StatusBadRequest)
		return
	}
	framework.Write(w, res, http.StatusOK)
}

// getDomainDetail get domain detail by domain id
// @Route {
// 		url = /:id
//		method = get
//		middleware = authz.Authenticate
//		resource = get_detail_domain:global
//		200 = orm.Domain
//		400 = controller.ErrorResponseSimple
//		401 = controller.ErrorResponseSimple
//		403 = controller.ErrorResponseSimple
// }
func (c *Controller) getDomainDetailGet(ctx context.Context, w http.ResponseWriter, r *http.Request) {

	res, err := c.getDomainDetail(ctx, r)
	if err != nil {
		framework.Write(w, err, http.StatusBadRequest)
		return
	}
	framework.Write(w, res, http.StatusOK)
}
