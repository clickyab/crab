package controllers

import (
	"context"
	"errors"
	"net/http"

	"fmt"
	"net/mail"
	"time"

	"strconv"

	"clickyab.com/crab/modules/campaign/orm"
	"clickyab.com/crab/modules/domain/middleware/domain"
	"clickyab.com/crab/modules/user/middleware/authz"
	"github.com/clickyab/services/assert"
	"github.com/clickyab/services/framework/controller"
	"github.com/rs/xmux"
)

// Controller is the controller for the campaign package
// @Route {
//		group = /campaign
//		middleware = domain.Access
// }
type Controller struct {
	controller.Base
}

// @Validate{
//}
type createCampaignPayload struct {
	orm.CreateCampaign
}

func (l *createCampaignPayload) ValidateExtra(ctx context.Context) error {

	if l.Campaign.Kind != orm.AppCampaign && l.Campaign.Kind != orm.WebCampaign {
		return fmt.Errorf("%s is not a valid campaign kind. choose from %s or %s", l.Campaign.Kind, orm.AppCampaign, orm.WebCampaign)
	}
	if l.Campaign.Type != orm.BannerType && l.Campaign.Type != orm.VastType && l.Campaign.Type != orm.NativeType {
		return fmt.Errorf("%s is not a valid campaign type. choose from %s, %s or %s", l.Campaign.Type, orm.BannerType, orm.VastType, orm.NativeType)
	}
	if l.Campaign.StartAt.Unix() < time.Now().Unix() {
		return errors.New("campaign should start in future")
	}
	if !l.Campaign.EndAt.IsZero() && l.Campaign.StartAt.Unix() > l.Campaign.EndAt.Unix() {
		return errors.New("campaign should end after start")
	}

	if l.Campaign.Budget < 0 {
		return errors.New("campaign budget should be a positive number")
	}

	if l.Campaign.CostType != orm.CPM && l.Campaign.CostType != orm.CPC && l.Campaign.CostType != orm.CPA {
		return fmt.Errorf("%s is not a valid cost type. choose from %s, %s or %s", l.Campaign.CostType, orm.CPC, orm.CPM, orm.CPA)
	}

	for m := range l.Attributes.Email {
		_, e := mail.ParseAddress(l.Attributes.Email[m])
		if e != nil {
			return fmt.Errorf("%s is not a valid email", l.Attributes.Email[m])
		}
	}

	return nil
}

type addResponse struct {
	ID int64 `json:"id"`
}

// add campaign
// @Route {
// 		url = /add
//		method = post
//		payload = createCampaignPayload
//		middleware = authz.Authenticate
//		200 = addResponse
//		400 = controller.ErrorResponseSimple
// }
func (c Controller) add(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	u := authz.MustGetUser(ctx)
	d := domain.MustGetDomain(ctx)
	p := c.MustGetPayload(ctx).(*createCampaignPayload)
	i, e := orm.NewOrmManager().AddCampaign(p.CreateCampaign, u, d)
	if e == orm.ErrInventoryID {
		c.BadResponse(w, e)
		return
	}
	assert.Nil(e)
	c.OKResponse(w, addResponse{
		i,
	})

}

// @Validate{
//}
type updateCampaignPayload struct {
	orm.UpdateCampaign
}

func (l *updateCampaignPayload) ValidateExtra(ctx context.Context) error {

	if !l.Campaign.StartAt.IsZero() && l.Campaign.StartAt.Unix() < time.Now().Unix() {
		return errors.New("campaign should start in future")
	}
	if !l.Campaign.EndAt.IsZero() && l.Campaign.StartAt.Unix() > l.Campaign.EndAt.Unix() {
		return errors.New("campaign should end after start")
	}

	if l.Campaign.Budget < 0 {
		return errors.New("campaign budget should be a positive number")
	}

	for m := range l.Attributes.Email {
		_, e := mail.ParseAddress(l.Attributes.Email[m])
		if e != nil {
			return fmt.Errorf("%s is not a valid email", l.Attributes.Email[m])
		}
	}

	return nil
}

// update campaign
// @Route {
// 		url = /update/:id
//		method = put
//		payload = updateCampaignPayload
//		middleware = authz.Authenticate
//		200 = controller.NormalResponse
//		400 = controller.ErrorResponseSimple
//		404 = controller.ErrorResponseSimple
//		resource = edit-campaign:self
// }
func (c Controller) update(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(xmux.Param(ctx, "id"), 10, 64)
	if err != nil {
		c.BadResponse(w, errors.New("id is not valid"))
	}
	u := authz.MustGetUser(ctx)

	d := domain.MustGetDomain(ctx)
	p := c.MustGetPayload(ctx).(*updateCampaignPayload)

	o := orm.NewOrmManager()
	ca, e := o.FindCampaignByID(id)
	if e != nil || ca.DomainID != d.ID || ca.UserID != u.ID {
		c.NotFoundResponse(w, nil)
		return
	}
	// TODO: check access
	//u.HasOn("edit-campaign",ca.UserID,[],d.ID)
	e = o.UpdateCampaignByPayload(id, p.UpdateCampaign)

	if e == orm.ErrInventoryID {
		c.BadResponse(w, e)
		return
	}
	assert.Nil(e)
	c.OKResponse(w, nil)

}
