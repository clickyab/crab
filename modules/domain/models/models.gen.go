// Code generated build with models DO NOT EDIT.

package models

import (
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

// Manager is the model manager for models package
type Manager struct {
	mysql.Manager
}

// NewModelsManager create and return a manager for this module
func NewModelsManager() *Manager {
	return &Manager{}
}

// NewModelsManagerFromTransaction create and return a manager for this module from a transaction
func NewModelsManagerFromTransaction(tx gorp.SqlExecutor) (*Manager, error) {
	m := &Manager{}
	err := m.Hijack(tx)
	if err != nil {
		return nil, err
	}

	return m, nil
}

// Initialize models package
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
	mysql.Register(NewModelsManager())
}
