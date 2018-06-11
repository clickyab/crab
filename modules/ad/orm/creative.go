package orm

import (
	"time"

	"fmt"

	campaignOrm "clickyab.com/crab/modules/campaign/orm"
	"clickyab.com/crab/modules/user/aaa"
	as "github.com/clickyab/services/array"
	"github.com/clickyab/services/assert"
	"github.com/clickyab/services/mysql"
	"gopkg.in/gorp.v2"
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
	ID              int64                  `json:"id" db:"id" structs:"id"`
	UserID          int64                  `json:"user_id" db:"user_id" structs:"user_id"`
	CampaignID      int64                  `json:"campaign_id" db:"campaign_id" structs:"campaign_id"`
	Status          CreativeStatusType     `json:"status" db:"status" structs:"status"`
	Type            CreativeTypes          `json:"type" db:"type" structs:"type"`
	URL             string                 `json:"url" db:"url" structs:"url"`
	Name            string                 `json:"name" db:"name" structs:"name"`
	MaxBid          mysql.NullInt64        `json:"max_bid" db:"max_bid" structs:"max_bid,string"`
	Attributes      mysql.GenericJSONField `json:"attributes" db:"attributes" structs:"attributes"`
	RejectReasonsID mysql.NullInt64        `json:"reject_reason_id" db:"reject_reasons_id" structs:"reject_reasons_id,string"`
	CreatedAt       time.Time              `json:"created_at" db:"created_at" structs:"created_at"`
	UpdatedAt       time.Time              `json:"updated_at" db:"updated_at"`
	ArchivedAt      mysql.NullTime         `json:"archived_at" db:"archived_at"`
	aaa.AuditExtraData
}

// PostInsert to set creative id in audit
func (c *Creative) PostInsert(s gorp.SqlExecutor) error {
	err := c.SetAuditEntity("creative", c.ID)
	if err != nil {
		return err
	}
	return c.AuditExtraData.PostInsert(s)
}

// PostUpdate to set creative id in audit
func (c *Creative) PostUpdate(s gorp.SqlExecutor) error {
	err := c.SetAuditEntity("creative", c.ID)
	if err != nil {
		return err
	}
	return c.AuditExtraData.PostUpdate(s)
}

// CreativeSaveResult to return creative and related assets after insert or update
type CreativeSaveResult struct {
	Creative *Creative                `json:"creative"`
	Assets   map[string][]interface{} `json:"assets"`
}

// CreativeCampaignResult CreativeCampaignResult
type CreativeCampaignResult struct {
	CampaignOwnerEmail  string                   `json:"campaign_owner_email"`
	ID                  int64                    `json:"id"`
	CampaignOwnerMobile string                   `json:"campaign_owner_mobile"`
	Title               string                   `json:"title"`
	Kind                campaignOrm.CampaignKind `json:"kind"`
	Creatives           []CreativeSaveResult
}

// AdUser creative user obj
type AdUser struct {
	Creative
	UserID int64 `json:"user_id" db:"user_id"`
}

// AdsUserSlice slice for creative user
type AdsUserSlice []AdUser

// ToRejectRequest user request struct to reject a creative
type ToRejectRequest struct {
	CreativeID     int64 `json:"creative_id"`
	RejectReasonID int64 `json:"reject_reason_id"`
}

// ChangeStatusReq user request struct
// @Validate{
//}
type ChangeStatusReq struct {
	Status         CreativeStatusType `json:"status" validate:"required"`
	CreativeID     int64              `json:"creative_id" validate:"required"`
	RejectReasonID int64              `json:"reject_reason_id,omitempty"`
	Creative       *Creative          `json:"-"`
}

// CreativeWithRelation creative object with reject reasons and users
type CreativeWithRelation struct {
	ID                 int64              `db:"creative_id"`
	Name               string             `db:"name"`
	UserID             int64              `db:"user_id"`
	UserFirstName      string             `db:"first_name"`
	UserLastName       string             `db:"last_name"`
	UserEmail          string             `db:"email"`
	CreativeStatus     CreativeStatusType `db:"creative_status"`
	RejectReasonID     mysql.NullInt64    `db:"reject_reason_id"`
	RejectReasonReason mysql.NullString   `db:"reason"`
	Message            string             `json:"-" db:"-"`
}

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

// MultiCreative MultiCreative
type MultiCreative struct {
	Creative *Creative
	Assets   []*Asset
}

