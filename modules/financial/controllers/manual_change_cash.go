package controllers

import (
	"context"
	"math"
	"net/http"

	"clickyab.com/crab/modules/domain/middleware/domain"
	ormDomain "clickyab.com/crab/modules/domain/orm"
	"clickyab.com/crab/modules/financial/errors"
	"clickyab.com/crab/modules/financial/orm"
	"clickyab.com/crab/modules/user/aaa"
	"clickyab.com/crab/modules/user/middleware/authz"
	"github.com/clickyab/services/permission"
	"github.com/clickyab/services/xlog"
)

// @Validate{
//}
type changeCashStatus struct {
	UserID        int64                 `json:"user_id" validate:"required"`
	Reason        orm.ChangeCashReasons `json:"reason" validate:"required"`
	Amount        int64                 `json:"amount" validate:"required"`
	Description   string                `json:"description"`
	currentDomain *ormDomain.Domain     `json:"-"`
	operatorUser  *aaa.User             `json:"-"`
	targetUser    *aaa.User             `json:"-"`
}

// ChangeCashResult to return number of cash change
type ChangeCashResult struct {
	TargetUserID int64 `json:"target_user_id"`
	Amount       int64 `json:"amount"`
}

func (p *changeCashStatus) ValidateExtra(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	if !p.Reason.IsValid() {
		return errors.InvalidReasonErr
	}
	currentUser := authz.MustGetUser(ctx)
	p.operatorUser = currentUser

	dm := domain.MustGetDomain(ctx)
	// check if user id is valid
	userObject, err := aaa.NewAaaManager().FindUserWithParentsByID(p.UserID, dm.ID)
	if err != nil {
		return errors.InvalidIDErr
	}
	p.targetUser = userObject
	// check if should inc the balance
	if p.Amount > 0 {
		// check if creator has enough money in her/his
		if p.Amount > p.operatorUser.Balance {
			return errors.NotEnoughBalance
		}
	} else {
		n := int64(math.Abs(float64(p.Amount)))
		if p.targetUser.Balance < n {
			return errors.UserBalanceLowErr
		}
	}
	p.currentDomain = dm
	return nil
}

// manualChangeCash to financial
// @Rest {
// 		url = /manual-change-cash
//		protected = true
// 		method = put
// 		resource = manual_change_cash:global
// }
func (c *Controller) manualChangeCash(ctx context.Context, r *http.Request, p *changeCashStatus) (*ChangeCashResult, error) {
	//check permission
	_, ok := p.operatorUser.HasOn("manual_change_cash", p.targetUser.ID, p.currentDomain.ID, false, false, permission.ScopeGlobal)
	if !ok {
		return nil, errors.AccessDenied
	}

	manualCash := &orm.ManualCashChange{
		DomainID:    p.currentDomain.ID,
		UserID:      p.UserID,
		OperatorID:  p.operatorUser.ID,
		Reason:      p.Reason,
		Amount:      p.Amount,
		Description: p.Description,
	}

	m := orm.NewOrmManager()
	err := m.ApplyManualCash(p.operatorUser, p.targetUser, manualCash)

	if err != nil {
		xlog.GetWithError(ctx, err).Debug("database error when applying manual cash:", err)
		return nil, errors.ApplyManualCashDbErr
	}

	result := &ChangeCashResult{
		TargetUserID: p.UserID,
		Amount:       p.Amount,
	}
	// return result to user
	return result, nil
}
