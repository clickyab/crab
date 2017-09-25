package controllers

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/mail"
	"reflect"
	"strconv"
	"strings"
	"time"

	asset "clickyab.com/crab/modules/asset/orm"
	"clickyab.com/crab/modules/campaign/orm"
	"clickyab.com/crab/modules/domain/middleware/domain"
	"clickyab.com/crab/modules/user/middleware/authz"
	"github.com/clickyab/services/array"
	"github.com/clickyab/services/assert"
	"github.com/clickyab/services/framework/controller"
	"github.com/clickyab/services/random"
	"github.com/rs/xmux"
	"github.com/sirupsen/logrus"
)

// Controller is the controller for the campaign package
// @Route {
//		group = /campaign
//		middleware = domain.Access
// }
type Controller struct {
	controller.Base
}

func validate(l orm.CampaignStatus) error {

	if l.StartAt.IsZero() {
		return errors.New("campaign should start in future")
	}
	if !l.EndAt.IsZero() && l.StartAt.Unix() > l.EndAt.Unix() {
		return errors.New("campaign should end after start")
	}
	var any bool
	tm := reflect.ValueOf(l.Schedule)
	for i := 0; i < tm.NumField(); i++ {
		if tm.Field(i).Bool() == true {
			any = true
			break
		}
	}
	if !any {
		return errors.New("at least one object in schedule should be true")
	}
	return nil
}

// @Validate{
//}
type createCampaignPayload struct {
	orm.CampaignBase
}

func (l *createCampaignPayload) ValidateExtra(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	if err := validate(l.CampaignStatus); err != nil {
		return err
	}
	if !l.Kind.IsValid() {
		return fmt.Errorf("%s is not a valid campaign kind. choose from %s or %s", l.Kind, orm.AppCampaign, orm.WebCampaign)

	}
	if !l.Type.IsValid() {
		return fmt.Errorf("%s is not a valid campaign type. choose from %s, %s or %s", l.Type, orm.BannerType, orm.VastType, orm.NativeType)
	}

	today, err := time.Parse("02-01-03", time.Now().Format("02-01-03"))
	assert.Nil(err)
	if l.StartAt.Unix() < today.Unix() {
		return errors.New("campaign should start in future")
	}

	return nil
}

// createBase campaign
// @Route {
// 		url = /create
//		method = post
//		payload = createCampaignPayload
//		middleware = authz.Authenticate
//		200 = orm.Campaign
//		400 = controller.ErrorResponseSimple
// }
func (c Controller) createBase(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	u := authz.MustGetUser(ctx)
	d := domain.MustGetDomain(ctx)
	p := c.MustGetPayload(ctx).(*createCampaignPayload)
	ca, err := orm.NewOrmManager().AddCampaign(p.CampaignBase, u, d)
	if err != nil {
		j, e := json.MarshalIndent(ca, " ", "  ")
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
		c.BadResponse(w, nil)
		return
	}
	c.OKResponse(w, ca)
}

// @Validate{
//}
type updateCampaignPayload struct {
	orm.CampaignStatus
}

func (l *updateCampaignPayload) ValidateExtra(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	return validate(l.CampaignStatus)
}

// updateBase campaign
// @Route {
// 		url = /base/:id
//		method = put
//		payload = updateCampaignPayload
//		middleware = authz.Authenticate
//		200 = orm.Campaign
//		400 = controller.ErrorResponseSimple
//		404 = controller.ErrorResponseSimple
//		resource = edit-campaign:self
// }
func (c Controller) updateBase(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(xmux.Param(ctx, "id"), 10, 64)
	if err != nil {
		c.BadResponse(w, errors.New("id is not valid"))
	}

	d := domain.MustGetDomain(ctx)
	p := c.MustGetPayload(ctx).(*updateCampaignPayload)

	o := orm.NewOrmManager()
	ca, e := o.FindCampaignByID(id)
	if e != nil || ca.DomainID != d.ID {
		c.NotFoundResponse(w, nil)
		return
	}
	// TODO: check access
	//u.HasOn("edit-campaign",ca.UserID,[],d.ID)
	err = o.UpdateCampaignByID(p.CampaignStatus, ca)
	if err == orm.ErrorStartDate {
		c.BadResponse(w, e)
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
		c.BadResponse(w, nil)
		return
	}

	c.OKResponse(w, ca)

}

