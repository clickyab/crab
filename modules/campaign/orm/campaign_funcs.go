package orm

import (
	"fmt"
	"time"

	"database/sql"

	"clickyab.com/crab/modules/campaign/errors"
	domainOrm "clickyab.com/crab/modules/domain/orm"
	"clickyab.com/crab/modules/user/aaa"
	"github.com/clickyab/services/assert"
	"github.com/fatih/structs"
)

const (
	// Foreign is filter for region is every where except iran
	Foreign = "foreign"
)

// AddCampaign for creating campaign with minimum info
func (m *Manager) AddCampaign(ca *Campaign, c CampaignBase, u *aaa.User, d *domainOrm.Domain) (Schedule, error) {
	var s Schedule

	ca.Title = c.Title
	ca.DomainID = d.ID
	ca.UserID = u.ID
	ca.Kind = c.Kind
	ca.Exchange = All
	ca.TLD = c.TLD
	ca.Strategy = CPC
	ca.Status = c.Status
	ca.StartAt = c.StartAt
	ca.EndAt = c.EndAt
	ca.Progress = ProgressInProgress

	err := m.Begin()
	assert.Nil(err)
	defer func() {
		if err == nil {
			assert.Nil(m.Commit())
			return
		}
		assert.Nil(m.Rollback())
	}()

	data := structs.Map(ca)
	err = ca.SetAuditDescribe(data, "create new campaign")
	if err != nil {
		return s, err
	}

	if err = m.CreateCampaign(ca); err != nil {
		return s, err
	}

	s = Schedule{
		CampaignID:    ca.ID,
		ScheduleSheet: c.Schedule,
	}
	err = m.CreateSchedule(&s)

	return s, err
}

// UpdateCampaignByID for updating campaign with minimum info
func (m *Manager) UpdateCampaignByID(c CampaignBase, ca *Campaign) error {
	ca.Status = c.Status
	if ca.StartAt != c.StartAt {
		if time.Now().Truncate(time.Hour * 24).After(c.StartAt) {
			return errors.StartTimeError
		}
	}
	ca.StartAt = c.StartAt
	ca.EndAt = c.EndAt
	ca.Title = c.Title
	ca.TLD = c.TLD

	err := m.Begin()
	assert.Nil(err)
	defer func() {
		if err == nil {
			assert.Nil(m.Commit())
			return
		}
		assert.Nil(m.Rollback())
	}()

	s, err := m.FindScheduleByCampaignID(ca.ID)
	if err != nil && err != sql.ErrNoRows {
		return err
	}

	if err == sql.ErrNoRows {
		s = &Schedule{}
	}
	s.ScheduleSheet = c.Schedule
	err = m.UpdateCampaign(ca)
	if err != nil {
		return err
	}

	err = m.UpdateSchedule(s)
	return err
}

// GetSchedule get campaign schedule data
func (m *Manager) GetSchedule(caID int64) (*Schedule, error) {
	s, err := m.FindScheduleByCampaignID(caID)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}

	return s, nil
}

// FindCampaignByIDDomain return the Campaign base on its id and domain id
func (m *Manager) FindCampaignByIDDomain(caID, dID int64) (*Campaign, error) {
	var res Campaign
	err := m.GetRDbMap().SelectOne(
		&res,
		fmt.Sprintf("SELECT %s FROM %s WHERE id=? AND domain_id=?", GetSelectFields(CampaignTableFull, ""), CampaignTableFull),
		caID,
		dID,
	)

	if err != nil {
		return nil, err
	}

	return &res, err
}
