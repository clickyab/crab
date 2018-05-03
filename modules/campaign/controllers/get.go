package controllers

import (
	"context"
	"database/sql"
	"net/http"
	"strconv"

	"github.com/clickyab/services/xlog"

	"clickyab.com/crab/modules/campaign/errors"
	"clickyab.com/crab/modules/campaign/orm"
	"clickyab.com/crab/modules/domain/middleware/domain"
	inventoryOrm "clickyab.com/crab/modules/inventory/orm"
	"clickyab.com/crab/modules/user/aaa"
	userErrors "clickyab.com/crab/modules/user/errors"
	"clickyab.com/crab/modules/user/middleware/authz"
	"github.com/rs/xmux"
)

type campaignGetResponse struct {
	orm.Campaign
	Receivers  []orm.Receiver                     `json:"receivers"`
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
func (c *Controller) get(ctx context.Context, r *http.Request) (*campaignGetResponse, error) {
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

	res := campaignGetResponse{
		Campaign: *campaign,
	}

	dbm := orm.NewOrmManager()

	sc, err := dbm.GetSchedule(campaign.ID)
	if err != nil {
		xlog.GetWithError(ctx, err).Debug("campaign get route: find scheduales error")
		return nil, errors.NotFoundSchedule
	}
	if sc != nil {
		res.Schedule = sc.ScheduleSheet
	}

	attrs, err := dbm.FindCampaignAttributesByCampaignID(campaign.ID)
	if err != nil && err != sql.ErrNoRows {
		xlog.GetWithError(ctx, err).Debug("campaign get route: find attributes error")
		return nil, errors.NotFoundAttributes
	}
	if attrs != nil {
		res.Attributes = *attrs
	}

	if campaign.InventoryID.Valid {
		invDBM := inventoryOrm.NewOrmManager()
		inv, err := invDBM.FindInventoryAndPubCount(campaign.InventoryID.Int64)
		if err != nil {
			xlog.GetWithError(ctx, err).Debug("campaign get route: find inventory error")
			return nil, errors.InventoryNotFound
		}

		res.Inventory = inv
	}

	recs := dbm.GetReportReceivers(campaign.ID)
	if len(recs) > 0 {
		res.Receivers = recs
	}

	return &res, nil
}
