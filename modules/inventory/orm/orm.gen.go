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

	// InventoryPublisherTableFull is the InventoryPublisher table name
	InventoryPublisherTableFull = "inventories_publishers"

	// PublisherTableFull is the Publisher table name
	PublisherTableFull = "publishers"
)

func GetSelectFields(tb string, alias string) string {
	if alias != "" {
		alias += "."
	}
	switch tb {

	case InventoryTableFull:
		return fmt.Sprintf(`%[1]sid,%[1]screated_at,%[1]supdated_at,%[1]suser_id,%[1]sdomain_id,%[1]slabel,%[1]sstatus`, alias)

	case InventoryPublisherTableFull:
		return fmt.Sprintf(`%[1]spublisher_id,%[1]sinventory_id`, alias)

	case PublisherTableFull:
		return fmt.Sprintf(`%[1]sid,%[1]sname,%[1]sdomain,%[1]scategories,%[1]ssupplier,%[1]skind,%[1]sstatus,%[1]screated_at,%[1]supdated_at,%[1]sdeleted_at`, alias)

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
		true,
		"ID",
	)

	m.AddTableWithName(
		InventoryPublisher{},
		InventoryPublisherTableFull,
	).SetKeys(
		false,
		"PublisherID",
		"InventoryID",
	)

	m.AddTableWithName(
		Publisher{},
		PublisherTableFull,
	).SetKeys(
		true,
		"ID",
	)

}
func init() {
	mysql.Register(NewOrmManager())
}
