// Code generated build with models DO NOT EDIT.

package orm

import (
	"fmt"

	"github.com/clickyab/services/mysql"
	gorp "gopkg.in/gorp.v2"
)

// AUTO GENERATED CODE. DO NOT EDIT!

const (
	// InfoTableFull is the Info table name
	InfoTableFull = "-"
)

func getSelectFields(tb string, alias string) string {
	if alias != "" {
		alias += "."
	}
	switch tb {

	case InfoTableFull:
		return fmt.Sprintf(`%[1]sdata1,%[1]sdata2,%[1]sdata3,%[1]sdata4,%[1]skeyfield1,%[1]skeyfield2,%[1]skeyfield3,%[1]skeyfield4`, alias)

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
		Info{},
		InfoTableFull,
	).SetKeys(
		false,
		"KeyField1",
	)

}
func init() {
	mysql.Register(NewOrmManager())
}
