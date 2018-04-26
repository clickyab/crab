package controllers

import (
	"context"
	"net/http"
	"strconv"

	"clickyab.com/crab/modules/domain/middleware/domain"
	"clickyab.com/crab/modules/financial/errors"
	"clickyab.com/crab/modules/financial/orm"
	"clickyab.com/crab/modules/user/aaa"
	"clickyab.com/crab/modules/user/middleware/authz"
	"github.com/rs/xmux"
)

// getPaymentTransaction get single payment transaction
// @Rest {
// 		url = /payment/:id
//		protected = true
// 		method = get
//		resource = make_payment:self
// }
func (c *Controller) getPaymentTransaction(ctx context.Context, r *http.Request) (*orm.OnlinePayment, error) {
	tID, err := strconv.ParseInt(xmux.Param(ctx, "id"), 10, 64)
	if err != nil {
		return nil, errors.InvalidIDErr
	}
	currentUser := authz.MustGetUser(ctx)
	dm := domain.MustGetDomain(ctx)

	transaction, err := orm.NewOrmManager().FindOnlinePaymentByID(tID)
	if err != nil {
		return nil, errors.NotFoundError(tID)
	}

	owner, err := aaa.NewAaaManager().FindUserWithParentsByID(transaction.UserID, dm.ID)
	if err != nil {
		return nil, errors.AccessDenied
	}

	_, ok := aaa.CheckPermOn(owner, currentUser, "make_payment", dm.ID)
	if !ok {
		return nil, errors.AccessDenied
	}

	return transaction, nil
}
