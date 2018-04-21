package saman

import (
	"github.com/clickyab/services/config"
)

var paymentMethod = config.RegisterString("crab.modules.financial.saman.method", "POST", "saman payment method")
var redirectURL = config.RegisterString("crab.modules.financial.saman.redirect", "", "saman redirect url")
var merchantID = config.RegisterInt64("crab.modules.financial.saman.merchant", 10903000, "saman merchant id")
var bankURL = config.RegisterString("crab.modules.financial.saman.url", "https://sep.shaparak.ir/payment.aspx", "saman bank url")

// Saman handle payment for saman gateway
type Saman struct {
	FAmount int64
	FResNum string
}

// InitPaymentResp init payment response
type InitPaymentResp struct {
	Params  map[string]interface{} `json:"params"`
	Method  string                 `json:"method"`
	BankURL string                 `json:"bank_url"`
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
func (s *Saman) InitPayment() *InitPaymentResp {
	return &InitPaymentResp{
		Params: map[string]interface{}{
			"MID":         merchantID.Int64(),
			"ResNum":      s.FResNum,
			"RedirectURL": redirectURL.String(),
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
func (s *Saman) RedirectURL() string {
	return redirectURL.String()
}
