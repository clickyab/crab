package controllers

import (
	"context"
	"database/sql"
	"net/http"

	"clickyab.com/crab/modules/campaign/errors"
	"clickyab.com/crab/modules/campaign/orm"
	inventoryOrm "clickyab.com/crab/modules/inventory/orm"
	"clickyab.com/crab/modules/user/aaa"
	"clickyab.com/crab/modules/user/middleware/authz"
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
	currentUser := authz.MustGetUser(ctx)
	_, ok := aaa.CheckPermOn(baseData.owner, currentUser, "edit_campaign", baseData.domain.ID)
	if !ok {
		return nil, errors.AccessDenied
	}

	baseData.campaign.Progress = orm.ProgressFinalized
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
