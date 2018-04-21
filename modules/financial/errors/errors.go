package errors

import (
	"github.com/clickyab/services/gettext/t9e"
)

var (
	// InvalidIDErr invalid id, error
	InvalidIDErr error = t9e.G("invalid id, please check your request data.")
	// UnsupportTypeError unsupported type. error
	UnsupportTypeError error = t9e.G("unsupported type.")
	// TypeError invalid bill type error
	TypeError error = t9e.G("invalid bill type. you can select %s or %s or %s please check your request data and try again", "banner", "vast", "native")
	// UpdateError when want update bill or related data
	UpdateError error = t9e.G("can't update bill main or related data")
	// CreateError can't create new bill
	CreateError error = t9e.G("db error! can't create new bill")
	// MinBankSnapErr min min bank snap error
	MinBankSnapErr error = t9e.G("minimum money not met")
	// DisableGateWayErr gateway is disabled
	DisableGateWayErr error = t9e.G("disabled gateway")
	// NotFoundGateway gateway not found
	NotFoundGateway error = t9e.G("gateway not found")
	// MakeOnlinePaymentErr error while making online payments error while making online payments
	MakeOnlinePaymentErr error = t9e.G("error while making online payments")
	//GateWayNotSupportedErr gate way not supported
	GateWayNotSupportedErr error = t9e.G("gateway not supported")
)

// NotFoundError maker
func NotFoundError(id int64) error {
	if id > 0 {
		return t9e.G("bill with identifier %s not found, please check your request data.", id)
	}

	return t9e.G("bill not found, please check your request data.")
}

// MinChargeError min charge not met
func MinChargeError(min int64) error {
	return t9e.G("min charge should be %d", min)
}

// InvalidError maker
func InvalidError(dataName string) error {
	return t9e.G("Invalid %s. please check your request data and try again", dataName)
}
