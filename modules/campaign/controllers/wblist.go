package controllers

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"clickyab.com/crab/modules/campaign/models"
	"github.com/clickyab/services/assert"
	"github.com/clickyab/services/random"
	"github.com/rs/xmux"
	"github.com/sirupsen/logrus"
)

// @Validate{
//}
type whiteBlackPayload struct {
	ListID int64 `json:"list_id" db:"-" validate:"required,min=1"`
}

// updateWhiteBlackList will update campaign white/black list
// @Route {
// 		url = /whiteblack/:id
//		method = put
//		payload = whiteBlackPayload
//		200 = models.Campaign
//		400 = controller.ErrorResponseSimple
//		404 = controller.ErrorResponseSimple
//		middleware = authz.Authenticate
// }
func (c *Controller) updateWhiteBlackList(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(xmux.Param(ctx, "id"), 10, 64)
	if err != nil {
		c.BadResponse(w, errors.New("id is not valid"))
	}
	p := c.MustGetPayload(ctx).(*whiteBlackPayload)

	if err != nil || id < 1 {
		c.BadResponse(w, errors.New("id is not valid"))
	}
	db := models.NewModelsManager()
	o, err := db.FindCampaignByID(id)
	if err != nil {
		c.NotFoundResponse(w, nil)
	}

	err = db.UpdateCampaignWhiteBlackList(p.ListID, o)
	if err == models.ErrInventoryID {
		c.BadResponse(w, err)
		return
	}
	if err != nil {
		j, e := json.MarshalIndent(o, "", "\t")
		assert.Nil(e)
		pj, e := json.MarshalIndent(p, "", "\t")
		assert.Nil(e)

		eid := <-random.ID
		logrus.WithField("error", err).
			WithField("payload", string(pj)).
			WithField("eid", eid).
			WithField("campaign", string(j)).
			Debug("update base campaign ")
		w.Header().Set("x-error-id", eid)
		c.BadResponse(w, errors.New("can not update white/black list"))
		return
	}
	c.OKResponse(w, o)
}

// deleteWhiteBlackList will update campaign white/black list
// @Route {
// 		url = /whiteblack/:id
//		method = delete
//		200 = controller.NormalResponse
//		400 = controller.ErrorResponseSimple
//		404 = controller.ErrorResponseSimple
//		middleware = authz.Authenticate
// }
func (c *Controller) deleteWhiteBlackList(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(xmux.Param(ctx, "id"), 10, 64)
	if err != nil {
		c.BadResponse(w, errors.New("id is not valid"))
	}
	db := models.NewModelsManager()
	o, err := db.FindCampaignByID(id)
	if err != nil {
		c.NotFoundResponse(w, nil)
	}

	err = db.DeleteCampaignWhiteBlackList(o)
	if err == models.ErrInventoryID {
		c.BadResponse(w, err)
		return
	}
	if err != nil {
		j, e := json.MarshalIndent(o, "", "\t")
		assert.Nil(e)

		eid := <-random.ID
		logrus.WithField("error", err).
			WithField("eid", eid).
			WithField("campaign", string(j)).
			Debug("update base campaign ")
		w.Header().Set("x-error-id", eid)
		c.BadResponse(w, errors.New("can not delete white/black list"))
		return
	}
	c.OKResponse(w, o)
}
