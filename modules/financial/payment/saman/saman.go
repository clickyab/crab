package saman

import (
	"bytes"
	"fmt"
	"net/http"
	"strconv"

	"path"

	"encoding/xml"
	"io/ioutil"

	"net/url"

	"clickyab.com/crab/modules/financial/orm"
	"clickyab.com/crab/modules/financial/payment"
	"github.com/clickyab/services/assert"
	"github.com/clickyab/services/config"
	"github.com/clickyab/services/framework"
	"github.com/clickyab/services/simplehash"
)

var (
	paymentMethod        = config.RegisterString("crab.modules.financial.saman.method", "POST", "saman payment method")
	merchantID           = config.RegisterInt64("crab.modules.financial.saman.merchant", 10903000, "saman merchant id")
	bankURL              = config.RegisterString("crab.modules.financial.saman.url", "https://sep.shaparak.ir/payment.aspx", "saman bank url")
	verifyURL            = config.RegisterString("crab.modules.financial.saman.verify", "https://sep.shaparak.ir/payments/referencepayment.asmx", "saman bank verify url")
	bankName             = config.RegisterString("crab.modules.financial.saman.name", "saman", "saman bank name")
	currencyTransferRate = config.RegisterInt("crab.modules.financial.saman.currency.rate", 10, "payment currency rate")

	verifyMethodErrors = map[int64]error{
		-1:  orm.VERIFYRequestData,
		-3:  orm.VERIFYParams,
		-4:  orm.VERIFYMerchant,
		-6:  orm.VERIFYTimeout,
		-7:  orm.VERIFYEmptyDigitalCode,
		-8:  orm.VERIFYLongParamsLength,
		-9:  orm.VERIFYIllegalCharsReturnAmount,
		-10: orm.VERIFYDigitalCodeInvalid,
		-11: orm.VERIFYLongParams,
		-12: orm.VERIFYNegativeReturnAmount,
		-13: orm.VERIFYReturnAmountMismatch,
		-14: orm.VERIFYNotFoundTransaction,
		-15: orm.VERIFYFloatReturnAmount,
		-16: orm.VERIFYInternalBankErr,
		-17: orm.VERIFYReturnSomeOfAmount,
		-18: orm.VERIFYIPPassReverseInvalid,
	}
	payMethodErrors = map[int64]error{
		-1: orm.PAYTransactionCancelled,
		79: orm.PAYReturnMoreThanTransaction,
		12: orm.PAYEarlyVerifyCalled,
		14: orm.PAYWrongCardNumber,
		15: orm.PAYCardNotFound,
		33: orm.PAYCardExpired,
		38: orm.PAYExceedWrongPassCard,
		55: orm.PAYPasswordCardWrong,
		61: orm.PAYAmountMoreThanValid,
		93: orm.PAYPanPinWrong,
		68: orm.PAYResponseTimeout,
		34: orm.PAYCCV2ExpDate,
		51: orm.PAYSufficientFunds,
		84: orm.PAYCardIssuerInvalid,
		96: orm.PAYOtherBankErr,
	}
)

const (
	verifyBody = `<?xml version="1.0" encoding="utf-8"?>
	<soap12:Envelope xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xmlns:xsd="http://www.w3.org/2001/XMLSchema" xmlns:soapenc="http://www.w3.org/2003/05/soap-encoding" xmlns:tns="urn:Foo" xmlns:types="urn:Foo/encodedTypes" xmlns:rpc="http://www.w3.org/2003/05/soap-rpc" xmlns:soap12="http://www.w3.org/2003/05/soap-envelope">
	<soap12:Body soap12:encodingStyle="http://www.w3.org/2003/05/soap-encoding">
	  <tns:verifyTransaction>
		<String_1 xsi:type="xsd:string">%s</String_1>
		<String_2 xsi:type="xsd:string">%d</String_2>
	  </tns:verifyTransaction>
	</soap12:Body>
	</soap12:Envelope>`
)

type verifyResponse struct {
	XMLName xml.Name `xml:"Envelope"`
	Result  string   `xml:"Body>verifyTransactionResponse>result"`
}

// Saman handle payment for saman gateway
type Saman struct {
	payment.CommonPay
}

// RedirectValidation validate redirect from bank
func (s *Saman) RedirectValidation(p payment.RedirectParams) error {
	if p.Attr["state"] != "OK" || p.StatusCode != 0 {
		return s.PaymentErr(p.StatusCode)
	}

	if p.Attr["mID"] != fmt.Sprint(merchantID.Int64()) {
		return orm.MerchantMismatchErr
	}
	return nil
}

