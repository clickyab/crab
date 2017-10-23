// Code generated build with enum DO NOT EDIT.

package orm

import (
	"database/sql/driver"

	"github.com/clickyab/services/array"
	"github.com/clickyab/services/trans"
)

// IsValid try to validate enum value on ths type
func (e ISPKind) IsValid() bool {
	return array.StringInArray(
		string(e),
		string(BothISPKind),
		string(CellularISPKind),
		string(ISPISPKind),
	)
}

// Scan convert the json array ino string slice
func (e *ISPKind) Scan(src interface{}) error {
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
	if !ISPKind(b).IsValid() {
		return trans.E("invaid value")
	}
	*e = ISPKind(b)
	return nil
}

// Value try to get the string slice representation in database
func (e ISPKind) Value() (driver.Value, error) {
	if !e.IsValid() {
		return nil, trans.E("invalid status")
	}
	return string(e), nil
}