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
	"github.com/clickyab/services/mysql"
	"github.com/rs/xmux"
)

// budget will update campaign finance stat=archive,start,pause
// @Route {
// 		url = /:id/:stat
//		method = patch
//		resource = change_campaign:self
//		400 = controller.ErrorResponseSimple
//		404 = controller.ErrorResponseSimple
//		403 = controller.ErrorResponseSimple
//		middleware = authz.Authenticate
// }
func (c *Controller) archive(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	currentUser := authz.MustGetUser(ctx)
	d := domain.MustGetDomain(ctx)
	id, err := strconv.ParseInt(xmux.Param(ctx, "id"), 10, 64)
	if err != nil {
		c.BadResponse(w, t9e.G("id not valid"))
		return
	}
	//get and validate action name
	stat := xmux.Param(ctx, "stat")
	if stat == "" || (stat != "archive" && stat != "start" && stat != "pause") {
		c.BadResponse(w, t9e.G("invalid stat detected"))
		return
	}

	// load campaign
	cpManager := orm.NewOrmManager()
	campaign, err := cpManager.FindCampaignByIDDomain(id, d.ID)
	if err != nil {
		c.NotFoundResponse(w, t9e.G("campaign not found"))
		return
	}
	userManager := aaa.NewAaaManager()
	owner, err := userManager.FindUserWithParentsByID(campaign.UserID, campaign.DomainID)
	assert.Nil(err)
	_, ok := aaa.CheckPermOn(owner, currentUser, "change_campaign", campaign.DomainID)
	if !ok {
		c.ForbiddenResponse(w, t9e.G("don't have access for this action"))
		return
	}
	// if campaign current mode is archive nothing can be done
	if campaign.ArchiveAt.Valid && campaign.ArchiveAt.Time.Before(time.Now()) {
		//nothing can do
		c.BadResponse(w, t9e.G("cant manipulate archived campaign"))
		return
	}

	if stat == "start" {
		campaign.Active = true
		assert.Nil(cpManager.UpdateCampaign(campaign))
	} else if stat == "pause" {
		campaign.Active = false
		assert.Nil(cpManager.UpdateCampaign(campaign))
	} else { // archive is selected
		campaign.ArchiveAt = mysql.NullTime{Valid: true, Time: time.Now()}
		assert.Nil(cpManager.UpdateCampaign(campaign))
	}
	c.OKResponse(w, nil)
}
