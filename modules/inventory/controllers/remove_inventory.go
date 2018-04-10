package controllers

import (
	"context"
	"net/http"
	"strconv"

	"clickyab.com/crab/modules/domain/middleware/domain"
	"clickyab.com/crab/modules/inventory/errors"
	"clickyab.com/crab/modules/inventory/orm"
	"clickyab.com/crab/modules/user/aaa"
	"clickyab.com/crab/modules/user/middleware/authz"
	"github.com/clickyab/services/assert"
	"github.com/rs/xmux"
)

//@Validate {
//}
type removeInventoryPayload struct {
	PubIDs []int64 `json:"pub_ids" db:"domains" validate:"required"`
}

// removePreset edit inventory
// @Rest {
// 		url = /removepub/:id
//		method = patch
//		protected = true
//		resource = edit_inventory:self
// }
func (ctrl *Controller) removePreset(ctx context.Context, r *http.Request, pl *removeInventoryPayload) (*orm.Inventory, error) {
	id, err := strconv.ParseInt(xmux.Param(ctx, "id"), 10, 64)
	if err != nil {
		return nil, errors.InvalidIDErr
	}
	currentUser := authz.MustGetUser(ctx)
	dm := domain.MustGetDomain(ctx)
	invManager := orm.NewOrmManager()
	currentInventory, err := invManager.FindInventoryByID(id)
	if err != nil || currentInventory.DomainID != dm.ID {
		return nil, errors.NotFoundError(currentInventory.ID)
	}
	owner, err := aaa.NewAaaManager().FindUserWithParentsByID(currentInventory.UserID, dm.ID)
	assert.Nil(err)
	_, ok := aaa.CheckPermOn(owner, currentUser, "edit_inventory", dm.ID)
	if !ok {
		return nil, errors.AccessDeniedErr
	}

	res, e := invManager.RemoveInventoryPub(currentInventory, pl.PubIDs)
	assert.Nil(e)
	return res, nil
}
