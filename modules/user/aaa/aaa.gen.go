// Code generated build with models DO NOT EDIT.

package aaa

import (
	"fmt"

	"github.com/clickyab/services/mysql"
	gorp "gopkg.in/gorp.v2"
)

// AUTO GENERATED CODE. DO NOT EDIT!

const (
	// AdvisorTableFull is the Advisor table name
	AdvisorTableFull = "advisors"

	// AuditLogDetailTableFull is the AuditLogDetail table name
	AuditLogDetailTableFull = "audit_log_details"

	// AuditLogTableFull is the AuditLog table name
	AuditLogTableFull = "audit_logs"

	// CorporationTableFull is the Corporation table name
	CorporationTableFull = "corporations"

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

	case AdvisorTableFull:
		return fmt.Sprintf(`%[1]suser_id,%[1]sadvisor_id,%[1]sdomain_id,%[1]screated_at`, alias)

	case AuditLogDetailTableFull:
		return fmt.Sprintf(`%[1]sid,%[1]saudit_log_id,%[1]sdata,%[1]screated_at,%[1]supdated_at`, alias)

	case AuditLogTableFull:
		return fmt.Sprintf(`%[1]sid,%[1]sdomain_id,%[1]suser_id,%[1]suser_perm,%[1]sperm_scope,%[1]saction,%[1]starget_model,%[1]starget_id,%[1]sowner_id,%[1]simpersonate,%[1]simpersonator_id,%[1]sdescription,%[1]screated_at,%[1]supdated_at`, alias)

	case CorporationTableFull:
		return fmt.Sprintf(`%[1]suser_id,%[1]slegal_name,%[1]slegal_register,%[1]seconomic_code`, alias)

	case RolePermissionTableFull:
		return fmt.Sprintf(`%[1]sid,%[1]srole_id,%[1]sscope,%[1]sperm,%[1]screated_at,%[1]supdated_at`, alias)

	case RoleUserTableFull:
		return fmt.Sprintf(`%[1]suser_id,%[1]srole_id,%[1]screated_at`, alias)

	case RoleTableFull:
		return fmt.Sprintf(`%[1]sid,%[1]sname,%[1]sdescription,%[1]sdomain_id,%[1]screated_at,%[1]supdated_at`, alias)

	case UserTableFull:
		return fmt.Sprintf(`%[1]sid,%[1]semail,%[1]spassword,%[1]saccess_token,%[1]savatar,%[1]sstatus,%[1]sold_password,%[1]scity_id,%[1]sland_line,%[1]scellphone,%[1]spostal_code,%[1]sfirst_name,%[1]slast_name,%[1]saddress,%[1]sgender,%[1]sssn,%[1]sbalance,%[1]sattributes,%[1]sfinancial_code,%[1]sadvantage,%[1]screated_at,%[1]supdated_at`, alias)

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
		Advisor{},
		AdvisorTableFull,
	).SetKeys(
		false,
		"UserID",
		"AdvisorID",
		"DomainID",
	)

	m.AddTableWithName(
		AuditLogDetail{},
		AuditLogDetailTableFull,
	).SetKeys(
		true,
		"ID",
	)

	m.AddTableWithName(
		AuditLog{},
		AuditLogTableFull,
	).SetKeys(
		true,
		"ID",
	)

	m.AddTableWithName(
		Corporation{},
		CorporationTableFull,
	).SetKeys(
		false,
		"UserID",
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
