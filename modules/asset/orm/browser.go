package orm

import (
	"fmt"
)

// Browser model in database
// @Model {
//		table = browsers
//		primary = true, id
//		find_by = id
//		list = yes
// }
type Browser struct {
	base
	Name string `json:"name" db:"name"`
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
