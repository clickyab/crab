package controllers

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"reflect"

	"clickyab.com/crab/modules/campaign/orm"
	"clickyab.com/crab/modules/domain/middleware/domain"
	"clickyab.com/crab/modules/user/aaa"
	"clickyab.com/crab/modules/user/middleware/authz"
	"github.com/clickyab/services/assert"
	"github.com/clickyab/services/framework/controller"
	"github.com/clickyab/services/mysql"
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

// @Validate{
//}
type createCampaignPayload struct {
	Kind    orm.CampaignKind `json:"kind"`
	Type    orm.CampaignType `json:"type"`
	Status  bool             `json:"status"`
	StartAt time.Time        `json:"start_at"`
	EndAt   time.Time        `json:"end_at"`
	Title   string           `json:"title" validate:"required,gt=5"`
	H00     string           `json:"h00" hour:""`
	H01     string           `json:"h01" hour:""`
	H02     string           `json:"h02" hour:""`
	H03     string           `json:"h03" hour:""`
	H04     string           `json:"h04" hour:""`
	H05     string           `json:"h05" hour:""`
	H06     string           `json:"h06" hour:""`
	H07     string           `json:"h07" hour:""`
	H08     string           `json:"h08" hour:""`
	H09     string           `json:"h09" hour:""`
	H10     string           `json:"h10" hour:""`
	H11     string           `json:"h11" hour:""`
	H12     string           `json:"h12" hour:""`
	H13     string           `json:"h13" hour:""`
	H14     string           `json:"h14" hour:""`
	H15     string           `json:"h15" hour:""`
	H16     string           `json:"h16" hour:""`
	H17     string           `json:"h17" hour:""`
	H18     string           `json:"h18" hour:""`
	H19     string           `json:"h19" hour:""`
	H20     string           `json:"h20" hour:""`
	H21     string           `json:"h21" hour:""`
	H22     string           `json:"h22" hour:""`
	H23     string           `json:"h23" hour:""`
}

func validateHours(m interface{}) bool {
	v := reflect.ValueOf(m)
	t := reflect.TypeOf(m)
	for i := 0; i < t.NumField(); i++ {
		if _, ok := t.Field(i).Tag.Lookup("hour"); ok {
			if reflect.Indirect(v).FieldByName(t.Field(i).Name).String() != "" {
				return true
			}
		}
	}
	return false
}

