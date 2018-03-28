package controllers

import (
	"context"
	"net/http"
	"strconv"

	"time"

	"clickyab.com/crab/modules/campaign/errors"
	"clickyab.com/crab/modules/campaign/orm"
	"clickyab.com/crab/modules/domain/middleware/domain"
	"clickyab.com/crab/modules/user/aaa"
	"clickyab.com/crab/modules/user/middleware/authz"
	"github.com/clickyab/services/assert"
	"github.com/clickyab/services/gettext/t9e"
	"github.com/rs/xmux"
)

// @Validate{
//}
type copyCampaignPayload struct {
	Title string `json:"title" validate:"gt=3"`
}

// copy a campaign by id
// @Rest {
// 		url = /:id
//		protected = true
// 		method = patch
//		resource = copy_campaign:self
// }
func (c Controller) copyCampaign(ctx context.Context, r *http.Request, p *copyCampaignPayload) (*orm.Campaign, error) {
	currentUser := authz.MustGetUser(ctx)
	d := domain.MustGetDomain(ctx)
	id, err := strconv.ParseInt(xmux.Param(ctx, "id"), 10, 64)
	if err != nil {
		return nil, errors.InvalidIDErr
	}

	cpManager := orm.NewOrmManager()
	campaign, err := cpManager.FindCampaignByIDDomain(id, d.ID)
	if err != nil {
		return campaign, errors.NotFoundError(id)
	}

	userManager := aaa.NewAaaManager()
	owner, err := userManager.FindUserWithParentsByID(campaign.UserID, campaign.DomainID)
	assert.Nil(err)
	_, ok := aaa.CheckPermOn(owner, currentUser, "copy_campaign", campaign.DomainID)
	if !ok {
		return campaign, t9e.G("access denied. you can't copy campaign")
	}

	// check for archive campaign
	if campaign.ArchivedAt.Valid && campaign.ArchivedAt.Time.Before(time.Now()) {
		return campaign, errors.ArchivedEditError
	}

	campaign.Title = p.Title
	campaign.ID = 0
	err = cpManager.CreateCampaign(campaign)
	if err != nil {
		return campaign, errors.DuplicateNameError
	}

	return campaign, nil
}
