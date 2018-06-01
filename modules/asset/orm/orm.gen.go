// Code generated build with models DO NOT EDIT.

package orm

import (
	"fmt"

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

	// PlatformTableFull is the Platform table name
	PlatformTableFull = "platforms"
)

func GetSelectFields(tb string, alias string) string {
	if alias != "" {
		alias += "."
	}
	switch tb {

	case BrowserTableFull:
		return fmt.Sprintf(`%[1]screated_at,%[1]supdated_at,%[1]sdeleted_at,%[1]sname`, alias)

	case CategoryTableFull:
		return fmt.Sprintf(`%[1]sname,%[1]sdescription,%[1]sdeleted_at`, alias)

	case ISPTableFull:
		return fmt.Sprintf(`%[1]sname,%[1]skind,%[1]sstatus,%[1]screated_at,%[1]supdated_at`, alias)

	case ManufacturerTableFull:
		return fmt.Sprintf(`%[1]screated_at,%[1]supdated_at,%[1]sstatus,%[1]sname`, alias)

	case OSTableFull:
		return fmt.Sprintf(`%[1]sname,%[1]sstatus,%[1]screated_at,%[1]supdated_at`, alias)

	case PlatformTableFull:
		return fmt.Sprintf(`%[1]sname,%[1]sstatus,%[1]screated_at,%[1]supdated_at`, alias)

	}
	return ""
}

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
		false,
		"Name",
	)

	m.AddTableWithName(
		ISP{},
		ISPTableFull,
	).SetKeys(
		false,
		"Name",
	)

	m.AddTableWithName(
		Manufacturer{},
		ManufacturerTableFull,
	).SetKeys(
		false,
		"Name",
	)

	m.AddTableWithName(
		OS{},
		OSTableFull,
	).SetKeys(
		false,
		"Name",
	)

	m.AddTableWithName(
		Platform{},
		PlatformTableFull,
	).SetKeys(
		false,
		"Name",
	)

}
func init() {
	mysql.Register(NewOrmManager())
}
