package orm

import (
	"time"

	"errors"

	"clickyab.com/crab/modules/domain/dmn"
	"clickyab.com/crab/modules/inventory/orm"
	"clickyab.com/crab/modules/user/aaa"
	"github.com/clickyab/services/assert"
	"github.com/clickyab/services/mysql"
)

// UserValidStatus is the user status
type (
	// UserValidStatus is the user status
	// @Enum{
	// }
	UserValidStatus string
)

const (
	// RegisteredUserStatus user registered
	RegisteredUserStatus UserValidStatus = "registered"
	// BlockedUserStatus user blocked
	BlockedUserStatus UserValidStatus = "blocked"
	// ActiveUserStatus user active
	ActiveUserStatus UserValidStatus = "active"
)

// CampaignKind is kind of campaign <web,app>
// @Enum{
// }
type CampaignKind string

const (
	// WebCampaign is web
	WebCampaign CampaignKind = "web"
	// AppCampaign is app
	AppCampaign = "app"
)

// CampaignType is type of campaign <vast,banner,native>
// @Enum{
// }
type CampaignType string

const (

	// BannerType of campaign
	BannerType CampaignType = "banner"
	// VastType   of campaign
	VastType CampaignType = "vast"
	// NativeType of campaign
	NativeType CampaignType = "native"
)

// CostType is type of campaign <cpm,cpc,cpa>
// @Enum{
// }
type CostType string

const (
	// CPM is cpm
	CPM CostType = "cpm"
	// CPC is cpc
	CPC CostType = "cpc"
	// CPA is cpa
	CPA CostType = "cpa"
)

// Campaign campaign model in database
// @Model {
//		table = campaigns
//		primary = true, id
//		find_by = id
//		list = yes
// }
type Campaign struct {
	ID           int64           `json:"id" db:"id"`
	CreatedAt    time.Time       `json:"created_at" db:"created_at"`
	UpdatedAt    time.Time       `json:"updated_at" db:"updated_at"`
	Active       bool            `json:"active" db:"active"`
	UserID       int64           `json:"user_id" db:"user_id"`
	DomainID     int64           `json:"domain_id" db:"domain_id"`
	Kind         CampaignKind    `json:"kind" db:"kind"`
	Type         CampaignType    `json:"type" db:"type"`
	Status       bool            `json:"status" db:"status"`
	StartAt      time.Time       `json:"start_at" db:"start_at"`
	EndAt        time.Time       `json:"end_at" db:"end_at"`
	Title        string          `json:"title" db:"title"`
	Budget       int64           `json:"budget" db:"budget"`
	DailyLimit   int64           `json:"daily_limit" db:"daily_limit"`
	CostType     CostType        `json:"cost_type" db:"cost_type"`
	CPCCost      int64           `json:"cpc_cost" db:"cpc_cost"`
	WhiteBlackID mysql.NullInt64 `json:"white_black_id" db:"white_black_id"`
	// WhiteBlackType true is whitelist
	WhiteBlackType  mysql.NullBool        `json:"white_black_id"db:"white_black_id"`
	WhiteBlackValue mysql.StringJSONArray `json:"white_blacklist"db:"white_blacklist"`
}

// CampaignAttributes model in database
// @Model {
//		table = campaign_attributes
//		primary = true, id
//		find_by = id, campaign_id
//		list = yes
// 		belong = Campaign:campaign_id
// }
type CampaignAttributes struct {
	ID           int64                 `json:"id" db:"id"`
	CampaignID   int64                 `json:"campaign_id" db:"campaign_id"`
	Email        mysql.StringJSONArray `json:"email" db:"email"`
	Device       mysql.StringJSONArray `json:"device" db:"device"`
	Manufacturer mysql.StringJSONArray `json:"manufacturer" db:"manufacturer"`
	OS           mysql.StringJSONArray `json:"os" db:"os"`
	Browser      mysql.StringJSONArray `json:"browser" db:"browser"`
	IAB          mysql.StringJSONArray `json:"iab" db:"iab"`
	Region       mysql.StringJSONArray `json:"region" db:"region"`
	Cellular     mysql.StringJSONArray `json:"cellular" db:"cellular"`
	ISP          mysql.StringJSONArray `json:"isp" db:"isp"`
}

// ErrInventoryID of insert or update campaign
var ErrInventoryID = errors.New("there is no inventory with this id")

