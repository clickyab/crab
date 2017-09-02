package controllers

import (
	"context"
	"net/http"

	"errors"
	"fmt"
	"strconv"

	"clickyab.com/crab/modules/inventory/orm"
	"clickyab.com/crab/modules/user/middleware/authz"
	"github.com/clickyab/services/assert"
	"github.com/clickyab/services/framework/controller"
	"github.com/clickyab/services/trans"
	"github.com/rs/xmux"
)

// Controller is the controller for the location package
// @Route {
// 		middleware = domain.Access
//		group = /inventory
// }
type Controller struct {
	controller.Base
}

type inventories []orm.Inventory

// Inventories return all user inventories
// @Route {
// 		url = /:user
//		method = get
//		200 = inventories
//		middleware = authz.Authenticate
// }
func (ctrl *Controller) Inventories(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	u := authz.MustGetUser(ctx)
	res, e := orm.NewOrmManager().FindInventoryByUserID(u.ID)
	assert.Nil(e)
	if res == nil {
		ctrl.NotFoundResponse(w, trans.E("User doesn't have any list"))
		return
	}
	ctrl.OKResponse(w, res)
}

// Inventory return a user inventory
// @Route {
// 		url = /:id
//		method = get
//		200 = orm.Inventory
//		middleware = authz.Authenticate
// }
func (ctrl *Controller) Inventory(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	id, e := strconv.ParseInt(xmux.Param(ctx, "id"), 10, 64)
	if e != nil {
		ctrl.BadResponse(w, errors.New(fmt.Sprintf("not valid id")))
		return
	}
	res, e := orm.NewOrmManager().FindInventoryByID(id)
	assert.Nil(e)
	if res == nil {
		ctrl.NotFoundResponse(w, trans.E("Inventory with id %s does not exists!", id))
		return
	}
	ctrl.OKResponse(w, res)
}

