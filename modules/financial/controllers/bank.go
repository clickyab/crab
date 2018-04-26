package controllers

import (
	"context"
	"net/http"

	"clickyab.com/crab/modules/financial/errors"

	"strconv"

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

var minChargeAmount = config.RegisterInt64("crab.modules.financial.min.real", 1000, "min charge amount")

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
	// TODO : implement other gateways later
	if gate.Name != "saman" {
		return nil, errors.GateWayNotSupportedErr
	}
	var pay = saman.Saman{
		FAmount: p.payAmount,
		FResNum: <-random.ID,
		FUserID: currentUser.ID,
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
	return pay.InitPayment(r), nil

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
	resNum := r.PostFormValue("ResNum")
	refNum := r.PostFormValue("RefNum")
	state := r.PostFormValue("State")
	traceNumber := r.PostFormValue("TRACENO")
	cID := r.PostFormValue("CID")
	securePan := r.PostFormValue("SecurePan")
	mID := r.PostFormValue("MID")
	stateCode := r.PostFormValue("StateCode")
	stateCodeInt, err := strconv.ParseInt(stateCode, 10, 0)
	assert.Nil(err)
	payHash := xmux.Param(ctx, "hash")

	bankObj := &saman.Saman{
		FResNum: resNum,
	}

	payOrm := orm.NewOrmManager()

	transaction, err := payOrm.FindInitPaymentByResNum(resNum)
	if err != nil {
		// Failed payment  redirect to failed page
		params := url.Values{}
		params.Set("success", "no")
		assert.Nil(bankObj.FrontRedirect(w, r, http.StatusFound, params))
		return
	}

	bankObj.FUserID = transaction.UserID
	bankObj.FAmount = transaction.Amount

	if state != "OK" || stateCodeInt != 0 {
		// Failed payment update transaction and redirect to failed page
		transaction.Status = orm.BackToSite
		transaction.BankStatus = mysql.NullInt64{Valid: true, Int64: stateCodeInt}
		transaction.Reason = mysql.NullString{Valid: true, String: bankObj.PaymentErr(stateCodeInt).Error()}
		assert.Nil(payOrm.UpdateOnlinePayment(transaction))
		params := url.Values{}
		params.Set("success", "no")
		params.Set("payment", fmt.Sprint(transaction.ID))
		assert.Nil(bankObj.FrontRedirect(w, r, http.StatusFound, params))
		return
	}

	// TODO add others when implement other banks
	if bankName != "saman" {
		transaction.Status = orm.BackToSite
		transaction.BankStatus = mysql.NullInt64{Valid: true, Int64: stateCodeInt}
		transaction.Reason = mysql.NullString{Valid: true, String: errors.NotSupportedBankErr.Error()}
		transaction.RefNum = mysql.NullString{Valid: refNum != "", String: refNum}
		assert.Nil(payOrm.UpdateOnlinePayment(transaction))
		params := url.Values{}
		params.Set("success", "no")
		params.Set("payment", fmt.Sprint(transaction.ID))
		assert.Nil(bankObj.FrontRedirect(w, r, http.StatusFound, params))
		return
	}

	if mID != fmt.Sprint(bankObj.MID()) {
		// Failed payment update transaction and redirect to failed page
		transaction.Status = orm.BackToSite
		transaction.BankStatus = mysql.NullInt64{Valid: true, Int64: stateCodeInt}
		transaction.Reason = mysql.NullString{Valid: true, String: errors.MerchantMismatchErr.Error()}
		transaction.RefNum = mysql.NullString{Valid: refNum != "", String: refNum}
		assert.Nil(payOrm.UpdateOnlinePayment(transaction))
		params := url.Values{}
		params.Set("success", "no")
		params.Set("payment", fmt.Sprint(transaction.ID))
		assert.Nil(bankObj.FrontRedirect(w, r, http.StatusFound, params))
		return
	}

	err = bankObj.VerifyPayment(resNum, refNum, payHash)
	if err != nil {
		// Failed payment update transaction and redirect to failed page
		transaction.Status = orm.BackToSite
		transaction.BankStatus = mysql.NullInt64{Valid: true, Int64: stateCodeInt}
		transaction.Reason = mysql.NullString{Valid: true, String: err.Error()}
		transaction.RefNum = mysql.NullString{Valid: refNum != "", String: refNum}
		assert.Nil(payOrm.UpdateOnlinePayment(transaction))
		params := url.Values{}
		params.Set("success", "no")
		params.Set("payment", fmt.Sprint(transaction.ID))
		assert.Nil(bankObj.FrontRedirect(w, r, http.StatusFound, params))
		return
	}
	// Success payment occurred charge user and redirect to success page
	transaction.Status = orm.Finalized
	transaction.RefNum = mysql.NullString{Valid: true, String: refNum}
	transaction.BankStatus = mysql.NullInt64{Valid: true, Int64: stateCodeInt}
	transaction.TraceNumber = mysql.NullString{Valid: true, String: traceNumber}
	transaction.CID = mysql.NullString{Valid: true, String: cID}
	transaction.Attr = map[string]interface{}{
		"SecurePan": securePan,
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
	params.Set("success", "yes")
	params.Set("payment", fmt.Sprint(transaction.ID))
	assert.Nil(bankObj.FrontRedirect(w, r, http.StatusMovedPermanently, params))
}
