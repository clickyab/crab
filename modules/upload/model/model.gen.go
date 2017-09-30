// Code generated build with models DO NOT EDIT.

package model

import (
	"github.com/clickyab/services/mysql"
	gorp "gopkg.in/gorp.v2"
)

// AUTO GENERATED CODE. DO NOT EDIT!

const (
	// UploadTableFull is the Upload table name
	UploadTableFull = "uploads"
)

// Manager is the model manager for model package
type Manager struct {
	mysql.Manager
}

// NewModelManager create and return a manager for this module
func NewModelManager() *Manager {
	return &Manager{}
}

// NewModelManagerFromTransaction create and return a manager for this module from a transaction
func NewModelManagerFromTransaction(tx gorp.SqlExecutor) (*Manager, error) {
	m := &Manager{}
	err := m.Hijack(tx)
	if err != nil {
		return nil, err
	}

	return m, nil
}

// Initialize model package
func (m *Manager) Initialize() {

	m.AddTableWithName(
		Upload{},
		UploadTableFull,
	).SetKeys(
		false,
		"ID",
	)

}
func init() {
	mysql.Register(NewModelManager())
}
