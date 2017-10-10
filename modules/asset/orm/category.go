package orm

// Category category model in database
// @Model {
//		table = categories
//		primary = false, name
//		find_by = name
//		list = yes
// }
type Category struct {
	Name        string `json:"name" db:"name"`
	Description string `json:"description" db:"description"`
	Active      bool   `json:"active" db:"active"`
}
