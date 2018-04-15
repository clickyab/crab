package controllers

import (
	"context"
	"net/http"

	"clickyab.com/crab/modules/financial/orm"
)

//TODO: implement payload struct

// registerSnap register new bank snap by advertiser
// @Rest {
// 		url = /add
//		protected = true
// 		method = post
//		resource = creat_bank_snap:self
// }
func (c *Controller) registerSnap(ctx context.Context, r *http.Request) (*orm.BankSnap, error) {
	//TODO: implement route ....

	return nil, nil
}
