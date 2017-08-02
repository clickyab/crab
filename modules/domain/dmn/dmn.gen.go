// Code generated build with models DO NOT EDIT.

package dmn

import (
	"github.com/clickyab/services/mysql"
	gorp "gopkg.in/gorp.v2"
)

// AUTO GENERATED CODE. DO NOT EDIT!

const (
	// DomainTableFull is the Domain table name
	DomainTableFull = "domains"
)

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

}
func init() {
	mysql.Register(NewDmnManager())
}
