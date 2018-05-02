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

	// CampaignReportReceiversTableFull is the CampaignReportReceivers table name
	CampaignReportReceiversTableFull = "campaign_report_receivers"

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
		return fmt.Sprintf(`%[1]scampaign_id,%[1]sdaily_id,%[1]shour_id,%[1]sfake_imp,%[1]sfake_click,%[1]simp,%[1]sclick,%[1]scpc,%[1]scpa,%[1]scpm,%[1]sconv,%[1]spublisher_id,%[1]screated_at,%[1]supdated_at`, alias)

	case CampaignReportReceiversTableFull:
		return fmt.Sprintf(`%[1]scampaign_id,%[1]suser_id,%[1]screated_at`, alias)

	case CampaignTableFull:
		return fmt.Sprintf(`%[1]sid,%[1]suser_id,%[1]sdomain_id,%[1]stitle,%[1]skind,%[1]sstatus,%[1]sprogress,%[1]sstart_at,%[1]send_at,%[1]stotal_budget,%[1]sdaily_budget,%[1]sstrategy,%[1]smax_bid,%[1]sexchange,%[1]sinventory_id,%[1]sinventory_type,%[1]sinventory_domains,%[1]stld,%[1]stoday_spend,%[1]stotal_spend,%[1]screated_at,%[1]supdated_at,%[1]sarchived_at`, alias)

	case ScheduleTableFull:
		return fmt.Sprintf(`%[1]sid,%[1]scampaign_id,%[1]supdated_at,h00,h01,h02,h03,h04,h05,h06,h07,h08,h09,h10,h11,h12,h13,h14,h15,h16,h17,h18,h19,h20,h21,h22,h23`, alias)

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
		"HourID",
		"PublisherID",
	)

	m.AddTableWithName(
		CampaignReportReceivers{},
		CampaignReportReceiversTableFull,
	).SetKeys(
		false,
		"CampaignID",
		"UserID",
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
