package controllers

import (
	"context"
	"net/http"
	"strconv"

	"clickyab.com/crab/modules/domain/middleware/domain"
	dmn "clickyab.com/crab/modules/domain/orm"
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
	PubIDs               []int64                  `json:"pub_ids" db:"domains" validate:"required"`
	currentInventoryPubs []orm.InventoryPublisher `json:"-"`
	currentInventory     *orm.Inventory           `json:"-"`
	currentDomain        *dmn.Domain              `json:"-"`
}

func (pl *removeInventoryPayload) ValidateExtra(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	id, err := strconv.ParseInt(xmux.Param(ctx, "id"), 10, 64)
	if err != nil {
		return errors.InvalidIDErr
	}
	dm := domain.MustGetDomain(ctx)
	pl.currentDomain = dm
	invManager := orm.NewOrmManager()
	currentInventory, err := invManager.FindInventoryByID(id)
	if err != nil || currentInventory.DomainID != dm.ID {
		return errors.NotFoundError(currentInventory.ID)
	}
	currentInventoryPubs := orm.NewOrmManager().ListInventoryPublishersWithFilter("inventory_id=?", currentInventory.ID)
	pl.currentInventoryPubs = currentInventoryPubs
	pl.currentInventory = currentInventory
	return nil
}

// removePreset edit inventory
// @Rest {
// 		url = /removepub/:id
//		method = patch
//		protected = true
//		resource = edit_inventory:self
// }
func (ctrl *Controller) removePreset(ctx context.Context, r *http.Request, pl *removeInventoryPayload) (*orm.Inventory, error) {

	currentUser := authz.MustGetUser(ctx)

	owner, err := aaa.NewAaaManager().FindUserWithParentsByID(pl.currentInventory.UserID, pl.currentDomain.ID)
	assert.Nil(err)
	_, ok := currentUser.HasOn("edit_inventory", owner.ID, pl.currentDomain.ID, false, false)
	if !ok {
		return nil, errors.AccessDeniedErr
	}

	oldPubIDs := func() []int64 {
		var oldIDs = make([]int64, len(pl.currentInventoryPubs))
		for i := range pl.currentInventoryPubs {
			oldIDs[i] = pl.currentInventoryPubs[i].PublisherID
		}
		return oldIDs
	}()

	res, e := orm.NewOrmManager().RemoveInventoryPub(pl.currentInventory, pl.PubIDs, oldPubIDs)
	assert.Nil(e)
	return res, nil
}
