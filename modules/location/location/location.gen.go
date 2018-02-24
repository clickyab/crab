// Code generated build with models DO NOT EDIT.

package location

import (
	"fmt"

	"github.com/clickyab/services/mysql"
	gorp "gopkg.in/gorp.v2"
)

// AUTO GENERATED CODE. DO NOT EDIT!

const (
	// CityTableFull is the City table name
	CityTableFull = "cities"

	// CountryTableFull is the Country table name
	CountryTableFull = "countries"

	// ProvinceTableFull is the Province table name
	ProvinceTableFull = "provinces"
)

func getSelectFields(tb string, alias string) string {
	if alias != "" {
		alias += "."
	}
	switch tb {

	case CityTableFull:
		return fmt.Sprintf(`%[1]sid,%[1]sname,%[1]sprovince_id`, alias)

	case CountryTableFull:
		return fmt.Sprintf(`%[1]sid,%[1]sname`, alias)

	case ProvinceTableFull:
		return fmt.Sprintf(`%[1]sid,%[1]sname,%[1]scountry_id`, alias)

	}
	return ""
}

// Manager is the model manager for location package
type Manager struct {
	mysql.Manager
}

// NewLocationManager create and return a manager for this module
func NewLocationManager() *Manager {
	return &Manager{}
}

// NewLocationManagerFromTransaction create and return a manager for this module from a transaction
func NewLocationManagerFromTransaction(tx gorp.SqlExecutor) (*Manager, error) {
	m := &Manager{}
	err := m.Hijack(tx)
	if err != nil {
		return nil, err
	}

	return m, nil
}

// Initialize location package
func (m *Manager) Initialize() {

	m.AddTableWithName(
		City{},
		CityTableFull,
	).SetKeys(
		true,
		"ID",
	)

	m.AddTableWithName(
		Country{},
		CountryTableFull,
	).SetKeys(
		true,
		"ID",
	)

	m.AddTableWithName(
		Province{},
		ProvinceTableFull,
	).SetKeys(
		true,
		"ID",
	)

}
func init() {
	mysql.Register(NewLocationManager())
}
