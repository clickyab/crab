package orm

import (
	"database/sql/driver"
	"time"

	"github.com/clickyab/services/array"
	"github.com/clickyab/services/mysql"
	"github.com/clickyab/services/trans"
)

type publisherType string

const (
	publisherTypeWeb = "web"
	publisherTypeAPP = "app"
)

// IsValid try to validate enum value on ths type
func (e publisherType) IsValid() bool {
	return array.StringInArray(
		string(e),
		string(publisherTypeWeb),
		string(publisherTypeAPP),
	)
}

// Scan convert the json array ino string slice
func (e *publisherType) Scan(src interface{}) error {
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
	if !publisherType(b).IsValid() {
		return trans.E("invaid value")
	}
	*e = publisherType(b)
	return nil
}

// Value try to get the string slice representation in database
func (e publisherType) Value() (driver.Value, error) {
	if !e.IsValid() {
		return nil, trans.E("invalid status")
	}
	return string(e), nil
}

// Inventory user_inventories model in database
// @Model {
//		table = user_inventories
//		primary = true, id
//		find_by = id, user_id
//		list = yes
// }
type Inventory struct {
	ID            int64                 `json:"id" db:"id"`
	CreatedAt     time.Time             `json:"created_at" db:"created_at"`
	UpdatedAt     time.Time             `json:"updated_at" db:"updated_at"`
	Active        bool                  `json:"active" db:"active"`
	UserID        int64                 `json:"user_id" db:"user_id"`
	Label         string                `json:"label" db:"label"`
	Domains       mysql.StringJSONArray `json:"domains" db:"domains"`
	Kind          bool                  `json:"kind" db:"kind"`
	PublisherType publisherType         `json:"publisher_type" db:"publisher_type"`
}
