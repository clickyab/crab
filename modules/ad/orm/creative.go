package orm

import (
	"time"

	"fmt"

	campaignOrm "clickyab.com/crab/modules/campaign/orm"
	"github.com/clickyab/services/assert"
	"github.com/clickyab/services/initializer"
)

// CreativeStatusType is the creative active status
type (
	// CreativeStatusType is the creative active status
	// @Enum{
	// }
	CreativeStatusType string
)

const (
	// RejectedCreativeStatus creative rejected status
	RejectedCreativeStatus CreativeStatusType = "rejected"
	// AcceptedCreativeStatus creative accepted status
	AcceptedCreativeStatus CreativeStatusType = "accepted"
	// PendingCreativeStatus creative pending status
	PendingCreativeStatus CreativeStatusType = "pending"
)

// CreativeTypes is the creative type
type (
	// CreativeTypes is the creative type
	// @Enum{
	// }
	CreativeTypes string
)

const (
	// CreativeBannerType banner
	CreativeBannerType CreativeTypes = "banner"
	// CreativeVastType vast
	CreativeVastType CreativeTypes = "vast"
	// CreativeNativeType native
	CreativeNativeType CreativeTypes = "native"
)

// BaseCreativeData base data of creative to use in some payloads
type BaseCreativeData struct {
	Type       CreativeTypes      `json:"-"`
	ID         int64              `json:"-"`
	Status     CreativeStatusType `json:"status" validation:"omitempty"`
	CampaignID int64              `json:"campaign_id" validation:"required"`
	URL        string             `json:"url" validation:"required"`
	MaxBid     int64              `json:"max_bid" validation="required,gt=0"`
	Attributes string             `json:"attributes" validation:"omitempty"`
}

// Creative model in database
// @Model {
//		table = creatives
//		primary = true, id
//		find_by = id,campaign_id
//		list = yes
// }
type Creative struct {
	ID         int64              `json:"id" db:"id"`
	CampaignID int64              `json:"campaign_id" db:"campaign_id"`
	Status     CreativeStatusType `json:"status" db:"status"`
	Type       CreativeTypes      `json:"type" db:"type"`
	URL        string             `json:"url" db:"url"`
	MaxBid     int64              `json:"max_bid" db:"max_bid"`
	Attributes string             `json:"attributes" db:"attributes"`
	CreatedAt  time.Time          `json:"created_at" db:"created_at"`
	UpdatedAt  time.Time          `json:"updated_at" db:"updated_at"`
	ArchivedAt time.Time          `json:"archived_at" db:"archived_at"`
}

// CreativeSaveResult to return creative and related assets after insert or update
type CreativeSaveResult struct {
	Creative Creative `json:"creative"`
	Assets   []Asset  `json:"assets"`
}

// AdUser creative user obj
type AdUser struct {
	Creative
	UserID int64 `json:"user_id" db:"user_id"`
}

// AdsUserSlice slice for creative user
type AdsUserSlice []AdUser

// GetAdsByCampaignID return the Creative base on its campaign id
func (m *Manager) GetAdsByCampaignID(cpID int64, d int64) ([]AdUser, int64) {
	var res []AdUser
	var userID int64
	_, err := m.GetRDbMap().Select(
		&res,
		fmt.Sprintf("SELECT a.*,c.user_id FROM %s AS a "+
			"INNER JOIN %s AS c ON c.id=a.campaign_id "+
			"WHERE a.campaign_id=? AND c.domain_id=?",
			CreativeTableFull,
			campaignOrm.CampaignTableFull,
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

// AddCreative insert new creative with related native assets
func (m *Manager) AddCreative(cr BaseCreativeData, assets []Asset) (CreativeSaveResult, error) {
	newCreative := Creative{
		CampaignID: cr.CampaignID,
		Status:     cr.Status,
		Type:       cr.Type,
		URL:        cr.URL,
		Attributes: cr.Attributes,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}

	res := CreativeSaveResult{
		Creative: newCreative,
		Assets:   assets,
	}

	// Start a new transaction
	err := m.Begin()
	if err != nil {
		return res, err
	}

	func(in interface{}) {
		if ii, ok := in.(initializer.Simple); ok {
			ii.Initialize()
		}
	}(newCreative)

	err = m.GetRDbMap().Insert(&newCreative)
	if err != nil {
		er := m.Rollback()
		assert.NotNil(er)

		return res, err
	}

	s := make([]interface{}, len(assets))
	for i := range assets {
		assets[i].CreativeID = newCreative.ID
		s[i] = &assets[i]
	}

	func(in interface{}) {
		if ii, ok := in.(initializer.Simple); ok {
			ii.Initialize()
		}
	}(s)

	err = m.GetRDbMap().Insert(s...)
	if err != nil {
		er := m.Rollback()
		assert.NotNil(er)

		return res, err
	}

	res = CreativeSaveResult{
		Creative: newCreative,
		Assets:   assets,
	}

	// if the commit is successful, a nil error is returned
	return res, m.Commit()
}

// EditCreative insert new creative with related assets
func (m *Manager) EditCreative(cr BaseCreativeData, assets []Asset) (CreativeSaveResult, error) {
	newData := Creative{
		ID:         cr.ID,
		CampaignID: cr.CampaignID,
		Status:     cr.Status,
		Type:       cr.Type,
		URL:        cr.URL,
		Attributes: cr.Attributes,
		UpdatedAt:  time.Now(),
	}

	res := CreativeSaveResult{
		Creative: newData,
		Assets:   assets,
	}

	// Start a new transaction
	err := m.Begin()
	if err != nil {
		return res, err
	}

	func(in interface{}) {
		if ii, ok := in.(initializer.Simple); ok {
			ii.Initialize()
		}
	}(newData)

	_, err = m.GetRDbMap().Update(&newData)
	if err != nil {
		er := m.Rollback()
		assert.NotNil(er)

		return res, err
	}

	err = m.DeleteAllCreativeAssets(newData.ID)
	if err != nil {
		er := m.Rollback()
		assert.NotNil(er)

		return res, err
	}

	s := make([]interface{}, len(assets))
	for i := range assets {
		assets[i].CreativeID = newData.ID
		s[i] = &assets[i]
	}

	func(in interface{}) {
		if ii, ok := in.(initializer.Simple); ok {
			ii.Initialize()
		}
	}(s)

	err = m.GetRDbMap().Insert(s...)
	if err != nil {
		er := m.Rollback()
		assert.NotNil(er)

		return res, err
	}

	res = CreativeSaveResult{
		Creative: newData,
		Assets:   assets,
	}

	// if the commit is successful, a nil error is returned
	return res, m.Commit()
}

// FindCreativeByIDAndType find creative with id and type
func (m *Manager) FindCreativeByIDAndType(crID int64, cType CreativeTypes) (Creative, error) {
	var res []Creative

	_, err := m.GetRDbMap().Select(
		&res,
		fmt.Sprintf("SELECT * FROM %s AS cr "+
			"WHERE cr.id=? AND cr.type=?",
			CreativeTableFull,
		),
		crID,
		cType,
	)

	if err != nil || len(res) == 0 {
		return Creative{}, err
	}

	return res[0], nil
}
