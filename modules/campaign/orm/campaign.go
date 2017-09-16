package orm

import (
	"time"

	"errors"

	"clickyab.com/crab/modules/domain/dmn"
	"clickyab.com/crab/modules/inventory/orm"
	"clickyab.com/crab/modules/user/aaa"
	"github.com/clickyab/services/assert"
	_ "github.com/clickyab/services/assert"
	"github.com/clickyab/services/mysql"
	_ "github.com/clickyab/services/mysql"
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
	BannerType CampaignType = "banner"
	VastType   CampaignType = "vast"
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
	Daily_limit  int64           `json:"daily_limit" db:"daily_limit"`
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


type CampaignPayload struct {
	Kind         CampaignKind `json:"kind"`
	Type         CampaignType `json:"type"`
	Status       bool         `json:"status"`
	StartAt      time.Time    `json:"start_at"`
	EndAt        time.Time    `json:"end_at"`
	Title        string       `json:"title"`
	Budget       int64        `json:"budget"`
	Daily_limit  int64        `json:"daily_limit"`
	CostType     CostType     `json:"cost_type"`
	CPCCost      int64        `json:"cpc_cost"`
	WhiteBlackID int64        `json:"white_black_id"`
	Email        []string     `json:"email"`
	Device       []string     `json:"device"`
	Manufacturer []string     `json:"manufacturer"`
	OS           []string     `json:"os"`
	Browser      []string     `json:"browser"`
	IAB          []string     `json:"iab"`
	Region       []string     `json:"region"`
	Cellular     []string     `json:"cellular"`
	ISP          []string     `json:"isp"`
	H00          bool         `json:"h00"`
	H01          bool         `json:"h01"`
	H02          bool         `json:"h02"`
	H03          bool         `json:"h03"`
	H04          bool         `json:"h04"`
	H05          bool         `json:"h05"`
	H06          bool         `json:"h06"`
	H07          bool         `json:"h07"`
	H08          bool         `json:"h08"`
	H09          bool         `json:"h09"`
	H10          bool         `json:"h10"`
	H11          bool         `json:"h11"`
	H12          bool         `json:"h12"`
	H13          bool         `json:"h13"`
	H14          bool         `json:"h14"`
	H15          bool         `json:"h15"`
	H16          bool         `json:"h16"`
	H17          bool         `json:"h17"`
	H18          bool         `json:"h18"`
	H19          bool         `json:"h19"`
	H20          bool         `json:"h20"`
	H21          bool         `json:"h21"`
	H22          bool         `json:"h22"`
	H23          bool         `json:"h23"`
}

func (m *Manager) AddCampaign(p CampaignPayload, u *aaa.User, d *dmn.Domain) (int64, error) {
	c := &Campaign{}
	now := time.Now()
	c.CreatedAt = now
	c.UpdatedAt = now
	c.Active = true
	c.UserID = u.ID
	c.DomainID = d.ID

	c.Kind = p.Kind
	c.Type = p.Type
	c.Status = p.Status
	c.StartAt = p.StartAt
	c.EndAt = p.EndAt
	c.Title = p.Title
	c.Budget = p.Budget
	c.Daily_limit = p.Daily_limit
	c.CostType = p.CostType
	c.CPCCost = p.CPCCost

	if p.WhiteBlackID > 1 {
		n, e := orm.NewOrmManager().FindWhiteBlackListByID(p.WhiteBlackID)
		if e != nil {
			return 0, errors.New("there is no inventory with this id")
		}
		c.WhiteBlackID = mysql.NullInt64{
			Valid: true,
			Int64: p.WhiteBlackID,
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
	a.Email = p.Email
	a.Device = p.Device
	a.Manufacturer = p.Manufacturer
	a.OS = p.OS
	a.Browser = p.Browser
	a.IAB = p.IAB
	a.Region = p.Region
	a.Cellular = p.Cellular
	a.ISP = p.ISP
	a.CampaignID = c.ID

	err = m.CreateCampaignAttributes(a)
	if err != nil {
		return 0, err
	}

	h := &Schedule{}
	h.CampaignID = c.ID
	h.UpdatedAt = now
	h.H00 = p.H00
	h.H01 = p.H01
	h.H02 = p.H02
	h.H03 = p.H03
	h.H04 = p.H04
	h.H05 = p.H05
	h.H06 = p.H06
	h.H07 = p.H07
	h.H08 = p.H08
	h.H09 = p.H09
	h.H10 = p.H10
	h.H11 = p.H11
	h.H12 = p.H12
	h.H13 = p.H13
	h.H14 = p.H14
	h.H15 = p.H15
	h.H16 = p.H16
	h.H17 = p.H17
	h.H18 = p.H18
	h.H19 = p.H19
	h.H20 = p.H20
	h.H21 = p.H21
	h.H22 = p.H22
	h.H23 = p.H23
	err = m.CreateSchedule(h)
	if err != nil {
		return 0, err
	}
	return c.ID, nil

}
