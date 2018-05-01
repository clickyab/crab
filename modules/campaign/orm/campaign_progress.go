package orm

import (
	"database/sql"
	"fmt"
	"time"

	"clickyab.com/crab/modules/user/aaa"
	"github.com/clickyab/services/assert"
)

// CampaignProgress campaign progress object
type CampaignProgress struct {
	TotalSpend  int64        `json:"total_spend" db:"total_spend"`
	Click       int64        `json:"click" db:"click"`
	Imp         int64        `json:"imp" db:"imp"`
	DailyBudget int64        `json:"daily_budget" db:"daily_budget"`
	TotalBudget int64        `json:"total_budget" db:"total_budget"`
	Ctr         float64      `json:"ctr" db:"ctr"`
	AvgCPC      float64      `json:"avg_cpc" db:"avg_cpc"`
	MaxBid      int64        `json:"max_bid" db:"max_bid"`
	OwnerEmail  string       `json:"owner_email" db:"owner_email"`
	Status      Status       `json:"status" db:"status"`
	Kind        CampaignKind `json:"kind" db:"kind"`
	StartAt     time.Time    `json:"start_at" db:"start_at"`
	EndAt       time.Time    `json:"end_at" db:"end_at"`
	Title       string       `json:"title" db:"title"`
}

// GetProgressData get progress bar data
func (m *Manager) GetProgressData(campaignID, domainID int64) CampaignProgress {
	var res CampaignProgress
	q := fmt.Sprintf(`
		SELECT
		c.total_spend 								AS total_spend,
		COALESCE(SUM(cd.click),0) 					AS click,
		COALESCE(SUM(cd.imp),0) 					AS imp,
		c.daily_budget 								AS daily_budget,
		c.total_budget 								AS total_budget,
		COALESCE((SUM(cd.click)/SUM(cd.imp))*10,0)  AS ctr,
		COALESCE(SUM(cd.cpc)/SUM(cd.click),0) 		AS avg_cpc,
		c.max_bid 									AS max_bid,
		owner.email 								AS owner_email,
		c.status 									AS status,
		c.kind 										AS kind,
		c.start_at 									AS start_at,
		c.title 									AS title,
		end_at 										AS end_at
		FROM %s AS c
		INNER JOIN %s AS owner ON owner.id=c.user_id
		LEFT JOIN %s AS cd ON cd.campaign_id=c.id
		WHERE c.id=? AND c.domain_id=? 
		GROUP BY c.id LIMIT 1`,
		CampaignTableFull,
		aaa.UserTableFull,
		CampaignDetailTableFull,
	)

	err := m.GetRDbMap().SelectOne(&res, q, campaignID, domainID)
	if err != sql.ErrNoRows {
		assert.Nil(err)
	}

	return res
}
