package controllers

import (
	"context"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	asset "clickyab.com/crab/modules/asset/orm"
	"clickyab.com/crab/modules/campaign/errors"
	"clickyab.com/crab/modules/campaign/orm"
	"clickyab.com/crab/modules/location/location"
	"clickyab.com/crab/modules/user/aaa"
	"clickyab.com/crab/modules/user/middleware/authz"
	"github.com/clickyab/services/array"
	"github.com/clickyab/services/gettext/t9e"
	"github.com/rs/xmux"
	"github.com/sirupsen/logrus"
)

// @Validate{
//}
type attributesPayload struct {
	orm.CampaignAttributes
	campaign *orm.Campaign `json:"-"`
}

func (l *attributesPayload) ValidateExtra(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	queryGen := func(t string, s []string) string {
		m := len(s)
		return fmt.Sprintf(`select count(name) as total from %s where name in (%s)`, t, strings.Repeat("?,", m)[:2*m-1])
	}

	if array.StringInArray(orm.Foreign, l.Region...) && len(l.Region) > 1 {
		return errors.InvalidError("region")
	}
	o := asset.NewOrmManager()

	// TODO: add other validation field
	stringArrays := map[string][]string{
		asset.ISPTableFull:          l.ISP,
		asset.OSTableFull:           l.OS,
		asset.BrowserTableFull:      l.Browser,
		asset.CategoryTableFull:     l.IAB,
		asset.ManufacturerTableFull: l.Manufacturer,
		asset.PlatformTableFull:     l.Device,
		location.ProvinceTableFull:  l.Region,
	}
	for i := range stringArrays {
		if len(stringArrays[i]) == 0 {
			delete(stringArrays, i)
		}
	}

	campaignIDInt, err := strconv.ParseInt(xmux.Param(ctx, "id"), 10, 64)
	if err != nil {
		return t9e.G("campaign id not valid")
	}
	db := orm.NewOrmManager()
	campaign, err := db.FindCampaignByID(campaignIDInt)
	if err != nil {
		return t9e.G("campaign not found")
	}
	l.campaign = campaign

	// TODO : extra validation based on campaign kind

	for k, v := range stringArrays {
		values := make([]interface{}, 0)
		for i := range v {
			values = append(values, v[i])
		}

		if t, err := o.GetRDbMap().SelectInt(queryGen(k, v), values...); err != nil || int64(len(v)) != t {
			logrus.Warn(err)
			return errors.InvalidError(k)
		}
	}

	return nil
}

// attributes will update campaign attribute
// @Rest {
// 		url = /attributes/:id
//		protected = true
// 		method = put
// }
func (c *Controller) attributes(ctx context.Context, r *http.Request, p *attributesPayload) (*orm.Campaign, error) {
	id, err := strconv.ParseInt(xmux.Param(ctx, "id"), 10, 64)
	if err != nil || id < 1 {
		return nil, errors.InvalidIDErr
	}

	db := orm.NewOrmManager()
	ca, err := db.FindCampaignByID(id)
	if err != nil {
		return ca, errors.NotFoundError(id)
	}

	// check access
	userManager := aaa.NewAaaManager()
	owner, err := userManager.FindUserWithParentsByID(p.campaign.UserID, id)
	if err != nil {
		return ca, t9e.G("can't find user with related domain")
	}

	currentUser := authz.MustGetUser(ctx)
	_, ok := aaa.CheckPermOn(owner, currentUser, "edit_attributes", id)
	if !ok {
		return ca, errors.AccessDenied
	}

	err = db.UpdateAttribute(p.CampaignAttributes, p.campaign)
	if err != nil {
		return ca, t9e.G("can't update campaign attributes")
	}

	return ca, nil
}
