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

	"clickyab.com/crab/modules/financial/errors"
	"clickyab.com/crab/modules/financial/payment"
	"github.com/clickyab/services/assert"
	"github.com/clickyab/services/config"
	"github.com/clickyab/services/framework"
	"github.com/clickyab/services/gettext/t9e"
	"github.com/clickyab/services/simplehash"
)

var (
	paymentMethod = config.RegisterString("crab.modules.financial.saman.method", "POST", "saman payment method")
	merchantID    = config.RegisterInt64("crab.modules.financial.saman.merchant", 10903000, "saman merchant id")
	bankURL       = config.RegisterString("crab.modules.financial.saman.url", "https://sep.shaparak.ir/payment.aspx", "saman bank url")
	verifyURL     = config.RegisterString("crab.modules.financial.saman.verify", "https://sep.shaparak.ir/payments/referencepayment.asmx", "saman bank verify url")
	bankName      = config.RegisterString("crab.modules.financial.saman.name", "saman", "saman bank name")

	verifyMethodErrors = map[int64]error{
		-1:  t9e.G("[BANK VERIFY ERR] error in processing request data (problem in one of the params or unsuccessful reverse transaction method"),
		-3:  t9e.G("[BANK VERIFY ERR] params have illegal characters"),
		-4:  t9e.G("[BANK VERIFY ERR] merchant authentication failed"),
		-6:  t9e.G("[BANK VERIFY ERR] money already back to customer or 30 min timeout reached"),
		-7:  t9e.G("[BANK VERIFY ERR] digital code is empty"),
		-8:  t9e.G("[BANK VERIFY ERR] length of params is too long"),
		-9:  t9e.G("[BANK VERIFY ERR] illegal characters in return amount"),
		-10: t9e.G("[BANK VERIFY ERR] digital code is not base 64"),
		-11: t9e.G("[BANK VERIFY ERR] params length is too short"),
		-12: t9e.G("[BANK VERIFY ERR] return amount is negative"),
		-13: t9e.G("[BANK VERIFY ERR] return amount for reverse is more than not returned amount in digital code"),
		-14: t9e.G("[BANK VERIFY ERR] transaction not found"),
		-15: t9e.G("[BANK VERIFY ERR] return amount is float"),
		-16: t9e.G("[BANK VERIFY ERR] internal bank error"),
		-17: t9e.G("[BANK VERIFY ERR] can not return some of the amount"),
		-18: t9e.G("[BANK VERIFY ERR] ip address is not valid or password of pass of reverse transaction method is wrong"),
	}
	payMethodErrors = map[int64]error{
		-1: t9e.G("[BANK PAY ERR] transaction cancelled by user"),
		79: t9e.G("[BANK PAY ERR] return amount is more than transaction amount"),
		12: t9e.G("[BANK PAY ERR] verify method called before finding transaction"),
		14: t9e.G("[BANK PAY ERR] card number is invalid"),
		15: t9e.G("[BANK PAY ERR] no such card owner"),
		33: t9e.G("[BANK PAY ERR] card is expired"),
		38: t9e.G("[BANK PAY ERR] password inserted wrong for 3 times"),
		55: t9e.G("[BANK PAY ERR] password of card is invalid"),
		61: t9e.G("[BANK PAY ERR] the amount is more than valid amount"),
		93: t9e.G("[BANK PAY ERR] transaction has been authorized but pin and pan is wrong"),
		68: t9e.G("[BANK PAY ERR] response timeout"),
		34: t9e.G("[BANK PAY ERR] ccv2 and expDate is wrong or empty"),
		51: t9e.G("[BANK PAY ERR] no sufficient user funds"),
		84: t9e.G("[BANK PAY ERR] card issuer is down"),
		96: t9e.G("[BANK PAY ERR] other bank errors caused this"),
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
	FAmount int64
	FResNum string
	FUserID int64
}

// VerifyPayment for saman bank
func (s *Saman) VerifyPayment(resNum, refNum, hash string) error {
	body := fmt.Sprintf(verifyBody, refNum, s.MID())
	client := &http.Client{}

	req, err := http.NewRequest("POST", s.VerifyURL(), bytes.NewBuffer([]byte(body)))
	assert.Nil(err)

	req.Header.Set("Content-Type", "application/soap+xml")
	req.Header.Set("Cache-Control", "no-cache")

	resp, err := client.Do(req)
	if err != nil {
		return errors.RequestVerifyErr
	}
	defer func() {
		assert.Nil(resp.Body.Close())
	}()

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return errors.RespVerifyErr
	}

	var verifyResponse verifyResponse
	err = xml.Unmarshal(respBody, &verifyResponse)
	assert.Nil(err)

	paid, err := strconv.ParseInt(verifyResponse.Result, 10, 0)
	assert.Nil(err)

	if paid < 0 {
		return s.VerifyErr(paid)
	}

	if s.Amount() != paid {
		return errors.PriceMismatchErr
	}

	//check hash
	generatedHash := simplehash.SHA1(fmt.Sprintf("%d-%s-%d-%s", s.UserID(), bankName.String(), paid, resNum))
	if generatedHash != hash {
		return errors.HashMismatchErr
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
			"Amount":      s.FAmount,
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
	key := simplehash.SHA1(fmt.Sprintf("%d-%s-%d-%s", s.UserID(), s.BankName(), s.Amount(), s.ResNum()))
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
	return t9e.G("[BANK PAY ERR] unexpected error occurred")
}

// VerifyErr verify error
func (s *Saman) VerifyErr(code int64) error {
	res, ok := verifyMethodErrors[code]
	if ok {
		return res
	}
	return t9e.G("[BANK VERIFY ERR] unexpected error occurred")
}
