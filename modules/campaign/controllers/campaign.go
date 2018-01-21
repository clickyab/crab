package controllers

import (
	"context"
	"encoding/json"
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
	"github.com/clickyab/services/gettext/t9e"
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
	Kind     orm.CampaignKind `json:"kind"`
	Type     orm.CampaignType `json:"type"`
	Status   bool             `json:"status"`
	StartAt  time.Time        `json:"start_at"`
	EndAt    mysql.NullTime   `json:"end_at"`
	Title    string           `json:"title" validate:"required,gt=5"`
	Schedule struct {
		H00 string `json:"h00" hour:""`
		H01 string `json:"h01" hour:""`
		H02 string `json:"h02" hour:""`
		H03 string `json:"h03" hour:""`
		H04 string `json:"h04" hour:""`
		H05 string `json:"h05" hour:""`
		H06 string `json:"h06" hour:""`
		H07 string `json:"h07" hour:""`
		H08 string `json:"h08" hour:""`
		H09 string `json:"h09" hour:""`
		H10 string `json:"h10" hour:""`
		H11 string `json:"h11" hour:""`
		H12 string `json:"h12" hour:""`
		H13 string `json:"h13" hour:""`
		H14 string `json:"h14" hour:""`
		H15 string `json:"h15" hour:""`
		H16 string `json:"h16" hour:""`
		H17 string `json:"h17" hour:""`
		H18 string `json:"h18" hour:""`
		H19 string `json:"h19" hour:""`
		H20 string `json:"h20" hour:""`
		H21 string `json:"h21" hour:""`
		H22 string `json:"h22" hour:""`
		H23 string `json:"h23" hour:""`
	} `json:"schedule"`
}

