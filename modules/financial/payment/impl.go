package payment

import "clickyab.com/crab/modules/financial/payment/saman"

// Payable handle payment
type Payable interface {
	InitPayment() *saman.InitPaymentResp
	GetPaymentURL() string
	GetPaymentMethod() string

	MID() int64
	RedirectURL() string
	Amount() int64
	ResNum() string
}
