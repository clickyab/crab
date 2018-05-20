package controllers

import (
	"context"
	"net/http"

	"clickyab.com/crab/modules/campaign/errors"
	"clickyab.com/crab/modules/campaign/orm"
	inv "clickyab.com/crab/modules/inventory/orm"
	"clickyab.com/crab/modules/user/aaa"
	"clickyab.com/crab/modules/user/middleware/authz"
	"github.com/clickyab/services/mysql"
	"github.com/fatih/structs"
)

// @Validate{
//}
type assignInventoryPayload struct {
	ID              int64              `json:"id"`
	InvState        orm.InventoryState `json:"state"`
	remove          bool               `json:"-"`
	targetInventory *inv.Inventory     `json:"-"`
	pubDomains      []string           `json:"-"`
	baseData        *BaseData          `json:"-"`
}

func (pl *assignInventoryPayload) ValidateExtra(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	res, err := CheckUserCamapignDomain(ctx)
	if err != nil {
		return err
	}
	pl.baseData = res

	// find target inventory
	invManager := inv.NewOrmManager()
	inventory, err := invManager.FindInventoryByIDDomain(pl.ID, pl.baseData.domain.ID)
	if err != nil {
		return errors.InvalidIDErr
	}
	if !pl.InvState.IsValid() {
		return errors.InventoryStateErr
	}
	// find publishers for the specified inventory
	domains := invManager.FindInventoryDomainsByInvID(inventory.ID)

	pl.targetInventory = inventory

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
	token := authz.MustGetToken(ctx)
	// check perm for campaign entity
	uScope, ok := aaa.CheckPermOn(p.baseData.owner, p.baseData.currentUser, "edit_campaign", p.baseData.domain.ID)
	if !ok {
		return nil, errors.AccessDenied
	}

	err := p.baseData.campaign.SetAuditUserData(p.baseData.currentUser.ID, token, p.baseData.domain.ID, "edit_campaign", uScope)
	if err != nil {
		return nil, err
	}

	if p.remove {
		p.baseData.campaign.InventoryID = mysql.NullInt64{Int64: 0, Valid: false}
		p.baseData.campaign.InventoryType = orm.NullInventoryState{Valid: false, InventoryState: p.InvState}
		p.baseData.campaign.InventoryDomains = make([]string, 0)

		d := structs.Map(p.baseData.campaign)
		err := p.baseData.campaign.SetAuditDescribe(d, "remove inventory from campaign")
		if err != nil {
			return nil, err
		}

		err = orm.NewOrmManager().UpdateCampaign(p.baseData.campaign)
		if err != nil {
			return nil, errors.UpdateCampaignErr
		}
		return p.baseData.campaign, nil
	}

	userManager := aaa.NewAaaManager()
	invOwner, err := userManager.FindUserWithParentsByID(p.targetInventory.UserID, p.baseData.domain.ID)
	if err != nil {
		return nil, err
	}

	//check perm for inventory entity
	_, ok = aaa.CheckPermOn(invOwner, invOwner, "edit_inventory", p.baseData.domain.ID)
	if !ok {
		return nil, errors.AccessDenied
	}

	p.baseData.campaign.InventoryID = mysql.NullInt64{Int64: p.targetInventory.ID, Valid: true}
	p.baseData.campaign.InventoryType = orm.NullInventoryState{Valid: true, InventoryState: p.InvState}
	p.baseData.campaign.InventoryDomains = p.pubDomains

	d := structs.Map(p.baseData.campaign)
	err = p.baseData.campaign.SetAuditDescribe(d, "add inventory to campaign")
	if err != nil {
		return nil, err
	}

	err = orm.NewOrmManager().UpdateCampaign(p.baseData.campaign)
	if err != nil {
		return nil, errors.UpdateCampaignErr
	}
	return p.baseData.campaign, nil

}
