package controllers

import (
	"time"

	"math/rand"

	"clickyab.com/crab/libs"
	"clickyab.com/crab/modules/ad/orm"
	"github.com/clickyab/services/assert"
)

// Seed apply seed to fill creatives details
func Seed(creative *orm.Creative) {
	now := time.Now()
	startTime := now.AddDate(0, 0, -30)
	calcTime := startTime

	for i := 0; i <= 30; i++ {
		calcTime = calcTime.AddDate(0, 0, 1)
		dateID, hourID := libs.TimeToIDHour(calcTime)

		res := &orm.CreativeDetail{
			CampaignID: creative.CampaignID,
			CreativeID: creative.ID,
			Imp:        randBetween(500, 1000),
			Click:      randBetween(10, 500),
			CPC:        randBetween(1000, 50000),
			Conv:       randBetween(0, 10),
			CPA:        randBetween(1000, 50000),
			CPM:        randBetween(1000, 50000),
			CreatedAt:  calcTime,
			UpdatedAt:  calcTime,
			DailyID:    dateID,
			HourID:     hourID,
		}
		assert.Nil(orm.NewOrmManager().CreateCreativeDetail(res))

		if calcTime.Unix() >= now.Unix() {
			break
		}

	}

}

func randBetween(min, max int64) int64 {
	rand.Seed(time.Now().Unix())
	return rand.Int63n(max-min) + min
}
