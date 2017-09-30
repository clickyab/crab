package models

// Category category model in database
// @Model {
//		table = categories
//		primary = false, name
//		find_by = name
//		list = yes
// }
type Category struct {
	baseName
}

// Categories is a unified slice of category
type Categories []Category

// ListActiveCategories find active browsers by name
func (m *Manager) ListActiveCategories() (Categories, error) {
	var res Categories
	if err := m.allActive(res, CategoryTableFull); err != nil {
		return nil, err
	}
	return res, nil
}
