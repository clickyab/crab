// Code generated build with models DO NOT EDIT.

package dmn

import (
	"fmt"

	"github.com/clickyab/services/mysql"
	gorp "gopkg.in/gorp.v2"
)

// AUTO GENERATED CODE. DO NOT EDIT!

const (
	// DomainTableFull is the Domain table name
	DomainTableFull = "domains"

	// DomainUserTableFull is the DomainUser table name
	DomainUserTableFull = "users_domains"
)

func getSelectFields(tb string, alias string) string {
	if alias != "" {
		alias += "."
	}
	switch tb {

	case DomainTableFull:
		return fmt.Sprintf(`%[1]sid,%[1]sname,%[1]sdescription,%[1]sattributes,%[1]sstatus,%[1]screated_at,%[1]supdated_at`, alias)

	case DomainUserTableFull:
		return fmt.Sprintf(`%[1]sdomain_id,%[1]sstatus,%[1]suser_id`, alias)

	}
	return ""
}

// Manager is the model manager for dmn package
type Manager struct {
	mysql.Manager
}

// NewDmnManager create and return a manager for this module
func NewDmnManager() *Manager {
	return &Manager{}
}

// NewDmnManagerFromTransaction create and return a manager for this module from a transaction
func NewDmnManagerFromTransaction(tx gorp.SqlExecutor) (*Manager, error) {
	m := &Manager{}
	err := m.Hijack(tx)
	if err != nil {
		return nil, err
	}

	return m, nil
}

// Initialize dmn package
func (m *Manager) Initialize() {

	m.AddTableWithName(
		Domain{},
		DomainTableFull,
	).SetKeys(
		true,
		"ID",
	)

	m.AddTableWithName(
		DomainUser{},
		DomainUserTableFull,
	).SetKeys(
		false,
		"UserID",
		"DomainID",
	)

}
func init() {
	mysql.Register(NewDmnManager())
}
