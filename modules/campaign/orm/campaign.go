package orm

import (
	"time"

	"errors"

	"database/sql"

	"clickyab.com/crab/modules/domain/dmn"
	"clickyab.com/crab/modules/user/aaa"
	"github.com/clickyab/services/assert"
)

const (
	// Foreign is filter for region is every where except iran
	Foreign = "foreign"
)

// ErrInventoryID of insert or update campaign
var ErrInventoryID = errors.New("there is no inventory with this id")

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
	ca.WhiteBlackType = ClickyabTyp
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

var (
	// ErrorStartDate should raise if campaign start date is not valid
	ErrorStartDate = errors.New("start date can't be past")
)

// UpdateCampaignByID for updating campaign with minimum info
func (m *Manager) UpdateCampaignByID(c CampaignStatus, ca *Campaign) error {

	ca.Status = c.Status
	if ca.StartAt != c.StartAt {
		today, err := time.Parse("02-01-03", time.Now().Format("02-01-03"))
		assert.Nil(err)
		if c.StartAt.Unix() < today.Unix() {
			return ErrorStartDate
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
	if err != nil {
		return err
	}
	s.ScheduleSheet = c.Schedule

	err = m.UpdateCampaign(ca)
	if err != nil {
		return err
	}

	m.attachAttribute(ca)
	err = m.UpdateSchedule(s)
	return err
}

// Finalize will mark campaign ready for publish
func (m *Manager) Finalize(ca *Campaign) {

	ca.Progress = ProgressFinalized
	assert.Nil(m.UpdateCampaign(ca))
	m.attachAttribute(ca)
	m.attachSchedule(ca)

}

func (m *Manager) attachSchedule(c *Campaign) {
	s, err := m.FindScheduleByCampaignID(c.ID)
	assert.Nil(err)
	c.Schedule = s.ScheduleSheet
}

func (m *Manager) attachAttribute(c *Campaign) {
	s, err := m.FindCampaignAttributesByCampaignID(c.ID)
	if err != sql.ErrNoRows {
		assert.Nil(err)
	}
	c.Attributes = s
}
