package add

import (
	"time"

	"github.com/clickyab/services/assert"
)

// AdActiveStatus is the ad active status
type (
	// AdActiveStatus is the ad active status
	// @Enum{
	// }
	AdActiveStatus string
)

const (
	// RejectedAdStatus ad rejected status
	RejectedAdStatus AdActiveStatus = "rejected"
	// AcceptedAdStatus ad accepted status
	AcceptedAdStatus AdActiveStatus = "accepted"
	// PendingAdStatus ad pending status
	PendingAdStatus AdActiveStatus = "pending"
)

// AdType is the ad type
type (
	// AdType is the ad type
	// @Enum{
	// }
	AdType string
)

const (
	// BannerAdType banner
	BannerAdType AdType = "banner"
	// VideoAdType video
	VideoAdType AdType = "video"
	// NativeAdType native
	NativeAdType AdType = "native"
)

// Ad Ad model in database
// @Model {
//		table = ads
//		primary = true, id
//		find_by = id,src
//		list = yes
// }
type Ad struct {
	ID         int64          `json:"id" db:"id"`
	CampaignID int64          `json:"campaign_id" db:"campaign_id"`
	Src        string         `json:"src" db:"src"`
	Mime       string         `json:"mime" db:"mime"`
	Target     string         `json:"target" db:"target"`
	Width      int            `json:"width" db:"width"`
	Height     int            `json:"height" db:"height"`
	Status     AdActiveStatus `json:"status" db:"status"`
	Type       AdType         `json:"type" db:"type"`
	CreatedAt  time.Time      `json:"created_at" db:"created_at"`
	UpdatedAt  time.Time      `json:"updated_at" db:"updated_at"`
}

// BannerMethod either create or update
type BannerMethod string

var (
	// CreateBannerMethod creation selected
	CreateBannerMethod BannerMethod = "create"
	// UpdateBannerMethod update selected
	UpdateBannerMethod BannerMethod = "update"
)

// CreateUpdateCampaignNormalBanner assign banner to campaign either create or update
func (m *Manager) CreateUpdateCampaignNormalBanner(ads []*Ad) ([]Ad, error) {
	var res []Ad
	err := m.Begin()
	assert.Nil(err)
	defer func() {
		if err != nil {
			assert.Nil(m.Rollback())
		} else {
			assert.Nil(m.Commit())
		}
	}()
	for i := range ads {
		if ads[i].ID == 0 {
			err = m.CreateAd(ads[i])
		} else {
			err = m.UpdateAd(ads[i])
		}
		if err != nil {
			return res, err
		}
		res = append(res, *ads[i])
	}
	return res, nil
}
