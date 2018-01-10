// Code generated build with models DO NOT EDIT.

package dmn

import (
	"fmt"

	"github.com/clickyab/services/mysql"
	gorp "gopkg.in/gorp.v2"
)

// AUTO GENERATED CODE. DO NOT EDIT!

const (
	// DomainUserTableFull is the DomainUser table name
	DomainUserTableFull = "domain_user"

	// DomainTableFull is the Domain table name
	DomainTableFull = "domains"
)

func getSelectFields(tb string, alias string) string {
	if alias != "" {
		alias += "."
	}
	switch tb {

	case DomainUserTableFull:
		return fmt.Sprintf(`%[1]sdomain_id,%[1]suser_id`, alias)

	case DomainTableFull:
		return fmt.Sprintf(`%[1]sid,%[1]sname,%[1]sdescription,%[1]sactive,%[1]screated_at,%[1]supdated_at`, alias)

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
		DomainUser{},
		DomainUserTableFull,
	).SetKeys(
		false,
		"UserID",
		"DomainID",
	)

	m.AddTableWithName(
		Domain{},
		DomainTableFull,
	).SetKeys(
		true,
		"ID",
	)

}
func init() {
	mysql.Register(NewDmnManager())
}
