// Code generated build with enum DO NOT EDIT.

package aaa

import (
	"database/sql/driver"

	"github.com/clickyab/services/array"
	"github.com/clickyab/services/trans"
)

// IsValid try to validate enum value on ths type
func (e ActiveStatus) IsValid() bool {
	return array.StringInArray(
		string(e),
		string(ActiveStatusYes),
		string(ActiveStatusNo),
	)
}

// Scan convert the json array ino string slice
func (e *ActiveStatus) Scan(src interface{}) error {
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
	if !ActiveStatus(b).IsValid() {
		return trans.E("invaid value")
	}
	*e = ActiveStatus(b)
	return nil
}

// Value try to get the string slice representation in database
func (e ActiveStatus) Value() (driver.Value, error) {
	if !e.IsValid() {
		return nil, trans.E("invalid status")
	}
	return string(e), nil
}

// IsValid try to validate enum value on ths type
func (e UserTyp) IsValid() bool {
	return array.StringInArray(
		string(e),
		string(PersonalUserTyp),
		string(CorporationUserTyp),
	)
}

// Scan convert the json array ino string slice
func (e *UserTyp) Scan(src interface{}) error {
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
	if !UserTyp(b).IsValid() {
		return trans.E("invaid value")
	}
	*e = UserTyp(b)
	return nil
}

// Value try to get the string slice representation in database
func (e UserTyp) Value() (driver.Value, error) {
	if !e.IsValid() {
		return nil, trans.E("invalid status")
	}
	return string(e), nil
}

// IsValid try to validate enum value on ths type
func (e UserValidStatus) IsValid() bool {
	return array.StringInArray(
		string(e),
		string(RegisteredUserStatus),
		string(BlockedUserStatus),
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
		return trans.E("unsupported type")
	}
	if !UserValidStatus(b).IsValid() {
		return trans.E("invaid value")
	}
	*e = UserValidStatus(b)
	return nil
}

// Value try to get the string slice representation in database
func (e UserValidStatus) Value() (driver.Value, error) {
	if !e.IsValid() {
		return nil, trans.E("invalid status")
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
		return trans.E("unsupported type")
	}
	if !GenderType(b).IsValid() {
		return trans.E("invaid value")
	}
	*e = GenderType(b)
	return nil
}

// Value try to get the string slice representation in database
func (e GenderType) Value() (driver.Value, error) {
	if !e.IsValid() {
		return nil, trans.E("invalid status")
	}
	return string(e), nil
}
