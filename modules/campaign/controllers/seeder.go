package controllers

import (
	"time"

	"math/rand"

	"clickyab.com/crab/libs"
	"clickyab.com/crab/modules/campaign/orm"
	orm2 "clickyab.com/crab/modules/inventory/orm"
	"github.com/clickyab/services/assert"
)

// Seed apply seed to fill campaign details
func Seed(campaign *orm.Campaign) {
	now := time.Now()
	startTime := now.AddDate(0, 0, -30)
	calcTime := startTime

	pubs := orm2.NewOrmManager().ListPublishersWithFilter("kind=? ORDER BY RAND() LIMIT 10", campaign.Kind)

	for i := 0; i <= 30; i++ {
		calcTime = calcTime.AddDate(0, 0, 1)
		dateID, hourID := libs.TimeToIDHour(calcTime)
		for i := range pubs {

			res := &orm.CampaignDetail{
				CampaignID:  campaign.ID,
				Imp:         randBetween(500, 1000),
				Click:       randBetween(10, 500),
				CPC:         randBetween(1000, 50000),
				Conv:        randBetween(0, 10),
				CPA:         randBetween(1000, 50000),
				CPM:         randBetween(1000, 50000),
				CreatedAt:   calcTime,
				UpdatedAt:   calcTime,
				DailyID:     dateID,
				HourID:      hourID,
				PublisherID: pubs[i].ID,
			}
			assert.Nil(orm.NewOrmManager().CreateCampaignDetail(res))
		}

		if calcTime.Unix() >= now.Unix() {
			break
		}

	}

}

func randBetween(min, max int64) int64 {
	return rand.Int63n(max-min) + min
}
