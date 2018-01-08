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

	// RegionTableFull is the Region table name
	RegionTableFull = "regions"
)

func getSelectFields(tb string, alias string) string {
	if alias != "" {
		alias += "."
	}
	switch tb {

	case BrowserTableFull:
		return fmt.Sprintf(`%[1]s&#34;created_at&#34;,%[1]s&#34;updated_at&#34;,%[1]s&#34;active&#34;,%[1]s&#34;name&#34;`, alias)

	case CategoryTableFull:
		return fmt.Sprintf(`%[1]s&#34;name&#34;,%[1]s&#34;description&#34;,%[1]s&#34;active&#34;`, alias)

	case ISPTableFull:
		return fmt.Sprintf(`%[1]s&#34;name&#34;,%[1]s&#34;kind&#34;,%[1]s&#34;active&#34;,%[1]s&#34;created_at&#34;,%[1]s&#34;updated_at&#34;`, alias)

	case ManufacturerTableFull:
		return fmt.Sprintf(`%[1]s&#34;created_at&#34;,%[1]s&#34;updated_at&#34;,%[1]s&#34;active&#34;,%[1]s&#34;name&#34;`, alias)

	case OSTableFull:
		return fmt.Sprintf(`%[1]s&#34;name&#34;,%[1]s&#34;active&#34;,%[1]s&#34;created_at&#34;,%[1]s&#34;updated_at&#34;`, alias)

	case PlatformTableFull:
		return fmt.Sprintf(`%[1]s&#34;name&#34;,%[1]s&#34;active&#34;,%[1]s&#34;created_at&#34;,%[1]s&#34;updated_at&#34;`, alias)

	case RegionTableFull:
		return fmt.Sprintf(`%[1]s&#34;id&#34;,%[1]s&#34;created_at&#34;,%[1]s&#34;updated_at&#34;,%[1]s&#34;active&#34;,%[1]s&#34;name&#34;`, alias)

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
