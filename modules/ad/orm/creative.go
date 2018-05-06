package orm

import (
	"time"

	"fmt"

	campaignOrm "clickyab.com/crab/modules/campaign/orm"
	"github.com/clickyab/services/assert"
	"github.com/clickyab/services/mysql"
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

// Creative model in database
// @Model {
//		table = creatives
//		primary = true, id
//		find_by = id,campaign_id
//		list = yes
// }
type Creative struct {
	ID         int64                  `json:"id" db:"id"`
	UserID     int64                  `json:"user_id" db:"user_id"`
	CampaignID int64                  `json:"campaign_id" db:"campaign_id"`
	Status     CreativeStatusType     `json:"status" db:"status"`
	Type       CreativeTypes          `json:"type" db:"type"`
	URL        string                 `json:"url" db:"url"`
	Name       string                 `json:"name" db:"name"`
	MaxBid     mysql.NullInt64        `json:"max_bid" db:"max_bid"`
	Attributes mysql.GenericJSONField `json:"attributes" db:"attributes"`
	CreatedAt  time.Time              `json:"created_at" db:"created_at"`
	UpdatedAt  time.Time              `json:"updated_at" db:"updated_at"`
	ArchivedAt mysql.NullTime         `json:"archived_at" db:"archived_at"`
}

// CreativeSaveResult to return creative and related assets after insert or update
type CreativeSaveResult struct {
	Creative *Creative                `json:"creative"`
	Assets   map[string][]interface{} `json:"assets"`
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
func (m *Manager) AddCreative(cr *Creative, assets []*Asset) (*CreativeSaveResult, error) {
	err := m.Begin()
	assert.Nil(err)
	defer func() {
		if err != nil {
			assert.Nil(m.Rollback())
		} else {
			assert.Nil(m.Commit())
		}
	}()

	err = m.CreateCreative(cr)
	if err != nil {
		return nil, err
	}

	var finalAsset = make([]Asset, 0)

	for i := range assets {
		assets[i].CreativeID = cr.ID
		err = m.CreateAsset(assets[i])
		if err != nil {
			return nil, err
		}
		finalAsset = append(finalAsset, *assets[i])
	}
	return &CreativeSaveResult{
		Creative: cr,
		Assets:   beautyAsset(finalAsset),
	}, nil
}

// EditCreative update creative and its assets
func (m *Manager) EditCreative(cr *Creative, assets []*Asset) (*CreativeSaveResult, error) {
	err := m.Begin()
	assert.Nil(err)
	defer func() {
		if err != nil {
			assert.Nil(m.Rollback())
		} else {
			assert.Nil(m.Commit())
		}
	}()

	err = m.UpdateCreative(cr)
	if err != nil {
		return nil, err
	}

	var finalAsset = make([]Asset, 0)

	err = m.DeleteAssetsByCreative(cr.ID)
	if err != nil {
		return nil, err
	}

	for i := range assets {
		assets[i].CreativeID = cr.ID
		err = m.CreateAsset(assets[i])
		if err != nil {
			return nil, err
		}
		finalAsset = append(finalAsset, *assets[i])
	}
	return &CreativeSaveResult{
		Creative: cr,
		Assets:   beautyAsset(finalAsset),
	}, nil
}

// DeleteAssetsByCreative delete assets by creative id
func (m *Manager) DeleteAssetsByCreative(id int64) error {
	q := fmt.Sprintf("DELETE FROM %s WHERE creative_id=?", AssetTableFull)
	_, err := m.GetWDbMap().Exec(q, id)
	return err
}

// FindCreativeByIDAndType find creative with id and type
func (m *Manager) FindCreativeByIDAndType(crID int64, cType CreativeTypes) (*Creative, error) {
	var res Creative

	err := m.GetRDbMap().SelectOne(
		&res,
		fmt.Sprintf("SELECT %s FROM %s AS cr "+
			"WHERE cr.id=? AND cr.type=?",
			getSelectFields(CreativeTableFull, "cr"),
			CreativeTableFull,
		),
		crID,
		cType,
	)
	if err != nil {
		return nil, err
	}
	return &res, nil
}

// SetCreativesStatus to Change status of all Creative with campaign id
func (m *Manager) SetCreativesStatus(campaignID int64, status CreativeStatusType) (int64, error) {
	q := fmt.Sprintf("UPDATE %s SET status=? WHERE campaign_id=?", CreativeTableFull)
	res, err := m.GetWDbMap().Exec(q, status, campaignID)
	if err != nil {
		return 0, err
	}
	rowEffected, err := res.RowsAffected()
	return rowEffected, err
}
