package controllers

import (
	"context"
	"net/http"
	"strconv"

	adManager "clickyab.com/crab/modules/ad/orm"
	"clickyab.com/crab/modules/campaign/errors"
	"clickyab.com/crab/modules/campaign/orm"
	"clickyab.com/crab/modules/domain/middleware/domain"
	"clickyab.com/crab/modules/user/aaa"
	"clickyab.com/crab/modules/user/middleware/authz"
	"github.com/rs/xmux"
)

// getCreativeByCampaign to get creative by id
// @Rest {
// 		url = /creative/:id
//		protected = true
// 		method = get
// 		resource = get_creative:self
// }
func (c Controller) getCreativeByCampaign(ctx context.Context, r *http.Request) (*adManager.CreativeCampaignResult, error) {
	idInt, err := strconv.ParseInt(xmux.Param(ctx, "id"), 10, 0)
	if err != nil {
		return nil, errors.InvalidIDErr
	}
	currentUser := authz.MustGetUser(ctx)
	dm := domain.MustGetDomain(ctx)

	//find campaign
	currentCampaign, err := orm.NewOrmManager().FindCampaignByIDDomain(idInt, dm.ID)
	if err != nil {
		return nil, errors.NotFoundError(idInt)
	}
	// find owner
	campaignOwner, err := aaa.NewAaaManager().FindUserByID(currentCampaign.UserID)
	if err != nil {
		return nil, errors.NotFoundError(currentCampaign.UserID)
	}
	_, ok := currentUser.HasOn("get_creative", currentCampaign.UserID, dm.ID, false, false)
	if !ok {
		return nil, errors.AccessDenied
	}

	var res = make([]*adManager.CreativeSaveResult, 0)

	// find creatives
	currentCreatives := adManager.NewOrmManager().ListCreativesWithFilter("campaign_id=?", currentCampaign.ID)
	for i := range currentCreatives {
		assets := adManager.NewOrmManager().FindAssetsBeautyByCreativeID(currentCreatives[i].ID)
		res = append(res, &adManager.CreativeSaveResult{
			Creative: &currentCreatives[i],
			Assets:   assets,
		})
	}

	return &adManager.CreativeCampaignResult{
		Creatives:          res,
		CampaignOwnerEmail: campaignOwner.Email,
	}, nil
}
