package orm

import (
	"database/sql"

	"github.com/clickyab/services/assert"
	"github.com/clickyab/services/mysql"
)

// CampaignAttributes model in database
// @Model {
//		table = campaign_attributes
//		primary = false, campaign_id
//		find_by = campaign_id
//		list = yes
// 		belong = Campaign:campaign_id
// }
type CampaignAttributes struct {
	CampaignID   int64                 `json:"-" db:"campaign_id"`
	Device       mysql.StringJSONArray `json:"device" db:"device"`
	Manufacturer mysql.StringJSONArray `json:"manufacturer" db:"manufacturer"`
	OS           mysql.StringJSONArray `json:"os" db:"os"`
	Browser      mysql.StringJSONArray `json:"browser" db:"browser"`
	IAB          mysql.StringJSONArray `json:"iab" db:"iab"`
	Region       mysql.StringJSONArray `json:"region" db:"region"`
	Cellular     mysql.StringJSONArray `json:"cellular" db:"cellular"`
	ISP          mysql.StringJSONArray `json:"isp" db:"isp"`
}

// UpdateAttribute will update campaign attributes
func (m *Manager) UpdateAttribute(attributes CampaignAttributes, ca *Campaign) error {

	_, err := m.FindCampaignAttributesByCampaignID(ca.ID)
	if err != sql.ErrNoRows {
		assert.Nil(err)
	}
	at := &attributes
	at.CampaignID = ca.ID

	if err != nil {
		err = m.CreateCampaignAttributes(at)
	} else {
		err = m.UpdateCampaignAttributes(at)
	}

	if err != nil {
		return err
	}

	ca.Attributes = at
	err = m.attachSchedule(ca)
	return err

}
