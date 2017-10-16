package controllers

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"reflect"
	"strconv"
	"time"

	"clickyab.com/crab/modules/campaign/orm"
	"clickyab.com/crab/modules/domain/middleware/domain"
	"clickyab.com/crab/modules/user/aaa"
	"clickyab.com/crab/modules/user/middleware/authz"
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

// get gets a campaign by id
// @Route {
// 		url = /:id
//		method = get
//		resource = get-campaign:self
//		200 = orm.Campaign
//		400 = controller.ErrorResponseSimple
//		404 = controller.ErrorResponseSimple
//		middleware = authz.Authenticate
// }
func (c *Controller) get(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	userDomain := domain.MustGetDomain(ctx)
	currentUser := authz.MustGetUser(ctx)
	id := xmux.Param(ctx, "id")
	campID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		c.BadResponse(w, errors.New("bad id"))
		return
	}

	campaign, err := orm.NewOrmManager().FindCampaignByIDDomain(campID, userDomain.ID)
	if err != nil {
		c.NotFoundResponse(w, nil)
		return
	}

	owner, err := aaa.NewAaaManager().FindUserWithParentsByID(campaign.UserID, userDomain.ID)
	assert.Nil(err)

	_, ok := aaa.CheckPermOn(owner, currentUser, "get-campaign", userDomain.ID)
	if !ok {
		c.ForbiddenResponse(w, errors.New("don't have access for this action"))
		return
	}

	c.OKResponse(w, campaign)
}
