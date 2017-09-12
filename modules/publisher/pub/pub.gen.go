// Code generated build with models DO NOT EDIT.

package pub

import (
	"github.com/clickyab/services/mysql"
	gorp "gopkg.in/gorp.v2"
)

// AUTO GENERATED CODE. DO NOT EDIT!

const (
	// PublisherTableFull is the Publisher table name
	PublisherTableFull = "publishers"
)

// Manager is the model manager for pub package
type Manager struct {
	mysql.Manager
}

// NewPubManager create and return a manager for this module
func NewPubManager() *Manager {
	return &Manager{}
}

// NewPubManagerFromTransaction create and return a manager for this module from a transaction
func NewPubManagerFromTransaction(tx gorp.SqlExecutor) (*Manager, error) {
	m := &Manager{}
	err := m.Hijack(tx)
	if err != nil {
		return nil, err
	}

	return m, nil
}

// Initialize pub package
func (m *Manager) Initialize() {

	m.AddTableWithName(
		Publisher{},
		PublisherTableFull,
	).SetKeys(
		true,
		"ID",
	)

}
func init() {
	mysql.Register(NewPubManager())
}
