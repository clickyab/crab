// Code generated build with models DO NOT EDIT.

package asst

import (
	"github.com/clickyab/services/mysql"
	gorp "gopkg.in/gorp.v2"
)

// AUTO GENERATED CODE. DO NOT EDIT!

const (
	// CategoryTableFull is the Category table name
	CategoryTableFull = "categories"

	// ISPTableFull is the ISP table name
	ISPTableFull = "isps"

	// OSTableFull is the OS table name
	OSTableFull = "oses"
)

// Manager is the model manager for asst package
type Manager struct {
	mysql.Manager
}

// NewAsstManager create and return a manager for this module
func NewAsstManager() *Manager {
	return &Manager{}
}

// NewAsstManagerFromTransaction create and return a manager for this module from a transaction
func NewAsstManagerFromTransaction(tx gorp.SqlExecutor) (*Manager, error) {
	m := &Manager{}
	err := m.Hijack(tx)
	if err != nil {
		return nil, err
	}

	return m, nil
}

// Initialize asst package
func (m *Manager) Initialize() {

	m.AddTableWithName(
		Category{},
		CategoryTableFull,
	).SetKeys(
		true,
		"ID",
	)

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
	mysql.Register(NewAsstManager())
}
