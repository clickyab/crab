package orm

import (
	"fmt"
	"time"

	"github.com/clickyab/services/assert"
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
	_, err := m.GetWDbMap().Exec(
		fmt.Sprintf("delete from %s "+
			"where campaign_id=?",
			CampaignReportReceiversTableFull,
		),
		caID,
	)

	return err
}

// Receiver for campaign reports receivers
type Receiver struct {
	ID    int64  `json:"id" db:"id" `
	Email string `json:"email" db:"email"`
}

// GetReportReceivers to delete all campaign reports receivers
func (m *Manager) GetReportReceivers(caID int64) []Receiver {
	r := make([]Receiver, 0)
	_, err := m.GetRDbMap().Select(&r, `select a.id, a.email from users as a join  campaign_report_receivers as b
on b.user_id=a.id  where b.campaign_id=?`, caID)
	assert.Nil(err)
	return r
}

// UpdateReportReceivers remove campaign report receivers and add new
func (m *Manager) UpdateReportReceivers(rec []int64, caID int64) error {
	err := m.DeleteAllCampaignReportReceivers(caID)
	if err != nil {
		return err
	}

	return m.AddReceivers(rec, caID)
}
