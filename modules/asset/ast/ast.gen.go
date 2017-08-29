// Code generated build with models DO NOT EDIT.

package ast

import (
	"github.com/clickyab/services/mysql"
	gorp "gopkg.in/gorp.v2"
)

// AUTO GENERATED CODE. DO NOT EDIT!

const (
	// ISPTableFull is the ISP table name
	ISPTableFull = "isps"

	// OSTableFull is the OS table name
	OSTableFull = "oses"
)

// Manager is the model manager for ast package
type Manager struct {
	mysql.Manager
}

// NewAstManager create and return a manager for this module
func NewAstManager() *Manager {
	return &Manager{}
}

// NewAstManagerFromTransaction create and return a manager for this module from a transaction
func NewAstManagerFromTransaction(tx gorp.SqlExecutor) (*Manager, error) {
	m := &Manager{}
	err := m.Hijack(tx)
	if err != nil {
		return nil, err
	}

	return m, nil
}

// Initialize ast package
func (m *Manager) Initialize() {

	m.AddTableWithName(
		ISP{},
		ISPTableFull,
	).SetKeys(
		true,
		"ID",
	)

	m.AddTableWithName(
		OS{},
		OSTableFull,
	).SetKeys(
		true,
		"ID",
	)

}
func init() {
	mysql.Register(NewAstManager())
}
