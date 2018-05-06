package controllers

import (
	"context"
	"net/http"

	"strconv"

	"clickyab.com/crab/modules/ad/errors"
	"clickyab.com/crab/modules/ad/orm"
	campOrm "clickyab.com/crab/modules/campaign/orm"
	"clickyab.com/crab/modules/domain/middleware/domain"
	orm2 "clickyab.com/crab/modules/domain/orm"
	"clickyab.com/crab/modules/user/aaa"
	"clickyab.com/crab/modules/user/middleware/authz"
	"github.com/rs/xmux"
)

// @Validate{
//}
type changeStatus struct {
	Status          orm.CreativeStatusType `json:"status" validate:"required"`
	currentCampaign *campOrm.Campaign      `json:"-"`
	currentDomain   *orm2.Domain           `json:"-"`
}

// CreativeStatusChangeResult to return number of changed creative {
type CreativeStatusChangeResult struct {
	CampaignID               int64 `json:"campaign"`
	NumberOfEffectedCreative int64 `json:"number_of_effected_creative"`
}

func (p *changeStatus) ValidateExtra(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	if !p.Status.IsValid() {
		return errors.InvalidStatusErr
	}
	dm := domain.MustGetDomain(ctx)
	p.currentDomain = dm
	idInt, err := strconv.ParseInt(xmux.Param(ctx, "id"), 10, 0)
	if err != nil {
		return errors.InvalidIDErr
	}
	// find campaign by id domain
	currentCampaign, err := campOrm.NewOrmManager().FindCampaignByIDDomain(idInt, dm.ID)
	if err != nil {
		return errors.InvalidIDErr
	}
	p.currentCampaign = currentCampaign
	return nil
}

// changeCampaignCreativeStatus to campaign
// @Rest {
// 		url = /campaign-creative-status/:id
//		protected = true
// 		method = patch
// 		resource = change_creative_status:self
// }
func (c Controller) changeCampaignCreativeStatus(ctx context.Context, r *http.Request, p *changeStatus) (*CreativeStatusChangeResult, error) {
	currentUser := authz.MustGetUser(ctx)
	//find campaign owner
	owner, err := aaa.NewAaaManager().FindUserWithParentsByID(p.currentCampaign.UserID, p.currentDomain.ID)
	if err != nil {
		return nil, errors.InvalidIDErr
	}
	//check permission
	_, ok := aaa.CheckPermOn(owner, currentUser, "change_creative_status", p.currentDomain.ID)
	if !ok {
		return nil, errors.AccessDenied
	}
	db := orm.NewOrmManager()
	rowEffectedCount, err := db.ChangeStatusOfAllCreativeInACampaign(p.currentCampaign.ID, p.Status)
	if err != nil {
		return nil, errors.UpdateStatusDbErr
	}
	// if any error return it
	return &CreativeStatusChangeResult{
		CampaignID:               p.currentCampaign.ID,
		NumberOfEffectedCreative: rowEffectedCount,
	}, nil

}
