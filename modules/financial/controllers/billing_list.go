package controllers

import (
	"context"
	"net/http"

	"clickyab.com/crab/modules/financial/orm"
)

//TODO: implement payload struct

// billingList get list of all user billings
// @Rest {
// 		url = /
//		protected = true
// 		method = get
//		resource = list_billing:self
// }
func (c *Controller) billingList(ctx context.Context, r *http.Request) (*orm.Billing, error) {
	//TODO: implement route ....

	return nil, nil
}
