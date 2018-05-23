package controllers

import (
	"context"
	"net/http"

	"clickyab.com/crab/modules/financial/errors"

	"fmt"

	"net/url"

	"clickyab.com/crab/modules/domain/middleware/domain"
	"clickyab.com/crab/modules/financial/orm"
	"clickyab.com/crab/modules/financial/payment"
	"clickyab.com/crab/modules/financial/payment/saman"
	"clickyab.com/crab/modules/user/middleware/authz"
	"github.com/clickyab/services/assert"
	"github.com/clickyab/services/config"
	"github.com/clickyab/services/mysql"
	"github.com/clickyab/services/random"
	"github.com/clickyab/services/xlog"
	"github.com/rs/xmux"
	"github.com/sirupsen/logrus"
)

var minChargeAmount = config.RegisterInt64("crab.modules.financial.min.real", 100, "min charge amount")

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
	vatAmount := (validChargeAmount * orm.VatPercent.Int64()) / 100
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
func (c *Controller) getPaymentData(ctx context.Context, r *http.Request, p *initPaymentPayload) (*payment.InitPaymentResp, error) {
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

	var pay = payment.CommonPay{
		FAmount: p.payAmount,
		FResNum: <-random.ID,
		FUserID: currentUser.ID,
		BankObj: nil,
	}

	// TODO : implement other gateways later
	if gate.Name == "saman" {
		pay.BankObj = &saman.Saman{CommonPay: pay}
	} else {
		return nil, errors.GateWayNotSupportedErr
	}

	// create online payment
	onlinePay := &orm.OnlinePayment{
		Status:    orm.Init,
		Amount:    pay.BankObj.Amount(),
		ResNum:    pay.BankObj.ResNum(),
		UserID:    currentUser.ID,
		DomainID:  dm.ID,
		GatewayID: gate.ID,
	}

	err = orm.NewOrmManager().CreateOnlinePayment(onlinePay)

	if err != nil {
		return nil, errors.MakeOnlinePaymentErr
	}
	return pay.BankObj.InitPayment(r), nil

}

// backFromBank return from bank
// @Route {
// 		url = /payment/return/:bank/:hash
//		method = post
// }
func (c *Controller) backFromBank(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	currentDomain := domain.MustGetDomain(ctx)

	// get data from bank redirect request
	bankName := xmux.Param(ctx, "bank")
	payHash := xmux.Param(ctx, "hash")

	payOrm := orm.NewOrmManager()

	l := &payment.CommonPay{}

	gateway, err := payOrm.FindActiveGatewayByName(bankName)
	if err != nil {
		params := url.Values{}
		params.Set("success", "false")
		assert.Nil(l.FrontRedirect(w, r, http.StatusFound, params))
		return
	}

	if gateway.Name == "saman" {
		l.BankObj = &saman.Saman{
			CommonPay: payment.CommonPay{},
		}
	} else {
		// Failed payment  redirect to failed page
		params := url.Values{}
		params.Set("success", "false")
		assert.Nil(l.BankObj.FrontRedirect(w, r, http.StatusFound, params))
		return
	}
	redirectParams := l.BankObj.GetParams(r)

	transaction, err := payOrm.FindInitPaymentByResNum(redirectParams.ResNum)
	if err != nil {
		// Failed payment  redirect to failed page
		params := url.Values{}
		params.Set("success", "false")
		assert.Nil(l.FrontRedirect(w, r, http.StatusFound, params))
		return
	}

	l.BankObj.SetAmount(transaction.Amount)
	l.BankObj.SetUserID(transaction.UserID)
	l.BankObj.SetResNum(transaction.ResNum)

	err = l.BankObj.HashVerification(l.BankObj.PayAmount(), payHash, redirectParams.ResNum)
	if err != nil {
		// Failed payment update transaction and redirect to failed page
		transaction.Status = orm.BackToSite
		transaction.BankStatus = mysql.NullInt64{Valid: true, Int64: redirectParams.StatusCode}
		transaction.ErrorReason = orm.NullBankReason{Valid: true, BankReason: orm.BankReasonState(err.Error())}
		transaction.RefNum = mysql.NullString{Valid: redirectParams.RefNum != "", String: redirectParams.RefNum}
		assert.Nil(payOrm.UpdateOnlinePayment(transaction))
		params := url.Values{}
		params.Set("success", "false")
		params.Set("payment", fmt.Sprint(transaction.ID))
		assert.Nil(l.BankObj.FrontRedirect(w, r, http.StatusFound, params))
		return
	}

	err = l.BankObj.RedirectValidation(redirectParams)
	if err != nil {
		transaction.Status = orm.BackToSite
		transaction.BankStatus = mysql.NullInt64{Valid: true, Int64: redirectParams.StatusCode}
		transaction.ErrorReason = orm.NullBankReason{Valid: true, BankReason: orm.BankReasonState(err.Error())}
		transaction.RefNum = mysql.NullString{Valid: redirectParams.RefNum != "", String: redirectParams.RefNum}
		assert.Nil(payOrm.UpdateOnlinePayment(transaction))
		params := url.Values{}
		params.Set("success", "false")
		params.Set("payment", fmt.Sprint(transaction.ID))
		assert.Nil(l.BankObj.FrontRedirect(w, r, http.StatusFound, params))
		return
	}

	err = l.BankObj.VerifyTransaction(redirectParams.ResNum, redirectParams.RefNum)
	if err != nil {
		// Failed payment update transaction and redirect to failed page
		transaction.Status = orm.BackToSite
		transaction.BankStatus = mysql.NullInt64{Valid: true, Int64: redirectParams.StatusCode}
		transaction.ErrorReason = orm.NullBankReason{Valid: true, BankReason: orm.BankReasonState(err.Error())}
		transaction.RefNum = mysql.NullString{Valid: redirectParams.RefNum != "", String: redirectParams.RefNum}
		assert.Nil(payOrm.UpdateOnlinePayment(transaction))
		params := url.Values{}
		params.Set("success", "false")
		params.Set("payment", fmt.Sprint(transaction.ID))
		assert.Nil(l.BankObj.FrontRedirect(w, r, http.StatusFound, params))
		return
	}

	// Success payment occurred charge user and redirect to success page
	transaction.Status = orm.Finalized
	transaction.RefNum = mysql.NullString{Valid: true, String: redirectParams.RefNum}
	transaction.BankStatus = mysql.NullInt64{Valid: true, Int64: redirectParams.StatusCode}
	transaction.TraceNumber = mysql.NullString{Valid: true, String: redirectParams.Attr["traceNumber"]}
	transaction.CID = mysql.NullString{Valid: true, String: redirectParams.Attr["cID"]}
	transaction.Attr = map[string]interface{}{
		"SecurePan": redirectParams.Attr["securePan"],
	}

	// charge user
	chargeAmount := (transaction.Amount * 100) / (100 + orm.VatPercent.Int64())
	err = payOrm.ChargeUser(transaction, currentDomain.ID, chargeAmount)
	if err != nil {
		f := logrus.Fields{
			"transaction_id": transaction.ID,
			"domain_id":      currentDomain.ID,
		}
		xlog.GetWithFields(ctx, f).Error("charge user failed")
	}
	assert.Nil(err)

	params := url.Values{}
	params.Set("success", "true")
	params.Set("payment", fmt.Sprint(transaction.ID))
	assert.Nil(l.BankObj.FrontRedirect(w, r, http.StatusMovedPermanently, params))
}
