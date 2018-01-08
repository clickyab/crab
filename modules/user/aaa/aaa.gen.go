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
		return fmt.Sprintf(`%[1]s&#34;id&#34;,%[1]s&#34;user_id&#34;,%[1]s&#34;legal_name&#34;,%[1]s&#34;legal_register&#34;,%[1]s&#34;economic_code&#34;`, alias)

	case ParentUserTableFull:
		return fmt.Sprintf(`%[1]s&#34;user_id&#34;,%[1]s&#34;parent_id&#34;,%[1]s&#34;domain_id&#34;,%[1]s&#34;created_at&#34;`, alias)

	case RolePermissionTableFull:
		return fmt.Sprintf(`%[1]s&#34;id&#34;,%[1]s&#34;role_id&#34;,%[1]s&#34;scope&#34;,%[1]s&#34;perm&#34;,%[1]s&#34;created_at&#34;,%[1]s&#34;updated_at&#34;`, alias)

	case RoleUserTableFull:
		return fmt.Sprintf(`%[1]s&#34;user_id&#34;,%[1]s&#34;role_id&#34;,%[1]s&#34;created_at&#34;`, alias)

	case RoleTableFull:
		return fmt.Sprintf(`%[1]s&#34;id&#34;,%[1]s&#34;name&#34;,%[1]s&#34;description&#34;,%[1]s&#34;domain_id&#34;,%[1]s&#34;created_at&#34;,%[1]s&#34;updated_at&#34;`, alias)

	case UserTableFull:
		return fmt.Sprintf(`%[1]s&#34;id&#34;,%[1]s&#34;email&#34;,%[1]s&#34;password&#34;,%[1]s&#34;access_token&#34;,%[1]s&#34;avatar&#34;,%[1]s&#34;status&#34;,%[1]s&#34;created_at&#34;,%[1]s&#34;updated_at&#34;,%[1]s&#34;old_password&#34;,%[1]s&#34;city_id&#34;,%[1]s&#34;land_line&#34;,%[1]s&#34;cellphone&#34;,%[1]s&#34;postal_code&#34;,%[1]s&#34;first_name&#34;,%[1]s&#34;last_name&#34;,%[1]s&#34;address&#34;,%[1]s&#34;gender&#34;,%[1]s&#34;ssn&#34;,%[1]s&#34;Corporation&#34;`, alias)

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
