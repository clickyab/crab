// Code generated build with restful DO NOT EDIT.

package controllers

import (
	"context"
	"net/http"

	"github.com/clickyab/services/framework"
)

// registerSnap register new bank snap by advertiser
// @Route {
// 		url = /add
//		method = post
//		payload = registerBankSnapPayload
//		middleware = authz.Authenticate
//		resource = create_bank_snap:self
//		200 = orm.BankSnap
//		400 = controller.ErrorResponseSimple
//		401 = controller.ErrorResponseSimple
//		403 = controller.ErrorResponseSimple
// }
func (c *Controller) registerSnapPost(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	pl := c.MustGetPayload(ctx).(*registerBankSnapPayload)
	res, err := c.registerSnap(ctx, r, pl)
	if err != nil {
		framework.Write(w, err, http.StatusBadRequest)
		return
	}
	framework.Write(w, res, http.StatusOK)
}

// billingList get list of all user billings
// @Route {
// 		url = /
//		method = get
//		middleware = authz.Authenticate
//		resource = list_billing:self
//		200 = orm.Billing
//		400 = controller.ErrorResponseSimple
//		401 = controller.ErrorResponseSimple
//		403 = controller.ErrorResponseSimple
// }
func (c *Controller) billingListGet(ctx context.Context, w http.ResponseWriter, r *http.Request) {

	res, err := c.billingList(ctx, r)
	if err != nil {
		framework.Write(w, err, http.StatusBadRequest)
		return
	}
	framework.Write(w, res, http.StatusOK)
}
