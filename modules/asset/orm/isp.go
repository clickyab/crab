package orm

import "time"

// ISP isp model in database
// @Model {
//		table = isps
//		primary = true, id
//		find_by = id,name
//		list = yes
// }
type ISP struct {
	ID        int64     `json:"id" db:"id"`
	Name      string    `json:"name" db:"name"`
	Active    bool      `json:"active" db:"active"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}