// AddMultiCreative insert new creative with related native assets
func (m *Manager) AddMultiCreative(cr []MultiCreative) ([]CreativeSaveResult, error) {
	err := m.Begin()
	assert.Nil(err)
	defer func() {
		if err != nil {
			assert.Nil(m.Rollback())
		} else {
			assert.Nil(m.Commit())
		}
	}()

	var creativeResult = make([]CreativeSaveResult, 0)

	for i := range cr {
		err = m.CreateCreative(cr[i].Creative)
		if err != nil {
			return nil, err
		}
		var finalAsset = make([]Asset, 0)

		for j := range cr[i].Assets {
			cr[i].Assets[j].CreativeID = cr[i].Creative.ID
			err = m.CreateAsset(cr[i].Assets[j])
			if err != nil {
				return nil, err
			}
			finalAsset = append(finalAsset, *cr[i].Assets[j])
		}
		creativeResult = append(creativeResult, CreativeSaveResult{
			Creative: cr[i].Creative,
			Assets:   beautyAsset(finalAsset),
		})
	}

	return creativeResult, nil
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
			GetSelectFields(CreativeTableFull, "cr"),
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

// FindCreativesByCampaign find creative by campaign id
func (m *Manager) FindCreativesByCampaign(campaignID int64) ([]*Creative, error) {
	var res []*Creative

	_, err := m.GetRDbMap().Select(
		&res,
		fmt.Sprintf("SELECT %s FROM %s "+
			"WHERE campaign_id=?",
			GetSelectFields(CreativeTableFull, ""),
			CreativeTableFull,
		),
		campaignID,
	)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// SetCampaignCreativesStatus to Change status of all Creative with campaign id
func (m *Manager) SetCampaignCreativesStatus(campaignID int64, status CreativeStatusType) (int64, error) {
	q := fmt.Sprintf("UPDATE %s SET status=? WHERE campaign_id=?", CreativeTableFull)
	res, err := m.GetWDbMap().Exec(q, status, campaignID)
	if err != nil {
		return 0, err
	}
	rowEffected, err := res.RowsAffected()
	return rowEffected, err
}

// ChangeCreativesStatus to change creatives status
func (m *Manager) ChangeCreativesStatus(req []ChangeStatusReq, campaignID int64) error {
	err := m.Begin()
	assert.Nil(err)
	defer func() {
		if err != nil {
			assert.Nil(m.Rollback())
		} else {
			assert.Nil(m.Commit())
		}
	}()
	var q string
	for _, creative := range req {
		if creative.Status == AcceptedCreativeStatus {
			q = fmt.Sprintf("UPDATE %s SET status=? WHERE id=? AND campaign_id=?", CreativeTableFull)
			_, err = m.GetWDbMap().Exec(q, AcceptedCreativeStatus, creative.CreativeID, campaignID)
		} else if creative.Status == RejectedCreativeStatus {
			q = fmt.Sprintf("UPDATE %s SET status=?, reject_reasons_id=? WHERE id=? AND campaign_id=?", CreativeTableFull)
			_, err = m.GetWDbMap().Exec(q, RejectedCreativeStatus, creative.RejectReasonID, creative.CreativeID, campaignID)
		}
		if err != nil {
			return err
		}
	}
	return nil
}

// GetCreativeWithIDRange get an array of creative with given ids, used to reduce query count
func (m *Manager) GetCreativeWithIDRange(idList []int64, campaignID int64) ([]*CreativeWithRelation, error) {
	creativesListStr := as.ArrayToString(idList, ",")
	var result []*CreativeWithRelation

	q := fmt.Sprintf("SELECT c.id as creative_id,c.name, c.status as creative_status,"+
		"u.id as user_id, u.first_name, u.last_name, u.email, c2.reason,"+
		"c2.id as reject_reason_id, c2.reason "+
		"FROM %s AS c "+
		"INNER JOIN users u on c.user_id = u.id "+
		"LEFT JOIN creative_reject_reasons c2 on c.reject_reasons_id = c2.id "+
		"WHERE c.id IN (%s) AND c.campaign_id=?",
		CreativeTableFull, creativesListStr)
	_, err := m.GetRDbMap().Select(
		&result,
		q,
		campaignID,
	)
	if err != nil {
		return nil, err
	}
	return result, nil
}
