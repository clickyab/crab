package orm

import (
	"time"

	"github.com/clickyab/services/initializer"
)

// CampaignReportReceivers campaign detail model in database
// @Model {
//		table = campaign_report_receivers
//		primary = false, campaign_id,user_id
//		find_by = campaign_id,user_id
// }
type CampaignReportReceivers struct {
	CampaignID int64     `json:"campaign_id" db:"campaign_id"`
	UserID     int64     `json:"user_id" db:"user_id"`
	CreatedAt  time.Time `json:"created_at" db:"created_at"`
}

// AddReceivers add campaign report receivers
func (m *Manager) AddReceivers(rec []int64, caID int64) error {
	s := make([]interface{}, len(rec))
	for i, uID := range rec {
		s[i] = &CampaignReportReceivers{
			CampaignID: caID,
			UserID:     uID,
			CreatedAt:  time.Now(),
		}
	}

	func(in interface{}) {
		if ii, ok := in.(initializer.Simple); ok {
			ii.Initialize()
		}
	}(s)

	return m.GetWDbMap().Insert(s...)
}

// DeleteAllCampaignReportReceivers to delete all campaign reports receivers
func (m *Manager) DeleteAllCampaignReportReceivers(caID int64) error {
	_, err := m.GetWDbMap().Exec("delete from campaign_report_receivers where campaign_id=?", caID)

	return err
}

// UpdateReportReceivers remove campaign report receivers and add new
func (m *Manager) UpdateReportReceivers(rec []int64, caID int64) error {
	err := m.DeleteAllCampaignReportReceivers(caID)
	if err != nil {
		return err
	}

	return m.AddReceivers(rec, caID)
}
