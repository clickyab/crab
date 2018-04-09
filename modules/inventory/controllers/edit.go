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
type addInventoryPayload struct {
	PubIDs               []int64                  `json:"pub_ids" validate:"required"`
	currentInventory     *orm.Inventory           `json:"-"`
	currentInventoryPubs []orm.InventoryPublisher `json:"-"`
}

//@Validate {
//}
type removeInventoryPayload struct {
	PubIDs []int64 `json:"pub_ids" db:"domains" validate:"required"`
}

func (pl *addInventoryPayload) ValidateExtra(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	id, err := strconv.ParseInt(xmux.Param(ctx, "id"), 10, 64)
	if err != nil {
		return errors.InvalidIDErr
	}
	currentInventory, err := orm.NewOrmManager().FindInventoryByID(id)
	if err != nil {
		return errors.NotFoundError(id)
	}
	pl.currentInventory = currentInventory
	currentInventoryPubs := orm.NewOrmManager().ListInventoryPublishersWithFilter("inventory_id=?", pl.currentInventory.ID)
	pl.currentInventoryPubs = currentInventoryPubs
	if len(pl.PubIDs)+len(pl.currentInventoryPubs) > maxPubInventoryCount.Int() {
		return errors.MaxPubLimit(maxPubInventoryCount.Int())
	}
	return nil
}

// addPreset edit inventory
// @Rest {
// 		url = /addpub/:id
//		method = patch
//		protected = true
//		resource = edit_inventory:self
// }
func (ctrl *Controller) addPreset(ctx context.Context, r *http.Request, pl *addInventoryPayload) (*orm.Inventory, error) {
	currentUser := authz.MustGetUser(ctx)
	dm := domain.MustGetDomain(ctx)
	invManager := orm.NewOrmManager()

	if pl.currentInventory.DomainID != dm.ID {
		return nil, errors.NotFoundError(pl.currentInventory.ID)
	}
	owner, err := aaa.NewAaaManager().FindUserWithParentsByID(pl.currentInventory.UserID, dm.ID)
	assert.Nil(err)
	_, ok := aaa.CheckPermOn(owner, currentUser, "edit_inventory", dm.ID)
	if !ok {
		return nil, errors.AccessDeniedErr
	}

	res, e := invManager.AddInventoryComplete(pl.currentInventory, pl.currentInventoryPubs, pl.PubIDs)
	assert.Nil(e)
	return res, nil
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

	currentInventoryPubs := invManager.ListInventoryPublishersWithFilter("inventory_id=?", currentInventory.ID)

	res, e := invManager.RemoveInventoryPub(currentInventory, currentInventoryPubs, pl.PubIDs)
	assert.Nil(e)
	return res, nil
}

//@Validate {
//}
type changeLabelPayload struct {
	Label string `json:"label" db:"label" validate:"gt=7"`
}

// removePreset edit inventory
// @Rest {
// 		url = /:id
//		method = put
//		protected = true
//		resource = edit_inventory:self
// }
func (ctrl *Controller) changeLabel(ctx context.Context, r *http.Request, pl *changeLabelPayload) (*orm.Inventory, error) {
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

	res, e := invManager.ChangeLabel(currentInventory, pl.Label)
	assert.Nil(e)
	return res, nil
}
