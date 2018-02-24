package orm

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"clickyab.com/crab/modules/user/aaa"
	"github.com/clickyab/services/assert"
	"github.com/clickyab/services/permission"
)

// CampaignDetail campaign detail model in database
// @Model {
//		table = campaign_detail
//		primary = false, campaign_id,daily_id
// }
type CampaignDetail struct {
	CampaignID int64     `json:"campaign_id" db:"campaign_id"`
	DailyID    int64     `json:"daily_id" db:"daily_id"`
	HourID     int64     `json:"hour_id" db:"hour_id"`
	FakeImp    int64     `json:"fake_imp" db:"fake_imp"`
	FakeClick  int64     `json:"fake_click" db:"fake_click"`
	Imp        int64     `json:"imp" db:"imp"`
	Click      int64     `json:"click" db:"click"`
	CPC        int64     `json:"cpc" db:"cpc"`
	CPA        int64     `json:"cpa" db:"cpa"`
	Conv       int64     `json:"conv" db:"conv"`
	CreatedAt  time.Time `json:"created_at" db:"created_at"`
	UpdatedAt  time.Time `json:"updated_at" db:"updated_at"`
}

// CampaignDailyDataTable is the campaign daily data in data table
// @DataTable {
//		url = /daily/:id
//		entity = campaigndaily
//		checkable = false
//		multiselect = false
//		entity = campaigndaily
//		view = campaign_list:self
//		controller = clickyab.com/crab/modules/campaign/controllers
//		fill = FillCampaignDailyDataTableArray
// }
type CampaignDailyDataTable struct {
	CreatedAt time.Time `json:"created_at" db:"created_at" type:"date" sort:"true"`
	CostType  CostType  `json:"-" db:"cost_type"`

	Imp   int64 `json:"imp" db:"imp" type:"number" sort:"true"`
	Click int64 `json:"click" db:"click" type:"number" sort:"true"`
	Conv  int64 `json:"conv" db:"conv" type:"number" sort:"true"`

	Cpm int64 `json:"cpm" db:"cpm"`
	Cpc int64 `json:"cpc" db:"cpc"`

	Spent int64   `json:"spent" db:"spent" sort:"true"`
	Cpa   int64   `json:"cpa" db:"cpa"`
	Ctr   float64 `json:"ctr" db:"ctr" sort:"true"`

	Actions string `db:"-" json:"_actions" visible:"false"`
}

// FillCampaignDailyDataTableArray is the function to handle
func (m *Manager) FillCampaignDailyDataTableArray(
	pc permission.InterfaceComplete,
	filters map[string]string,
	dateRange map[string]string,
	search map[string]string,
	contextparams map[string]string,
	sort, order string, p, c int) (CampaignDailyDataTableArray, int64) {

	var params []interface{}
	var res CampaignDailyDataTableArray
	var where []string

	val, ok := contextparams["id"]
	if !ok {
		return res, 0
	}
	id, err := strconv.ParseInt(val, 10, 64)
	if err != nil {
		return res, 0
	}

	countQuery := fmt.Sprintf(`SELECT COUNT(cd.daily_id) FROM %s AS cd
	INNER JOIN %s AS cp ON (cp.id=cd.campaign_id AND cp.domain_id=?)`,
		CampaignDetailTableFull, CampaignTableFull)

	query := fmt.Sprintf(`SELECT cd.imp AS imp,
	COALESCE(SUM(cd.click),0) AS click,
	COALESCE(SUM(cd.conv),0) AS conv,
	COALESCE(SUM(cd.cpa),0) AS cpa,
	COALESCE(SUM(cd.cpm),0) AS cpm,
	COALESCE(SUM(cd.cpc),0) AS cpc,
	cd.created_at,
	cp.cost_type,
	COALESCE(SUM(cd.click)*10/SUM(cd.imp),0) AS ctr,
	(CASE WHEN cp.cost_type="cpc" THEN COALESCE(SUM(cd.cpc),0) WHEN cp.cost_type="cpm" THEN COALESCE(SUM(cd.cpm),0) WHEN cp.cost_type="cpa" THEN COALESCE(SUM(cd.cpa),0) END) AS spent
	FROM %s AS cd
	INNER JOIN %s AS cp ON (cp.id=cd.campaign_id AND cp.domain_id=?)`,
		CampaignDetailTableFull, CampaignTableFull)

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
			where = append(where, fmt.Sprintf(`%s BETWEEN "%s" AND "%s"`, dateRangeField, fromTime.Truncate(time.Hour*24).Format("2006-01-02 00:00:00"), toTime.Truncate(time.Hour*24).Format("2006-01-02 00:00:00")))
		}
	}

	params = append(params, pc.GetDomainID())
	params = append(params, id)
	where = append(where, "cd.campaign_id=?")
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

	wh := " WHERE "

	//check for perm
	if len(where) > 0 {
		query += wh
		countQuery += wh
	}
	query += strings.Join(where, " AND ")
	countQuery += strings.Join(where, " AND ")

	countQuery += " GROUP BY cd.daily_id "
	query += " GROUP BY cd.daily_id "

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
	return res, count
}
