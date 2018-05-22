package controllers

import (
	"context"
	"net/http"
	"strconv"

	"clickyab.com/crab/modules/campaign/errors"
	"clickyab.com/crab/modules/campaign/orm"
	"clickyab.com/crab/modules/domain/middleware/domain"
	"clickyab.com/crab/modules/user/aaa"
	userErrors "clickyab.com/crab/modules/user/errors"
	"clickyab.com/crab/modules/user/middleware/authz"
	"github.com/rs/xmux"
)

// getCampaignProgress getCampaignProgress
// @Rest {
// 		url = /progress/:id
//		protected = true
// 		method = get
//		resource = get_campaign:self
// }
func (c Controller) getCampaignProgress(ctx context.Context, r *http.Request) (*orm.CampaignProgress, error) {
	idInt, err := strconv.ParseInt(xmux.Param(ctx, "id"), 10, 0)
	if err != nil {
		return nil, errors.InvalidIDErr
	}
	currentUser := authz.MustGetUser(ctx)
	dm := domain.MustGetDomain(ctx)
	cpManger := orm.NewOrmManager()
	campaign, err := orm.NewOrmManager().FindCampaignByIDDomain(idInt, dm.ID)
	if err != nil {
		return nil, errors.NotFoundError(idInt)
	}

	owner, err := aaa.NewAaaManager().FindUserWithParentsByID(campaign.UserID, dm.ID)
	if err != nil {
		return nil, userErrors.NotFoundWithDomainError(dm.DomainBase)
	}

	// check access
	_, ok := aaa.CheckPermOn(owner, currentUser, "get_campaign", dm.ID)
	if !ok {
		return nil, errors.AccessDenied
	}
	res := cpManger.GetProgressData(campaign.ID, dm.ID)
	return &res, nil

}
