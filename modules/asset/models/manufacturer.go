package models

// Manufacturer model in database
// @Model {
//		table = manufacturers
//		primary = false, name
//		find_by = name
//		list = yes
// }
type Manufacturer struct {
	baseName
}

// Manufacturers is a unified slice of manufacturer
type Manufacturers []Manufacturer

// ListActiveManufacturers find active browsers by name
func (m *Manager) ListActiveManufacturers() (Manufacturers, error) {
	var res Manufacturers
	err := m.allActive(res, ManufacturerTableFull)
	if err != nil {
		return nil, err
	}
	return res, nil
}
