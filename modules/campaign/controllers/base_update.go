package controllers

import (
	"context"
	"net/http"
	"time"

	"clickyab.com/crab/modules/campaign/errors"
	"clickyab.com/crab/modules/campaign/orm"
	"clickyab.com/crab/modules/user/aaa"
	"github.com/clickyab/services/mysql"
	"github.com/clickyab/services/xlog"
	"github.com/fatih/structs"
)

// @Validate{
//}
type campaignBase struct {
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
	baseData *BaseData `json:"-"`
}

func (l *campaignBase) ValidateExtra(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	res, err := CheckUserCamapignDomain(ctx)
	if err != nil {
		return err
	}
	l.baseData = res

	if l.StartAt.IsZero() {
		return errors.CampaignStartTimeError
	}
	if l.EndAt.Valid && l.StartAt.Before(l.EndAt.Time) {
		return errors.EndTimeError
	}

	if !validateHours(l.Schedule) {
		return errors.TimeScheduleError
	}
	return nil
}

type updateResult struct {
	orm.Campaign
	Schedule orm.Schedule `json:"schedule"`
}

// updateBase campaign
// @Rest {
// 		url = /base/:id
//		protected = true
// 		method = put
//		resource = edit_campaign:self
// }
func (c Controller) updateBase(ctx context.Context, r *http.Request, p *campaignBase) (*updateResult, error) {
	db := orm.NewOrmManager()

	// check access
	uScope, ok := aaa.CheckPermOn(p.baseData.owner, p.baseData.currentUser, "edit_campaign", p.baseData.domain.ID)
	if !ok {
		return nil, errors.AccessDenied
	}

	err := p.baseData.campaign.SetAuditUserData(p.baseData.currentUser.ID, false, 0, "edit_campaign", uScope)
	if err != nil {
		return nil, err
	}

	baseData := orm.CampaignBase{
		Title:   p.Title,
		TLD:     p.TLD,
		Status:  p.Status,
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
	}

	d := structs.Map(baseData)
	err = p.baseData.campaign.SetAuditDescribe(d, "edit campaign base data")
	if err != nil {
		return nil, err
	}

	err = db.UpdateCampaignByID(baseData, p.baseData.campaign)
	if err != nil {
		if err.Error() == errors.StartTimeError.Error() {
			return nil, err
		}
		xlog.GetWithError(ctx, err).Debug("can't update base info of campaign")
		return nil, errors.UpdateError
	}

	sc, err := db.GetSchedule(p.baseData.campaign.ID)
	if err != nil {
		return nil, errors.NotFoundSchedule
	}

	res := updateResult{
		Campaign: *p.baseData.campaign,
		Schedule: *sc,
	}
	return &res, nil
}
