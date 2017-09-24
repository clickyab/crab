// Code generated build with enum DO NOT EDIT.

package add

import (
	"database/sql/driver"

	"github.com/clickyab/services/array"
	"github.com/clickyab/services/trans"
)

// IsValid try to validate enum value on ths type
func (e AdActiveStatus) IsValid() bool {
	return array.StringInArray(
		string(e),
		string(RejectedAdStatus),
		string(AcceptedAdStatus),
		string(PendingAdStatus),
	)
}

// Scan convert the json array ino string slice
func (e *AdActiveStatus) Scan(src interface{}) error {
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
	if !AdActiveStatus(b).IsValid() {
		return trans.E("invaid value")
	}
	*e = AdActiveStatus(b)
	return nil
}

// Value try to get the string slice representation in database
func (e AdActiveStatus) Value() (driver.Value, error) {
	if !e.IsValid() {
		return nil, trans.E("invalid status")
	}
	return string(e), nil
}

// IsValid try to validate enum value on ths type
func (e AdType) IsValid() bool {
	return array.StringInArray(
		string(e),
		string(BannerAdType),
		string(VideoAdType),
		string(NativeAdType),
	)
}

// Scan convert the json array ino string slice
func (e *AdType) Scan(src interface{}) error {
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
	if !AdType(b).IsValid() {
		return trans.E("invaid value")
	}
	*e = AdType(b)
	return nil
}

// Value try to get the string slice representation in database
func (e AdType) Value() (driver.Value, error) {
	if !e.IsValid() {
		return nil, trans.E("invalid status")
	}
	return string(e), nil
}
