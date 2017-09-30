package controllers

import (
	"context"
	"net/http"

	"strconv"

	"fmt"

	"clickyab.com/crab/modules/inventory/models"
	"clickyab.com/crab/modules/user/middleware/authz"
	"github.com/clickyab/services/assert"
	"github.com/clickyab/services/framework/controller"
	"github.com/clickyab/services/trans"
	"github.com/rs/xmux"
)

// Inventory is the controller for the location package
// @Route {
// 		middleware = domain.Access
//		group = /inventory
// }
type Inventory struct {
	controller.Base
}

// return all user presets for filtering publishers (black/white list)
// @Route {
// 		url = /presets
//		method = get
//		200 = models.Presets
//		404 = controller.ErrorResponseSimple
//		middleware = authz.Authenticate
// }
func (c *Inventory) presets(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	u := authz.MustGetUser(ctx)
	res := models.NewModelsManager().ListPresetsWithFilter(
		"user_id = ? ", u.ID)
	if len(res) == 0 {
		c.NotFoundResponse(w, trans.E("User doesn't have any list"))
		return
	}
	c.OKResponse(w, res)
}

// user preset
// @Route {
// 		url = /preset/:id
//		method = get
//		200 = models.Preset
//		404 = controller.ErrorResponseSimple
//		middleware = authz.Authenticate
// }
func (c *Inventory) preset(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	id, e := strconv.ParseInt(xmux.Param(ctx, "id"), 10, 64)
	if e != nil {
		c.BadResponse(w, trans.E("not valid id"))
		return
	}
	res, e := models.NewModelsManager().FindPresetByID(id)
	if e != nil {
		c.NotFoundResponse(w, fmt.Errorf("inventory with id %d does not exists", id))
		return
	}
	c.OKResponse(w, res)
}

//@validate {
//}
type presetsPayload struct {
	models.PresetData
}

// addPreset get a new whitelist blacklist for user
// @Route {
// 		url = /preset
//		method = post
//		200 = models.Preset
//		middleware = authz.Authenticate
//		payload = presetsPayload
// }
func (c *Inventory) addPreset(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	pl := c.MustGetPayload(ctx).(*presetsPayload)
	u := authz.MustGetUser(ctx)
	d := &models.Preset{
		Active:     true,
		PresetData: pl.PresetData,
		UserID:     u.ID,
	}
	e := models.NewModelsManager().CreatePreset(d)
	assert.Nil(e)
	c.OKResponse(w, d)
}