func validateHours(m interface{}) bool {
	v := reflect.ValueOf(m)
	t := reflect.TypeOf(m)
	for i := 0; i < t.NumField(); i++ {
		if t.Field(i).Type.Kind() == reflect.Struct {
			if validateHours(v.Field(i).Interface()) {
				return true
			}
			continue
		}
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
		return t9e.G("campaign should start in future")
	}
	if l.EndAt.Valid && l.StartAt.Unix() > l.EndAt.Time.Unix() {
		return t9e.G("campaign should end after start")
	}

	if !validateHours(l.Schedule) {
		return t9e.G("at least one object in schedule should be true")
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
		return t9e.G("campaign should start in future")
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
					String: p.Schedule.H00,
					Valid:  p.Schedule.H00 != "",
				}, H01: mysql.NullString{
					String: p.Schedule.H01,
					Valid:  p.Schedule.H01 != "",
				}, H02: mysql.NullString{
					String: p.Schedule.H02,
					Valid:  p.Schedule.H02 != "",
				}, H03: mysql.NullString{
					String: p.Schedule.H03,
					Valid:  p.Schedule.H03 != "",
				}, H04: mysql.NullString{
					String: p.Schedule.H04,
					Valid:  p.Schedule.H04 != "",
				}, H05: mysql.NullString{
					String: p.Schedule.H05,
					Valid:  p.Schedule.H05 != "",
				}, H06: mysql.NullString{
					String: p.Schedule.H06,
					Valid:  p.Schedule.H06 != "",
				}, H07: mysql.NullString{
					String: p.Schedule.H07,
					Valid:  p.Schedule.H07 != "",
				}, H08: mysql.NullString{
					String: p.Schedule.H08,
					Valid:  p.Schedule.H08 != "",
				}, H09: mysql.NullString{
					String: p.Schedule.H09,
					Valid:  p.Schedule.H09 != "",
				}, H10: mysql.NullString{
					String: p.Schedule.H10,
					Valid:  p.Schedule.H10 != "",
				}, H11: mysql.NullString{
					String: p.Schedule.H11,
					Valid:  p.Schedule.H11 != "",
				}, H12: mysql.NullString{
					String: p.Schedule.H12,
					Valid:  p.Schedule.H12 != "",
				}, H13: mysql.NullString{
					String: p.Schedule.H13,
					Valid:  p.Schedule.H13 != "",
				}, H14: mysql.NullString{
					String: p.Schedule.H14,
					Valid:  p.Schedule.H14 != "",
				}, H15: mysql.NullString{
					String: p.Schedule.H15,
					Valid:  p.Schedule.H15 != "",
				}, H16: mysql.NullString{
					String: p.Schedule.H16,
					Valid:  p.Schedule.H16 != "",
				}, H17: mysql.NullString{
					String: p.Schedule.H17,
					Valid:  p.Schedule.H17 != "",
				}, H18: mysql.NullString{
					String: p.Schedule.H18,
					Valid:  p.Schedule.H18 != "",
				}, H19: mysql.NullString{
					String: p.Schedule.H19,
					Valid:  p.Schedule.H19 != "",
				}, H20: mysql.NullString{
					String: p.Schedule.H20,
					Valid:  p.Schedule.H20 != "",
				}, H21: mysql.NullString{
					String: p.Schedule.H21,
					Valid:  p.Schedule.H21 != "",
				}, H22: mysql.NullString{
					String: p.Schedule.H22,
					Valid:  p.Schedule.H22 != "",
				}, H23: mysql.NullString{
					String: p.Schedule.H23,
					Valid:  p.Schedule.H23 != "",
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
	Status   bool           `json:"status"`
	StartAt  time.Time      `json:"start_at"`
	EndAt    mysql.NullTime `json:"end_at"`
	Title    string         `json:"title"  validate:"required,gt=5"`
	Schedule struct {
		H00 string `json:"h00" hour:""`
		H01 string `json:"h01" hour:""`
		H02 string `json:"h02" hour:""`
		H03 string `json:"h03" hour:""`
		H04 string `json:"h04" hour:""`
		H05 string `json:"h05" hour:""`
		H06 string `json:"h06" hour:""`
		H07 string `json:"h07" hour:""`
		H08 string `json:"h08" hour:""`
		H09 string `json:"h09" hour:""`
		H10 string `json:"h10" hour:""`
		H11 string `json:"h11" hour:""`
		H12 string `json:"h12" hour:""`
		H13 string `json:"h13" hour:""`
		H14 string `json:"h14" hour:""`
		H15 string `json:"h15" hour:""`
		H16 string `json:"h16" hour:""`
		H17 string `json:"h17" hour:""`
		H18 string `json:"h18" hour:""`
		H19 string `json:"h19" hour:""`
		H20 string `json:"h20" hour:""`
		H21 string `json:"h21" hour:""`
		H22 string `json:"h22" hour:""`
		H23 string `json:"h23" hour:""`
	} `json:"schedule"`
}

func (l *campaignStatus) ValidateExtra(ctx context.Context, w http.ResponseWriter, r *http.Request) error {

	if l.StartAt.IsZero() {
		return t9e.G("campaign should start in future")
	}
	if l.EndAt.Valid && l.StartAt.Unix() > l.EndAt.Time.Unix() {
		return t9e.G("campaign should end after start")
	}

	if !validateHours(l.Schedule) {
		return t9e.G("at least one object in schedule should be true")
	}
	return nil
}

// updateBase campaign
// @Route {
// 		url = /base/:id
//		method = put
//		payload = campaignStatus
//		middleware = authz.Authenticate
//		200 = orm.Campaign
//		400 = controller.ErrorResponseSimple
//		404 = controller.ErrorResponseSimple
//		resource = edit-campaign:self
// }
func (c Controller) updateBase(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(xmux.Param(ctx, "id"), 10, 64)
	if err != nil {
		c.BadResponse(w, t9e.G("id is not valid"))
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
				String: p.Schedule.H00,
				Valid:  p.Schedule.H00 != "",
			}, H01: mysql.NullString{
				String: p.Schedule.H01,
				Valid:  p.Schedule.H01 != "",
			}, H02: mysql.NullString{
				String: p.Schedule.H02,
				Valid:  p.Schedule.H02 != "",
			}, H03: mysql.NullString{
				String: p.Schedule.H03,
				Valid:  p.Schedule.H03 != "",
			}, H04: mysql.NullString{
				String: p.Schedule.H04,
				Valid:  p.Schedule.H04 != "",
			}, H05: mysql.NullString{
				String: p.Schedule.H05,
				Valid:  p.Schedule.H05 != "",
			}, H06: mysql.NullString{
				String: p.Schedule.H06,
				Valid:  p.Schedule.H06 != "",
			}, H07: mysql.NullString{
				String: p.Schedule.H07,
				Valid:  p.Schedule.H07 != "",
			}, H08: mysql.NullString{
				String: p.Schedule.H08,
				Valid:  p.Schedule.H08 != "",
			}, H09: mysql.NullString{
				String: p.Schedule.H09,
				Valid:  p.Schedule.H09 != "",
			}, H10: mysql.NullString{
				String: p.Schedule.H10,
				Valid:  p.Schedule.H10 != "",
			}, H11: mysql.NullString{
				String: p.Schedule.H11,
				Valid:  p.Schedule.H11 != "",
			}, H12: mysql.NullString{
				String: p.Schedule.H12,
				Valid:  p.Schedule.H12 != "",
			}, H13: mysql.NullString{
				String: p.Schedule.H13,
				Valid:  p.Schedule.H13 != "",
			}, H14: mysql.NullString{
				String: p.Schedule.H14,
				Valid:  p.Schedule.H14 != "",
			}, H15: mysql.NullString{
				String: p.Schedule.H15,
				Valid:  p.Schedule.H15 != "",
			}, H16: mysql.NullString{
				String: p.Schedule.H16,
				Valid:  p.Schedule.H16 != "",
			}, H17: mysql.NullString{
				String: p.Schedule.H17,
				Valid:  p.Schedule.H17 != "",
			}, H18: mysql.NullString{
				String: p.Schedule.H18,
				Valid:  p.Schedule.H18 != "",
			}, H19: mysql.NullString{
				String: p.Schedule.H19,
				Valid:  p.Schedule.H19 != "",
			}, H20: mysql.NullString{
				String: p.Schedule.H20,
				Valid:  p.Schedule.H20 != "",
			}, H21: mysql.NullString{
				String: p.Schedule.H21,
				Valid:  p.Schedule.H21 != "",
			}, H22: mysql.NullString{
				String: p.Schedule.H22,
				Valid:  p.Schedule.H22 != "",
			}, H23: mysql.NullString{
				String: p.Schedule.H23,
				Valid:  p.Schedule.H23 != "",
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
		c.BadResponse(w, t9e.G("id is not valid"))
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
// 		url = /get/:id
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
		c.BadResponse(w, t9e.G("bad id"))
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
		c.ForbiddenResponse(w, t9e.G("don't have access for this action"))
		return
	}

	c.OKResponse(w, campaign)
}
