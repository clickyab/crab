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
)

//@Validate {
//}
type createInventoryPayload struct {
	Label           string          `json:"label"  validate:"gt=7"`
	PubIDs          []int64         `json:"pub_ids"  validate:"required"`
	validPublishers []orm.Publisher `json:"-"`
}

func (pl *createInventoryPayload) ValidateExtra(ctx context.Context, w http.ResponseWriter, r *http.Request) error {

	if len(pl.PubIDs) > maxPubInventoryCount.Int() {
		return errors.MaxPubLimit(maxPubInventoryCount.Int())
	}

	if len(pl.PubIDs) == 0 {
		return errors.EmptyPublisherSelectedErr
	}
	validPublishers := orm.NewOrmManager().GetValidPubs(pl.PubIDs)

	if len(validPublishers) == 0 {
		return errors.EmptyPublisherSelectedErr
	}
	pl.validPublishers = validPublishers
	return nil
}

// createPreset get a new whitelist blacklist for user
// @Rest {
// 		url = /create
//		method = post
//		protected = true
//		resource = create_inventory:self
// }
func (ctrl *Controller) createPreset(ctx context.Context, r *http.Request, pl *createInventoryPayload) (*orm.Inventory, error) {
	currentUser := authz.MustGetUser(ctx)
	dm := domain.MustGetDomain(ctx)
	invManager := orm.NewOrmManager()

	var validPubIDs []int64

	for j := range pl.validPublishers {
		validPubIDs = append(validPubIDs, pl.validPublishers[j].ID)
	}

	res, e := invManager.CreateInventoryComplete(pl.Label, validPubIDs, dm.ID, currentUser.ID)
	assert.Nil(e)
	return res, nil
}

//@Validate {
//}
type duplicateInventoryPayload struct {
	ID int64 `json:"id"  validate:"required"`
}

// duplicate a new whitelist blacklist for user
// @Rest {
// 		url = /duplicate
//		method = post
//		protected = true
//		resource = duplicate_inventory:self
// }
func (ctrl *Controller) duplicate(ctx context.Context, r *http.Request, pl *duplicateInventoryPayload) (*orm.Inventory, error) {
	currentUser := authz.MustGetUser(ctx)
	dm := domain.MustGetDomain(ctx)

	m := orm.NewOrmManager()
	a, err := m.FindInventoryByIDDomain(pl.ID, dm.ID)

	if err != nil {
		return nil, errors.InvalidIDErr
	}

	owner, err := aaa.NewAaaManager().FindUserWithParentsByID(a.ID, dm.ID)
	if err != nil {
		return nil, errors.InvalidIDErr
	}
	_, ok := currentUser.HasOn("duplicate_inventory", owner.ID, dm.ID, false, false)
	if !ok {
		return nil, errors.AccessDeniedErr
	}

	res, e := m.Duplicate(pl.ID)
	assert.Nil(e)
	return res, nil
}
