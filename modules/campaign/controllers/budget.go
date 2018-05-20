package controllers

import (
	"context"
	"net/http"

	"github.com/fatih/structs"

	"clickyab.com/crab/modules/campaign/errors"
	"clickyab.com/crab/modules/campaign/orm"
	"clickyab.com/crab/modules/user/aaa"
	userError "clickyab.com/crab/modules/user/errors"
	"clickyab.com/crab/modules/user/middleware/authz"
	"github.com/clickyab/services/xlog"
)

// @Validate{
//}
type budgetPayload struct {
	TotalBudget int64            `json:"total_budget" validate:"required,gt=0"`
	DailyBudget int64            `json:"daily_budget" validate:"required,gt=0"`
	Strategy    orm.Strategy     `json:"strategy" validate:"required"`
	MaxBid      int64            `json:"max_bid" validate:"required,gt=0"`
	Exchange    orm.ExchangeType `json:"exchange" validate:"required"`
	Receivers   []int64          `json:"receivers" validate:"omitempty"`
	baseData    *BaseData        `json:"-"`
}

func (l *budgetPayload) ValidateExtra(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	res, err := CheckUserCamapignDomain(ctx)
	if err != nil {
		return err
	}
	l.baseData = res

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
//		resource = edit_budget:self
// }
func (c *Controller) budget(ctx context.Context, r *http.Request, p *budgetPayload) (*orm.Campaign, error) {
	db := orm.NewOrmManager()
	token := authz.MustGetToken(ctx)
	// check access
	uScope, ok := aaa.CheckPermOn(p.baseData.owner, p.baseData.currentUser, "edit_budget", p.baseData.domain.ID)
	if !ok {
		return nil, errors.AccessDenied
	}

	err := p.baseData.campaign.SetAuditUserData(p.baseData.currentUser.ID, token, p.baseData.domain.ID, "edit_campaign", uScope)
	if err != nil {
		return nil, err
	}

	p.baseData.campaign.TotalBudget = p.TotalBudget
	p.baseData.campaign.DailyBudget = p.DailyBudget
	p.baseData.campaign.Strategy = p.Strategy
	p.baseData.campaign.MaxBid = p.MaxBid

	d := structs.Map(p.baseData.campaign)
	err = p.baseData.campaign.SetAuditDescribe(d, "edit campaign budget data")
	if err != nil {
		return nil, err
	}

	err = db.UpdateCampaign(p.baseData.campaign)
	if err != nil {
		xlog.GetWithError(ctx, err).Debug("update base campaign")

		return nil, errors.UpdateError
	}

	err = db.UpdateReportReceivers(p.Receivers, p.baseData.campaign.ID)
	if err != nil {
		xlog.GetWithError(ctx, err).Debug("add campaign report receivers error")

		return nil, errors.UpdateError
	}

	return p.baseData.campaign, nil
}
