package controllers

import (
	"context"
	"net/http"

	"strconv"

	"clickyab.com/crab/modules/ad/errors"
	adOrm "clickyab.com/crab/modules/ad/orm"
	"clickyab.com/crab/modules/ad/services"
	orm3 "clickyab.com/crab/modules/campaign/orm"
	"clickyab.com/crab/modules/domain/middleware/domain"
	orm2 "clickyab.com/crab/modules/domain/orm"
	"clickyab.com/crab/modules/user/middleware/authz"
	"github.com/clickyab/services/permission"
	"github.com/clickyab/services/xlog"
	"github.com/rs/xmux"
)

// @Validate{
//}
type changeStatusPayload struct {
	CreativesStatus []adOrm.ChangeStatusReq `json:"creatives_status" validation:"required"`
	currentDomain   *orm2.Domain            `json:"-"`
	currentCampaign *orm3.Campaign          `json:"-"`
}

// ChangeStatusResult return result of bulk changing status of creatives
type ChangeStatusResult struct {
	CreativesStatus []adOrm.ChangeStatusReq `json:"creatives_status"`
}

func (p *changeStatusPayload) ValidateExtra(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	idInt, err := strconv.ParseInt(xmux.Param(ctx, "id"), 10, 0)
	if err != nil {
		return errors.InvalidIDErr
	}
	dm := domain.MustGetDomain(ctx)
	p.currentDomain = dm
	// find campaign by id domain
	currentCampaign, err := orm3.NewOrmManager().FindCampaignByIDDomain(idInt, dm.ID)
	if err != nil {
		return errors.InvalidIDErr
	}
	orm := adOrm.NewOrmManager()
	for i := range p.CreativesStatus {
		p.CreativesStatus[i].Creative, err = orm.FindCreativeByID(p.CreativesStatus[i].CreativeID)
		if err != nil {
			return errors.CreativeNotFoundErr
		}
	}
	p.currentCampaign = currentCampaign
	return nil
}

// changeCreativesStatus bulk approve reject creatives status of a campaign
// @Rest {
// 		url = /change-creatives-status/:id
//		protected = true
// 		method = put
// 		resource = change_creatives_status:superGlobal
// }
func (c *Controller) changeCreativesStatus(ctx context.Context, r *http.Request, p *changeStatusPayload) (*ChangeStatusResult, error) {
	currentUser := authz.MustGetUser(ctx)
	token := authz.MustGetToken(ctx)
	// apply approve or reject
	prob := services.ChangeCreativesStatus(p.CreativesStatus, currentUser.ID, p.currentDomain.ID, token, permission.ScopeSuperGlobal)
	if prob != nil {
		xlog.GetWithError(ctx, prob).Debug("database error when change creatives status")
		return nil, errors.UpdateStatusDbErr
	}
	result := ChangeStatusResult{
		CreativesStatus: p.CreativesStatus,
	}

	var ids []int64
	for _, r := range p.CreativesStatus {
		ids = append(ids, r.CreativeID)
	}
	err := services.SendChangeStatusMessage(ids, p.currentCampaign.ID)
	if err != nil {
		xlog.GetWithError(ctx, err).Debug("send notify email for creative status change failed")
		return nil, errors.SendNotifyEmailErr
	}
	return &result, nil
}
