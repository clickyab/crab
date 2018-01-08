// Code generated build with models DO NOT EDIT.

package orm

import (
	"fmt"

	"github.com/clickyab/services/mysql"
	gorp "gopkg.in/gorp.v2"
)

// AUTO GENERATED CODE. DO NOT EDIT!

const (
	// InventoryTableFull is the Inventory table name
	InventoryTableFull = "inventories"

	// WhiteBlackListTableFull is the WhiteBlackList table name
	WhiteBlackListTableFull = "user_wlbl_presets"
)

func getSelectFields(tb string, alias string) string {
	if alias != "" {
		alias += "."
	}
	switch tb {

	case InventoryTableFull:
		return fmt.Sprintf(`%[1]s&#34;id&#34;,%[1]s&#34;created_at&#34;,%[1]s&#34;updated_at&#34;,%[1]s&#34;active&#34;,%[1]s&#34;name&#34;,%[1]s&#34;domain&#34;,%[1]s&#34;cat&#34;,%[1]s&#34;publisher&#34;,%[1]s&#34;kind&#34;,%[1]s&#34;status&#34;`, alias)

	case WhiteBlackListTableFull:
		return fmt.Sprintf(`%[1]s&#34;id&#34;,%[1]s&#34;created_at&#34;,%[1]s&#34;updated_at&#34;,%[1]s&#34;active&#34;,%[1]s&#34;user_id&#34;,%[1]s&#34;domain_id&#34;,%[1]s&#34;label&#34;,%[1]s&#34;domains&#34;,%[1]s&#34;publisher_type&#34;`, alias)

	}
	return ""
}

// Manager is the model manager for orm package
type Manager struct {
	mysql.Manager
}

// NewOrmManager create and return a manager for this module
func NewOrmManager() *Manager {
	return &Manager{}
}

// NewOrmManagerFromTransaction create and return a manager for this module from a transaction
func NewOrmManagerFromTransaction(tx gorp.SqlExecutor) (*Manager, error) {
	m := &Manager{}
	err := m.Hijack(tx)
	if err != nil {
		return nil, err
	}

	return m, nil
}

// Initialize orm package
func (m *Manager) Initialize() {

	m.AddTableWithName(
		Inventory{},
		InventoryTableFull,
	).SetKeys(
		false,
		"ID",
	)

	m.AddTableWithName(
		WhiteBlackList{},
		WhiteBlackListTableFull,
	).SetKeys(
		true,
		"ID",
	)

}
func init() {
	mysql.Register(NewOrmManager())
}
