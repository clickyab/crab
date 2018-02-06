package orm

import (
	"time"

	"database/sql"

	"clickyab.com/crab/modules/campaign/errors"
	"clickyab.com/crab/modules/domain/dmn"
	"clickyab.com/crab/modules/user/aaa"
	"github.com/clickyab/services/assert"
)

const (
	// Foreign is filter for region is every where except iran
	Foreign = "foreign"
)

// AddCampaign for creating campaign with minimum info
func (m *Manager) AddCampaign(c CampaignBase, u *aaa.User, d *dmn.Domain) (*Campaign, error) {
	ca := &Campaign{
		base: base{
			Active: true,
		},
		DomainID: d.ID,
		UserID:   u.ID,
		CampaignFinance: CampaignFinance{

			CostType: CPC,
		},
		CampaignBaseType: CampaignBaseType{

			Type: c.Type,
			Kind: c.Kind,
		},
		CampaignStatus: CampaignStatus{

			Status:  c.Status,
			StartAt: c.StartAt,
			EndAt:   c.EndAt,
			Title:   c.Title,
		},
		Progress: ProgressInProgress,
	}
	switch c.Kind {
	case WebCampaign:
		ca.webMaxBid(c)
	case AppCampaign:
		ca.appMaxBid(c)
	}
	err := m.Begin()
	assert.Nil(err)
	defer func() {
		if err == nil {
			assert.Nil(m.Commit())
			return
		}
		assert.Nil(m.Rollback())
	}()

	if err = m.CreateCampaign(ca); err != nil {
		return nil, err
	}
	s := &Schedule{
		CampaignID:    ca.ID,
		ScheduleSheet: c.Schedule,
	}
	ca.Schedule = s.ScheduleSheet
	err = m.CreateSchedule(s)
	return ca, err
}

// UpdateCampaignByID for updating campaign with minimum info
func (m *Manager) UpdateCampaignByID(c CampaignStatus, ca *Campaign) error {

	ca.Status = c.Status
	if ca.StartAt != c.StartAt {
		if c.StartAt.Unix() < time.Now().Unix() {
			return errors.StartTimeError
		}
	}
	ca.StartAt = c.StartAt
	ca.EndAt = c.EndAt
	ca.Title = c.Title

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

	err = m.attachAttribute(ca)
	if err != nil {
		return err
	}

	err = m.UpdateSchedule(s)
	return err
}

// Finalize will mark campaign ready for publish
func (m *Manager) Finalize(ca *Campaign) error {

	ca.Progress = ProgressFinalized
	err := m.UpdateCampaign(ca)
	if err != nil {
		return err
	}

	err = m.attachAttribute(ca)
	if err != nil {
		return err
	}

	err = m.attachSchedule(ca)

	return err
}

func (m *Manager) attachSchedule(c *Campaign) error {
	s, err := m.FindScheduleByCampaignID(c.ID)
	if err != sql.ErrNoRows {
		return err
	}

	c.Schedule = s.ScheduleSheet

	return nil
}

func (m *Manager) attachAttribute(c *Campaign) error {
	s, err := m.FindCampaignAttributesByCampaignID(c.ID)
	if err != sql.ErrNoRows {
		return err
	}

	c.Attributes = s

	return nil
}
