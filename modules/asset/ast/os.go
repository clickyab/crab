package ast

import "time"

type (
	// ActiveStatus is the yes/no status
	// @Enum{
	// }
	ActiveStatus string
)

const (
	// ActiveStatusYes active
	ActiveStatusYes ActiveStatus = "yes"
	// ActiveStatusNo not active
	ActiveStatusNo ActiveStatus = "no"
)

// OS os model in database
// @Model {
//		table = oses
//		primary = true, id
//		find_by = id,name
// }
type OS struct {
	ID        int64        `json:"id" db:"id"`
	Name      string       `json:"name" db:"name"`
	Status    ActiveStatus `json:"status" db:"status"`
	CreatedAt time.Time    `json:"created_at" db:"created_at"`
	UpdatedAt time.Time    `json:"updated_at" db:"updated_at"`
}
