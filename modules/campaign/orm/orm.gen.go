// Code generated build with models DO NOT EDIT.

package orm

import (
	"fmt"

	"github.com/clickyab/services/mysql"
	gorp "gopkg.in/gorp.v2"
)

// AUTO GENERATED CODE. DO NOT EDIT!

const (
	// CampaignAttributesTableFull is the CampaignAttributes table name
	CampaignAttributesTableFull = "campaign_attributes"

	// CampaignDetailTableFull is the CampaignDetail table name
	CampaignDetailTableFull = "campaign_detail"

	// CampaignTableFull is the Campaign table name
	CampaignTableFull = "campaigns"

	// ScheduleTableFull is the Schedule table name
	ScheduleTableFull = "schedules"
)

func getSelectFields(tb string, alias string) string {
	if alias != "" {
		alias += "."
	}
	switch tb {

	case CampaignAttributesTableFull:
		return fmt.Sprintf(`%[1]scampaign_id,%[1]sdevice,%[1]smanufacturer,%[1]sos,%[1]sbrowser,%[1]siab,%[1]sregion,%[1]scellular,%[1]sisp`, alias)

	case CampaignDetailTableFull:
		return fmt.Sprintf(`%[1]scampaign_id,%[1]sdaily_id,%[1]sfake_imp,%[1]sfake_click,%[1]simp,%[1]sclick,%[1]scpc,%[1]sconv,%[1]screated_at,%[1]supdated_at`, alias)

	case CampaignTableFull:
		return fmt.Sprintf(`%[1]suser_id,%[1]sdomain_id,%[1]sexchange,%[1]swhite_black_id,%[1]swhite_black_type,%[1]swhite_black_value,%[1]sprogress,%[1]sAttributes`, alias)

	case ScheduleTableFull:
		return fmt.Sprintf(`%[1]sid,%[1]scampaign_id,%[1]supdated_at`, alias)

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
		CampaignAttributes{},
		CampaignAttributesTableFull,
	).SetKeys(
		false,
		"CampaignID",
	)

	m.AddTableWithName(
		CampaignDetail{},
		CampaignDetailTableFull,
	).SetKeys(
		false,
		"CampaignID",
		"DailyID",
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
