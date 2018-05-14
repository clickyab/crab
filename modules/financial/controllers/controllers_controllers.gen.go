// Code generated build with restful DO NOT EDIT.

package controllers

import (
	"context"
	"net/http"

	"github.com/clickyab/services/framework"
)

// getPaymentData get payment data
// @Route {
// 		url = /payment/init
//		method = post
//		payload = initPaymentPayload
//		middleware = authz.Authenticate
//		resource = make_payment:self
//		200 = payment.InitPaymentResp
//		400 = controller.ErrorResponseSimple
//		401 = controller.ErrorResponseSimple
//		403 = controller.ErrorResponseSimple
// }
func (c *Controller) getPaymentDataPost(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	pl := c.MustGetPayload(ctx).(*initPaymentPayload)
	res, err := c.getPaymentData(ctx, r, pl)
	if err != nil {
		framework.Write(w, err, http.StatusBadRequest)
		return
	}
	framework.Write(w, res, http.StatusOK)
}

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

// getGateways get payment data
// @Route {
// 		url = /gateways
//		method = get
//		middleware = authz.Authenticate
//		200 = getGateResp
//		400 = controller.ErrorResponseSimple
//		401 = controller.ErrorResponseSimple
// }
func (c *Controller) getGatewaysGet(ctx context.Context, w http.ResponseWriter, r *http.Request) {

	res, err := c.getGateways(ctx, r)
	if err != nil {
		framework.Write(w, err, http.StatusBadRequest)
		return
	}
	framework.Write(w, res, http.StatusOK)
}

// getPaymentTransaction get single payment transaction
// @Route {
// 		url = /payment/:id
//		method = get
//		middleware = authz.Authenticate
//		resource = make_payment:self
//		200 = orm.OnlinePayment
//		400 = controller.ErrorResponseSimple
//		401 = controller.ErrorResponseSimple
//		403 = controller.ErrorResponseSimple
// }
func (c *Controller) getPaymentTransactionGet(ctx context.Context, w http.ResponseWriter, r *http.Request) {

	res, err := c.getPaymentTransaction(ctx, r)
	if err != nil {
		framework.Write(w, err, http.StatusBadRequest)
		return
	}
	framework.Write(w, res, http.StatusOK)
}

// manualChangeCash to financial
// @Route {
// 		url = /manual-change-cash
//		method = put
//		payload = changeCashStatus
//		middleware = authz.Authenticate
//		resource = manual_change_cash:global
//		200 = ChangeCashResult
//		400 = controller.ErrorResponseSimple
//		401 = controller.ErrorResponseSimple
//		403 = controller.ErrorResponseSimple
// }
func (c *Controller) manualChangeCashPut(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	pl := c.MustGetPayload(ctx).(*changeCashStatus)
	res, err := c.manualChangeCash(ctx, r, pl)
	if err != nil {
		framework.Write(w, err, http.StatusBadRequest)
		return
	}
	framework.Write(w, res, http.StatusOK)
}
