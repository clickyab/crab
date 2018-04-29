package controllers

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	asset "clickyab.com/crab/modules/asset/orm"
	"clickyab.com/crab/modules/campaign/errors"
	"clickyab.com/crab/modules/campaign/orm"
	"clickyab.com/crab/modules/location/location"
	"clickyab.com/crab/modules/user/aaa"
	"github.com/clickyab/services/array"
	"github.com/clickyab/services/mysql"
	"github.com/fatih/structs"
)

// @Validate{
//}
type attributesPayload struct {
	Device       mysql.StringJSONArray `json:"device" validate:"omitempty"`
	Manufacturer mysql.StringJSONArray `json:"manufacturer" validate:"omitempty"`
	OS           mysql.StringJSONArray `json:"os" validate:"omitempty"`
	Browser      mysql.StringJSONArray `json:"browser" validate:"omitempty"`
	IAB          mysql.StringJSONArray `json:"iab" validate:"omitempty"`
	Region       mysql.StringJSONArray `json:"region" validate:"omitempty"`
	Cellular     mysql.StringJSONArray `json:"cellular" validate:"omitempty"`
	ISP          mysql.StringJSONArray `json:"isp" validate:"omitempty"`
	baseData     *BaseData             `json:"-"`
}

func (l *attributesPayload) ValidateExtra(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	res, err := CheckUserCamapignDomain(ctx)
	if err != nil {
		return err
	}
	l.baseData = res

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

	// TODO : extra validation based on campaign kind

	for k, v := range stringArrays {
		values := make([]interface{}, 0)
		for i := range v {
			values = append(values, v[i])
		}

		if t, err := o.GetRDbMap().SelectInt(queryGen(k, v), values...); err != nil || int64(len(v)) != t {
			return errors.InvalidError(k)
		}
	}

	return nil
}

type attributesResult struct {
	orm.Campaign
	orm.CampaignAttributes
}

// attributes will update campaign attribute
// @Rest {
// 		url = /attributes/:id
//		protected = true
// 		method = put
//		resource = edit_attributes:self
// }
func (c *Controller) attributes(ctx context.Context, r *http.Request, p *attributesPayload) (*attributesResult, error) {
	db := orm.NewOrmManager()

	// check access
	uScope, ok := aaa.CheckPermOn(p.baseData.owner, p.baseData.currentUser, "edit_attributes", p.baseData.domain.ID)
	if !ok {
		return nil, errors.AccessDenied
	}

	err := p.baseData.campaign.SetAuditUserData(p.baseData.currentUser.ID, false, 0, "edit_attributes", uScope)
	if err != nil {
		return nil, err
	}

	attrs := orm.CampaignAttributes{
		CampaignID:   p.baseData.campaign.ID,
		Device:       p.Device,
		Manufacturer: p.Manufacturer,
		OS:           p.OS,
		Browser:      p.Browser,
		IAB:          p.IAB,
		Region:       p.Region,
		Cellular:     p.Cellular,
		ISP:          p.ISP,
	}

	d := structs.Map(attrs)
	err = p.baseData.campaign.SetAuditDescribe(d, "update campaign attributes")
	if err != nil {
		return nil, err
	}

	attr, err := db.AttachCampaignAttributes(attrs)
	if err != nil {
		return nil, errors.UpdateError
	}

	res := attributesResult{
		Campaign:           *p.baseData.campaign,
		CampaignAttributes: attr,
	}

	return &res, nil
}
