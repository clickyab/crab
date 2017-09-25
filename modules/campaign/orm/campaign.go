package orm

import (
	"time"

	"errors"

	"database/sql"

	"clickyab.com/crab/modules/domain/dmn"
	"clickyab.com/crab/modules/inventory/orm"
	"clickyab.com/crab/modules/user/aaa"
	"github.com/clickyab/services/assert"
	"github.com/clickyab/services/config"
	"github.com/clickyab/services/mysql"
)

const (
	// Foreign is filter for region is every where except iran
	Foreign = "foreign"
)

var (
	//defaultWebBannerCPM = config.RegisterInt64("crab.modules.campaigns.defaults.web.banner.cpm", 0, "default web banner cpm")
	defaultWebBannerCPC = config.RegisterInt64("crab.modules.campaigns.defaults.web.banner.cpc", 250, "default web banner cpc")
	//defaultWebBannerCPA = config.RegisterInt64("crab.modules.campaigns.defaults.web.banner.cpa", 0, "default web banner cpa")

	//defaultWebVastCPM = config.RegisterInt64("crab.modules.campaigns.defaults.web.vast.cpm", 0, "default web vast cpm")
	defaultWebVastCPC = config.RegisterInt64("crab.modules.campaigns.defaults.web.vast.cpc", 200, "default web vast cpc")
	//defaultWebVastCPA = config.RegisterInt64("crab.modules.campaigns.defaults.web.vast.cpa", 0, "default web vast cpa")

	//defaultWebNativeCPM = config.RegisterInt64("crab.modules.campaigns.defaults.web.native.cpm", 0, "default web native cpm")
	defaultWebNativeCPC = config.RegisterInt64("crab.modules.campaigns.defaults.web.native.cpc", 150, "default web native cpc")
	//defaultWebNativeCPA = config.RegisterInt64("crab.modules.campaigns.defaults.web.native.cpa", 0, "default web native cpa")

	//defaultAppBannerCPM = config.RegisterInt64("crab.modules.campaigns.defaults.app.banner.cpm", 0, "default app banner cpm")
	defaultAppBannerCPC = config.RegisterInt64("crab.modules.campaigns.defaults.app.banner.cpc", 700, "default app banner cpc")
	//defaultAppBannerCPA = config.RegisterInt64("crab.modules.campaigns.defaults.app.banner.cpa", 0, "default app banner cpa")

	//defaultAppVastCPM = config.RegisterInt64("crab.modules.campaigns.defaults.app.vast.cpm", 0, "default app vast cpm")
	defaultAppVastCPC = config.RegisterInt64("crab.modules.campaigns.defaults.app.vast.cpc", 700, "default app vast cpc")
	//defaultAppVastCPA = config.RegisterInt64("crab.modules.campaigns.defaults.app.vast.cpa", 0, "default app vast cpa")

	//defaultAppNativeCPM = config.RegisterInt64("crab.modules.campaigns.defaults.app.native.cpm", 0, "default app native cpm")
	defaultAppNativeCPC = config.RegisterInt64("crab.modules.campaigns.defaults.app.native.cpc", 700, "default app native cpc")
	//defaultAppNativeCPA = config.RegisterInt64("crab.modules.campaigns.defaults.app.native.cpa", 0, "default app native cpa")
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

// Progress is progress of campaign
// @Enum{
// }
type Progress string

const (
	// ProgressInProgress is inprogress
	ProgressInProgress Progress = "inprogress"
	// ProgressFinalized is finalized
	ProgressFinalized Progress = "finalized"
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

type base struct {
	ID        int64     `json:"id" db:"id"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
	Active    bool      `json:"active" db:"active"`
}

// Campaign campaign model in database
// @Model {
//		table = campaigns
//		primary = true, id
//		find_by = id
//		list = yes
// }
type Campaign struct {
	base
	CampaignBaseType
	CampaignStatus
	CampaignFinance
	UserID       int64           `json:"user_id" db:"user_id"`
	DomainID     int64           `json:"domain_id" db:"domain_id"`
	WhiteBlackID mysql.NullInt64 `json:"-" db:"white_black_id"`
	// WhiteBlackType true is whitelist
	WhiteBlackType  mysql.NullBool        `json:"-" db:"white_black_type"`
	WhiteBlackValue mysql.StringJSONArray `json:"-" db:"white_black_value"`
	Progress        Progress              `json:"-" db:"progress"`
	ListID          int64                 `json:"white_black_id,omitempty" db:"-"`
	Attributes      *CampaignAttributes   `json:"attributes,omitempty" db:"-"`
}

// CampaignFinance is the financial
type CampaignFinance struct {
	Budget      int64                 `json:"budget" db:"budget"`
	DailyLimit  int64                 `json:"daily_limit" db:"daily_limit"`
	CostType    CostType              `json:"cost_type" db:"cost_type"`
	MaxBid      int64                 `json:"max_bid" db:"max_bid"`
	NotifyEmail mysql.StringJSONArray `json:"notify_email" db:"notify_email"`
}

// CampaignBaseType is fundamental data of campaign
type CampaignBaseType struct {
	Kind CampaignKind `json:"kind" db:"kind"`
	Type CampaignType `json:"type" db:"type"`
}

// CampaignStatus update campaign (stage one)
type CampaignStatus struct {
	Status   bool          `json:"status" db:"status"`
	StartAt  time.Time     `json:"start_at" db:"start_at"`
	EndAt    time.Time     `json:"end_at" db:"end_at"`
	Title    string        `json:"title" db:"title" validate:"required,gt=5"`
	Schedule ScheduleSheet `json:"schedule" db:"-"`
}

// CampaignBase is minimum data for creating campaign (stage one)
type CampaignBase struct { // stage one create
	CampaignBaseType
	CampaignStatus
}

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

// ErrInventoryID of insert or update campaign
var ErrInventoryID = errors.New("there is no inventory with this id")

// AddCampaign for creating campaign with minimum info
func (m *Manager) AddCampaign(c CampaignBase, u *aaa.User, d *dmn.Domain) (*Campaign, error) {
	ca := &Campaign{
		base: base{
			Active: true,
		},
		DomainID: d.ID,
		UserID:   u.ID,
		CampaignFinance: CampaignFinance{

			CostType: CPC,
		},
		CampaignBaseType: CampaignBaseType{

			Type: c.Type,
			Kind: c.Kind,
		},
		CampaignStatus: CampaignStatus{

			Status:  c.Status,
			StartAt: c.StartAt,
			EndAt:   c.EndAt,
			Title:   c.Title,
		},
		Progress: ProgressInProgress,
	}
	switch c.Kind {
	case WebCampaign:
		ca.webMaxBid(c)
	case AppCampaign:
		ca.appMaxBid(c)
	}
	err := m.Begin()
	assert.Nil(err)
	defer func() {
		if err == nil {
			assert.Nil(m.Commit())
			return
		}
		assert.Nil(m.Rollback())
	}()

	if err = m.CreateCampaign(ca); err != nil {
		return nil, err
	}
	s := &Schedule{
		CampaignID:    ca.ID,
		ScheduleSheet: c.Schedule,
	}
	ca.Schedule = s.ScheduleSheet
	err = m.CreateSchedule(s)
	return ca, err
}

var (
	// ErrorStartDate should raise if campaign start date is not valid
	ErrorStartDate = errors.New("start date can't be past")
)

// UpdateCampaignByID for updating campaign with minimum info
func (m *Manager) UpdateCampaignByID(c CampaignStatus, ca *Campaign) error {

	ca.Status = c.Status
	if ca.StartAt != c.StartAt {
		today, err := time.Parse("02-01-03", time.Now().Format("02-01-03"))
		assert.Nil(err)
		if c.StartAt.Unix() < today.Unix() {
			return ErrorStartDate
		}
	}
	ca.StartAt = c.StartAt
	ca.EndAt = c.EndAt
	ca.Title = c.Title

	err := m.Begin()
	assert.Nil(err)
	defer func() {
		if err == nil {
			assert.Nil(m.Commit())
			return
		}
		assert.Nil(m.Rollback())
	}()

	s, err := m.FindScheduleByCampaignID(ca.ID)
	if err != nil {
		return err
	}
	s.ScheduleSheet = c.Schedule

	err = m.UpdateCampaign(ca)
	if err != nil {
		return err
	}

	m.attachAttribute(ca)
	err = m.UpdateSchedule(s)
	return err
}

// UpdateCampaignFinance for updating campaign finance
func (m *Manager) UpdateCampaignFinance(c CampaignFinance, ca *Campaign) error {

	ca.Budget = c.Budget
	ca.DailyLimit = c.DailyLimit
	ca.CostType = c.CostType
	ca.MaxBid = c.MaxBid
	ca.NotifyEmail = c.NotifyEmail
	err := m.UpdateCampaign(ca)
	if err != nil {
		return err
	}
	m.attachSchedule(ca)
	m.attachAttribute(ca)
	return nil
}

// UpdateCampaignWhiteBlackList update white/black list
func (m *Manager) UpdateCampaignWhiteBlackList(w int64, ca *Campaign, u *aaa.User) error {

	var l *orm.WhiteBlackList
	var err error
	if w == -1 {
		l = &orm.WhiteBlackList{
			ID: 0,
		}
	} else {
		l, err = orm.NewOrmManager().FindWhiteBlackListByID(w)
		if err != nil {
			return err
		}
	}
	ca.WhiteBlackID = mysql.NullInt64{
		Valid: l.ID > 0,
		Int64: l.ID,
	}
	ca.WhiteBlackType = mysql.NullBool{
		Valid: l.ID > 0,
		Bool:  l.Kind,
	}
	ca.WhiteBlackValue = l.Domains
	err = m.UpdateCampaign(ca)
	if err != nil {
		return err
	}
	m.attachSchedule(ca)
	m.attachAttribute(ca)
	return nil
}

// UpdateAttribute will update campaign attributes
func (m *Manager) UpdateAttribute(attributes CampaignAttributes, ca *Campaign) error {

	at, err := m.FindCampaignAttributesByCampaignID(ca.ID)
	if err != sql.ErrNoRows {
		assert.Nil(err)
	}
	at = &attributes
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
	m.attachSchedule(ca)
	return nil

}

// Finalize will mark campaign ready for publish
func (m *Manager) Finalize(ca *Campaign) {

	ca.Progress = ProgressFinalized
	assert.Nil(m.UpdateCampaign(ca))
	m.attachAttribute(ca)
	m.attachSchedule(ca)

}

func (ca *Campaign) webMaxBid(c CampaignBase) {
	switch c.Type {
	case BannerType:
		ca.MaxBid = defaultWebBannerCPC.Int64()
	case VastType:
		ca.MaxBid = defaultWebVastCPC.Int64()
	case NativeType:
		ca.MaxBid = defaultWebNativeCPC.Int64()
	}
}
func (ca *Campaign) appMaxBid(c CampaignBase) {
	switch c.Type {
	case BannerType:
		ca.MaxBid = defaultAppBannerCPC.Int64()
	case VastType:
		ca.MaxBid = defaultAppVastCPC.Int64()
	case NativeType:
		ca.MaxBid = defaultAppNativeCPC.Int64()
	}
}

func (m *Manager) attachSchedule(c *Campaign) {
	s, err := m.FindScheduleByCampaignID(c.ID)
	assert.Nil(err)
	c.Schedule = s.ScheduleSheet
}

func (m *Manager) attachAttribute(c *Campaign) {
	s, err := m.FindCampaignAttributesByCampaignID(c.ID)
	if err != sql.ErrNoRows {
		assert.Nil(err)
	}
	c.Attributes = s
}
