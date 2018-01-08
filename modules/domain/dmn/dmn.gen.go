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
		return fmt.Sprintf(`%[1]s&#34;domain_id&#34;,%[1]s&#34;user_id&#34;`, alias)

	case DomainTableFull:
		return fmt.Sprintf(`%[1]s&#34;id&#34;,%[1]s&#34;name&#34;,%[1]s&#34;description&#34;,%[1]s&#34;active&#34;,%[1]s&#34;created_at&#34;,%[1]s&#34;updated_at&#34;`, alias)

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
