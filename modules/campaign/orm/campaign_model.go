package orm

import (
	"encoding/json"
	"time"

	"strings"

	"database/sql/driver"

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
	AppCampaign CampaignKind = "app"
)

// Progress is progress of campaign
// @Enum{
// }
type Progress string

const (
	// ProgressInProgress is in progress
	ProgressInProgress Progress = "inprogress"
	// ProgressFinalized is finalized
	ProgressFinalized Progress = "finalized"
)

// InventoryState is whether black or white list selected
// @Enum{
// }
type InventoryState string

// NullInventoryState to make a nullable enum for InventoryState
type NullInventoryState struct {
	Valid          bool
	InventoryState InventoryState
}

// MarshalJSON try to marshaling to json
func (nt NullInventoryState) MarshalJSON() ([]byte, error) {
	if nt.Valid {
		return json.Marshal(nt.InventoryState)
	}
	return []byte("null"), nil
}

// UnmarshalJSON try to unmarshal dae from input
func (nt *NullInventoryState) UnmarshalJSON(b []byte) error {
	text := strings.ToLower(string(b))
	if text == "null" {
		nt.Valid = false
		nt.InventoryState = InventoryState("")
		return nil
	}

	err := json.Unmarshal(b, &nt.InventoryState)
	if err != nil {
		return err
	}

	nt.Valid = true
	return nil
}

// Scan implements the Scanner interface.
func (nt *NullInventoryState) Scan(value interface{}) error {
	nt.InventoryState, nt.Valid = value.(InventoryState)
	return nil
}

// Value implements the driver Valuer interface.
func (nt NullInventoryState) Value() (driver.Value, error) {
	if !nt.Valid {
		return nil, nil
	}
	return string(nt.InventoryState), nil
}

const (
	// WhiteInventory white list selected
	WhiteInventory InventoryState = "white_list"
	// BlackInventory black list selected
	BlackInventory InventoryState = "black_list"
)

// Strategy is type of campaign <cpm,cpc,cpa>
// @Enum{
// }
type Strategy string

const (
	// CPM is cpm
	CPM Strategy = "cpm"
	// CPC is cpc
	CPC Strategy = "cpc"
	// CPA is cpa
	CPA Strategy = "cpa"
	wh  string   = " WHERE "
)

// Status is campaign status start/pause
// @Enum{
// }
type Status string

const (
	// StartStatus is start status
	StartStatus Status = "start"
	// PauseStatus is pause status
	PauseStatus Status = "pause"
)

// ExchangeType is campaign selected exchange
// @Enum{
// }
type ExchangeType string

const (
	// Clickyab exchange
	Clickyab ExchangeType = "clickyab"
	// AllExceptClickyab all exchanges without clickyab
	AllExceptClickyab ExchangeType = "all_except_clickyab"
	// All exchanges
	All ExchangeType = "all"
)

// Campaign campaign model in database
// @Model {
//		table = campaigns
//		primary = true, id
//		find_by = id
//		list = yes
// }
type Campaign struct {
	ID               int64                 `json:"id" db:"id"`
	UserID           int64                 `json:"user_id" db:"user_id"`
	DomainID         int64                 `json:"domain_id" db:"domain_id"`
	Title            string                `json:"title" db:"title"`
	Kind             CampaignKind          `json:"kind" db:"kind"`
	Status           Status                `json:"status" db:"status"`
	Progress         Progress              `json:"progress" db:"progress"`
	StartAt          time.Time             `json:"start_at" db:"start_at"`
	EndAt            mysql.NullTime        `json:"end_at" db:"end_at"`
	TotalBudget      int64                 `json:"total_budget" db:"total_budget"`
	DailyBudget      int64                 `json:"daily_budget" db:"daily_budget"`
	Strategy         Strategy              `json:"strategy" db:"strategy"`
	MaxBid           int64                 `json:"max_bid" db:"max_bid"`
	Exchange         ExchangeType          `json:"exchange" db:"exchange"`
	InventoryID      mysql.NullInt64       `json:"inventory_id" db:"inventory_id"`
	InventoryType    NullInventoryState    `json:"inventory_type" db:"inventory_type"` // InventoryType black_list or white_list
	InventoryDomains mysql.StringJSONArray `json:"-" db:"inventory_domains"`
	TLD              string                `json:"tld" db:"tld"`
	TodaySpend       int64                 `json:"today_spend" db:"today_spend"`
	TotalSpend       int64                 `json:"total_spend" db:"total_spend"`
	CreatedAt        time.Time             `json:"created_at" db:"created_at"`
	UpdatedAt        time.Time             `json:"updated_at" db:"updated_at"`
	ArchivedAt       mysql.NullTime        `json:"archived_at" db:"archived_at"`
}

