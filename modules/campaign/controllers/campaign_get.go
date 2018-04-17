package controllers

import (
	"context"
	"database/sql"
	"net/http"
	"strconv"

	"clickyab.com/crab/modules/campaign/errors"
	"clickyab.com/crab/modules/campaign/orm"
	"clickyab.com/crab/modules/domain/middleware/domain"
	inventoryOrm "clickyab.com/crab/modules/inventory/orm"
	"clickyab.com/crab/modules/user/aaa"
	userErrors "clickyab.com/crab/modules/user/errors"
	"clickyab.com/crab/modules/user/middleware/authz"
	"github.com/rs/xmux"
)

type response struct {
	orm.Campaign
	Inventory  inventoryOrm.InventoryWithPubCount `json:"inventory"`
	Schedule   orm.ScheduleSheet                  `json:"schedule"`
	Attributes orm.CampaignAttributes             `json:"attributes"`
}

// get gets a campaign by id
// @Rest {
// 		url = /get/:id
//		protected = true
// 		method = get
//		resource = get_campaign:self
// }
func (c *Controller) get(ctx context.Context, r *http.Request) (*response, error) {
	userDomain := domain.MustGetDomain(ctx)
	currentUser := authz.MustGetUser(ctx)
	id := xmux.Param(ctx, "id")
	campID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return nil, errors.InvalidIDErr
	}

	campaign, err := orm.NewOrmManager().FindCampaignByIDDomain(campID, userDomain.ID)
	if err != nil {
		return nil, errors.NotFoundError(campID)
	}

	owner, err := aaa.NewAaaManager().FindUserWithParentsByID(campaign.UserID, userDomain.ID)
	if err != nil {
		return nil, userErrors.NotFoundWithDomainError(userDomain.Name)
	}

	// check access
	_, ok := aaa.CheckPermOn(owner, currentUser, "get_campaign", userDomain.ID)
	if !ok {
		return nil, errors.AccessDenied
	}

	res := response{
		Campaign: *campaign,
	}

	dbm := orm.NewOrmManager()

	sc, err := dbm.GetSchedule(campaign.ID)
	if err != nil {
		return nil, errors.NotFoundSchedule
	}
	if sc != nil {
		res.Schedule = sc.ScheduleSheet
	}

	attrs, err := dbm.FindCampaignAttributesByCampaignID(campaign.ID)
	if err != nil && err != sql.ErrNoRows {
		return nil, errors.NotFoundAttributes
	}
	if attrs != nil {
		res.Attributes = *attrs
	}

	if campaign.InventoryID.Valid {
		invDBM := inventoryOrm.NewOrmManager()
		inv, err := invDBM.FindInventoryAndPubCount(campaign.InventoryID.Int64)
		if err != nil {
			return nil, errors.InventoryNotFound
		}

		res.Inventory = inv
	}

	return &res, nil
}
