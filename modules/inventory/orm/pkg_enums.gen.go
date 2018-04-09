// Code generated build with enum DO NOT EDIT.

package orm

import (
	"database/sql/driver"

	"github.com/clickyab/services/array"
	"github.com/clickyab/services/gettext/t9e"
)

// IsValid try to validate enum value on ths type
func (e PublisherType) IsValid() bool {
	return array.StringInArray(
		string(e),
		string(PublisherTypeWeb),
		string(PublisherTypeAPP),
	)
}

// Scan convert the json array ino string slice
func (e *PublisherType) Scan(src interface{}) error {
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
	if !PublisherType(b).IsValid() {
		return t9e.G("invalid value")
	}
	*e = PublisherType(b)
	return nil
}

// Value try to get the string slice representation in database
func (e PublisherType) Value() (driver.Value, error) {
	if !e.IsValid() {
		return nil, t9e.G("invalid status")
	}
	return string(e), nil
}

// IsValid try to validate enum value on ths type
func (e InventoryStatus) IsValid() bool {
	return array.StringInArray(
		string(e),
		string(EnableInventoryStatus),
		string(DisableInventoryStatus),
	)
}

// Scan convert the json array ino string slice
func (e *InventoryStatus) Scan(src interface{}) error {
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
	if !InventoryStatus(b).IsValid() {
		return t9e.G("invalid value")
	}
	*e = InventoryStatus(b)
	return nil
}

// Value try to get the string slice representation in database
func (e InventoryStatus) Value() (driver.Value, error) {
	if !e.IsValid() {
		return nil, t9e.G("invalid status")
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
		return t9e.G("unsupported type")
	}
	if !Status(b).IsValid() {
		return t9e.G("invalid value")
	}
	*e = Status(b)
	return nil
}

// Value try to get the string slice representation in database
func (e Status) Value() (driver.Value, error) {
	if !e.IsValid() {
		return nil, t9e.G("invalid status")
	}
	return string(e), nil
}
