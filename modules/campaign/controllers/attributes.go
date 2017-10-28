package controllers

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	asset "clickyab.com/crab/modules/asset/orm"
	"clickyab.com/crab/modules/campaign/orm"
	"github.com/clickyab/services/array"
	"github.com/clickyab/services/assert"
	"github.com/clickyab/services/mysql"
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
		return fmt.Sprintf(`select count(id) as total from %s where name in (%s)`, t, strings.Repeat("?,", m)[:2*m-1])
	}

	if array.StringInArray(orm.Foreign, l.Region...) && len(l.Region) > 1 {
		return errors.New("region is not valid")
	}
	o := asset.NewOrmManager()

	// TODO: add other validation field
	stringArrays := map[string]mysql.StringJSONArray{asset.ISPTableFull: l.ISP, asset.OSTableFull: l.OS, asset.BrowserTableFull: l.Browser, asset.CategoryTableFull: l.IAB, asset.ManufacturerTableFull: l.Manufacturer}
	for i := range stringArrays {
		if len(stringArrays[i]) == 0 {
			delete(stringArrays, i)
		}
	}

	for i := range stringArrays {
		if t, err := o.GetRDbMap().SelectInt(queryGen(i, stringArrays[i]), stringArrays[i]); err != nil && int64(len(stringArrays[i])) != t {
			return fmt.Errorf("%s is not valid", i)
		}
	}

	return nil
}

// attributes will update campaign attribute
// @Route {
// 		url = /attributes/:id
//		method = put
//		payload = attributesPayload
//		200 = campaignResponse
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
	c.OKResponse(w, createResponse(o))
}
