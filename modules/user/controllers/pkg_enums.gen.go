// Code generated build with enum DO NOT EDIT.

package user

import (
	"database/sql/driver"

	"github.com/clickyab/services/array"
	"github.com/clickyab/services/gettext/t9e"
)

// IsValid try to validate enum value on ths type
func (e AccountType) IsValid() bool {
	return array.StringInArray(
		string(e),
		string(PersonalUser),
		string(CorporationUser),
	)
}

// Scan convert the json array ino string slice
func (e *AccountType) Scan(src interface{}) error {
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
	if !AccountType(b).IsValid() {
		return t9e.G("invalid value")
	}
	*e = AccountType(b)
	return nil
}

// Value try to get the string slice representation in database
func (e AccountType) Value() (driver.Value, error) {
	if !e.IsValid() {
		return nil, t9e.G("invalid status")
	}
	return string(e), nil
}
