package orm

import (
	"time"

	"github.com/clickyab/services/mysql"
)

// PublisherType is the pub type
type (
	// PublisherType is the pub type
	// @Enum{
	// }
	PublisherType string
)

const (
	// PublisherTypeWeb web pub type
	PublisherTypeWeb PublisherType = "web"
	// PublisherTypeAPP web pub type
	PublisherTypeAPP PublisherType = "app"
)

// WhiteBlackList user_wlbl_presets model in database
// @Model {
//		table = user_wlbl_presets
//		primary = true, id
//		find_by = id
//		list = yes
// }
type WhiteBlackList struct {
	ID        int64                 `json:"id" db:"id"`
	CreatedAt time.Time             `json:"created_at" db:"created_at"`
	UpdatedAt time.Time             `json:"updated_at" db:"updated_at"`
	Active    bool                  `json:"active" db:"active"`
	UserID    int64                 `json:"user_id" db:"user_id"`
	Label     string                `json:"label" db:"label"`
	Domains   mysql.StringJSONArray `json:"domains" db:"domains"`
	// Kind shows if it's a white list (true) or blacklist (false)
	Kind          bool          `json:"kind" db:"kind"`
	PublisherType PublisherType `json:"publisher_type" db:"publisher_type"`
}