// VerifyTransaction for saman bank
func (s *Saman) VerifyTransaction(resNum, refNum string) error {
	body := fmt.Sprintf(verifyBody, refNum, s.MID())
	client := &http.Client{}

	req, err := http.NewRequest("POST", s.VerifyURL(), bytes.NewBuffer([]byte(body)))
	assert.Nil(err)

	req.Header.Set("Content-Type", "application/soap+xml")
	req.Header.Set("Cache-Control", "no-cache")

	resp, err := client.Do(req)
	if err != nil {
		return orm.RequestVerifyErr
	}
	defer func() {
		assert.Nil(resp.Body.Close())
	}()

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return orm.RespVerifyErr
	}

	var verifyResponse verifyResponse
	err = xml.Unmarshal(respBody, &verifyResponse)
	assert.Nil(err)

	paid, err := strconv.ParseInt(verifyResponse.Result, 10, 0)
	assert.Nil(err)

	if paid < 0 {
		return s.VerifyErr(paid)
	}

	if s.PayAmount() != paid {
		return orm.PriceMismatchErr
	}
	return nil
}

// HashVerification double verification
func (s *Saman) HashVerification(paid int64, hash, resNum string) error {
	generatedHash := simplehash.SHA1(fmt.Sprintf("%d-%s-%d-%s", s.UserID(), bankName.String(), paid, resNum))
	if generatedHash != hash {
		return orm.HashMismatchErr
	}
	return nil
}

// UserID get user id
func (s *Saman) UserID() int64 {
	return s.FUserID
}

// BankName return bank name
func (s *Saman) BankName() string {
	return bankName.String()
}

// Amount return amount
func (s *Saman) Amount() int64 {
	return s.FAmount
}

// GetPaymentURL return payment url
func (s *Saman) GetPaymentURL() string {
	return bankURL.String()
}

// GetPaymentMethod return payment method
func (s *Saman) GetPaymentMethod() string {
	return paymentMethod.String()
}

// ResNum return ResNum
func (s *Saman) ResNum() string {
	return s.FResNum
}

// InitPayment init payment and  return payment params
func (s *Saman) InitPayment(r *http.Request) *payment.InitPaymentResp {
	return &payment.InitPaymentResp{
		Params: map[string]interface{}{
			"MID":         merchantID.Int64(),
			"ResNum":      s.FResNum,
			"RedirectURL": s.RedirectURL(r),
			"Amount":      s.PayAmount(),
		},
		Method:  paymentMethod.String(),
		BankURL: bankURL.String(),
	}

}

// MID return MID
func (s *Saman) MID() int64 {
	return merchantID.Int64()
}

// RedirectURL return RedirectURL
func (s *Saman) RedirectURL(r *http.Request) string {
	key := simplehash.SHA1(fmt.Sprintf("%d-%s-%d-%s", s.UserID(), s.BankName(), s.PayAmount(), s.ResNum()))
	res := &url.URL{
		Host:   r.Host,
		Scheme: framework.Scheme(r),
		Path:   path.Join("api/financial/payment/return", s.BankName(), key),
	}
	return res.String()
}

// VerifyURL return verify url
func (s *Saman) VerifyURL() string {
	return verifyURL.String()
}

// PaymentErr payment error
func (s *Saman) PaymentErr(code int64) error {
	res, ok := payMethodErrors[code]
	if ok {
		return res
	}
	return orm.PayNotSupported
}

// VerifyErr verify error
func (s *Saman) VerifyErr(code int64) error {
	res, ok := verifyMethodErrors[code]
	if ok {
		return res
	}
	return orm.VERIFYNotSupported
}

// GetParams get bank redirect params
func (s *Saman) GetParams(r *http.Request) payment.RedirectParams {
	stateCode := r.PostFormValue("StateCode")
	stateCodeInt, err := strconv.ParseInt(stateCode, 10, 0)
	assert.Nil(err)
	res := payment.RedirectParams{
		RefNum:     r.PostFormValue("RefNum"),
		StatusCode: stateCodeInt,
		ResNum:     r.PostFormValue("ResNum"),
		Attr: map[string]string{
			"state":       r.PostFormValue("State"),
			"traceNumber": r.PostFormValue("TRACENO"),
			"cID":         r.PostFormValue("CID"),
			"securePan":   r.PostFormValue("SecurePan"),
			"mID":         r.PostFormValue("MID"),
		},
	}

	return res
}

// SetAmount set amount in underlay struct
func (s *Saman) SetAmount(amount int64) {
	s.FAmount = amount
}

// SetUserID set user id
func (s *Saman) SetUserID(userID int64) {
	s.FUserID = userID
}

// SetResNum set res num
func (s *Saman) SetResNum(res string) {
	s.FResNum = res
}

// PayAmount return bank pay amount
func (s *Saman) PayAmount() int64 {
	return s.Amount() * currencyTransferRate.Int64()
}
