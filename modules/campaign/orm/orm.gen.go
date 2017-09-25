// Code generated build with models DO NOT EDIT.

package orm

import (
	"github.com/clickyab/services/mysql"
	gorp "gopkg.in/gorp.v2"
)

// AUTO GENERATED CODE. DO NOT EDIT!

const (
	// CampaignAttributesTableFull is the CampaignAttributes table name
	CampaignAttributesTableFull = "campaign_attributes"

	// CampaignTableFull is the Campaign table name
	CampaignTableFull = "campaigns"

	// ScheduleTableFull is the Schedule table name
	ScheduleTableFull = "schedules"
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
		CampaignAttributes{},
		CampaignAttributesTableFull,
	).SetKeys(
		false,
		"CampaignID",
	)

	m.AddTableWithName(
		Campaign{},
		CampaignTableFull,
	).SetKeys(
		true,
		"ID",
	)

	m.AddTableWithName(
		Schedule{},
		ScheduleTableFull,
	).SetKeys(
		true,
		"ID",
	)

}
func init() {
	mysql.Register(NewOrmManager())
}