// CampaignBase is minimum data for creating campaign (stage one)
type CampaignBase struct { // stage one create
	Status   Status         `json:"status" db:"status"`
	Progress Progress       `json:"progress" db:"progress"`
	StartAt  time.Time      `json:"start_at" db:"start_at"`
	EndAt    mysql.NullTime `json:"end_at" db:"end_at"`
	Title    string         `json:"title" db:"title"`
	Kind     CampaignKind   `json:"kind" db:"kind"`
	TLD      string         `json:"tld" db:"tld"`
	Schedule ScheduleSheet  `json:"schedule" db:"-"`
}

// CampaignDataTable is the campaign full data in data table
// @DataTable {
//		url = /list
//		entity = campaign
//		view = campaign_list:self
//		checkable = false
//		multiselect = false
//		datefilter = cp.created_at
//		controller = clickyab.com/crab/modules/campaign/controllers
//		fill = FillCampaignDataTableArray
//		_edit = campaign_edit:self
//		_copy = campaign_copy:self
// }
type CampaignDataTable struct {
	ID        int64     `json:"id" db:"id" type:"number"`
	CreatedAt time.Time `json:"created_at" db:"created_at" type:"date" sort:"true" map:"cp.created_at"`
	Active    bool      `json:"active" db:"active" type:"bool"`

	Kind CampaignKind `json:"kind" db:"kind" type:"enum" filter:"true" map:"cp.kind"`

	Status  bool           `json:"status" db:"status" type:"bool"`
	StartAt time.Time      `json:"start_at" db:"start_at" type:"date" sort:"true"`
	EndAt   mysql.NullTime `json:"end_at" db:"end_at" type:"date"`
	Title   string         `json:"title" db:"title" type:"string" search:"true" map:"cp.title"`

	Budget     int64    `json:"budget" db:"budget" type:"number"`
	DailyLimit int64    `json:"daily_limit" db:"daily_limit" type:"number"`
	CostType   Strategy `json:"cost_type" db:"cost_type" type:"enum" filter:"true" map:"cp.cost_type"`
	MaxBid     int64    `json:"max_bid" db:"max_bid" type:"number" sort:"true"`

	AvgCPC     float64 `json:"avg_cpc" db:"avg_cpc" graph:"avg_cpc,Avg. CPC,line,false"`
	AvgCPM     float64 `json:"avg_cpm" db:"avg_cpm"`
	Ctr        float64 `json:"ctr" db:"ctr" graph:"ctr,CTR,line,false"`
	TotalImp   int64   `json:"total_imp" db:"total_imp" graph:"imp,Total Impression,bar,true"`
	TotalClick int64   `json:"total_click" db:"total_click" graph:"click,Click,line,true"`
	TotalConv  int64   `json:"total_conv" db:"total_conv"`
	TotalCpc   int64   `json:"total_cpc" db:"total_cpc"`
	TotalCpm   int64   `json:"total_cpm" db:"total_cpm"`

	TotalSpent int64 `json:"total_spent" db:"-" graph:"total_spent,Total spent,line,false"`

	TodayImp   int64   `json:"today_imp" db:"today_imp"`
	TodayClick int64   `json:"today_click" db:"today_click"`
	TodayCtr   float64 `json:"today_ctr" db:"today_ctr"`

	ParentIDs   []int64          `db:"-" json:"-" visible:"false"`
	ParentEmail mysql.NullString `db:"parent_email" json:"parent_email"`
	OwnerEmail  string           `db:"owner_email" json:"owner_email" type:"string" search:"true" map:"owner.email"`
	OwnerID     int64            `db:"owner_id" json:"owner_id" visible:"false"`
	DomainID    int64            `db:"domain_id" json:"domain_id"`
	Actions     string           `db:"-" json:"_actions" visible:"false"`
}

// CampaignGraph is the campaign full data in data table
// @Graph {
//		url = /graph/all
//		entity = chart
//		view = campaign_graph:self
//		key = ID
//		controller = clickyab.com/crab/modules/campaign/controllers
//		fill = FillCampaignGraph
// }
type CampaignGraph struct {
	OwnerEmail string       `db:"owner_email" json:"owner_email" type:"string" search:"true" map:"owner.email"`
	Kind       CampaignKind `json:"kind" db:"kind" type:"enum" filter:"true" map:"cp.kind"`
	Title      string       `json:"title" db:"title" type:"string" search:"true" map:"cp.title"`

	ID         int64   `json:"id" db:"id" type:"number"`
	AvgCPC     float64 `json:"avg_cpc" db:"avg_cpc" graph:"avg_cpc,Avg. CPC,line,false,4"`
	AvgCPM     float64 `json:"avg_cpm" db:"avg_cpm" graph:"avg_cpm,Avg. CPM,line,false,5"`
	Ctr        float64 `json:"ctr" db:"ctr" graph:"ctr,CTR,line,false,3"`
	TotalImp   int64   `json:"total_imp" db:"total_imp" graph:"imp,Total Impression,bar,true,2"`
	TotalClick int64   `json:"total_click" db:"total_click" graph:"click,Click,line,true,1"`
	TotalSpent int64   `json:"total_spent" db:"total_spent" graph:"total_spent,Total spent,line,false,6"`
}
