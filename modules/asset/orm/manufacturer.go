package orm

import "time"

// Manufacturer model in database
// @Model {
//		table = manufacturers
//		primary = false, brand
//		find_by = brand
//		list = yes
// }
type Manufacturer struct {
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
	Active    bool      `json:"active" db:"active"`
	Brand     string    `json:"name" db:"brand"`
}
