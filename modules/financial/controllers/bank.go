package controllers

import (
	"context"
	"net/http"

	"clickyab.com/crab/modules/financial/errors"

	"clickyab.com/crab/modules/domain/middleware/domain"
	"clickyab.com/crab/modules/financial/orm"
	"clickyab.com/crab/modules/financial/payment/saman"
	"clickyab.com/crab/modules/user/middleware/authz"
	"github.com/clickyab/services/config"
	"github.com/clickyab/services/random"
)

var vatPercent = config.RegisterInt("crab.modules.financial.vat", 9, "vat percent")
var minChargeAmount = config.RegisterInt64("crab.modules.financial.min.real", 500000, "min charge amount")

// @Validate{
//}
type initPaymentPayload struct {
	ChargeAmount int64 `json:"charge_amount" validate:"required"`
	GateWay      int64 `json:"gate_way" validate:"required"`

	payAmount    int64 `json:"-"`
	chargeAmount int64 `json:"-"`
	vatAmount    int64 `json:"-"`
}

func (pl *initPaymentPayload) ValidateExtra(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	if pl.ChargeAmount < minChargeAmount.Int64() {
		return errors.MinChargeError(minChargeAmount.Int64())
	}
	validChargeAmount := (pl.ChargeAmount / 100) * 100
	vatAmount := (validChargeAmount * vatPercent.Int64()) / 100
	payAmount := validChargeAmount + vatAmount

	pl.payAmount = payAmount
	pl.vatAmount = vatAmount
	pl.chargeAmount = validChargeAmount

	return nil
}

// getPaymentData get payment data
// @Rest {
// 		url = /payment/init
//		protected = true
// 		method = post
//		resource = make_payment:self
// }
func (c *Controller) getPaymentData(ctx context.Context, r *http.Request, p *initPaymentPayload) (*saman.InitPaymentResp, error) {
	currentUser := authz.MustGetUser(ctx)
	dm := domain.MustGetDomain(ctx)
	// find gateway
	gate, err := orm.NewOrmManager().FindGatewayByID(p.GateWay)
	if err != nil {
		return nil, errors.NotFoundGateway
	}
	if gate.Status == "disable" {
		return nil, errors.DisableGateWayErr
	}
	// TODO : implement other gateways later
	if gate.Name != "saman" {
		return nil, errors.GateWayNotSupportedErr
	}
	var pay = saman.Saman{
		FAmount: p.payAmount,
		FResNum: <-random.ID,
	}

	// create online payment
	onlinePay := &orm.OnlinePayment{
		Status:    orm.Init,
		Amount:    pay.Amount(),
		ResNum:    pay.ResNum(),
		UserID:    currentUser.ID,
		DomainID:  dm.ID,
		GatewayID: gate.ID,
	}

	err = orm.NewOrmManager().CreateOnlinePayment(onlinePay)

	if err != nil {
		return nil, errors.MakeOnlinePaymentErr
	}
	return pay.InitPayment(), nil

}
