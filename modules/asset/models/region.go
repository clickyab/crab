package models

// Region model in database
// @Model {
//		table = regions
//		primary = false, id
//		list = yes
// }
type Region struct {
	ID string `json:"id" db:"id"`
	baseName
}

// Regions is a unified slice of Region
type Regions []Region

// ListActiveRegions find active browsers by name
func (m *Manager) ListActiveRegions() (Regions, error) {
	var res Regions
	if err := m.allActive(res, OSTableFull); err != nil {
		return nil, err
	}
	return res, nil
}
