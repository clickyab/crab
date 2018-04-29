// Code generated build with enum DO NOT EDIT.

package aaa

import (
	"database/sql/driver"

	"github.com/clickyab/services/array"
	"github.com/clickyab/services/gettext/t9e"
)

// IsValid try to validate enum value on ths type
func (e AuditActionType) IsValid() bool {
	return array.StringInArray(
		string(e),
		string(InsertNewData),
		string(UpdateData),
		string(DeleteData),
	)
}

// Scan convert the json array ino string slice
func (e *AuditActionType) Scan(src interface{}) error {
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
	if !AuditActionType(b).IsValid() {
		return t9e.G("invalid value")
	}
	*e = AuditActionType(b)
	return nil
}

// Value try to get the string slice representation in database
func (e AuditActionType) Value() (driver.Value, error) {
	if !e.IsValid() {
		return nil, t9e.G("invalid status")
	}
	return string(e), nil
}

// IsValid try to validate enum value on ths type
func (e UserValidStatus) IsValid() bool {
	return array.StringInArray(
		string(e),
		string(RegisteredUserStatus),
		string(BlockedUserStatus),
		string(ActiveUserStatus),
	)
}

// Scan convert the json array ino string slice
func (e *UserValidStatus) Scan(src interface{}) error {
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
	if !UserValidStatus(b).IsValid() {
		return t9e.G("invalid value")
	}
	*e = UserValidStatus(b)
	return nil
}

// Value try to get the string slice representation in database
func (e UserValidStatus) Value() (driver.Value, error) {
	if !e.IsValid() {
		return nil, t9e.G("invalid status")
	}
	return string(e), nil
}

// IsValid try to validate enum value on ths type
func (e GenderType) IsValid() bool {
	return array.StringInArray(
		string(e),
		string(MaleGender),
		string(FemaleGender),
		string(NotSpecifiedGender),
	)
}

// Scan convert the json array ino string slice
func (e *GenderType) Scan(src interface{}) error {
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
	if !GenderType(b).IsValid() {
		return t9e.G("invalid value")
	}
	*e = GenderType(b)
	return nil
}

// Value try to get the string slice representation in database
func (e GenderType) Value() (driver.Value, error) {
	if !e.IsValid() {
		return nil, t9e.G("invalid status")
	}
	return string(e), nil
}
