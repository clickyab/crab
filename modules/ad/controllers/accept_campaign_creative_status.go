package controllers

import (
	"context"
	"net/http"

	"strconv"

	"clickyab.com/crab/modules/ad/errors"
	"clickyab.com/crab/modules/ad/orm"
	"clickyab.com/crab/modules/ad/services"
	campOrm "clickyab.com/crab/modules/campaign/orm"
	"clickyab.com/crab/modules/domain/middleware/domain"
	orm2 "clickyab.com/crab/modules/domain/orm"
	"clickyab.com/crab/modules/user/middleware/authz"
	"github.com/clickyab/services/permission"
	"github.com/clickyab/services/xlog"
	"github.com/rs/xmux"
)

// @Validate{
//}
type changeStatus struct {
	currentCampaign *campOrm.Campaign `json:"-"`
	currentDomain   *orm2.Domain      `json:"-"`
	creatives       []*orm.Creative   `json:"-"`
}

// CreativeStatusChangeResult to return number of changed creative {
type CreativeStatusChangeResult struct {
	CampaignID        int64 `json:"campaign"`
	EffectedCreatives int64 `json:"effected_creatives"`
}

func (p *changeStatus) ValidateExtra(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
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
	p.creatives, err = orm.NewOrmManager().FindCreativesByCampaign(currentCampaign.ID)
	if err != nil {
		return errors.CreativeNotFoundErr
	}
	p.currentCampaign = currentCampaign
	return nil
}

// acceptCampaignCreativeStatus to campaign
// @Rest {
// 		url = /accept-campaign-creative/:id
//		protected = true
// 		method = patch
// 		resource = change_creative_status:superGlobal
// }
func (c Controller) acceptCampaignCreativeStatus(ctx context.Context, r *http.Request, p *changeStatus) (*CreativeStatusChangeResult, error) {
	currentUser := authz.MustGetUser(ctx)
	token := authz.MustGetToken(ctx)
	err := services.AcceptCreatives(p.creatives, currentUser.ID, p.currentDomain.ID, token, permission.ScopeSuperGlobal)
	if err != nil {
		xlog.GetWithError(ctx, err).Errorf("campaign creatives accept error from db: campaignId:%d", p.currentCampaign.ID)
		return nil, errors.UpdateStatusDbErr
	}
	var ids []int64
	for _, r := range p.creatives {
		ids = append(ids, r.ID)
	}
	err = services.SendChangeStatusMessage(ids, p.currentCampaign.ID)
	if err != nil {
		xlog.GetWithError(ctx, err).Debug("send notify email for creative accept failed")
		return nil, errors.SendNotifyEmailErr
	}

	res := &CreativeStatusChangeResult{
		CampaignID:        p.currentCampaign.ID,
		EffectedCreatives: int64(len(p.creatives)),
	}
	return res, nil
}
