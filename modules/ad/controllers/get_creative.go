package controllers

import (
	"context"
	"net/http"

	"strconv"

	"clickyab.com/crab/modules/ad/errors"
	"clickyab.com/crab/modules/ad/orm"
	cErr "clickyab.com/crab/modules/campaign/errors"
	cManager "clickyab.com/crab/modules/campaign/orm"
	"clickyab.com/crab/modules/domain/middleware/domain"
	"clickyab.com/crab/modules/user/aaa"
	"clickyab.com/crab/modules/user/middleware/authz"
	"github.com/rs/xmux"
)

// getCreative to get crative by id
// @Rest {
// 		url = /creative/:id
//		protected = true
// 		method = get
// 		resource = get_creative:self
// }
func (c Controller) getCreative(ctx context.Context, r *http.Request) (*orm.CreativeSaveResult, error) {
	idInt, err := strconv.ParseInt(xmux.Param(ctx, "id"), 10, 0)
	if err != nil {
		return nil, errors.InvalidIDErr
	}
	currentUser := authz.MustGetUser(ctx)
	dm := domain.MustGetDomain(ctx)
	// find creative
	currentCreative, err := orm.NewOrmManager().FindCreativeByID(idInt)
	if err != nil {
		return nil, errors.InvalidIDErr
	}
	currentOwner, err := aaa.NewAaaManager().FindUserWithParentsByID(currentCreative.UserID, dm.ID)
	if err != nil {
		return nil, errors.InvalidIDErr
	}

	_, ok := aaa.CheckPermOn(currentOwner, currentUser, "get_creative", dm.ID)
	if !ok {
		return nil, errors.AccessDenied
	}

	//find campaign
	currentCampaign, err := cManager.NewOrmManager().FindCampaignByIDDomain(currentCreative.CampaignID, dm.ID)
	if err != nil {
		return nil, cErr.NotFoundError(currentCreative.CampaignID)
	}
	campiagnOwner, err := aaa.NewAaaManager().FindUserWithParentsByID(currentCampaign.UserID, dm.ID)
	if err != nil {
		return nil, cErr.NotFoundError(currentCreative.CampaignID)
	}
	_, ok = aaa.CheckPermOn(campiagnOwner, currentUser, "get_creative", dm.ID)
	if !ok {
		return nil, errors.AccessDenied
	}
	assets := orm.NewOrmManager().FindAssetsBeautyByCreativeID(currentCreative.ID)
	return &orm.CreativeSaveResult{
		Creative: currentCreative,
		Assets:   assets,
	}, nil
}
