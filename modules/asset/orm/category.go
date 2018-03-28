package orm

import "github.com/go-sql-driver/mysql"

// Category category model in database
// @Model {
//		table = categories
//		primary = false, name
//		find_by = name
//		list = yes
// }
type Category struct {
	Name        string         `json:"name" db:"name"`
	Description string         `json:"description" db:"description"`
	DeletedAt   mysql.NullTime `json:"deleted_at" db:"deleted_at"`
}
