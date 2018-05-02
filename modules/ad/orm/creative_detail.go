package orm

import (
	"time"
)

// CreativeDetail creative detail model in database
// @Model {
//		table = creative_detail
//		primary = false, creative_id,hour_id
// }
type CreativeDetail struct {
	CampaignID int64     `json:"campaign_id" db:"campaign_id"`
	CreativeID int64     `json:"creative_id" db:"creative_id"`
	DailyID    int64     `json:"daily_id" db:"daily_id"`
	HourID     int64     `json:"hour_id" db:"hour_id"`
	FakeImp    int64     `json:"fake_imp" db:"fake_imp"`
	FakeClick  int64     `json:"fake_click" db:"fake_click"`
	Imp        int64     `json:"imp" db:"imp"`
	Click      int64     `json:"click" db:"click"`
	CPC        int64     `json:"cpc" db:"cpc"`
	CPA        int64     `json:"cpa" db:"cpa"`
	CPM        int64     `json:"cpm" db:"cpm"`
	Conv       int64     `json:"conv" db:"conv"`
	CreatedAt  time.Time `json:"created_at" db:"created_at"`
	UpdatedAt  time.Time `json:"updated_at" db:"updated_at"`
}
