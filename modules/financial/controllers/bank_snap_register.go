package controllers

import (
	"context"
	"net/http"

	"clickyab.com/crab/modules/domain/middleware/domain"
	"clickyab.com/crab/modules/financial/errors"
	"clickyab.com/crab/modules/financial/orm"
	"clickyab.com/crab/modules/user/middleware/authz"
	"github.com/clickyab/services/config"
)

var (
	minRegisterBankSnap = config.RegisterInt64("crab.modules.financial.min.snap", 500000, "min snap bank money value")
	vatPercentage       = config.RegisterInt("crab.modules.financial.vat", 9, "vat percentage")
)

// @Validate{
//}
type registerBankSnapPayload struct {
	Amount      int64 `json:"amount" validate:"required"`
	TraceNumber int64 `json:"trace_number" validate:"required"`
}

// ValidateExtra validate registerBankSnapPayload struct
func (pl *registerBankSnapPayload) ValidateExtra(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	if pl.Amount < minRegisterBankSnap.Int64() {
		return errors.MinBankSnapErr
	}
	return nil
}

// registerSnap register new bank snap by advertiser
// @Rest {
// 		url = /add
//		protected = true
// 		method = post
//		resource = create_bank_snap:self
// }
func (c *Controller) registerSnap(ctx context.Context, r *http.Request, p *registerBankSnapPayload) (*orm.BankSnap, error) {
	currentUser := authz.MustGetUser(ctx)
	dm := domain.MustGetDomain(ctx)
	vatMoney := vatPercentage.Int64() * p.Amount / 100
	bank := &orm.BankSnap{
		PayAmount:   p.Amount,
		VAT:         vatMoney,
		Amount:      p.Amount - vatMoney,
		UserID:      currentUser.ID,
		DomainID:    dm.ID,
		Status:      orm.PendingStatus,
		TraceNumber: p.TraceNumber,
	}
	err := orm.NewOrmManager().CreateBankSnap(bank)
	if err != nil {
		return nil, errors.CreateError
	}
	return bank, nil
}