// @Validate{
//}
type budgetPayload struct {
	orm.CampaignFinance
}

func (l *budgetPayload) ValidateExtra(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	for _, m := range l.NotifyEmail {
		if _, err := mail.ParseAddress(m); err != nil {
			return err
		}
	}
	if !l.CostType.IsValid() {
		return fmt.Errorf("cost type %s is not valid. options are %s,%s or %s", l.CostType, orm.CPC, orm.CPM, orm.CPA)
	}
	if l.Budget < 0 || l.DailyLimit < 0 || l.MaxBid < 0 {
		return errors.New("budget, daily limit and max bid can not be a negative number")
	}

	return nil
}

// finance will update campaign finance
// @Route {
// 		url = /budget/:id
//		method = put
//		payload = budgetPayload
//		200 = orm.Campaign
//		400 = controller.ErrorResponseSimple
//		404 = controller.ErrorResponseSimple
//		middleware = authz.Authenticate
// }
func (c *Controller) budget(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(xmux.Param(ctx, "id"), 10, 64)
	p := c.MustGetPayload(ctx).(*budgetPayload)

	if err != nil || id < 1 {
		c.BadResponse(w, errors.New("id is not valid"))
	}
	db := orm.NewOrmManager()
	o, err := db.FindCampaignByID(id)
	if err != nil {
		c.NotFoundResponse(w, nil)
	}

	err = db.UpdateCampaignFinance(p.CampaignFinance, o)

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
		c.BadResponse(w, errors.New("can not update budget"))
		return
	}
	c.OKResponse(w, o)
}

// @Validate{
//}
type wblistPayload struct {
	ListID int64 `json:"list_id" db:"-" validate:"required"`
}

func (l *wblistPayload) ValidateExtra(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	if l.ListID == 0 || l.ListID < -1 {
		return errors.New("value is not valid! list id should be positive number or -1 for delete current list")
	}
	return nil
}

// wblist will update campaign white/black list
// @Route {
// 		url = /wblist/:id
//		method = put
//		payload = wblistPayload
//		200 = orm.Campaign
//		400 = controller.ErrorResponseSimple
//		404 = controller.ErrorResponseSimple
//		middleware = authz.Authenticate
// }
func (c *Controller) wblist(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(xmux.Param(ctx, "id"), 10, 64)
	p := c.MustGetPayload(ctx).(*wblistPayload)
	u := authz.MustGetUser(ctx)

	if err != nil || id < 1 {
		c.BadResponse(w, errors.New("id is not valid"))
	}
	db := orm.NewOrmManager()
	o, err := db.FindCampaignByID(id)
	if err != nil {
		c.NotFoundResponse(w, nil)
	}

	err = db.UpdateCampaignWhiteBlackList(p.ListID, o, u)
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
		c.BadResponse(w, errors.New("can not update white/black list"))
		return
	}
	c.OKResponse(w, o)
}

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

// finalize
// @Route {
// 		url = /finalize/:id
//		method = put
//		200 = controller.NormalResponse
//		400 = controller.ErrorResponseSimple
//		404 = controller.ErrorResponseSimple
//		middleware = authz.Authenticate
// }
func (c *Controller) finalize(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(xmux.Param(ctx, "id"), 10, 64)

	if err != nil || id < 1 {
		c.BadResponse(w, errors.New("id is not valid"))
	}
	db := orm.NewOrmManager()
	ca, err := db.FindCampaignByID(id)
	if err != nil {
		c.NotFoundResponse(w, nil)
	}

	db.Finalize(ca)
	c.OKResponse(w, nil)
}
