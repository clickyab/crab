// Code generated build with enum DO NOT EDIT.

package orm

import (
	"database/sql/driver"

	"github.com/clickyab/services/array"
	"github.com/clickyab/services/gettext/t9e"
)

// IsValid try to validate enum value on ths type
func (e BankSnapCheckStatus) IsValid() bool {
	return array.StringInArray(
		string(e),
		string(PendingStatus),
		string(AcceptedStatus),
		string(RejectedStatus),
	)
}

// Scan convert the json array ino string slice
func (e *BankSnapCheckStatus) Scan(src interface{}) error {
	var b []byte
	switch src.(type) {
	case []byte:
		b = src.([]byte)
	case string:
		b = []byte(src.(string))
	case nil:
		b = make([]byte, 0)
	default:
		return t9e.G("unsupported type")
	}
	if !BankSnapCheckStatus(b).IsValid() {
		return t9e.G("invalid value")
	}
	*e = BankSnapCheckStatus(b)
	return nil
}

// Value try to get the string slice representation in database
func (e BankSnapCheckStatus) Value() (driver.Value, error) {
	if !e.IsValid() {
		return nil, t9e.G("invalid status")
	}
	return string(e), nil
}

// IsValid try to validate enum value on ths type
func (e PayModels) IsValid() bool {
	return array.StringInArray(
		string(e),
		string(OnlinePaymentModel),
		string(BankSnapModel),
		string(ManualCashChangeModel),
	)
}

// Scan convert the json array ino string slice
func (e *PayModels) Scan(src interface{}) error {
	var b []byte
	switch src.(type) {
	case []byte:
		b = src.([]byte)
	case string:
		b = []byte(src.(string))
	case nil:
		b = make([]byte, 0)
	default:
		return t9e.G("unsupported type")
	}
	if !PayModels(b).IsValid() {
		return t9e.G("invalid value")
	}
	*e = PayModels(b)
	return nil
}

// Value try to get the string slice representation in database
func (e PayModels) Value() (driver.Value, error) {
	if !e.IsValid() {
		return nil, t9e.G("invalid status")
	}
	return string(e), nil
}

// IsValid try to validate enum value on ths type
func (e GatewayStatus) IsValid() bool {
	return array.StringInArray(
		string(e),
		string(Disabled),
		string(Enabled),
	)
}

// Scan convert the json array ino string slice
func (e *GatewayStatus) Scan(src interface{}) error {
	var b []byte
	switch src.(type) {
	case []byte:
		b = src.([]byte)
	case string:
		b = []byte(src.(string))
	case nil:
		b = make([]byte, 0)
	default:
		return t9e.G("unsupported type")
	}
	if !GatewayStatus(b).IsValid() {
		return t9e.G("invalid value")
	}
	*e = GatewayStatus(b)
	return nil
}

// Value try to get the string slice representation in database
func (e GatewayStatus) Value() (driver.Value, error) {
	if !e.IsValid() {
		return nil, t9e.G("invalid status")
	}
	return string(e), nil
}

// IsValid try to validate enum value on ths type
func (e DefaultType) IsValid() bool {
	return array.StringInArray(
		string(e),
		string(IsDefault),
		string(IsNotDefault),
	)
}

// Scan convert the json array ino string slice
func (e *DefaultType) Scan(src interface{}) error {
	var b []byte
	switch src.(type) {
	case []byte:
		b = src.([]byte)
	case string:
		b = []byte(src.(string))
	case nil:
		b = make([]byte, 0)
	default:
		return t9e.G("unsupported type")
	}
	if !DefaultType(b).IsValid() {
		return t9e.G("invalid value")
	}
	*e = DefaultType(b)
	return nil
}

// Value try to get the string slice representation in database
func (e DefaultType) Value() (driver.Value, error) {
	if !e.IsValid() {
		return nil, t9e.G("invalid status")
	}
	return string(e), nil
}

// IsValid try to validate enum value on ths type
func (e ChangeCashReasons) IsValid() bool {
	return array.StringInArray(
		string(e),
		string(Gift),
		string(ManualPay),
		string(Refund),
	)
}

// Scan convert the json array ino string slice
func (e *ChangeCashReasons) Scan(src interface{}) error {
	var b []byte
	switch src.(type) {
	case []byte:
		b = src.([]byte)
	case string:
		b = []byte(src.(string))
	case nil:
		b = make([]byte, 0)
	default:
		return t9e.G("unsupported type")
	}
	if !ChangeCashReasons(b).IsValid() {
		return t9e.G("invalid value")
	}
	*e = ChangeCashReasons(b)
	return nil
}

// Value try to get the string slice representation in database
func (e ChangeCashReasons) Value() (driver.Value, error) {
	if !e.IsValid() {
		return nil, t9e.G("invalid status")
	}
	return string(e), nil
}

// IsValid try to validate enum value on ths type
func (e ManualCashChangeStatus) IsValid() bool {
	return array.StringInArray(
		string(e),
		string(WaitToCheckCashChange),
		string(VerifiedCashChange),
		string(RejectedCashChange),
	)
}

// Scan convert the json array ino string slice
func (e *ManualCashChangeStatus) Scan(src interface{}) error {
	var b []byte
	switch src.(type) {
	case []byte:
		b = src.([]byte)
	case string:
		b = []byte(src.(string))
	case nil:
		b = make([]byte, 0)
	default:
		return t9e.G("unsupported type")
	}
	if !ManualCashChangeStatus(b).IsValid() {
		return t9e.G("invalid value")
	}
	*e = ManualCashChangeStatus(b)
	return nil
}

// Value try to get the string slice representation in database
func (e ManualCashChangeStatus) Value() (driver.Value, error) {
	if !e.IsValid() {
		return nil, t9e.G("invalid status")
	}
	return string(e), nil
}

// IsValid try to validate enum value on ths type
func (e OnlinePaymentStatus) IsValid() bool {
	return array.StringInArray(
		string(e),
		string(Init),
		string(BackToSite),
		string(Finalized),
	)
}

// Scan convert the json array ino string slice
func (e *OnlinePaymentStatus) Scan(src interface{}) error {
	var b []byte
	switch src.(type) {
	case []byte:
		b = src.([]byte)
	case string:
		b = []byte(src.(string))
	case nil:
		b = make([]byte, 0)
	default:
		return t9e.G("unsupported type")
	}
	if !OnlinePaymentStatus(b).IsValid() {
		return t9e.G("invalid value")
	}
	*e = OnlinePaymentStatus(b)
	return nil
}

// Value try to get the string slice representation in database
func (e OnlinePaymentStatus) Value() (driver.Value, error) {
	if !e.IsValid() {
		return nil, t9e.G("invalid status")
	}
	return string(e), nil
}
