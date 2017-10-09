// Code generated build with models DO NOT EDIT.

package orm

import (
	"github.com/clickyab/services/mysql"
	gorp "gopkg.in/gorp.v2"
)

// AUTO GENERATED CODE. DO NOT EDIT!

const (
	// BrowserTableFull is the Browser table name
	BrowserTableFull = "browsers"

	// CategoryTableFull is the Category table name
	CategoryTableFull = "categories"

	// ISPTableFull is the ISP table name
	ISPTableFull = "isps"

	// ManufacturerTableFull is the Manufacturer table name
	ManufacturerTableFull = "manufacturers"

	// OSTableFull is the OS table name
	OSTableFull = "oses"

	// RegionTableFull is the Region table name
	RegionTableFull = "regions"
)

// Manager is the model manager for orm package
type Manager struct {
	mysql.Manager
}

// NewOrmManager create and return a manager for this module
func NewOrmManager() *Manager {
	return &Manager{}
}

// NewOrmManagerFromTransaction create and return a manager for this module from a transaction
func NewOrmManagerFromTransaction(tx gorp.SqlExecutor) (*Manager, error) {
	m := &Manager{}
	err := m.Hijack(tx)
	if err != nil {
		return nil, err
	}

	return m, nil
}

// Initialize orm package
func (m *Manager) Initialize() {

	m.AddTableWithName(
		Browser{},
		BrowserTableFull,
	).SetKeys(
		false,
		"Name",
	)

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
		Manufacturer{},
		ManufacturerTableFull,
	).SetKeys(
		false,
		"Brand",
	)

	m.AddTableWithName(
		OS{},
		OSTableFull,
	).SetKeys(
		true,
		"ID",
	)

	m.AddTableWithName(
		Region{},
		RegionTableFull,
	).SetKeys(
		true,
		"ID",
	)

}
func init() {
	mysql.Register(NewOrmManager())
}
