package controllers

import (
	"context"
	"net/http"

	"time"

	"clickyab.com/crab/modules/campaign/errors"
	"clickyab.com/crab/modules/campaign/orm"
	"clickyab.com/crab/modules/user/aaa"
	"clickyab.com/crab/modules/user/middleware/authz"
	"github.com/clickyab/services/framework/controller"
	"github.com/fatih/structs"
)

// @Validate{
//}
type changeCampaignStatus struct {
	Status   orm.Status `json:"status" validate:"required"`
	baseData *BaseData  `json:"-"`
}

func (pl *changeCampaignStatus) ValidateExtra(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	if !pl.Status.IsValid() {
		return errors.InvalidCampaignStatusError
	}

	res, err := CheckUserCamapignDomain(ctx)
	if err != nil {
		return err
	}
	pl.baseData = res

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
	token := authz.MustGetToken(ctx)
	uScope, ok := aaa.CheckPermOn(pl.baseData.owner, pl.baseData.currentUser, "change_campaign_status", pl.baseData.campaign.DomainID)
	if !ok {
		return nil, errors.AccessDenied
	}

	err := pl.baseData.campaign.SetAuditUserData(pl.baseData.currentUser.ID, token, pl.baseData.campaign.DomainID, "edit_campaign", uScope)
	if err != nil {
		return nil, err
	}

	// if campaign current mode is archive nothing can be done
	if pl.baseData.campaign.ArchivedAt.Valid && pl.baseData.campaign.ArchivedAt.Time.Before(time.Now()) {
		//nothing can do
		return nil, errors.ChangeArchiveError
	}

	pl.baseData.campaign.Status = pl.Status

	d := structs.Map(pl.baseData.campaign)
	err = pl.baseData.campaign.SetAuditDescribe(d, "change campaign status")
	if err != nil {
		return nil, err
	}

	db := orm.NewOrmManager()
	err = db.UpdateCampaign(pl.baseData.campaign)
	if err != nil {
		return nil, errors.UpdateCampaignErr
	}

	return nil, nil
}
