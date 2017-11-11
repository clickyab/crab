package controllers

import (
	"context"
	"errors"
	"net/http"
	"strconv"

	"clickyab.com/crab/modules/ad/add"
	"clickyab.com/crab/modules/domain/middleware/domain"
	"clickyab.com/crab/modules/user/aaa"
	"clickyab.com/crab/modules/user/middleware/authz"
	"github.com/clickyab/services/assert"
	"github.com/rs/xmux"
)

// getCampaignAds get all campaign ads
// @Route {
// 		url = /:id/ad
//		method = get
//		middleware = authz.Authenticate
//		resource = get_banner:self
//		200 = add.AdsUserSlice
//		400 = controller.ErrorResponseSimple
//		404 = controller.ErrorResponseSimple
// }
func (c Controller) getCampaignAds(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	campaignIDInt, err := strconv.ParseInt(xmux.Param(ctx, "id"), 10, 64)
	if err != nil {
		c.BadResponse(w, errors.New("campaign id not valid"))
		return
	}
	d := domain.MustGetDomain(ctx)
	adManager := add.NewAddManager()
	ads, ownerID := adManager.GetAdsByCampaignID(campaignIDInt, d.ID)
	if len(ads) == 0 {
		c.NotFoundResponse(w, errors.New("no ads found"))
		return
	}
	currentUser := authz.MustGetUser(ctx)
	userManager := aaa.NewAaaManager()
	owner, err := userManager.FindUserWithParentsByID(ownerID, d.ID)
	assert.Nil(err)
	_, ok := aaa.CheckPermOn(owner, currentUser, "get_banner", d.ID)
	if !ok {
		c.ForbiddenResponse(w, errors.New("dont have access for this action"))
		return
	}
	if len(ads) == 0 {
		c.OKResponse(w, []string{})
		return
	}
	c.OKResponse(w, ads)
}
