package controllers

import (
	"context"
	"net/http"
	"strconv"
	"time"

	"reflect"

	"clickyab.com/crab/modules/campaign/errors"
	"clickyab.com/crab/modules/campaign/orm"
	"clickyab.com/crab/modules/domain/middleware/domain"
	"clickyab.com/crab/modules/user/aaa"
	"clickyab.com/crab/modules/user/middleware/authz"
	"github.com/clickyab/services/framework/controller"
	"github.com/clickyab/services/gettext/t9e"
	"github.com/clickyab/services/mysql"
	"github.com/clickyab/services/xlog"
	"github.com/rs/xmux"
)

// Controller is the controller for the user package
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
	Status   orm.Status       `json:"status" validate:"required"`
	StartAt  time.Time        `json:"start_at"`
	EndAt    mysql.NullTime   `json:"end_at" validate:"omitempty"`
	Title    string           `json:"title" validate:"required,gt=5"`
	TLD      string           `json:"tld" validate:"required"`
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
		return errors.StartTimeError
	}
	if l.EndAt.Valid && l.StartAt.Unix() > l.EndAt.Time.Unix() {
		return errors.EndTimeError
	}

	if !validateHours(l.Schedule) {
		return errors.TimeScheduleError
	}

	if !l.Kind.IsValid() {
		return errors.KindError

	}

	if l.StartAt.Before(time.Now().Truncate(time.Hour)) {
		return errors.StartTimeError
	}

	return nil
}

// createBase campaign
// @Rest {
// 		url = /create
//		protected = true
// 		method = post
// }
func (c Controller) createBase(ctx context.Context, r *http.Request, p *createCampaignPayload) (*orm.Campaign, error) {
	u := authz.MustGetUser(ctx)
	d := domain.MustGetDomain(ctx)
	ca, err := orm.NewOrmManager().AddCampaign(orm.CampaignBase{
		CampaignBaseType: orm.CampaignBaseType{
			Kind: p.Kind,
		},
		CampaignStatus: orm.CampaignStatus{
			TLD:     p.TLD,
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
		xlog.GetWithError(ctx, err).Debug("can't insert new campaign")

		return nil, t9e.G("can't create new campaign, db error")
	}

	return ca, nil
}

// @Validate{
//}
type campaignStatus struct {
	Status   orm.Status     `json:"status" validate:"required"`
	StartAt  time.Time      `json:"start_at"`
	EndAt    mysql.NullTime `json:"end_at"`
	Title    string         `json:"title"  validate:"required,gt=5"`
	TLD      string         `json:"tld"`
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
		return errors.StartTimeError
	}
	if l.EndAt.Valid && l.StartAt.Unix() > l.EndAt.Time.Unix() {
		return errors.EndTimeError
	}

	if !validateHours(l.Schedule) {
		return errors.TimeScheduleError
	}
	return nil
}

// updateBase campaign
// @Rest {
// 		url = /base/:id
//		protected = true
// 		method = put
//		resource = edit_campaign:self
// }
func (c Controller) updateBase(ctx context.Context, r *http.Request, p *campaignStatus) (*orm.Campaign, error) {
	id, err := strconv.ParseInt(xmux.Param(ctx, "id"), 10, 64)
	if err != nil {
		return nil, errors.InvalidIDErr
	}
	d := domain.MustGetDomain(ctx)

	o := orm.NewOrmManager()
	ca, e := o.FindCampaignByID(id)
	if e != nil || ca.DomainID != d.ID || ca != nil {
		return nil, errors.NotFoundError(id)
	}

	// check access
	userManager := aaa.NewAaaManager()
	owner, err := userManager.FindUserWithParentsByID(ca.UserID, id)
	if err != nil {
		return ca, t9e.G("can't find user with related domain")
	}

	currentUser := authz.MustGetUser(ctx)
	_, ok := aaa.CheckPermOn(owner, currentUser, "edit_campaign", id)
	if !ok {
		return ca, errors.AccessDenied
	}

	err = o.UpdateCampaignByID(orm.CampaignStatus{
		Status:  p.Status,
		Title:   p.Title,
		StartAt: p.StartAt,
		EndAt:   p.EndAt,
		TLD:     p.TLD,
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
	if err != nil {
		if err != errors.StartTimeError {
			xlog.GetWithError(ctx, err).Debug("can't update base info of campaign")
		}
		return nil, t9e.G("can't update base info of campaign")
	}

	return ca, nil
}

// finalize
// @Rest {
// 		url = /finalize/:id
//		protected = true
// 		method = put
// }
func (c *Controller) finalize(ctx context.Context, r *http.Request) (*orm.Campaign, error) {
	id, err := strconv.ParseInt(xmux.Param(ctx, "id"), 10, 64)

	if err != nil || id < 1 {
		return nil, errors.InvalidIDErr
	}
	db := orm.NewOrmManager()
	ca, err := db.FindCampaignByID(id)
	if err != nil {
		return ca, errors.NotFoundError(id)
	}

	err = db.Finalize(ca)
	if err != nil {
		return ca, t9e.G("can't finalize campaign data")
	}

	return ca, nil
}

// get gets a campaign by id
// @Rest {
// 		url = /get/:id
//		protected = true
// 		method = get
//		resource = get_campaign:self
// }
func (c *Controller) get(ctx context.Context, r *http.Request) (*orm.Campaign, error) {
	userDomain := domain.MustGetDomain(ctx)
	currentUser := authz.MustGetUser(ctx)
	id := xmux.Param(ctx, "id")
	campID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return nil, errors.InvalidIDErr
	}

	campaign, err := orm.NewOrmManager().FindCampaignByIDDomain(campID, userDomain.ID)
	if err != nil {
		return campaign, errors.NotFoundError(campID)
	}

	owner, err := aaa.NewAaaManager().FindUserWithParentsByID(campaign.UserID, userDomain.ID)
	if err != nil {
		return campaign, t9e.G("can't find user with id %s and domain id %s", campaign.UserID, userDomain.ID)
	}

	// check access
	_, ok := aaa.CheckPermOn(owner, currentUser, "get_campaign", userDomain.ID)
	if !ok {
		return campaign, t9e.G("access denied. you can't get campaign information")
	}

	return campaign, nil
}
