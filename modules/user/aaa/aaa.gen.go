// Code generated build with models DO NOT EDIT.

package aaa

import (
	"fmt"

	"github.com/clickyab/services/mysql"
	gorp "gopkg.in/gorp.v2"
)

// AUTO GENERATED CODE. DO NOT EDIT!

const (
	// CorporationTableFull is the Corporation table name
	CorporationTableFull = "corporations"

	// ParentUserTableFull is the ParentUser table name
	ParentUserTableFull = "parent_user"

	// RolePermissionTableFull is the RolePermission table name
	RolePermissionTableFull = "role_permission"

	// RoleUserTableFull is the RoleUser table name
	RoleUserTableFull = "role_user"

	// RoleTableFull is the Role table name
	RoleTableFull = "roles"

	// UserTableFull is the User table name
	UserTableFull = "users"
)

func getSelectFields(tb string, alias string) string {
	if alias != "" {
		alias += "."
	}
	switch tb {

	case CorporationTableFull:
		return fmt.Sprintf(`%[1]sid,%[1]suser_id,%[1]slegal_name,%[1]slegal_register,%[1]seconomic_code`, alias)

	case ParentUserTableFull:
		return fmt.Sprintf(`%[1]suser_id,%[1]sparent_id,%[1]sdomain_id,%[1]screated_at`, alias)

	case RolePermissionTableFull:
		return fmt.Sprintf(`%[1]sid,%[1]srole_id,%[1]sscope,%[1]sperm,%[1]screated_at,%[1]supdated_at`, alias)

	case RoleUserTableFull:
		return fmt.Sprintf(`%[1]suser_id,%[1]srole_id,%[1]screated_at`, alias)

	case RoleTableFull:
		return fmt.Sprintf(`%[1]sid,%[1]sname,%[1]sdescription,%[1]sdomain_id,%[1]screated_at,%[1]supdated_at`, alias)

	case UserTableFull:
		return fmt.Sprintf(`%[1]sid,%[1]semail,%[1]spassword,%[1]saccess_token,%[1]savatar,%[1]sstatus,%[1]screated_at,%[1]supdated_at,%[1]sold_password,%[1]scity_id,%[1]sland_line,%[1]scellphone,%[1]spostal_code,%[1]sfirst_name,%[1]slast_name,%[1]saddress,%[1]sgender,%[1]sssn,%[1]sCorporation`, alias)

	}
	return ""
}

// Manager is the model manager for aaa package
type Manager struct {
	mysql.Manager
}

// NewAaaManager create and return a manager for this module
func NewAaaManager() *Manager {
	return &Manager{}
}

// NewAaaManagerFromTransaction create and return a manager for this module from a transaction
func NewAaaManagerFromTransaction(tx gorp.SqlExecutor) (*Manager, error) {
	m := &Manager{}
	err := m.Hijack(tx)
	if err != nil {
		return nil, err
	}

	return m, nil
}

// Initialize aaa package
func (m *Manager) Initialize() {

	m.AddTableWithName(
		Corporation{},
		CorporationTableFull,
	).SetKeys(
		true,
		"ID",
	)

	m.AddTableWithName(
		ParentUser{},
		ParentUserTableFull,
	).SetKeys(
		false,
		"UserID",
		"ParentID",
		"DomainID",
	)

	m.AddTableWithName(
		RolePermission{},
		RolePermissionTableFull,
	).SetKeys(
		true,
		"ID",
	)

	m.AddTableWithName(
		RoleUser{},
		RoleUserTableFull,
	).SetKeys(
		false,
		"UserID",
		"RoleID",
	)

	m.AddTableWithName(
		Role{},
		RoleTableFull,
	).SetKeys(
		true,
		"ID",
	)

	m.AddTableWithName(
		User{},
		UserTableFull,
	).SetKeys(
		true,
		"ID",
	)

}
func init() {
	mysql.Register(NewAaaManager())
}
