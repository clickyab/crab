package controllers

import (
	"context"
	"fmt"
	"net/http"
	"strconv"
	"strings"

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

func (pl *addInventoryPayload) ValidateExtra(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	id, err := strconv.ParseInt(xmux.Param(ctx, "id"), 10, 64)
	if err != nil {
		return errors.InvalidIDErr
	}
	if len(pl.PubIDs) == 0 {
		return errors.EmptyPublisherSelectedErr
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

	bind := strings.TrimRight(strings.Repeat("?,", len(pl.PubIDs)), ",")
	validPublishers := invManager.ListPublishersWithFilter(fmt.Sprintf("id IN (%s)", bind),
		func(PubIDs []int64) []interface{} {
			var res = make([]interface{}, len(PubIDs))
			for i := range PubIDs {
				res[i] = PubIDs[i]
			}
			return res
		}(pl.PubIDs)...,
	)

	if len(validPublishers) == 0 {
		return nil, errors.EmptyPublisherSelectedErr
	}

	var validPubIDs []int64

	for j := range validPublishers {
		validPubIDs = append(validPubIDs, validPublishers[j].ID)
	}

	res, e := invManager.AddInventoryPub(pl.currentInventory, pl.currentInventoryPubs, validPubIDs)
	assert.Nil(e)
	return res, nil
}
