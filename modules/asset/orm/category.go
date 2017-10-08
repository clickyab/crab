package orm

// Category category model in database
// @Model {
//		table = categories
//		primary = true, id
//		find_by = id,name
//		list = yes
// }
type Category struct {
	ID          int64  `json:"id" db:"id"`
	Name        string `json:"name" db:"name"`
	Description string `json:"description" db:"description"`
	Active      bool   `json:"active" db:"active"`
}
