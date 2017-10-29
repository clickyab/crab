package orm

import "time"

// Manufacturer model in database
// @Model {
//		table = manufacturers
//		primary = false, name
//		find_by = name
//		list = yes
// }
type Manufacturer struct {
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
	Active    bool      `json:"active" db:"active"`
	Name      string    `json:"name" db:"name"`
}
