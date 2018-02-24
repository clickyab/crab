// Code generated build with enum DO NOT EDIT.

package orm

import (
	"database/sql/driver"

	"github.com/clickyab/services/array"
	"github.com/clickyab/services/gettext/t9e"
)

// IsValid try to validate enum value on ths type
func (e CampaignKind) IsValid() bool {
	return array.StringInArray(
		string(e),
		string(WebCampaign),
		string(AppCampaign),
	)
}

// Scan convert the json array ino string slice
func (e *CampaignKind) Scan(src interface{}) error {
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
	if !CampaignKind(b).IsValid() {
		return t9e.G("invalid value")
	}
	*e = CampaignKind(b)
	return nil
}

// Value try to get the string slice representation in database
func (e CampaignKind) Value() (driver.Value, error) {
	if !e.IsValid() {
		return nil, t9e.G("invalid status")
	}
	return string(e), nil
}

// IsValid try to validate enum value on ths type
func (e CampaignType) IsValid() bool {
	return array.StringInArray(
		string(e),
		string(BannerType),
		string(VastType),
		string(NativeType),
	)
}

// Scan convert the json array ino string slice
func (e *CampaignType) Scan(src interface{}) error {
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
	if !CampaignType(b).IsValid() {
		return t9e.G("invalid value")
	}
	*e = CampaignType(b)
	return nil
}

// Value try to get the string slice representation in database
func (e CampaignType) Value() (driver.Value, error) {
	if !e.IsValid() {
		return nil, t9e.G("invalid status")
	}
	return string(e), nil
}

// IsValid try to validate enum value on ths type
func (e Progress) IsValid() bool {
	return array.StringInArray(
		string(e),
		string(ProgressInProgress),
		string(ProgressFinalized),
	)
}

// Scan convert the json array ino string slice
func (e *Progress) Scan(src interface{}) error {
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
	if !Progress(b).IsValid() {
		return t9e.G("invalid value")
	}
	*e = Progress(b)
	return nil
}

// Value try to get the string slice representation in database
func (e Progress) Value() (driver.Value, error) {
	if !e.IsValid() {
		return nil, t9e.G("invalid status")
	}
	return string(e), nil
}

// IsValid try to validate enum value on ths type
func (e CostType) IsValid() bool {
	return array.StringInArray(
		string(e),
		string(CPM),
		string(CPC),
		string(CPA),
	)
}

// Scan convert the json array ino string slice
func (e *CostType) Scan(src interface{}) error {
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
	if !CostType(b).IsValid() {
		return t9e.G("invalid value")
	}
	*e = CostType(b)
	return nil
}

// Value try to get the string slice representation in database
func (e CostType) Value() (driver.Value, error) {
	if !e.IsValid() {
		return nil, t9e.G("invalid status")
	}
	return string(e), nil
}
