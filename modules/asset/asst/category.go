package asst

// Category category model in database
// @Model {
//		table = categories
//		primary = true, id
//		find_by = id,name
// }
type Category struct {
	ID          int64  `json:"id" db:"id"`
	Name        string `json:"name" db:"name"`
	Description string `json:"description" db:"description"`
	Status      bool   `json:"status" db:"status"`
}
