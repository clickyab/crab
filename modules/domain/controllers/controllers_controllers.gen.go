// Code generated build with restful DO NOT EDIT.

package controllers

import (
	"context"
	"net/http"

	"github.com/clickyab/services/framework"
)

// changeDomainStatus change domain status by id, status can be enable or disable
// @Route {
// 		url = /change-domain-status/:id
//		method = put
//		payload = changeDomainStatusPayload
//		middleware = authz.Authenticate
//		resource = change_domain_status:superGlobal
//		200 = orm.Domain
//		400 = controller.ErrorResponseSimple
//		401 = controller.ErrorResponseSimple
//		403 = controller.ErrorResponseSimple
// }
func (c *Controller) changeDomainStatusPut(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	pl := c.MustGetPayload(ctx).(*changeDomainStatusPayload)
	res, err := c.changeDomainStatus(ctx, r, pl)
	if err != nil {
		framework.Write(w, err, http.StatusBadRequest)
		return
	}
	framework.Write(w, res, http.StatusOK)
}

// createDomain to domain
// @Route {
// 		url = /create
//		method = post
//		payload = createDomainPayload
//		middleware = authz.Authenticate
//		resource = create_domain:superGlobal
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
//		resource = edit_domain:superGlobal
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
// 		url = /get/:id
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

// getDomainConfig get domain config by domain name
// @Route {
// 		url = /config/:name
//		method = get
//		200 = domainConfig
//		400 = controller.ErrorResponseSimple
// }
func (c *Controller) getDomainConfigGet(ctx context.Context, w http.ResponseWriter, r *http.Request) {

	res, err := c.getDomainConfig(ctx, r)
	if err != nil {
		framework.Write(w, err, http.StatusBadRequest)
		return
	}
	framework.Write(w, res, http.StatusOK)
}

// getGlobalConfig get global config
// @Route {
// 		url = /super-global-config
//		method = get
//		middleware = authz.Authenticate
//		resource = get_global_config:superGlobal
//		200 = orm.UserConfig
//		400 = controller.ErrorResponseSimple
//		401 = controller.ErrorResponseSimple
//		403 = controller.ErrorResponseSimple
// }
func (c *Controller) getGlobalConfigGet(ctx context.Context, w http.ResponseWriter, r *http.Request) {

	res, err := c.getGlobalConfig(ctx, r)
	if err != nil {
		framework.Write(w, err, http.StatusBadRequest)
		return
	}
	framework.Write(w, res, http.StatusOK)
}
