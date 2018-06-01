package orm

import (
	"fmt"
	"time"

	"clickyab.com/crab/modules/financial/errors"
	"github.com/clickyab/services/assert"
	"github.com/clickyab/services/mysql"
	gom "github.com/go-sql-driver/mysql"
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
	ID        int64          `json:"id" db:"id"`
	Name      string         `json:"name" db:"name"`
	Status    GatewayStatus  `json:"status" db:"status"`
	IsDefault DefaultType    `json:"is_default" db:"is_default"`
	CreatedAt time.Time      `json:"created_at" db:"created_at"`
	UpdatedAt mysql.NullTime `json:"updated_at" db:"updated_at"`
}

// FindActiveGatewayByName return the Gateway base on name
func (m *Manager) FindActiveGatewayByName(name string) (*Gateway, error) {
	var res Gateway
	err := m.GetRDbMap().SelectOne(
		&res,
		fmt.Sprintf("SELECT %s FROM %s WHERE name=? AND status=?", GetSelectFields(GatewayTableFull, ""), GatewayTableFull),
		name,
		Enabled,
	)

	if err != nil {
		return nil, err
	}

	return &res, nil
}

// notDefaultBut set all gateways to not default but on with id
func (m *Manager) notDefaultBut(gatewayID int64) (int64, error) {
	q := fmt.Sprintf("UPDATE %s SET is_default=? WHERE id!=?", GatewayTableFull)
	res, err := m.GetWDbMap().Exec(q, IsNotDefault, gatewayID)
	if err != nil {
		return 0, err
	}
	rowEffected, err := res.RowsAffected()
	return rowEffected, err
}

// JustDefault set a gateway to default and others to not default
func (m *Manager) JustDefault(gateway *Gateway) error {
	err := m.Begin()
	assert.Nil(err)
	defer func() {
		if err != nil {
			assert.Nil(m.Rollback())
		} else {
			assert.Nil(m.Commit())
		}
	}()
	gateway.IsDefault = IsDefault
	err = m.UpdateGateway(gateway)
	if err != nil {
		mysqlError, ok := err.(*gom.MySQLError)
		if !ok {
			return errors.EditGatewayErr
		}
		if mysqlError.Number == 1062 {
			return errors.GatewayAlreadyExistErr
		}
	}
	_, err = m.notDefaultBut(gateway.ID)
	return err
}
