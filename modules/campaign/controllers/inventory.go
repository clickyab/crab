package controllers

import (
	"context"
	"net/http"
	"strconv"

	"clickyab.com/crab/modules/campaign/errors"
	"clickyab.com/crab/modules/campaign/orm"
	"clickyab.com/crab/modules/domain/middleware/domain"
	dmn "clickyab.com/crab/modules/domain/orm"
	inv "clickyab.com/crab/modules/inventory/orm"
	"clickyab.com/crab/modules/user/aaa"
	"clickyab.com/crab/modules/user/middleware/authz"
	"github.com/clickyab/services/mysql"
	"github.com/rs/xmux"
)

// @Validate{
//}
type assignInventoryPayload struct {
	ID              int64              `json:"id"`
	InvState        orm.InventoryState `json:"state"`
	remove          bool               `json:"-"`
	currentCampaign *orm.Campaign      `json:"-"`
	targetInventory *inv.Inventory     `json:"-"`
	currentDomain   *dmn.Domain        `json:"-"`
	pubDomains      []string           `json:"-"`
}

func (pl *assignInventoryPayload) ValidateExtra(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	dm := domain.MustGetDomain(ctx)

	//find target campaign
	campaignIDInt, err := strconv.ParseInt(xmux.Param(ctx, "id"), 10, 64)
	if err != nil {
		return errors.InvalidIDErr
	}
	currentCampaign, err := orm.NewOrmManager().FindCampaignByIDDomain(campaignIDInt, dm.ID)
	if err != nil {
		return errors.InvalidIDErr
	}

	pl.currentCampaign = currentCampaign
	pl.currentDomain = dm

	if pl.ID == 0 {
		pl.remove = true
		return nil
	}

	// find target inventory
	invManager := inv.NewOrmManager()
	inventory := invManager.FindInventoryByIDDomain(pl.ID, dm.ID)
	if len(inventory) != 1 {
		return errors.InvalidIDErr
	}
	if !pl.InvState.IsValid() {
		return errors.InventoryStateErr
	}
	// find publishers for the specified inventory
	domains := invManager.FindInventoryDomainsByInvID(inventory[0].ID)

	pl.targetInventory = &inventory[0]

	pl.pubDomains = domains

	return nil
}

// assignInventory in campaign
// @Rest {
// 		url = /inventory/:id
//		protected = true
// 		method = put
//		resource = edit_campaign:self
// }
func (c Controller) assignInventory(ctx context.Context, r *http.Request, p *assignInventoryPayload) (*orm.Campaign, error) {
	currentUser := authz.MustGetUser(ctx)
	userManager := aaa.NewAaaManager()
	campaignOwner, err := userManager.FindUserWithParentsByID(p.currentCampaign.UserID, p.currentDomain.ID)
	if err != nil {
		return nil, err
	}

	// check perm for campaign entity
	_, ok := aaa.CheckPermOn(campaignOwner, currentUser, "edit_campaign", p.currentDomain.ID)
	if !ok {
		return nil, errors.AccessDenied
	}

	if p.remove {
		p.currentCampaign.InventoryID = mysql.NullInt64{Int64: 0, Valid: false}
		p.currentCampaign.InventoryType = orm.NullInventoryState{Valid: false, InventoryState: p.InvState}
		p.currentCampaign.InventoryDomains = p.pubDomains

		err = orm.NewOrmManager().UpdateCampaign(p.currentCampaign)
		if err != nil {
			return nil, errors.UpdateCampaignErr
		}
		return p.currentCampaign, nil
	}

	invOwner, err := userManager.FindUserWithParentsByID(p.targetInventory.UserID, p.currentDomain.ID)
	if err != nil {
		return nil, err
	}

	//check perm for inventory entity
	_, ok = aaa.CheckPermOn(invOwner, invOwner, "edit_inventory", p.currentDomain.ID)
	if !ok {
		return nil, errors.AccessDenied
	}

	p.currentCampaign.InventoryID = mysql.NullInt64{Int64: p.targetInventory.ID, Valid: true}
	p.currentCampaign.InventoryType = orm.NullInventoryState{Valid: true, InventoryState: p.InvState}
	p.currentCampaign.InventoryDomains = p.pubDomains

	err = orm.NewOrmManager().UpdateCampaign(p.currentCampaign)
	if err != nil {
		return nil, errors.UpdateCampaignErr
	}
	return p.currentCampaign, nil

}
