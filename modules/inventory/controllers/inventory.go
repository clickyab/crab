package controllers

import (
	"context"
	"net/http"

	"clickyab.com/crab/modules/domain/middleware/domain"
	"clickyab.com/crab/modules/inventory/errors"
	"clickyab.com/crab/modules/inventory/orm"
	"clickyab.com/crab/modules/user/aaa"
	"clickyab.com/crab/modules/user/middleware/authz"
	"github.com/clickyab/services/assert"
	"github.com/clickyab/services/config"
	"github.com/clickyab/services/framework/controller"
)

var maxPubInventoryCount = config.RegisterInt("crab.modules.inventory.pub.count", 250, "max publisher in single inventory")

// Controller is the controller for the location package
// @Route {
// 		middleware = domain.Access
//		group = /inventory
// }
type Controller struct {
	controller.Base
}

//@Validate {
//}
type createInventoryPayload struct {
	Label  string  `json:"label"  validate:"gt=7"`
	PubIDs []int64 `json:"pub_ids"  validate:"required"`
}

func (pl *createInventoryPayload) ValidateExtra(ctx context.Context, w http.ResponseWriter, r *http.Request) error {

	if len(pl.PubIDs) > maxPubInventoryCount.Int() {
		return errors.MaxPubLimit(maxPubInventoryCount.Int())
	}
	return nil
}

// createPreset get a new whitelist blacklist for user
// @Rest {
// 		url = /create
//		method = post
//		protected = true
//		resource = add_inventory:self
// }
func (ctrl *Controller) createPreset(ctx context.Context, r *http.Request, pl *createInventoryPayload) (*orm.Inventory, error) {
	currentUser := authz.MustGetUser(ctx)
	dm := domain.MustGetDomain(ctx)
	_, ok := aaa.CheckPermOn(currentUser, currentUser, "add_inventory", dm.ID)
	if !ok {
		return nil, errors.AccessDeniedErr
	}
	res, e := orm.NewOrmManager().CreateInventoryComplete(pl.Label, pl.PubIDs, dm.ID, currentUser.ID)
	assert.Nil(e)
	return res, nil
}
