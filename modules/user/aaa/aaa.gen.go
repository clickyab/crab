// Code generated build with models DO NOT EDIT.

package aaa

import (
	"github.com/clickyab/services/mysql"
	gorp "gopkg.in/gorp.v2"
)

// AUTO GENERATED CODE. DO NOT EDIT!

const (
	// ConsularCustomerTableFull is the ConsularCustomer table name
	ConsularCustomerTableFull = "consular_customer"

	// RolePermissionTableFull is the RolePermission table name
	RolePermissionTableFull = "role_permission"

	// RoleUserTableFull is the RoleUser table name
	RoleUserTableFull = "role_user"

	// RoleTableFull is the Role table name
	RoleTableFull = "roles"

	// UserCorporationTableFull is the UserCorporation table name
	UserCorporationTableFull = "user_corporation"

	// UserPersonalTableFull is the UserPersonal table name
	UserPersonalTableFull = "user_personal"

	// UserTableFull is the User table name
	UserTableFull = "users"
)

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
		ConsularCustomer{},
		ConsularCustomerTableFull,
	).SetKeys(
		false,
		"ConsularyID",
		"CustomerID",
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
		UserCorporation{},
		UserCorporationTableFull,
	).SetKeys(
		false,
		"UserID",
	)

	m.AddTableWithName(
		UserPersonal{},
		UserPersonalTableFull,
	).SetKeys(
		false,
		"UserID",
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
