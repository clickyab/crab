package controllers

import (
	"context"
	"database/sql"
	"net/http"

	"clickyab.com/crab/modules/campaign/errors"
	"clickyab.com/crab/modules/campaign/orm"
	inventoryOrm "clickyab.com/crab/modules/inventory/orm"
	"clickyab.com/crab/modules/user/middleware/authz"
	"github.com/fatih/structs"
)

type finalizeResult struct {
	orm.Campaign
	inventory inventoryOrm.InventoryWithPubCount `json:"inventory"`
}

// finalize
// @Rest {
// 		url = /finalize/:id
//		protected = true
// 		method = put
//		resource = edit_campaign:self
// }
func (c *Controller) finalize(ctx context.Context, r *http.Request) (*finalizeResult, error) {
	token := authz.MustGetToken(ctx)
	baseData, err := CheckUserCampaignDomain(ctx)
	if err != nil {
		return nil, err
	}

	db := orm.NewOrmManager()
	// check access
	uScope, ok := baseData.currentUser.HasOn("edit_campaign", baseData.owner.ID, baseData.domain.ID, false, false)
	if !ok {
		return nil, errors.AccessDenied
	}

	err = baseData.campaign.SetAuditUserData(baseData.currentUser.ID, token, baseData.domain.ID, "edit_campaign", uScope)
	if err != nil {
		return nil, err
	}

	baseData.campaign.Progress = orm.ProgressFinalized

	d := structs.Map(baseData.campaign)
	err = baseData.campaign.SetAuditDescribe(d, "finalize camapaign")
	if err != nil {
		return nil, err
	}

	err = db.UpdateCampaign(baseData.campaign)
	if err != nil {
		return nil, errors.UpdateError
	}

	invDBM := inventoryOrm.NewOrmManager()
	inv, err := invDBM.FindInventoryAndPubCount(baseData.campaign.ID)
	if err != nil {
		if err != sql.ErrNoRows {
			res := finalizeResult{
				Campaign: *baseData.campaign,
			}

			return &res, nil
		}

		return nil, errors.InventoryNotFound
	}

	res := finalizeResult{
		Campaign:  *baseData.campaign,
		inventory: inv,
	}

	return &res, nil
}
