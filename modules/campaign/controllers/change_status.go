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
	"github.com/rs/xmux"
)

// @Validate{
//}
type changeCampaignStatus struct {
	Status orm.Status `json:"status" validate:"required"`
}

func (pl *changeCampaignStatus) ValidateExtra(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	if !pl.Status.IsValid() {
		return errors.InvalidCampaignStatusError
	}
	return nil
}

// changeStatus will update campaign finance status=start,pause
// @Rest {
// 		url = /status/:id
//		protected = true
// 		method = patch
//		resource = change_campaign_status:self
// }
func (c *Controller) changeStatus(ctx context.Context, r *http.Request, pl *changeCampaignStatus) (*controller.NormalResponse, error) {
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
	_, ok := aaa.CheckPermOn(owner, currentUser, "change_campaign_status", campaign.DomainID)
	if !ok {
		return nil, errors.AccessDenied
	}
	// if campaign current mode is archive nothing can be done
	if campaign.ArchivedAt.Valid && campaign.ArchivedAt.Time.Before(time.Now()) {
		//nothing can do
		return nil, errors.ChangeArchiveError
	}

	campaign.Status = pl.Status
	err = cpManager.UpdateCampaign(campaign)
	if err != nil {
		return nil, errors.UpdateCampaignErr
	}

	return nil, nil
}
