package orm

import "time"

// OS os model in database
// @Model {
//		table = oses
//		primary = true, id
//		find_by = id,name
//		list = yes
// }
type OS struct {
	ID        int64     `json:"id" db:"id"`
	Name      string    `json:"name" db:"name"`
	Active    bool      `json:"active" db:"active"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}
