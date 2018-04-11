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
	"github.com/clickyab/services/config"
	"github.com/clickyab/services/permission"
)

var (
	defaultCTR = config.RegisterFloat64("crab.modules.campaign.ctr", .1, "default ctr in the system")
)

const (
	// Foreign is filter for region is every where except iran
	Foreign = "foreign"
)

// AddCampaign for creating campaign with minimum info
func (m *Manager) AddCampaign(c CampaignBase, u *aaa.User, d *domainOrm.Domain) (Campaign, Schedule, error) {
	ca := Campaign{
		Title:    c.Title,
		DomainID: d.ID,
		UserID:   u.ID,
		Kind:     c.Kind,
		Exchange: All,
		TLD:      c.TLD,
		Strategy: CPC,
		Status:   c.Status,
		StartAt:  c.StartAt,
		EndAt:    c.EndAt,
		Progress: ProgressInProgress,
	}
	var s Schedule

	err := m.Begin()
	assert.Nil(err)
	defer func() {
		if err == nil {
			assert.Nil(m.Commit())
			return
		}
		assert.Nil(m.Rollback())
	}()

	if err = m.CreateCampaign(&ca); err != nil {
		return ca, s, err
	}

	s = Schedule{
		CampaignID:    ca.ID,
		ScheduleSheet: c.Schedule,
	}
	err = m.CreateSchedule(&s)

	return ca, s, err
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
	LEFT JOIN %s AS parent ON parent.id=pu.parent_id
	LEFT JOIN %s AS cd ON cd.campaign_id=cp.id `,
		CampaignTableFull, aaa.UserTableFull, aaa.AdvisorTableFull, aaa.UserTableFull, CampaignDetailTableFull)

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

// FillCampaignDataTableArray is the function to handle
func (m *Manager) FillCampaignDataTableArray(
	pc permission.InterfaceComplete,
	filters map[string]string,
	dateRange map[string]string,
	search map[string]string,
	contextparams map[string]string,
	sort, order string, p, c int) (CampaignDataTableArray, int64) {
	var params []interface{}
	var res CampaignDataTableArray
	var where []string
	todayInt := libs.TimeToID(time.Now())
	countQuery := fmt.Sprintf(`SELECT COUNT(cp.id) FROM %s AS cp
	INNER JOIN %s AS owner ON owner.id=cp.user_id
	LEFT JOIN %s AS pu ON (pu.user_id=owner.id AND cp.domain_id=?)
	LEFT JOIN %s AS parent ON parent.id=pu.parent_id
	LEFT JOIN %s AS cd ON cd.campaign_id=cp.id
	LEFT JOIN %s AS ycd ON (ycd.campaign_id=cp.id AND ycd.daily_id=%d)`,
		CampaignTableFull, aaa.UserTableFull, aaa.AdvisorTableFull, aaa.UserTableFull, CampaignDetailTableFull, CampaignDetailTableFull, todayInt)
	query := fmt.Sprintf(`SELECT cp.id AS id,
	cp.title,
	cp.kind,
	cp.daily_limit,
	cp.type,
	cp.status,
	cp.max_bid,
	cp.cost_type,
	cp.budget,
	cp.start_at,
	cp.end_at,
	cp.created_at,
	cp.domain_id AS domain_id,
	owner.email AS owner_email,
	owner.id AS owner_id,
	parent.email AS parent_email,
	COALESCE(AVG(cd.cpc),0) AS avg_cpc,
	COALESCE(AVG(cd.cpm),0)AS avg_cpm,
	COALESCE(SUM(cd.click),0) AS total_click,
	COALESCE(SUM(cd.imp),0) AS total_imp,
	COALESCE(SUM(cd.conv),0) AS total_conv,
	COALESCE(SUM(cd.cpc),0) AS total_cpc,
	COALESCE(SUM(cd.cpm),0) AS total_cpm,
	COALESCE(ycd.imp,0) AS today_imp,
	COALESCE(ycd.click,0) AS today_click
	FROM %s AS cp INNER JOIN %s AS owner ON owner.id=cp.user_id
	LEFT JOIN %s AS pu ON (pu.user_id=owner.id AND cp.domain_id=?)
	LEFT JOIN %s AS parent ON parent.id=pu.parent_id
	LEFT JOIN %s AS cd ON cd.campaign_id=cp.id
	LEFT JOIN %s AS ycd ON (ycd.campaign_id=cp.id AND ycd.daily_id=%d)`,
		CampaignTableFull, aaa.UserTableFull, aaa.AdvisorTableFull, aaa.UserTableFull, CampaignDetailTableFull, CampaignDetailTableFull, todayInt)

	//check for date range
	var dateRangeField string
	var from string
	var to string
	for key, val := range dateRange {
		dateRangeArr := strings.Split(key, "-")
		if len(dateRangeArr) == 2 {
			dateRangeField = dateRangeArr[1]
			if dateRangeArr[0] == "from" {
				from = val
			}
			if dateRangeArr[0] == "to" {
				to = val
			}
		}
	}
	if dateRangeField != "" && from != "" && to != "" {
		fromTime, err1 := time.Parse(time.RFC3339, from)
		toTime, err2 := time.Parse(time.RFC3339, to)

		if err1 == nil && err2 == nil {
			where = append(where,
				fmt.Sprintf(`%s BETWEEN "%s" AND "%s"`, dateRangeField,
					fromTime.Truncate(time.Hour*24).Format("2006-01-02 00:00:00"),
					toTime.Truncate(time.Hour*24).Format("2006-01-02 00:00:00")))
		}
	}

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
		countQuery += wh
	}
	query += strings.Join(where, " AND ")
	countQuery += strings.Join(where, " AND ")

	countQuery += " GROUP BY cp.id "
	query += " GROUP BY cp.id "

	limit := c
	offset := (p - 1) * c
	if sort != "" {
		query += fmt.Sprintf(" ORDER BY %s %s ", sort, order)
	}
	query += fmt.Sprintf(" LIMIT %d OFFSET %d ", limit, offset)
	count, err := m.GetRDbMap().SelectInt(countQuery, params...)
	assert.Nil(err)

	_, err = m.GetRDbMap().Select(&res, query, params...)
	assert.Nil(err)
	for i := range res {
		res[i].Ctr = calculateCtr(res[i].TotalImp, res[i].TotalClick, defaultCTR.Float64())
		res[i].TodayCtr = calculateCtr(res[i].TodayImp, res[i].TodayClick, defaultCTR.Float64())
		res[i].TotalSpent = func() int64 {
			if res[i].CostType == CPC {
				return res[i].TotalCpc
			}
			return res[i].TotalCpm
		}()
	}

	return res, count
}

func calculateCtr(imp, click int64, def float64) float64 {
	if imp == 0 || click == 0 {
		return def
	}
	return float64(click) * 10 / float64(imp)
}
