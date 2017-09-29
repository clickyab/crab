package models

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

// Presets is slice of preset
type Presets []Preset

// Preset is model table for users_presets in database
// @Model {
//		table = users_presets
//		primary = true, id
//		find_by = id
//		list = yes
// }
type Preset struct {
	ID        int64     `json:"id" db:"id"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
	Active    bool      `json:"active" db:"active"`
	UserID    int64     `json:"user_id" db:"user_id"`
	PresetData
}

// PresetData is core data that required for making preset
type PresetData struct {
	Label   string                `json:"label" db:"label" validate:"gt=7"`
	Domains mysql.StringJSONArray `json:"domains" db:"domains" validate:"gt=1"`
	// Kind shows if it's a white list (true) or blacklist (false)
	Kind          bool          `json:"kind" db:"kind"`
	PublisherType PublisherType `json:"publisher_type" db:"publisher_type" valid:"eg='web'|eg='app'"`
}
