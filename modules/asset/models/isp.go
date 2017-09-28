package models

// ISP isp model in database
// @Model {
//		table = isps
//		primary = false, name
//		find_by = name
//		list = yes
// }
type ISP struct {
	baseName
}

// ISPs is a unified slice of ISP
type ISPs []ISP

// ListActiveISPs find active browsers by name
func (m *Manager) ListActiveISPs() (ISPs, error) {
	var res ISPs
	if err := m.allActive(res, ISPTableFull); err != nil {
		return nil, err
	}
	return res, nil
}
