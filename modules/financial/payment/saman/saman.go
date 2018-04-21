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
	FPayAmount    int64
	FChargeAmount int64
	FVatAmount    int64
	FResNum       string
}

// GetPaymentURL return payment url
func (s *Saman) GetPaymentURL() string {
	return bankURL.String()
}

// GetPaymentMethod return payment method
func (s *Saman) GetPaymentMethod() string {
	return paymentMethod.String()
}

// PayAmount return PayAmount
func (s *Saman) PayAmount() int64 {
	return s.FPayAmount
}

// ChargeAmount return charge amount
func (s *Saman) ChargeAmount() int64 {
	return s.FChargeAmount
}

// VatAmount return VatAmount
func (s *Saman) VatAmount() int64 {
	return s.FVatAmount
}

// ResNum return ResNum
func (s *Saman) ResNum() string {
	return s.FResNum
}

// GetPaymentParams return payment params
func (s *Saman) GetPaymentParams() map[string]interface{} {
	return map[string]interface{}{
		"MID":         merchantID.Int64(),
		"ResNum":      s.FResNum,
		"RedirectURL": redirectURL.String(),
		"Amount":      s.FPayAmount,
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
