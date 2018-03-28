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
//		primary = false, name
//		find_by = name
//		list = yes
// }
type Browser struct {
	CreatedAt time.Time      `json:"created_at" db:"created_at"`
	UpdatedAt time.Time      `json:"updated_at" db:"updated_at"`
	DeletedAt mysql.NullTime `json:"deleted_at" db:"deleted_at"`
	Status    AssetStatus    `json:"status" db:"status"`
	Name      string         `json:"name" db:"name"`
}

// ListActiveBrowsers find active browsers by name
func (m *Manager) ListActiveBrowsers() ([]Browser, error) {
	var res []Browser
	q := fmt.Sprintf("SELECT * FROM %s WHERE active=?", BrowserTableFull)
	_, err := m.GetRDbMap().Select(&res, q, true)
	if err != nil {
		return nil, err
	}
	return res, nil

}
