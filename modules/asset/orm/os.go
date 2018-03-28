package orm

import "time"

// OS os model in database
// @Model {
//		table = oses
//		primary = false, name
//		find_by = name
//		list = yes
// }
type OS struct {
	Name      string      `json:"name" db:"name"`
	Status    AssetStatus `json:"status" db:"status"`
	CreatedAt time.Time   `json:"created_at" db:"created_at"`
	UpdatedAt time.Time   `json:"updated_at" db:"updated_at"`
}
