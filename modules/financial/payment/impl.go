package payment

// Payable handle payment
type Payable interface {
	GetPaymentParams() map[string]interface{}
	GetPaymentURL() string
	GetPaymentMethod() string

	MID() int64
	RedirectURL() string
	PayAmount() int64
	ChargeAmount() int64
	VatAmount() int64
}
