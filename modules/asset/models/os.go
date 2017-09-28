package models

// OS os model in database
// @Model {
//		table = oses
//		primary = false, name
//		find_by = name
//		list = yes
// }
type OS struct {
	baseName
}

// OSes is a unified slice of OS
type OSes []OS

// ListActiveOSes find active browsers by name
func (m *Manager) ListActiveOSes() (OSes, error) {
	var res OSes
	if err := m.allActive(res, OSTableFull); err != nil {
		return nil, err
	}
	return res, nil
}
