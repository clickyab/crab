// Code generated build with enum DO NOT EDIT.

package pub

import (
	"database/sql/driver"

	"github.com/clickyab/services/array"
	"github.com/clickyab/services/trans"
)

// IsValid try to validate enum value on ths type
func (e PubType) IsValid() bool {
	return array.StringInArray(
		string(e),
		string(AppPubType),
		string(WebPubType),
	)
}

// Scan convert the json array ino string slice
func (e *PubType) Scan(src interface{}) error {
	var b []byte
	switch src.(type) {
	case []byte:
		b = src.([]byte)
	case string:
		b = []byte(src.(string))
	case nil:
		b = make([]byte, 0)
	default:
		return trans.E("unsupported type")
	}
	if !PubType(b).IsValid() {
		return trans.E("invaid value")
	}
	*e = PubType(b)
	return nil
}

// Value try to get the string slice representation in database
func (e PubType) Value() (driver.Value, error) {
	if !e.IsValid() {
		return nil, trans.E("invalid status")
	}
	return string(e), nil
}

// IsValid try to validate enum value on ths type
func (e Status) IsValid() bool {
	return array.StringInArray(
		string(e),
		string(ActiveStatus),
		string(PendingStatus),
		string(BlockedStatus),
	)
}

// Scan convert the json array ino string slice
func (e *Status) Scan(src interface{}) error {
	var b []byte
	switch src.(type) {
	case []byte:
		b = src.([]byte)
	case string:
		b = []byte(src.(string))
	case nil:
		b = make([]byte, 0)
	default:
		return trans.E("unsupported type")
	}
	if !Status(b).IsValid() {
		return trans.E("invaid value")
	}
	*e = Status(b)
	return nil
}

// Value try to get the string slice representation in database
func (e Status) Value() (driver.Value, error) {
	if !e.IsValid() {
		return nil, trans.E("invalid status")
	}
	return string(e), nil
}
