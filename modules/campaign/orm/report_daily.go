package orm

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"clickyab.com/crab/libs"
	"clickyab.com/crab/modules/ad/errors"
	"clickyab.com/crab/modules/time/tim"
	"clickyab.com/crab/modules/user/aaa"
	"github.com/clickyab/services/assert"
	"github.com/clickyab/services/permission"
)

// CampaignDaily is the campaign daily data in data table
// @DataTable {
//		url = /daily/:id
//		entity = CampaignDaily
//		checkable = false
//		datefilter = daily_id
//		multiselect = false
//		view = campaign_daily:self
//		controller = clickyab.com/crab/modules/campaign/controllers
//		fill = FillDaily
// }
type CampaignDaily struct {
	Date           time.Time `sort:"true" type:"date" json:"date" db:"basis"`
	Impression     int64     `sort:"true" type:"number" json:"impression" db:"impression"`
	Click          int64     `sort:"true" type:"number" json:"click" db:"click"`
	ECTR           float64   `sort:"true" type:"number" json:"ectr" db:"ectr"`
	ECPC           float64   `sort:"true" type:"number" json:"ecpc" db:"ecpc"`
	ECPM           float64   `sort:"true" type:"number" json:"ecpm" db:"ecpm"`
	Spend          int64     `sort:"true" type:"number" json:"spend" db:"spend"`
	Conversion     int64     `sort:"true" type:"number" visible:"false" json:"conversion" db:"conversion"`
	ConversionRate float64   `sort:"true" type:"number" visible:"false" json:"conversion_rate" db:"conversion_rate"`
	CPA            int64     `sort:"true" type:"number" visible:"false" json:"cpa" db:"cpa"`

	OwnerID   int64   `db:"-" json:"-" visible:"false"`
	DomainID  int64   `db:"-" json:"-" `
	ParentIDs []int64 `db:"-" json:"-" visible:"false"`
	Actions   string  `db:"-" json:"-" visible:"false"`
}

// FillDaily is the function to handle
func (m *Manager) FillDaily(
	pc permission.InterfaceComplete,
	filters map[string]string,
	from string,
	to string,
	search map[string]string,
	contextparams map[string]string,
	sort, order string, p, c int) (CampaignDailyArray, int64, error) {

	var params []interface{}
	var res CampaignDailyArray
	var where []string
	var whereLike []string

	val, ok := contextparams["id"]
	if !ok {
		return nil, 0, errors.DBError
	}
	intVal, err := strconv.ParseInt(val, 10, 0)
	if err != nil {
		return nil, 0, errors.DBError
	}

	where = append(where, "c.id=?")
	params = append(params, intVal)

	for field, value := range filters {
		where = append(where, fmt.Sprintf("%s=?", field))
		params = append(params, value)
	}

	//check for date filter
	if from != "" && to != "" {
		fromArr := strings.Split(from, "*")
		toArr := strings.Split(to, "*")
		fromTime, err := time.Parse("2006-01-02 15:04:05", fromArr[1])
		if err != nil {
			return nil, 0, errors.DBError
		}
		toTime, err := time.Parse("2006-01-02 15:04:05", toArr[1])
		if err != nil {
			return nil, 0, errors.DBError
		}
		where = append(where, fmt.Sprintf(`cd.%s BETWEEN ? AND ?`, fromArr[0]))
		params = append(params, libs.TimeToID(fromTime), libs.TimeToID(toTime))
	}

	//check for domain
	where = append(where, fmt.Sprintf("%s=?", "c.domain_id"))
	params = append(params, pc.GetDomainID())

	highestScope := pc.GetCurrentScope()
	if highestScope == permission.ScopeSelf {
		// find current user childes
		childes := pc.GetChildesPerm(permission.ScopeSelf, "campaign_publisher", pc.GetDomainID())
		childes = append(childes, pc.GetID())
		where = append(where, fmt.Sprintf("c.user_id IN (%s)",
			func() string {
				return strings.TrimRight(strings.Repeat("?,", len(childes)), ",")
			}(),
		),
		)
		for i := range childes {
			params = append(params, childes[i])
		}

	}

	for column, val := range search {
		whereLike = append(whereLike, fmt.Sprintf("%s LIKE ?", column))
		params = append(params, "%"+val+"%")
	}

	if len(whereLike) > 0 {
		wl := "(" + strings.Join(whereLike, " OR ") + ")"
		where = append(where, wl)
	}

	var conds string
	if len(where) > 0 {
		conds += " WHERE "
	}
	conds += strings.Join(where, " AND ")

	countQuery := fmt.Sprintf(`SELECT COUNT(da.basis) FROM %s AS c
		INNER JOIN %s AS owner ON owner.id=c.user_id
		LEFT JOIN %s AS cd ON cd.campaign_id=c.id
		JOIN %s AS da ON da.id= cd.daily_id %s GROUP BY da.basis`,

		CampaignTableFull,
		aaa.UserTableFull,
		CampaignDetailTableFull,
		tim.DateTableFull,
		conds,
	)

	query := fmt.Sprintf(`SELECT
		da.basis  											AS basis,
		COALESCE(SUM(cd.imp),0) 							AS impression,
		COALESCE(SUM(cd.click),0) 							AS click,
		COALESCE(AVG(cd.cpc),0) 							AS ecpc,
		COALESCE((SUM(cd.click)/SUM(cd.imp))*10,0) 			AS ectr,
		COALESCE((SUM(cd.cpc)+SUM(cd.cpm)+SUM(cd.cpa)),0) 	AS spend,
		COALESCE(AVG(cd.cpm),0) 							AS ecpm,
		COALESCE(SUM(cd.conv),0) 							AS conversion,
		COALESCE(SUM(cd.cpa),0) 							AS cpa,
		COALESCE((SUM(cd.conv)*100)/SUM(cd.click),0) 		AS conversion_rate
		FROM %s AS c
		INNER JOIN %s AS owner ON owner.id=c.user_id
		LEFT JOIN %s AS cd ON cd.campaign_id=c.id 
		JOIN %s AS da ON da.id= cd.daily_id  %s GROUP BY da.basis`,
		CampaignTableFull,
		aaa.UserTableFull,
		CampaignDetailTableFull,
		tim.DateTableFull,
		conds,
	)

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

	return res, count, nil
}
