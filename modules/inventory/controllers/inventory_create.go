package controllers

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"clickyab.com/crab/modules/domain/middleware/domain"
	"clickyab.com/crab/modules/inventory/errors"
	"clickyab.com/crab/modules/inventory/orm"
	"clickyab.com/crab/modules/user/aaa"
	"clickyab.com/crab/modules/user/middleware/authz"
	"github.com/clickyab/services/assert"
)

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
	invManager := orm.NewOrmManager()

	bind := strings.TrimRight(strings.Repeat("?,", len(pl.PubIDs)), ",")
	validPublishers := invManager.ListPublishersWithFilter(fmt.Sprintf("id IN (%s)", bind),
		func() []interface{} {
			var res = make([]interface{}, len(pl.PubIDs))
			for i := range pl.PubIDs {
				res[i] = pl.PubIDs[i]
			}
			return res
		}()...,
	)

	if len(validPublishers) == 0 {
		return nil, errors.EmptyPublisherSelectedErr
	}

	var validPubIDs []int64

	for j := range validPublishers {
		validPubIDs = append(validPubIDs, validPublishers[j].ID)
	}

	res, e := invManager.CreateInventoryComplete(pl.Label, validPubIDs, dm.ID, currentUser.ID)
	assert.Nil(e)
	return res, nil
}
