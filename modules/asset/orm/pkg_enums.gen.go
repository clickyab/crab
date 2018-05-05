// Code generated build with enum DO NOT EDIT.

package orm

import (
	"database/sql/driver"

	"github.com/clickyab/services/array"
	"github.com/clickyab/services/gettext/t9e"
)

// IsValid try to validate enum value on ths type
func (e AssetStatus) IsValid() bool {
	return array.StringInArray(
		string(e),
		string(EnableAssetStatus),
		string(DisableAssetStatus),
	)
}

// Scan convert the json array ino string slice
func (e *AssetStatus) Scan(src interface{}) error {
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
	if !AssetStatus(b).IsValid() {
		return t9e.G("invalid value")
	}
	*e = AssetStatus(b)
	return nil
}

// Value try to get the string slice representation in database
func (e AssetStatus) Value() (driver.Value, error) {
	if !e.IsValid() {
		return nil, t9e.G("invalid status")
	}
	return string(e), nil
}

// IsValid try to validate enum value on ths type
func (e ModelKind) IsValid() bool {
	return array.StringInArray(
		string(e),
		string(CampaignModel),
		string(PublisherModel),
	)
}

// Scan convert the json array ino string slice
func (e *ModelKind) Scan(src interface{}) error {
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
	if !ModelKind(b).IsValid() {
		return t9e.G("invalid value")
	}
	*e = ModelKind(b)
	return nil
}

// Value try to get the string slice representation in database
func (e ModelKind) Value() (driver.Value, error) {
	if !e.IsValid() {
		return nil, t9e.G("invalid status")
	}
	return string(e), nil
}

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
		return t9e.G("unsupported type")
	}
	if !ISPKind(b).IsValid() {
		return t9e.G("invalid value")
	}
	*e = ISPKind(b)
	return nil
}

// Value try to get the string slice representation in database
func (e ISPKind) Value() (driver.Value, error) {
	if !e.IsValid() {
		return nil, t9e.G("invalid status")
	}
	return string(e), nil
}