// UpdateCampaignByPayload will update a campaign
func (m *Manager) UpdateCampaignByPayload(id int64, p UpdateCampaign) error {
	c, e := m.FindCampaignByID(id)
	if e != nil {
		return e
	}
	now := time.Now()
	c.Status = p.Campaign.Status
	c.StartAt = p.Campaign.StartAt
	c.EndAt = p.Campaign.EndAt
	c.Title = p.Campaign.Title
	c.Budget = p.Campaign.Budget
	c.DailyLimit = p.Campaign.DailyLimit
	c.UpdatedAt = now
	c.CPCCost = p.Campaign.CPCCost
	if p.Campaign.WhiteBlackID == 0 {
		c.WhiteBlackID.Valid = false
		c.WhiteBlackValue = []string{}
		c.WhiteBlackType.Valid = false
	} else if p.Campaign.WhiteBlackID != c.WhiteBlackID.Int64 {
		n, e := orm.NewOrmManager().FindWhiteBlackListByID(p.Campaign.WhiteBlackID)
		if e != nil {
			return ErrInventoryID
		}
		c.WhiteBlackID = mysql.NullInt64{
			Valid: true,
			Int64: p.Campaign.WhiteBlackID,
		}
		c.WhiteBlackType = mysql.NullBool{
			Valid: true,
			Bool:  n.Kind,
		}
		c.WhiteBlackValue = n.Domains
	}

	h, e := m.FindScheduleByCampaignID(id)
	if e != nil {
		return e
	}

	h.UpdatedAt = now
	h.H00 = p.Schedule.H00
	h.H01 = p.Schedule.H01
	h.H02 = p.Schedule.H02
	h.H03 = p.Schedule.H03
	h.H04 = p.Schedule.H04
	h.H05 = p.Schedule.H05
	h.H06 = p.Schedule.H06
	h.H07 = p.Schedule.H07
	h.H08 = p.Schedule.H08
	h.H09 = p.Schedule.H09
	h.H10 = p.Schedule.H10
	h.H11 = p.Schedule.H11
	h.H12 = p.Schedule.H12
	h.H13 = p.Schedule.H13
	h.H14 = p.Schedule.H14
	h.H15 = p.Schedule.H15
	h.H16 = p.Schedule.H16
	h.H17 = p.Schedule.H17
	h.H18 = p.Schedule.H18
	h.H19 = p.Schedule.H19
	h.H20 = p.Schedule.H20
	h.H21 = p.Schedule.H21
	h.H22 = p.Schedule.H22
	h.H23 = p.Schedule.H23

	a, e := m.FindCampaignAttributesByCampaignID(c.ID)
	if e != nil {
		return e
	}
	a.Email = p.Attributes.Email
	a.Device = p.Attributes.Device
	a.Manufacturer = p.Attributes.Manufacturer
	a.OS = p.Attributes.OS
	a.Browser = p.Attributes.Browser
	a.IAB = p.Attributes.IAB
	a.Region = p.Attributes.Region
	a.Cellular = p.Attributes.Cellular
	a.ISP = p.Attributes.ISP
	err := m.Begin()
	assert.Nil(err)
	defer func() {
		if err != nil {
			m.Rollback()
			return
		}
		m.Commit()
	}()
	err = m.UpdateCampaign(c)
	if err != nil {
		return err
	}
	err = m.UpdateCampaignAttributes(a)
	if err != nil {
		return err
	}
	err = m.UpdateSchedule(h)
	if err != nil {
		return err
	}
	return nil

}

// AddCampaign will add new campaign to campaign table
func (m *Manager) AddCampaign(p CreateCampaign, u *aaa.User, d *dmn.Domain) (int64, error) {
	c := &Campaign{}
	now := time.Now()
	c.CreatedAt = now
	c.UpdatedAt = now
	c.Active = true
	c.UserID = u.ID
	c.DomainID = d.ID

	c.Kind = p.Campaign.Kind
	c.Type = p.Campaign.Type
	c.Status = p.Campaign.Status
	c.StartAt = p.Campaign.StartAt
	c.EndAt = p.Campaign.EndAt
	c.Title = p.Campaign.Title
	c.Budget = p.Campaign.Budget
	c.DailyLimit = p.Campaign.DailyLimit
	c.CostType = p.Campaign.CostType
	c.CPCCost = p.Campaign.CPCCost

	if p.Campaign.WhiteBlackID > 0 {
		n, e := orm.NewOrmManager().FindWhiteBlackListByID(p.Campaign.WhiteBlackID)
		if e != nil {
			return 0, ErrInventoryID
		}
		c.WhiteBlackID = mysql.NullInt64{
			Valid: true,
			Int64: p.Campaign.WhiteBlackID,
		}
		c.WhiteBlackType = mysql.NullBool{
			Valid: true,
			Bool:  n.Kind,
		}
		c.WhiteBlackValue = n.Domains
	}

	err := m.Begin()
	assert.Nil(err)
	defer func() {
		if err != nil {
			m.Rollback()
			return
		}
		m.Commit()
	}()

	err = m.CreateCampaign(c)
	if err != nil {
		return 0, err
	}
	a := &CampaignAttributes{}
	a.Email = p.Attributes.Email
	a.Device = p.Attributes.Device
	a.Manufacturer = p.Attributes.Manufacturer
	a.OS = p.Attributes.OS
	a.Browser = p.Attributes.Browser
	a.IAB = p.Attributes.IAB
	a.Region = p.Attributes.Region
	a.Cellular = p.Attributes.Cellular
	a.ISP = p.Attributes.ISP
	a.CampaignID = c.ID

	err = m.CreateCampaignAttributes(a)
	if err != nil {
		return 0, err
	}

	h := &Schedule{}
	h.CampaignID = c.ID
	h.UpdatedAt = now
	h.H00 = p.Schedule.H00
	h.H01 = p.Schedule.H01
	h.H02 = p.Schedule.H02
	h.H03 = p.Schedule.H03
	h.H04 = p.Schedule.H04
	h.H05 = p.Schedule.H05
	h.H06 = p.Schedule.H06
	h.H07 = p.Schedule.H07
	h.H08 = p.Schedule.H08
	h.H09 = p.Schedule.H09
	h.H10 = p.Schedule.H10
	h.H11 = p.Schedule.H11
	h.H12 = p.Schedule.H12
	h.H13 = p.Schedule.H13
	h.H14 = p.Schedule.H14
	h.H15 = p.Schedule.H15
	h.H16 = p.Schedule.H16
	h.H17 = p.Schedule.H17
	h.H18 = p.Schedule.H18
	h.H19 = p.Schedule.H19
	h.H20 = p.Schedule.H20
	h.H21 = p.Schedule.H21
	h.H22 = p.Schedule.H22
	h.H23 = p.Schedule.H23
	err = m.CreateSchedule(h)
	if err != nil {
		return 0, err
	}
	return c.ID, nil

}
