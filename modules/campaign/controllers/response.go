package controllers

import (
	"time"

	"clickyab.com/crab/modules/campaign/orm"
	"github.com/clickyab/services/mysql"
)

type campaignResponse struct {
	ID          int64                   `json:"id" `
	CreatedAt   time.Time               `json:"created_at" `
	UpdatedAt   time.Time               `json:"updated_at" `
	Kind        orm.CampaignKind        `json:"kind" `
	Type        orm.CampaignType        `json:"type" `
	Budget      int64                   `json:"budget" `
	DailyLimit  int64                   `json:"daily_limit" `
	CostType    orm.CostType            `json:"cost_type" `
	MaxBid      int64                   `json:"max_bid" `
	NotifyEmail mysql.StringJSONArray   `json:"notify_email" `
	UserID      int64                   `json:"user_id" `
	DomainID    int64                   `json:"domain_id" `
	ListID      int64                   `json:"white_black_id,omitempty"`
	Attributes  *orm.CampaignAttributes `json:"attributes,omitempty"`
	Schedule    schedule                `json:"schedule"`
}

type schedule struct {
	H00 string `json:"h00" `
	H01 string `json:"h01" `
	H02 string `json:"h02" `
	H03 string `json:"h03" `
	H04 string `json:"h04" `
	H05 string `json:"h05" `
	H06 string `json:"h06" `
	H07 string `json:"h07" `
	H08 string `json:"h08" `
	H09 string `json:"h09" `
	H10 string `json:"h10" `
	H11 string `json:"h11" `
	H12 string `json:"h12" `
	H13 string `json:"h13" `
	H14 string `json:"h14" `
	H15 string `json:"h15" `
	H16 string `json:"h16" `
	H17 string `json:"h17" `
	H18 string `json:"h18" `
	H19 string `json:"h19" `
	H20 string `json:"h20" `
	H21 string `json:"h21" `
	H22 string `json:"h22" `
	H23 string `json:"h23" `
}

func createResponse(c *orm.Campaign) campaignResponse {
	return campaignResponse{
		ID:          c.ID,
		CreatedAt:   c.CreatedAt,
		UpdatedAt:   c.UpdatedAt,
		Kind:        c.Kind,
		Type:        c.Type,
		Budget:      c.Budget,
		DailyLimit:  c.DailyLimit,
		CostType:    c.CostType,
		MaxBid:      c.MaxBid,
		NotifyEmail: c.NotifyEmail,
		UserID:      c.UserID,
		DomainID:    c.DomainID,
		ListID:      c.ListID,
		Attributes:  c.Attributes,
		Schedule: schedule{
			H00: c.Schedule.H00.String,
			H01: c.Schedule.H01.String,
			H02: c.Schedule.H02.String,
			H03: c.Schedule.H03.String,
			H04: c.Schedule.H04.String,
			H05: c.Schedule.H05.String,
			H06: c.Schedule.H06.String,
			H07: c.Schedule.H07.String,
			H08: c.Schedule.H08.String,
			H09: c.Schedule.H09.String,
			H10: c.Schedule.H10.String,
			H11: c.Schedule.H11.String,
			H12: c.Schedule.H12.String,
			H13: c.Schedule.H13.String,
			H14: c.Schedule.H14.String,
			H15: c.Schedule.H15.String,
			H16: c.Schedule.H16.String,
			H17: c.Schedule.H17.String,
			H18: c.Schedule.H18.String,
			H19: c.Schedule.H19.String,
			H20: c.Schedule.H20.String,
			H21: c.Schedule.H21.String,
			H22: c.Schedule.H22.String,
			H23: c.Schedule.H23.String,
		},
	}

}
