package controllers

import (
	"context"
	"net/http"
	"strconv"

	"clickyab.com/crab/modules/ad/add"
	AdsErr "clickyab.com/crab/modules/ad/errors"
	"clickyab.com/crab/modules/campaign/errors"
	"clickyab.com/crab/modules/domain/middleware/domain"
	"clickyab.com/crab/modules/user/aaa"
	"clickyab.com/crab/modules/user/middleware/authz"
	"github.com/clickyab/services/gettext/t9e"
	"github.com/rs/xmux"
)

type sliceAds []add.AdUser

// getCampaignAds get all campaign ads
// @Rest {
// 		url = /get/:id/ad
//		protected = true
// 		method = get
//		resource = get_banner:self
// }
func (c Controller) getCampaignAds(ctx context.Context, r *http.Request) (sliceAds, error) {
	campaignIDInt, err := strconv.ParseInt(xmux.Param(ctx, "id"), 10, 64)
	if err != nil {
		return nil, errors.InvalidIDErr
	}
	d := domain.MustGetDomain(ctx)
	adManager := add.NewAddManager()
	ads, ownerID := adManager.GetAdsByCampaignID(campaignIDInt, d.ID)
	if len(ads) == 0 {
		return ads, AdsErr.AdNotFound(campaignIDInt)
	}
	currentUser := authz.MustGetUser(ctx)
	userManager := aaa.NewAaaManager()
	owner, err := userManager.FindUserWithParentsByID(ownerID, d.ID)
	if err != nil {
		return ads, err
	}

	_, ok := aaa.CheckPermOn(owner, currentUser, "get_banner", d.ID)
	if !ok {
		return ads, t9e.G("access denied. can't get campaign banner data")
	}

	return ads, nil
}
