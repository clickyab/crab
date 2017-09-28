package controllers

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	asset "clickyab.com/crab/modules/asset/models"
	"clickyab.com/crab/modules/campaign/orm"
	"github.com/clickyab/services/array"
	"github.com/clickyab/services/assert"
	"github.com/clickyab/services/random"
	"github.com/rs/xmux"
	"github.com/sirupsen/logrus"
)

// @Validate{
//}
type attributesPayload struct {
	orm.CampaignAttributes
}

func (l *attributesPayload) ValidateExtra(ctx context.Context, w http.ResponseWriter, r *http.Request) error {

	queryGen := func(t string, s []string) string {
		m := len(s)
		return fmt.Sprintf(`SELECT count(id) AS total FROM %s WHERE name IN (%s) AND active = 1`, t, strings.Repeat("?,", m)[:2*m-1])
	}

	if array.StringInArray(orm.Foreign, l.Region...) && len(l.Region) > 1 {
		return errors.New("region is not valid")
	}
	o := asset.NewModelsManager()

	if t, err := o.GetRDbMap().SelectInt(queryGen(asset.ISPTableFull, l.ISP), l.ISP); err != nil && int64(len(l.ISP)) != t {
		return errors.New("isp is not valid")
	}
	if t, err := o.GetRDbMap().SelectInt(queryGen(asset.OSTableFull, l.OS), l.OS); err != nil && int64(len(l.OS)) != t {
		return errors.New("os is not valid")
	}
	if t, err := o.GetRDbMap().SelectInt(queryGen(asset.BrowserTableFull, l.Browser), l.Browser); err != nil && int64(len(l.Browser)) != t {
		return errors.New("browsers is not valid")
	}
	if t, err := o.GetRDbMap().SelectInt(queryGen(asset.CategoryTableFull, l.IAB), l.IAB); err != nil && int64(len(l.IAB)) != t {
		return errors.New("iab is not valid")
	}
	if t, err := o.GetRDbMap().SelectInt(queryGen(asset.ManufacturerTableFull, l.Manufacturer), l.Manufacturer); err != nil && int64(len(l.Manufacturer)) != t {
		return errors.New("manufacturer is not valid")
	}
	// TODO: Validate other fields

	return nil
}

// attributes will update campaign attribute
// @Route {
// 		url = /attributes/:id
//		method = put
//		payload = attributesPayload
//		200 = orm.Campaign
//		400 = controller.ErrorResponseSimple
//		404 = controller.ErrorResponseSimple
//		middleware = authz.Authenticate
// }
func (c *Controller) attributes(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(xmux.Param(ctx, "id"), 10, 64)
	p := c.MustGetPayload(ctx).(*attributesPayload)

	if err != nil || id < 1 {
		c.BadResponse(w, errors.New("id is not valid"))
	}
	db := orm.NewOrmManager()
	o, err := db.FindCampaignByID(id)
	if err != nil {
		c.NotFoundResponse(w, nil)
	}

	err = db.UpdateAttribute(p.CampaignAttributes, o)

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
		c.BadResponse(w, errors.New("can not update attributes"))
		return
	}
	c.OKResponse(w, o)
}
