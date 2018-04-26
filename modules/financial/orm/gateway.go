package orm

import (
	"fmt"
	"time"
)

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
	IsDefault DefaultType   `json:"is_default" db:"is_default"`
	CreatedAt time.Time     `json:"created_at" db:"created_at"`
	UpdatedAt time.Time     `json:"updated_at" db:"updated_at"`
}

// FindActiveGatewayByName return the Gateway base on name
func (m *Manager) FindActiveGatewayByName(name string) (*Gateway, error) {
	var res Gateway
	err := m.GetRDbMap().SelectOne(
		&res,
		fmt.Sprintf("SELECT %s FROM %s WHERE name=? AND status=?", getSelectFields(GatewayTableFull, ""), GatewayTableFull),
		name,
		Enabled,
	)

	if err != nil {
		return nil, err
	}

	return &res, nil
}
