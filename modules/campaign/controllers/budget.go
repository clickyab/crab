package controllers

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/mail"
	"strconv"

	"clickyab.com/crab/modules/campaign/orm"
	"github.com/clickyab/services/assert"
	"github.com/clickyab/services/random"
	"github.com/rs/xmux"
	"github.com/sirupsen/logrus"
)

// @Validate{
//}
type budgetPayload struct {
	orm.CampaignFinance
}

func (l *budgetPayload) ValidateExtra(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	for _, m := range l.NotifyEmail {
		if _, err := mail.ParseAddress(m); err != nil {
			return err
		}
	}
	if !l.CostType.IsValid() {
		return fmt.Errorf("cost type %s is not valid. options are %s,%s or %s", l.CostType, orm.CPC, orm.CPM, orm.CPA)
	}
	if l.Budget < 0 || l.DailyLimit < 0 || l.MaxBid < 0 {
		return errors.New("budget, daily limit and max bid can not be a negative number")
	}

	return nil
}

// budget will update campaign finance
// @Route {
// 		url = /budget/:id
//		method = put
//		payload = budgetPayload
//		200 = orm.Campaign
//		400 = controller.ErrorResponseSimple
//		404 = controller.ErrorResponseSimple
//		middleware = authz.Authenticate
// }
func (c *Controller) budget(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(xmux.Param(ctx, "id"), 10, 64)
	p := c.MustGetPayload(ctx).(*budgetPayload)

	if err != nil || id < 1 {
		c.BadResponse(w, errors.New("id is not valid"))
	}
	db := orm.NewOrmManager()
	o, err := db.FindCampaignByID(id)
	if err != nil {
		c.NotFoundResponse(w, nil)
	}

	err = db.UpdateCampaignBudget(p.CampaignFinance, o)

	if err != nil {
		j, e := json.MarshalIndent(o, " ", "  ")
		assert.Nil(e)
		pj, e := json.MarshalIndent(p, " ", "  ")
		assert.Nil(e)

		eid := <-random.ID
		logrus.WithField("error", err).
			WithField("payload", string(pj)).
			WithField("eid", eid).
			WithField("campaign", string(j)).
			Debug("update base campaign ")
		w.Header().Set("x-error-id", eid)
		c.BadResponse(w, errors.New("can not update budget"))
		return
	}
	c.OKResponse(w, o)
}
