package orm

import (
	"fmt"
	"time"

	"github.com/clickyab/services/mysql"
)

// AssetStatus is the user asset status
type (
	// AssetStatus is the user asset status
	// @Enum{
	// }
	AssetStatus string
)

const (
	// EnableAssetStatus enable asset
	EnableAssetStatus AssetStatus = "enable"
	// DisableAssetStatus disable asset
	DisableAssetStatus AssetStatus = "disable"
)

// Browser model in database
// @Model {
//		table = browsers
//		primary = true, id
//		find_by = name,id
//		list = yes
// }
type Browser struct {
	ID        int64          `json:"id" db:"id"`
	CreatedAt time.Time      `json:"created_at" db:"created_at"`
	UpdatedAt time.Time      `json:"updated_at" db:"updated_at"`
	DeletedAt mysql.NullTime `json:"deleted_at" db:"deleted_at"`
	Name      string         `json:"name" db:"name"`
	Mobile    bool           `json:"mobile" db:"mobile"`
	Desktop   bool           `json:"desktop" db:"desktop"`
	Tablet    bool           `json:"tablet" db:"tablet"`
	Other     bool           `json:"other" db:"other"`
}

// ListActiveBrowsers find active browsers by name
func (m *Manager) ListActiveBrowsers() ([]Browser, error) {
	var res []Browser
	q := fmt.Sprintf("SELECT %s FROM %s", GetSelectFields(BrowserTableFull, ""), BrowserTableFull)
	_, err := m.GetRDbMap().Select(&res, q)
	if err != nil {
		return nil, err
	}
	return res, nil

}
