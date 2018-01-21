package controllers

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"

	"clickyab.com/crab/modules/campaign/orm"
	"github.com/clickyab/services/assert"
	"github.com/clickyab/services/gettext/t9e"
	"github.com/clickyab/services/random"
	"github.com/rs/xmux"
	"github.com/sirupsen/logrus"
)

// @Validate{
//}
type whiteBlackPayload struct {
	ListID   int64 `json:"list_id"`
	WhiteTyp *bool `json:"white_typ" validate:"required"` //true for white false for black
	Exchange *bool `json:"exchange" validate:"required"`
}

// updateWhiteBlackList will update campaign white/black list
// @Route {
// 		url = /wb/:id
//		method = put
//		payload = whiteBlackPayload
//		200 = orm.Campaign
//		400 = controller.ErrorResponseSimple
//		404 = controller.ErrorResponseSimple
//		middleware = authz.Authenticate
// }
func (c *Controller) updateWhiteBlackList(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(xmux.Param(ctx, "id"), 10, 64)
	if err != nil {
		c.BadResponse(w, t9e.G("id is not valid"))
	}
	p := c.MustGetPayload(ctx).(*whiteBlackPayload)

	if err != nil || id < 1 {
		c.BadResponse(w, t9e.G("id is not valid"))
	}
	db := orm.NewOrmManager()
	o, err := db.FindCampaignByID(id)
	if err != nil {
		c.NotFoundResponse(w, nil)
	}

	err = db.UpdateCampaignWhiteBlackList(p.ListID, p.Exchange, p.WhiteTyp, o)
	if err == orm.ErrInventoryID {
		c.BadResponse(w, err)
		return
	}
	if err != nil {
		j, e := json.MarshalIndent(o, " ", "  ")
		assert.Nil(e)
		pj, e := json.MarshalIndent(p, " ", "  ")
		assert.Nil(e)

		eid := <-random.ID
		logrus.WithField("error", err).
			WithField("payload", string(pj)).
			WithField("eid", eid).
			WithField("campaign", string(j)).
			Debug("update base campaign ")
		w.Header().Set("x-error-id", eid)
		c.BadResponse(w, t9e.G("can not update white/black list"))
		return
	}
	c.OKResponse(w, o)
}

// deleteWhiteBlackList will update campaign white/black list
// @Route {
// 		url = /wblist/:id
//		method = delete
// 		200 = orm.Campaign
//		400 = controller.ErrorResponseSimple
//		404 = controller.ErrorResponseSimple
//		middleware = authz.Authenticate
// }
func (c *Controller) deleteWhiteBlackList(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(xmux.Param(ctx, "id"), 10, 64)
	if err != nil {
		c.BadResponse(w, t9e.G("id is not valid"))
	}
	db := orm.NewOrmManager()
	o, err := db.FindCampaignByID(id)
	if err != nil {
		c.NotFoundResponse(w, nil)
	}

	err = db.DeleteCampaignWhiteBlackList(o)
	if err == orm.ErrInventoryID {
		c.BadResponse(w, err)
		return
	}
	if err != nil {
		j, e := json.MarshalIndent(o, " ", "  ")
		assert.Nil(e)

		eid := <-random.ID
		logrus.WithField("error", err).
			WithField("eid", eid).
			WithField("campaign", string(j)).
			Debug("update base campaign ")
		w.Header().Set("x-error-id", eid)
		c.BadResponse(w, t9e.G("can not delete white/black list"))
		return
	}
	c.OKResponse(w, o)
}
