package orm

import "time"

// Region model in database
// @Model {
//		table = regions
//		primary = true, id
//		list = yes
// }
type Region struct {
	ID        string    `json:"id" db:"id"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
	Active    bool      `json:"active" db:"active"`
	Name      string    `json:"name" db:"name"`
}