func (l *createCampaignPayload) ValidateExtra(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	if l.StartAt.IsZero() {
		return errors.New("campaign should start in future")
	}
	if !l.EndAt.IsZero() && l.StartAt.Unix() > l.EndAt.Unix() {
		return errors.New("campaign should end after start")
	}

	if !validateHours(l) {
		return errors.New("at least one object in schedule should be true")
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
//		200 = campaignResponse
//		400 = controller.ErrorResponseSimple
// }
func (c Controller) createBase(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	u := authz.MustGetUser(ctx)
	d := domain.MustGetDomain(ctx)
	p := c.MustGetPayload(ctx).(*createCampaignPayload)
	ca, err := orm.NewOrmManager().AddCampaign(orm.CampaignBase{
		CampaignBaseType: orm.CampaignBaseType{
			Kind: p.Kind,
			Type: p.Type,
		},
		CampaignStatus: orm.CampaignStatus{
			Status:  p.Status,
			Title:   p.Title,
			StartAt: p.StartAt,
			EndAt:   p.EndAt,
			Schedule: orm.ScheduleSheet{
				H00: mysql.NullString{
					String: p.H00,
					Valid:  p.H00 != "",
				}, H01: mysql.NullString{
					String: p.H01,
					Valid:  p.H01 != "",
				}, H02: mysql.NullString{
					String: p.H02,
					Valid:  p.H02 != "",
				}, H03: mysql.NullString{
					String: p.H03,
					Valid:  p.H03 != "",
				}, H04: mysql.NullString{
					String: p.H04,
					Valid:  p.H04 != "",
				}, H05: mysql.NullString{
					String: p.H05,
					Valid:  p.H05 != "",
				}, H06: mysql.NullString{
					String: p.H06,
					Valid:  p.H06 != "",
				}, H07: mysql.NullString{
					String: p.H07,
					Valid:  p.H07 != "",
				}, H08: mysql.NullString{
					String: p.H08,
					Valid:  p.H08 != "",
				}, H09: mysql.NullString{
					String: p.H09,
					Valid:  p.H09 != "",
				}, H10: mysql.NullString{
					String: p.H10,
					Valid:  p.H10 != "",
				}, H11: mysql.NullString{
					String: p.H11,
					Valid:  p.H11 != "",
				}, H12: mysql.NullString{
					String: p.H12,
					Valid:  p.H12 != "",
				}, H13: mysql.NullString{
					String: p.H13,
					Valid:  p.H13 != "",
				}, H14: mysql.NullString{
					String: p.H14,
					Valid:  p.H14 != "",
				}, H15: mysql.NullString{
					String: p.H15,
					Valid:  p.H15 != "",
				}, H16: mysql.NullString{
					String: p.H16,
					Valid:  p.H16 != "",
				}, H17: mysql.NullString{
					String: p.H17,
					Valid:  p.H17 != "",
				}, H18: mysql.NullString{
					String: p.H18,
					Valid:  p.H18 != "",
				}, H19: mysql.NullString{
					String: p.H19,
					Valid:  p.H19 != "",
				}, H20: mysql.NullString{
					String: p.H20,
					Valid:  p.H20 != "",
				}, H21: mysql.NullString{
					String: p.H21,
					Valid:  p.H21 != "",
				}, H22: mysql.NullString{
					String: p.H22,
					Valid:  p.H22 != "",
				}, H23: mysql.NullString{
					String: p.H23,
					Valid:  p.H23 != "",
				},
			},
		},
	}, u, d)
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
type campaignStatus struct {
	Status  bool      `json:"status" `
	StartAt time.Time `json:"start_at" `
	EndAt   time.Time `json:"end_at" `
	Title   string    `json:"title"  validate:"required,gt=5"`
	H00     string    `json:"h00" hour:""`
	H01     string    `json:"h01" hour:""`
	H02     string    `json:"h02" hour:""`
	H03     string    `json:"h03" hour:""`
	H04     string    `json:"h04" hour:""`
	H05     string    `json:"h05" hour:""`
	H06     string    `json:"h06" hour:""`
	H07     string    `json:"h07" hour:""`
	H08     string    `json:"h08" hour:""`
	H09     string    `json:"h09" hour:""`
	H10     string    `json:"h10" hour:""`
	H11     string    `json:"h11" hour:""`
	H12     string    `json:"h12" hour:""`
	H13     string    `json:"h13" hour:""`
	H14     string    `json:"h14" hour:""`
	H15     string    `json:"h15" hour:""`
	H16     string    `json:"h16" hour:""`
	H17     string    `json:"h17" hour:""`
	H18     string    `json:"h18" hour:""`
	H19     string    `json:"h19" hour:""`
	H20     string    `json:"h20" hour:""`
	H21     string    `json:"h21" hour:""`
	H22     string    `json:"h22" hour:""`
	H23     string    `json:"h23" hour:""`
}

func (l *campaignStatus) ValidateExtra(ctx context.Context, w http.ResponseWriter, r *http.Request) error {

	if l.StartAt.IsZero() {
		return errors.New("campaign should start in future")
	}
	if !l.EndAt.IsZero() && l.StartAt.Unix() > l.EndAt.Unix() {
		return errors.New("campaign should end after start")
	}

	if !validateHours(l) {
		return errors.New("at least one object in schedule should be true")
	}
	return nil
}

// updateBase campaign
// @Route {
// 		url = /base/:id
//		method = put
//		payload = campaignStatus
//		middleware = authz.Authenticate
//		200 = campaignResponse
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
	p := c.MustGetPayload(ctx).(*campaignStatus)

	o := orm.NewOrmManager()
	ca, e := o.FindCampaignByID(id)
	if e != nil || ca.DomainID != d.ID {
		c.NotFoundResponse(w, nil)
		return
	}
	// TODO: check access
	//u.HasOn("edit-campaign",ca.UserID,[],d.ID)
	err = o.UpdateCampaignByID(orm.CampaignStatus{
		Status:  p.Status,
		Title:   p.Title,
		StartAt: p.StartAt,
		EndAt:   p.EndAt,
		Schedule: orm.ScheduleSheet{
			H00: mysql.NullString{
				String: p.H00,
				Valid:  p.H00 != "",
			}, H01: mysql.NullString{
				String: p.H01,
				Valid:  p.H01 != "",
			}, H02: mysql.NullString{
				String: p.H02,
				Valid:  p.H02 != "",
			}, H03: mysql.NullString{
				String: p.H03,
				Valid:  p.H03 != "",
			}, H04: mysql.NullString{
				String: p.H04,
				Valid:  p.H04 != "",
			}, H05: mysql.NullString{
				String: p.H05,
				Valid:  p.H05 != "",
			}, H06: mysql.NullString{
				String: p.H06,
				Valid:  p.H06 != "",
			}, H07: mysql.NullString{
				String: p.H07,
				Valid:  p.H07 != "",
			}, H08: mysql.NullString{
				String: p.H08,
				Valid:  p.H08 != "",
			}, H09: mysql.NullString{
				String: p.H09,
				Valid:  p.H09 != "",
			}, H10: mysql.NullString{
				String: p.H10,
				Valid:  p.H10 != "",
			}, H11: mysql.NullString{
				String: p.H11,
				Valid:  p.H11 != "",
			}, H12: mysql.NullString{
				String: p.H12,
				Valid:  p.H12 != "",
			}, H13: mysql.NullString{
				String: p.H13,
				Valid:  p.H13 != "",
			}, H14: mysql.NullString{
				String: p.H14,
				Valid:  p.H14 != "",
			}, H15: mysql.NullString{
				String: p.H15,
				Valid:  p.H15 != "",
			}, H16: mysql.NullString{
				String: p.H16,
				Valid:  p.H16 != "",
			}, H17: mysql.NullString{
				String: p.H17,
				Valid:  p.H17 != "",
			}, H18: mysql.NullString{
				String: p.H18,
				Valid:  p.H18 != "",
			}, H19: mysql.NullString{
				String: p.H19,
				Valid:  p.H19 != "",
			}, H20: mysql.NullString{
				String: p.H20,
				Valid:  p.H20 != "",
			}, H21: mysql.NullString{
				String: p.H21,
				Valid:  p.H21 != "",
			}, H22: mysql.NullString{
				String: p.H22,
				Valid:  p.H22 != "",
			}, H23: mysql.NullString{
				String: p.H23,
				Valid:  p.H23 != "",
			}}}, ca)
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

	c.OKResponse(w, createResponse(ca))

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
//		200 = campaignResponse
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

	c.OKResponse(w, createResponse(campaign))
}
