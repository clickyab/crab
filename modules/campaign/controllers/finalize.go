package controllers

import (
	"context"
	"database/sql"
	"net/http"

	"clickyab.com/crab/modules/campaign/errors"
	"clickyab.com/crab/modules/campaign/orm"
	inventoryOrm "clickyab.com/crab/modules/inventory/orm"
	"clickyab.com/crab/modules/user/aaa"
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
	baseData, err := CheckUserCamapignDomain(ctx)
	if err != nil {
		return nil, err
	}

	db := orm.NewOrmManager()
	// check access
	uScope, ok := aaa.CheckPermOn(baseData.owner, baseData.currentUser, "edit_campaign", baseData.domain.ID)
	if !ok {
		return nil, errors.AccessDenied
	}

	err = baseData.campaign.SetAuditUserData(baseData.currentUser.ID, false, 0, "edit_campaign", uScope)
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
