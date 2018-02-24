package controllers

import (
	"context"
	"net/http"
	"strconv"

	"time"

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

// assignNormalBanner assignNormalBanner module is banner type (banner/native)
// @Route {
// 		url = /:id
//		method = patch
//		payload = copyCampaignPayload
//		resource = copy_campaign:self
//		middleware = authz.Authenticate
//		200 = orm.Campaign
//		400 = controller.ErrorResponseSimple
//		404 = controller.ErrorResponseSimple
//		403 = controller.ErrorResponseSimple
// }
func (c Controller) copyCampaign(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	currentUser := authz.MustGetUser(ctx)
	d := domain.MustGetDomain(ctx)
	id, err := strconv.ParseInt(xmux.Param(ctx, "id"), 10, 64)
	if err != nil {
		c.BadResponse(w, t9e.G("id not valid"))
	}
	p := c.MustGetPayload(ctx).(*copyCampaignPayload)
	cpManager := orm.NewOrmManager()
	campaign, err := cpManager.FindCampaignByIDDomain(id, d.ID)
	if err != nil {
		c.NotFoundResponse(w, t9e.G("campaign not found"))
		return
	}
	userManager := aaa.NewAaaManager()
	owner, err := userManager.FindUserWithParentsByID(campaign.UserID, campaign.DomainID)
	assert.Nil(err)
	_, ok := aaa.CheckPermOn(owner, currentUser, "copy_campaign", campaign.DomainID)
	if !ok {
		c.ForbiddenResponse(w, t9e.G("don't have access for this action"))
		return
	}

	// check for archive campaign
	if campaign.ArchiveAt.Valid && campaign.ArchiveAt.Time.Before(time.Now()) {
		c.BadResponse(w, t9e.G("cant copy the archived campaign"))
		return
	}

	campaign.Title = p.Title
	campaign.ID = 0
	err = cpManager.CreateCampaign(campaign)
	if err != nil {
		c.BadResponse(w, t9e.G("cant create campaign duplicate name error"))
		return
	}

	c.OKResponse(w, campaign)

}
