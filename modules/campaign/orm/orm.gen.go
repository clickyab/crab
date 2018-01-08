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
		return fmt.Sprintf(`%[1]s&#34;campaign_id&#34;,%[1]s&#34;device&#34;,%[1]s&#34;manufacturer&#34;,%[1]s&#34;os&#34;,%[1]s&#34;browser&#34;,%[1]s&#34;iab&#34;,%[1]s&#34;region&#34;,%[1]s&#34;cellular&#34;,%[1]s&#34;isp&#34;`, alias)

	case CampaignTableFull:
		return fmt.Sprintf(`%[1]s&#34;user_id&#34;,%[1]s&#34;domain_id&#34;,%[1]s&#34;exchange&#34;,%[1]s&#34;white_black_id&#34;,%[1]s&#34;white_black_type&#34;,%[1]s&#34;white_black_value&#34;,%[1]s&#34;progress&#34;,%[1]s&#34;Attributes&#34;`, alias)

	case ScheduleTableFull:
		return fmt.Sprintf(`%[1]s&#34;id&#34;,%[1]s&#34;campaign_id&#34;,%[1]s&#34;updated_at&#34;`, alias)

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
