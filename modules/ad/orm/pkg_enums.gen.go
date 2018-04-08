// Code generated build with enum DO NOT EDIT.

package orm

import (
	"database/sql/driver"

	"github.com/clickyab/services/array"
	"github.com/clickyab/services/gettext/t9e"
)

// IsValid try to validate enum value on ths type
func (e AssetTypes) IsValid() bool {
	return array.StringInArray(
		string(e),
		string(AssetImageType),
		string(AssetVideoType),
		string(AssetTextType),
		string(AssetNumberType),
	)
}

// Scan convert the json array ino string slice
func (e *AssetTypes) Scan(src interface{}) error {
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
	if !AssetTypes(b).IsValid() {
		return t9e.G("invalid value")
	}
	*e = AssetTypes(b)
	return nil
}

// Value try to get the string slice representation in database
func (e AssetTypes) Value() (driver.Value, error) {
	if !e.IsValid() {
		return nil, t9e.G("invalid status")
	}
	return string(e), nil
}

// IsValid try to validate enum value on ths type
func (e CreativeStatusType) IsValid() bool {
	return array.StringInArray(
		string(e),
		string(RejectedCreativeStatus),
		string(AcceptedCreativeStatus),
		string(PendingCreativeStatus),
	)
}

// Scan convert the json array ino string slice
func (e *CreativeStatusType) Scan(src interface{}) error {
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
	if !CreativeStatusType(b).IsValid() {
		return t9e.G("invalid value")
	}
	*e = CreativeStatusType(b)
	return nil
}

// Value try to get the string slice representation in database
func (e CreativeStatusType) Value() (driver.Value, error) {
	if !e.IsValid() {
		return nil, t9e.G("invalid status")
	}
	return string(e), nil
}

// IsValid try to validate enum value on ths type
func (e CreativeTypes) IsValid() bool {
	return array.StringInArray(
		string(e),
		string(CreativeBannerType),
		string(CreativeVastType),
		string(CreativeNativeType),
	)
}

// Scan convert the json array ino string slice
func (e *CreativeTypes) Scan(src interface{}) error {
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
	if !CreativeTypes(b).IsValid() {
		return t9e.G("invalid value")
	}
	*e = CreativeTypes(b)
	return nil
}

// Value try to get the string slice representation in database
func (e CreativeTypes) Value() (driver.Value, error) {
	if !e.IsValid() {
		return nil, t9e.G("invalid status")
	}
	return string(e), nil
}
