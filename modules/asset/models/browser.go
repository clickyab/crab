package models

// Browser model in database
// @Model {
//		table = browsers
//		primary = false, name
//		find_by = name
//		list = yes
// }
type Browser struct {
	baseName
}

// Browsers is a unified slice of browser
type Browsers []Browser

// ListActiveBrowsers find active browsers by name
func (m *Manager) ListActiveBrowsers() (Browsers, error) {
	var res Browsers
	if err := m.allActive(res, BrowserTableFull); err != nil {
		return nil, err
	}
	return res, nil

}
