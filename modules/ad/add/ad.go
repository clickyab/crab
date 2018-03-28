package add

import (
	"encoding/json"
	"time"

	"fmt"

	"database/sql/driver"

	"clickyab.com/crab/modules/ad/errors"
	"clickyab.com/crab/modules/campaign/orm"
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

// NativeAdAttr native ad attribute
type NativeAdAttr struct {
	Title string `json:"title"`
}

// AdAttr ad attribute
type AdAttr struct {
	Native *NativeAdAttr `json:"native,omitempty"`
}

// Scan for add attr
func (b *AdAttr) Scan(src interface{}) error {
	var c []byte

	switch src.(type) {
	case []byte:
		c = src.([]byte)
	case string:
		c = []byte(src.(string))
	default:
		return errors.UnsupportTypeError
	}

	return json.Unmarshal(c, b)
}

// Value for ad attr
func (b AdAttr) Value() (driver.Value, error) {
	e, err := json.Marshal(b)
	if err != nil {
		return nil, err
	}
	return e, nil
}

// Ad Ad model in database
// @Model {
//		table = ads
//		primary = true, id
//		find_by = id,src
//		list = yes
// }
type Ad struct {
	ID         int64          `json:"id" db:"id"`
	Title      string         `json:"title" db:"title"`
	CampaignID int64          `json:"campaign_id" db:"campaign_id"`
	Src        string         `json:"src" db:"src"`
	Mime       string         `json:"mime" db:"mime"`
	Target     string         `json:"target" db:"target"`
	Width      int            `json:"width" db:"width"`
	Height     int            `json:"height" db:"height"`
	Status     AdActiveStatus `json:"status" db:"status"`
	Type       AdType         `json:"type" db:"type"`
	Attr       AdAttr         `json:"attr" db:"attr"`
	CreatedAt  time.Time      `json:"created_at" db:"created_at"`
	UpdatedAt  time.Time      `json:"updated_at" db:"updated_at"`
}

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

// AdUser ad user obj
type AdUser struct {
	Ad
	UserID int64 `json:"user_id" db:"user_id"`
}

// AdsUserSlice slice for ad user
type AdsUserSlice []AdUser

// GetAdsByCampaignID return the Ad base on its campaign id
func (m *Manager) GetAdsByCampaignID(cpID int64, d int64) ([]AdUser, int64) {
	var res []AdUser
	var userID int64
	_, err := m.GetRDbMap().Select(
		&res,
		fmt.Sprintf("SELECT a.*,c.user_id FROM %s AS a "+
			"INNER JOIN %s AS c ON c.id=a.campaign_id "+
			"WHERE a.campaign_id=? AND c.domain_id=?",
			AdTableFull,
			orm.CampaignTableFull,
		),
		cpID,
		d,
	)
	assert.Nil(err)
	if len(res) > 0 {
		userID = res[0].UserID
	}
	return res, userID
}
