package orm

import (
	"encoding/json"
	"time"

	"strings"

	"database/sql/driver"

	"clickyab.com/crab/modules/user/aaa"
	"github.com/clickyab/services/gettext/t9e"
	"github.com/clickyab/services/mysql"
	"gopkg.in/gorp.v2"
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
	var b []byte
	switch value.(type) {
	case []byte:
		b = value.([]byte)
	case string:
		b = []byte(value.(string))
	case nil:
		b = make([]byte, 0)
	default:
		return t9e.G("unsupported type")
	}
	nt.InventoryState, nt.Valid = InventoryState(b), InventoryState(b) != ""
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
	wh           = " WHERE "
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
	ID               int64                 `json:"id" db:"id" structs:"id"`
	UserID           int64                 `json:"user_id" db:"user_id" structs:"user_id"`
	DomainID         int64                 `json:"domain_id" db:"domain_id" structs:"domain_id"`
	Title            string                `json:"title" db:"title" structs:"title"`
	Kind             CampaignKind          `json:"kind" db:"kind" structs:"kind"`
	Status           Status                `json:"status" db:"status" structs:"status"`
	Progress         Progress              `json:"progress" db:"progress" structs:"progress"`
	StartAt          time.Time             `json:"start_at" db:"start_at" structs:"start_at"`
	EndAt            mysql.NullTime        `json:"end_at" db:"end_at" structs:"end_at,string"`
	TotalBudget      int64                 `json:"total_budget" db:"total_budget" structs:"total_budget"`
	DailyBudget      int64                 `json:"daily_budget" db:"daily_budget" structs:"daily_budget"`
	Strategy         Strategy              `json:"strategy" db:"strategy" structs:"strategy"`
	MaxBid           int64                 `json:"max_bid" db:"max_bid" structs:"max_bid"`
	Exchange         ExchangeType          `json:"exchange" db:"exchange" structs:"exchange"`
	InventoryID      mysql.NullInt64       `json:"inventory_id" db:"inventory_id" structs:"inventory_id,string"`
	InventoryType    NullInventoryState    `json:"inventory_type" db:"inventory_type" structs:"inventory_type"` // InventoryType black_list or white_list
	InventoryDomains mysql.StringJSONArray `json:"-" db:"inventory_domains"`
	TLD              mysql.NullString      `json:"tld" db:"tld"`
	TodaySpend       int64                 `json:"today_spend" db:"today_spend"`
	TotalSpend       int64                 `json:"total_spend" db:"total_spend"`
	CreatedAt        time.Time             `json:"created_at" db:"created_at"`
	UpdatedAt        time.Time             `json:"updated_at" db:"updated_at"`
	ArchivedAt       mysql.NullTime        `json:"archived_at" db:"archived_at"`
	aaa.AuditExtraData
}

// PostInsert to set campaign id in audit
func (ca *Campaign) PostInsert(s gorp.SqlExecutor) error {
	err := ca.SetAuditEntity("campaign", ca.ID)
	if err != nil {
		return err
	}

	return ca.AuditExtraData.PostInsert(s)
}

// CampaignBase is minimum data for creating campaign (stage one)
type CampaignBase struct { // stage one create
	Status   Status           `json:"status" db:"status"`
	Progress Progress         `json:"progress" db:"progress"`
	StartAt  time.Time        `json:"start_at" db:"start_at"`
	EndAt    mysql.NullTime   `json:"end_at" db:"end_at"`
	Title    string           `json:"title" db:"title"`
	Kind     CampaignKind     `json:"kind" db:"kind"`
	TLD      mysql.NullString `json:"tld" db:"tld"`
	Schedule ScheduleSheet    `json:"schedule" db:"-"`
}
