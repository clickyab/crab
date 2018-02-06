package controllers

import (
	"context"
	"fmt"
	"net/http"
	"net/mail"
	"strconv"

	"clickyab.com/crab/modules/campaign/errors"
	"clickyab.com/crab/modules/campaign/orm"
	"github.com/clickyab/services/gettext/t9e"
	"github.com/clickyab/services/xlog"
	"github.com/rs/xmux"
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

	return nil
}

// budget will update campaign finance
// @Rest {
// 		url = /budget/:id
//		protected = true
// 		method = put
// }
func (c *Controller) budget(ctx context.Context, r *http.Request, p *budgetPayload) (*orm.Campaign, error) {
	id, err := strconv.ParseInt(xmux.Param(ctx, "id"), 10, 64)
	if err != nil || id < 1 {
		return nil, errors.InvalidIDErr
	}

	db := orm.NewOrmManager()
	o, err := db.FindCampaignByID(id)
	if err != nil {
		return nil, errors.NotFoundError(id)
	}

	err = db.UpdateCampaignBudget(p.CampaignFinance, o)

	if err != nil {
		xlog.GetWithError(ctx, err).Debug("update base campaign")

		return nil, t9e.G("can't update campaign budget")
	}

	return o, nil
}
