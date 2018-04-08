package controllers

import (
	"context"
	"net/http"
	"strconv"

	"clickyab.com/crab/modules/campaign/errors"
	"clickyab.com/crab/modules/campaign/orm"
	"clickyab.com/crab/modules/user/aaa"
	userError "clickyab.com/crab/modules/user/errors"
	"clickyab.com/crab/modules/user/middleware/authz"
	"github.com/clickyab/services/gettext/t9e"
	"github.com/clickyab/services/xlog"
	"github.com/rs/xmux"
)

// @Validate{
//}
type budgetPayload struct {
	orm.CampaignFinance
	Exchange  orm.ExchangeType `json:"exchange" db:"exchange" validate:"required"`
	Receivers []int64          `json:"notify_users" validate:"required"`
}

func (l *budgetPayload) ValidateExtra(ctx context.Context, w http.ResponseWriter, r *http.Request) error {

	if !l.Strategy.IsValid() {
		return errors.InvalidStrategyError
	}

	if len(l.Receivers) > 0 {
		db := aaa.NewAaaManager()

		for _, uID := range l.Receivers {
			_, err := db.FindUserByID(uID)
			if err != nil {
				return userError.NotFoundError(uID)
			}
		}
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
	ca, err := db.FindCampaignByID(id)
	if err != nil {
		return nil, errors.NotFoundError(id)
	}

	// check access
	userManager := aaa.NewAaaManager()
	owner, err := userManager.FindUserWithParentsByID(ca.UserID, ca.DomainID)
	if err != nil {
		return ca, t9e.G("can't find user with related domain")
	}

	currentUser := authz.MustGetUser(ctx)
	_, ok := aaa.CheckPermOn(owner, currentUser, "edit_budget", ca.DomainID)
	if !ok {
		return ca, errors.AccessDenied
	}

	err = db.UpdateCampaignBudget(p.CampaignFinance, ca)
	if err != nil {
		xlog.GetWithError(ctx, err).Debug("update base campaign")

		return nil, t9e.G("can't update campaign budget")
	}

	err = db.UpdateReportReceivers(p.Receivers, ca.ID)
	if err != nil {
		xlog.GetWithError(ctx, err).Debug("add campaign report receivers error")

		return nil, t9e.G("can't add/update campaign report receivers")
	}

	return ca, nil
}
