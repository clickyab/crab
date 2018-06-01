package controllers

import (
	"context"
	"net/http"
	"strconv"

	cManager "clickyab.com/crab/modules/campaign/orm"
	"clickyab.com/crab/modules/domain/middleware/domain"
	dManager "clickyab.com/crab/modules/domain/orm"
	"clickyab.com/crab/modules/inventory/errors"
	"clickyab.com/crab/modules/inventory/orm"
	"clickyab.com/crab/modules/user/middleware/authz"
	"github.com/clickyab/services/assert"
	"github.com/rs/xmux"
)

//@Validate {
//}
type changeStatusPayload struct {
	Status           orm.InventoryStatus `json:"status" validate:"required"`
	currentInventory *orm.Inventory      `json:"-"`
	currentDomain    *dManager.Domain    `json:"-"`
}

func (pl *changeStatusPayload) ValidateExtra(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	if !pl.Status.IsValid() {
		return errors.InvalidInventoryStatusErr
	}
	currentDomain := domain.MustGetDomain(ctx)
	pl.currentDomain = currentDomain
	id, err := strconv.ParseInt(xmux.Param(ctx, "id"), 10, 64)
	if err != nil {
		return errors.InvalidIDErr
	}
	inventory, err := orm.NewOrmManager().FindInventoryByIDDomain(id, currentDomain.ID)
	if err != nil {
		return errors.NotFoundError(id)
	}
	pl.currentInventory = inventory
	// get all start status campaign related to this inventory
	campaigns := cManager.NewOrmManager().ListCampaignsWithFilter("inventory_id=? AND status=? AND domain_id=?", inventory.ID, cManager.StartStatus, currentDomain.ID)
	if len(campaigns) > 0 {
		return errors.InventoryHasStartCampErr
	}
	return nil
}

// inventoryChangeStatus change inventory status
// @Rest {
// 		url = /inventory/:id
//		method = patch
//		protected = true
//		resource = edit_inventory:self
// }
func (ctrl *Controller) inventoryChangeStatus(ctx context.Context, r *http.Request, pl *changeStatusPayload) (*orm.Inventory, error) {
	currentUser := authz.MustGetUser(ctx)
	_, ok := currentUser.HasOn("edit_inventory", pl.currentInventory.UserID, pl.currentDomain.ID, false, false)
	if !ok {
		return nil, errors.AccessDeniedErr
	}
	pl.currentInventory.Status = pl.Status
	assert.Nil(orm.NewOrmManager().UpdateInventory(pl.currentInventory))
	return pl.currentInventory, nil

}
