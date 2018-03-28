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
	"github.com/clickyab/services/gettext/t9e"
	"github.com/clickyab/services/mysql"
	"github.com/rs/xmux"
)

// budget will update campaign finance stat=archive,start,pause
// @Rest {
// 		url = /:id/:stat
//		protected = true
// 		method = patch
//		resource = change_campaign:self
// }
func (c *Controller) archive(ctx context.Context, r *http.Request) (*controller.NormalResponse, error) {
	currentUser := authz.MustGetUser(ctx)
	d := domain.MustGetDomain(ctx)
	id, err := strconv.ParseInt(xmux.Param(ctx, "id"), 10, 64)
	if err != nil {
		return nil, errors.InvalidIDErr
	}
	//get and validate action name
	stat := xmux.Param(ctx, "stat")
	if stat == "" || (stat != "archive" && stat != "start" && stat != "pause") {
		return nil, errors.InvalidCampaignStatusError
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
	_, ok := aaa.CheckPermOn(owner, currentUser, "change_campaign", campaign.DomainID)
	if !ok {
		return nil, t9e.G("access denied. you can't change campaign status")
	}
	// if campaign current mode is archive nothing can be done
	if campaign.ArchivedAt.Valid && campaign.ArchivedAt.Time.Before(time.Now()) {
		//nothing can do
		return nil, t9e.G("can't manipulate status")
	}

	if stat == "start" {
		campaign.Status = orm.StartStatus
		assert.Nil(cpManager.UpdateCampaign(campaign))
	} else if stat == "pause" {
		campaign.Status = orm.PauseStatus
		assert.Nil(cpManager.UpdateCampaign(campaign))
	} else { // archive is selected
		campaign.ArchivedAt = mysql.NullTime{Valid: true, Time: time.Now()}
		assert.Nil(cpManager.UpdateCampaign(campaign))
	}
	err = cpManager.UpdateCampaign(campaign)
	if err != nil {
		return nil, t9e.G("can't update campaign data")
	}

	return nil, nil
}
