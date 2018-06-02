package controllers

import (
	"context"
	"net/http"
	"strconv"

	domainErrors "clickyab.com/crab/modules/domain/errors"
	"clickyab.com/crab/modules/domain/orm"
	"clickyab.com/crab/modules/financial/errors"
	payOrm "clickyab.com/crab/modules/financial/orm"
	"clickyab.com/crab/modules/user/aaa"
	"clickyab.com/crab/modules/user/middleware/authz"
	"github.com/rs/xmux"
)

type chargeOwner struct {
	Amount int64 `json:"amount"`
}

// chargeOwner charge owner by domain less user
// @Rest {
// 		url = /charge/whitelabel/:id
//		protected = true
// 		method = post
//		resource = charge_owner:superGlobal
// }
func (c *Controller) chargeOwner(ctx context.Context, r *http.Request, p *chargeOwner) (*payOrm.ManualCashChange, error) {
	currentUser := authz.MustGetUser(ctx)
	idInt, err := strconv.ParseInt(xmux.Param(ctx, "id"), 10, 0)
	if err != nil {
		return nil, errors.InvalidIDErr
	}
	// find whitelabel
	targetDomain, err := orm.NewOrmManager().FindDomainByID(idInt)
	if err != nil {
		return nil, domainErrors.DomainNotFoundError(idInt)
	}
	//find domain owner
	owner, err := aaa.NewAaaManager().FindDomainOwner(targetDomain.ID)
	if err != nil {
		return nil, errors.OwnerNotFoundErr
	}
	ch := &payOrm.ManualCashChange{
		Amount:      p.Amount,
		DomainID:    targetDomain.ID,
		Description: "",
		Reason:      payOrm.ManualPay,
		OperatorID:  currentUser.ID,
		UserID:      owner.ID,
	}
	// charge user
	err = payOrm.NewOrmManager().ApplyOwnerManualCash(currentUser, owner, ch)
	if err != nil {
		return nil, errors.ApplyManualCashDbErr
	}
	if err != nil {
		return nil, err
	}
	return ch, nil

}
