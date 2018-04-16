package orm

import "time"

// GatewayStatus status
// @Enum{
// }
type GatewayStatus string

const (
	// Disabled gateway
	Disabled GatewayStatus = "disable"
	// Enabled gateway
	Enabled GatewayStatus = "enable"
)

// DefaultType status
// @Enum{
// }
type DefaultType string

const (
	// IsDefault gateway
	IsDefault DefaultType = "yes"
	// IsNotDefault gateway
	IsNotDefault DefaultType = "no"
)

// Gateway for payments
// @Model {
//		table = gateways
//		primary = true, id
//		find_by = id
//		list = yes
// }
type Gateway struct {
	ID        int64         `json:"id" db:"id"`
	Name      string        `json:"name" db:"name"`
	Status    GatewayStatus `json:"status" db:"status"`
	Default   DefaultType   `json:"default" db:"default"`
	CreatedAt time.Time     `json:"created_at" db:"created_at"`
	UpdatedAt time.Time     `json:"created_at" db:"created_at"`
}
