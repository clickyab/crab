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

// get gets a campaign by id
// @Rest {
// 		url = /get/:id
//		protected = true
// 		method = get
//		resource = get_campaign:self
// }
func (c *Controller) get(ctx context.Context, r *http.Request) (*orm.Campaign, error) {
	userDomain := domain.MustGetDomain(ctx)
	currentUser := authz.MustGetUser(ctx)
	id := xmux.Param(ctx, "id")
	campID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return nil, errors.InvalidIDErr
	}

	campaign, err := orm.NewOrmManager().FindCampaignByIDDomain(campID, userDomain.ID)
	if err != nil {
		return campaign, errors.NotFoundError(campID)
	}

	owner, err := aaa.NewAaaManager().FindUserWithParentsByID(campaign.UserID, userDomain.ID)
	if err != nil {
		return campaign, userErrors.NotFoundWithDomainError(userDomain.Name)
	}

	// check access
	_, ok := aaa.CheckPermOn(owner, currentUser, "get_campaign", userDomain.ID)
	if !ok {
		return campaign, errors.AccessDenied
	}

	return campaign, nil
}
