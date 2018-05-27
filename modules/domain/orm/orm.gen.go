// Code generated build with models DO NOT EDIT.

package orm

import (
	"fmt"

	"github.com/clickyab/services/mysql"
	gorp "gopkg.in/gorp.v2"
)

// AUTO GENERATED CODE. DO NOT EDIT!

const (
	// DomainTableFull is the Domain table name
	DomainTableFull = "domains"

	// DomainUserTableFull is the DomainUser table name
	DomainUserTableFull = "users_domains"
)

func getSelectFields(tb string, alias string) string {
	if alias != "" {
		alias += "."
	}
	switch tb {

	case DomainTableFull:
		return fmt.Sprintf(`%[1]sid,%[1]sdomain_base,%[1]stitle,%[1]slogo,%[1]stheme,%[1]sdescription,%[1]sattributes,%[1]sstatus,%[1]smin_total_budget,%[1]smin_daily_budget,%[1]smin_web_native_cpc,%[1]smin_web_banner_cpc,%[1]smin_web_vast_cpc,%[1]smin_app_native_cpc,%[1]smin_app_banner_cpc,%[1]smin_app_vast_cpc,%[1]smin_web_cpc,%[1]smin_app_cpc,%[1]smin_web_native_cpm,%[1]smin_web_banner_cpm,%[1]smin_web_vast_cpm,%[1]smin_app_native_cpm,%[1]smin_app_banner_cpm,%[1]smin_app_vast_cpm,%[1]smin_web_cpm,%[1]smin_app_cpm,%[1]sadvantage,%[1]screated_at,%[1]supdated_at`, alias)

	case DomainUserTableFull:
		return fmt.Sprintf(`%[1]sdomain_id,%[1]sstatus,%[1]suser_id`, alias)

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
		Domain{},
		DomainTableFull,
	).SetKeys(
		true,
		"ID",
	)

	m.AddTableWithName(
		DomainUser{},
		DomainUserTableFull,
	).SetKeys(
		false,
		"UserID",
		"DomainID",
	)

}
func init() {
	mysql.Register(NewOrmManager())
}
