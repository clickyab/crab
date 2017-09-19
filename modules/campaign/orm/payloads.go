package orm

import (
	"time"
)

// CreateCampaign is model for create campaign
type CreateCampaign struct {
	Campaign struct {
		Kind     CampaignKind `json:"kind" validate:"required"`
		Type     CampaignType `json:"type" validate:"required"`
		CostType CostType     `json:"cost_type"`
		common
	} `json:"campaign"`
	Attributes attributes `json:"attributes"`
	Schedule   schedule   `json:"schedule"`
}

// UpdateCampaign is model for Update campaign
type UpdateCampaign struct {
	Campaign struct {
		common
	} `json:"campaign"`
	Attributes attributes `json:"attributes"`
	Schedule   schedule   `json:"schedule"`
}

type attributes struct {
	Email        []string `json:"email"`
	Device       []string `json:"device"`
	Manufacturer []string `json:"manufacturer"`
	OS           []string `json:"os"`
	Browser      []string `json:"browser"`
	IAB          []string `json:"iab"`
	Region       []string `json:"region"`
	Cellular     []string `json:"cellular"`
	ISP          []string `json:"isp"`
}
type common struct {
	Status       bool      `json:"status"`
	StartAt      time.Time `json:"start_at"`
	EndAt        time.Time `json:"end_at"`
	Title        string    `json:"title" validate:"gt=5"`
	Budget       int64     `json:"budget"`
	DailyLimit   int64     `json:"daily_limit"`
	CPCCost      int64     `json:"cpc_cost"`
	WhiteBlackID int64     `json:"white_black_id"`
}

type schedule struct {
	H00 bool `json:"h00"`
	H01 bool `json:"h01"`
	H02 bool `json:"h02"`
	H03 bool `json:"h03"`
	H04 bool `json:"h04"`
	H05 bool `json:"h05"`
	H06 bool `json:"h06"`
	H07 bool `json:"h07"`
	H08 bool `json:"h08"`
	H09 bool `json:"h09"`
	H10 bool `json:"h10"`
	H11 bool `json:"h11"`
	H12 bool `json:"h12"`
	H13 bool `json:"h13"`
	H14 bool `json:"h14"`
	H15 bool `json:"h15"`
	H16 bool `json:"h16"`
	H17 bool `json:"h17"`
	H18 bool `json:"h18"`
	H19 bool `json:"h19"`
	H20 bool `json:"h20"`
	H21 bool `json:"h21"`
	H22 bool `json:"h22"`
	H23 bool `json:"h23"`
}
