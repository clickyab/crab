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
	"github.com/clickyab/services/framework/controller"
	"github.com/clickyab/services/mysql"
	"github.com/rs/xmux"
)

// archive will archive campaign
// @Rest {
// 		url = /archive/:id
//		protected = true
// 		method = patch
//		resource = archive_campaign:self
// }
func (c *Controller) archive(ctx context.Context, r *http.Request) (*controller.NormalResponse, error) {
	currentUser := authz.MustGetUser(ctx)
	d := domain.MustGetDomain(ctx)
	id, err := strconv.ParseInt(xmux.Param(ctx, "id"), 10, 64)
	if err != nil {
		return nil, errors.InvalidIDErr
	}

	// load campaign
	cpManager := orm.NewOrmManager()
	campaign, err := cpManager.FindCampaignByIDDomain(id, d.ID)
	if err != nil {
		return nil, errors.NotFoundError(id)
	}
	userManager := aaa.NewAaaManager()
	owner, err := userManager.FindUserWithParentsByID(campaign.UserID, campaign.DomainID)
	assert.Nil(err)
	_, ok := aaa.CheckPermOn(owner, currentUser, "archive_campaign", campaign.DomainID)
	if !ok {
		return nil, errors.AccessDenied
	}
	// if campaign current mode is archive nothing can be done
	if campaign.ArchivedAt.Valid && campaign.ArchivedAt.Time.Before(time.Now()) {
		//nothing can do
		return nil, errors.ChangeArchiveError
	}

	campaign.ArchivedAt = mysql.NullTime{Valid: true, Time: time.Now()}

	err = cpManager.UpdateCampaign(campaign)
	if err != nil {
		return nil, errors.UpdateCampaignErr
	}

	return nil, nil
}
