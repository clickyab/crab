package orm

import (
	"time"

	"fmt"

	"github.com/clickyab/services/mysql"
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
	Exchange     bool            `json:"exchange" db:"exchange"`
	WhiteBlackID mysql.NullInt64 `json:"-" db:"white_black_id"`
	// WhiteBlackType true is whitelist
	WhiteBlackType  mysql.NullBool           `json:"-" db:"white_black_type"`
	WhiteBlackValue mysql.StringMapJSONArray `json:"-" db:"white_black_value"`
	Progress        Progress                 `json:"-" db:"progress"`
	ListID          int64                    `json:"white_black_id,omitempty" db:"-"`
	Attributes      *CampaignAttributes      `json:"attributes,omitempty" db:"-"`
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
	Status   bool           `json:"status" db:"status"`
	StartAt  time.Time      `json:"start_at" db:"start_at"`
	EndAt    mysql.NullTime `json:"end_at" db:"end_at"`
	Title    string         `json:"title" db:"title" `
	Schedule ScheduleSheet  `json:"schedule" db:"-"`
}

// CampaignBase is minimum data for creating campaign (stage one)
type CampaignBase struct { // stage one create
	CampaignBaseType
	CampaignStatus
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

// FindCampaignByIDDomain return the Campaign base on its id and domain id
func (m *Manager) FindCampaignByIDDomain(id, d int64) (*Campaign, error) {
	var res Campaign
	err := m.GetRDbMap().SelectOne(
		&res,
		fmt.Sprintf("SELECT * FROM %s WHERE id=? AND domain_id=?", CampaignTableFull),
		id,
		d,
	)

	if err != nil {
		return nil, err
	}

	m.attachAttribute(&res)
	m.attachSchedule(&res)
	return &res, nil
}
