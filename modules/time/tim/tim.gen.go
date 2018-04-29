// Code generated build with models DO NOT EDIT.

package tim

import (
	"fmt"

	"github.com/clickyab/services/mysql"
	gorp "gopkg.in/gorp.v2"
)

// AUTO GENERATED CODE. DO NOT EDIT!

const (
	// DateTableFull is the Date table name
	DateTableFull = "date_table"

	// HourTableFull is the Hour table name
	HourTableFull = "hour_table"
)

func getSelectFields(tb string, alias string) string {
	if alias != "" {
		alias += "."
	}
	switch tb {

	case DateTableFull:
		return fmt.Sprintf(`%[1]sid,%[1]syear,%[1]smonth,%[1]sday_int_64,%[1]sj_year,%[1]sj_month,%[1]sj_day,%[1]sextra,%[1]sbasis`, alias)

	case HourTableFull:
		return fmt.Sprintf(`%[1]sid,%[1]syear,%[1]smonth,%[1]sday_int_64,%[1]shour,%[1]sj_year,%[1]sj_month,%[1]sj_day,%[1]sdate_id,%[1]sextra,%[1]sbasis`, alias)

	}
	return ""
}

// Manager is the model manager for tim package
type Manager struct {
	mysql.Manager
}

// NewTimManager create and return a manager for this module
func NewTimManager() *Manager {
	return &Manager{}
}

// NewTimManagerFromTransaction create and return a manager for this module from a transaction
func NewTimManagerFromTransaction(tx gorp.SqlExecutor) (*Manager, error) {
	m := &Manager{}
	err := m.Hijack(tx)
	if err != nil {
		return nil, err
	}

	return m, nil
}

// Initialize tim package
func (m *Manager) Initialize() {

	m.AddTableWithName(
		Date{},
		DateTableFull,
	).SetKeys(
		true,
		"ID",
	)

	m.AddTableWithName(
		Hour{},
		HourTableFull,
	).SetKeys(
		true,
		"ID",
	)

}
func init() {
	mysql.Register(NewTimManager())
}
