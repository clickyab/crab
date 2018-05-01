package orm

import (
	"fmt"
	"strings"
	"time"

	"database/sql"

	"clickyab.com/crab/libs"
	"clickyab.com/crab/modules/campaign/errors"
	domainOrm "clickyab.com/crab/modules/domain/orm"
	"clickyab.com/crab/modules/user/aaa"
	"github.com/clickyab/services/assert"
	"github.com/clickyab/services/permission"
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
		fmt.Sprintf("SELECT * FROM %s WHERE id=? AND domain_id=?", CampaignTableFull),
		caID,
		dID,
	)

	if err != nil {
		return nil, err
	}

	return &res, err
}

// FillCampaignGraph is the function to handle
func (m *Manager) FillCampaignGraph(
	pc permission.InterfaceComplete,
	filters map[string]string,
	search map[string]string,
	contextparams map[string]string,
	from, to time.Time) []CampaignGraph {
	res := make([]CampaignGraph, 0)

	query := fmt.Sprintf(`SELECT cd.daily_id as id,
	COALESCE(AVG(cd.cpc),0) AS avg_cpc,
	COALESCE(AVG(cd.cpm),0) AS avg_cpm,
	COALESCE(SUM(cd.click),0) AS total_click,
	COALESCE(SUM(cd.imp),0) AS total_imp,
	COALESCE((SUM(cd.click)/SUM(cd.imp))*10,0) AS ctr,
	COALESCE(SUM(cd.cpc)+SUM(cd.cpm),0) AS total_spent
	FROM %s AS cp INNER JOIN %s AS owner ON owner.id=cp.user_id
	LEFT JOIN %s AS pu ON (pu.user_id=owner.id AND cp.domain_id=?)
	LEFT JOIN %s AS parent ON parent.id=pu.advisor_id
	LEFT JOIN %s AS cd ON cd.campaign_id=cp.id `,
		CampaignTableFull, aaa.UserTableFull, aaa.AdvisorTableFull, aaa.UserTableFull,
		CampaignDetailTableFull)

	var where []string

	where = append(where, fmt.Sprintf(`%s BETWEEN %d AND %d`, "cd.daily_id",
		libs.TimeToID(from),
		libs.TimeToID(to)))
	var params []interface{}
	params = append(params, pc.GetDomainID())

	for field, value := range filters {
		where = append(where, fmt.Sprintf("%s=?", field))
		params = append(params, value)
	}
	for column, val := range search {
		where = append(where, fmt.Sprintf("%s LIKE ?", column))
		params = append(params, "%"+val+"%")
	}
	currentUserID := pc.GetID()
	highestScope := pc.GetCurrentScope()

	// find current user childes
	userManager := aaa.NewAaaManager()
	childes := userManager.GetUserChildesIDDomain(currentUserID, pc.GetDomainID())
	childes = append(childes, currentUserID)
	// self or parent
	if highestScope == permission.ScopeSelf {
		//check if parent or owner
		where = append(where, fmt.Sprintf("cp.user_id IN (%s)",
			func() string {
				return strings.TrimRight(strings.Repeat("?,", len(childes)), ",")
			}(),
		),
		)
		for i := range childes {
			params = append(params, childes[i])
		}

	}
	//check for perm
	if len(where) > 0 {
		query += wh
	}
	query += strings.Join(where, " AND ")

	query += " GROUP BY cd.daily_id"
	_, err := m.GetRDbMap().Select(&res, query, params...)
	assert.Nil(err)

	return res
}

// CampaignProgress campaign progress object
type CampaignProgress struct {
	TotalSpend  int64        `json:"total_spend" db:"total_spend"`
	Click       int64        `json:"click" db:"click"`
	Imp         int64        `json:"imp" db:"imp"`
	DailyBudget int64        `json:"daily_budget" db:"daily_budget"`
	TotalBudget int64        `json:"total_budget" db:"total_budget"`
	Ctr         float64      `json:"ctr" db:"ctr"`
	AvgCPC      float64      `json:"avg_cpc" db:"avg_cpc"`
	MaxBid      int64        `json:"max_bid" db:"max_bid"`
	OwnerEmail  string       `json:"owner_email" db:"owner_email"`
	Status      Status       `json:"status" db:"status"`
	Kind        CampaignKind `json:"kind" db:"kind"`
}

// GetProgressData get progress bar data
func (m *Manager) GetProgressData(campaignID, domainID int64) CampaignProgress {
	var res CampaignProgress
	q := fmt.Sprintf(`SELECT
							c.total_spend AS total_spend,
							COALESCE(SUM(cd.click),0) AS click,
							COALESCE(SUM(cd.imp),0) AS imp,
							c.daily_budget AS daily_budget,
							c.total_budget AS total_budget,
							COALESCE((SUM(cd.click)/SUM(cd.imp))*10,0) AS ctr,
							COALESCE(SUM(cd.cpc)/SUM(cd.click),0) AS avg_cpc,
							c.max_bid AS max_bid,
							owner.email AS owner_email,
							c.status AS status,
							c.kind AS kind
							FROM %s AS c
							INNER JOIN %s AS owner ON owner.id=c.user_id
							LEFT JOIN %s AS cd ON cd.campaign_id=c.id WHERE c.id=? AND c.domain_id=? GROUP BY c.id LIMIT 1
							`, CampaignTableFull, aaa.UserTableFull, CampaignDetailTableFull)

	assert.Nil(m.GetRDbMap().SelectOne(&res, q, campaignID, domainID))
	return res
}
