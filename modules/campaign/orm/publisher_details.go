package orm

import (
	"fmt"
	"strconv"
	"strings"

	"time"

	"clickyab.com/crab/libs"
	"clickyab.com/crab/modules/ad/errors"
	"clickyab.com/crab/modules/user/aaa"
	"github.com/clickyab/services/assert"
	"github.com/clickyab/services/mysql"
	"github.com/clickyab/services/permission"
)

// PublisherDetails is the campaign daily data in data table
// @DataTable {
//		url = /publisher-details/:id
//		entity = publisherDetails
//		checkable = false
//		datefilter = daily_id
//		multiselect = false
//		view = campaign_publisher:self
//		controller = clickyab.com/crab/modules/campaign/controllers
//		fill = FillCampaignPublisher
// }
type PublisherDetails struct {
	Domain         mysql.NullString `type:"string" search:"true" json:"domain" db:"domain"`
	Impression     int64            `sort:"true" type:"number" json:"impression" db:"impression"`
	Click          int64            `sort:"true" type:"number" json:"click" db:"click"`
	ECPC           float64          `sort:"true" type:"number" json:"ecpc" db:"ecpc"`
	ECTR           float64          `sort:"true" type:"number" json:"ectr" db:"ectr"`
	ECPM           float64          `sort:"true" type:"number" json:"ecpm" db:"ecpm"`
	Spend          int64            `sort:"true" type:"number" json:"spend" db:"spend"`
	Conversion     int64            `sort:"true" type:"number" json:"conversion" db:"conversion"`
	ConversionRate float64          `sort:"true" type:"number" json:"conversion_rate" db:"conversion_rate"`
	CPA            float64          `sort:"true" type:"number" json:"cpa" db:"cpa"`

	OwnerID   int64   `db:"-" json:"-" visible:"false"`
	DomainID  int64   `db:"-" json:"-"`
	ParentIDs []int64 `db:"-" json:"-" visible:"false"`
	Actions   string  `db:"-" json:"_actions" visible:"false"`
}

// FillCampaignPublisher is the function to handle
func (m *Manager) FillCampaignPublisher(
	pc permission.InterfaceComplete,
	filters map[string]string,
	from string,
	to string,
	search map[string]string,
	contextparams map[string]string,
	sort, order string, p, c int) (PublisherDetailsArray, int64, error) {

	var params []interface{}
	var res PublisherDetailsArray
	var where []string
	var whereLike []string

	//add inventory
	val, ok := contextparams["id"]
	if !ok {
		return nil, 0, errors.DBError
	}
	intVal, err := strconv.ParseInt(val, 10, 0)
	if err != nil {
		return nil, 0, errors.DBError
	}

	where = append(where, "cd.campaign_id=?")
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
		conds += wh
	}
	conds += strings.Join(where, " AND ")

	countQuery := fmt.Sprintf(`SELECT COUNT(p.id) FROM %s AS c
		INNER JOIN %s AS owner ON owner.id=c.user_id
		LEFT JOIN %s AS cd ON cd.campaign_id=c.id
		LEFT JOIN %s AS p ON p.id=cd.publisher_id %s `,
		CampaignTableFull,
		aaa.UserTableFull,
		CampaignDetailTableFull,
		"publishers",
		conds,
	)
	countQuery += " GROUP BY cd.publisher_id"
	count, err := m.GetRDbMap().SelectInt(countQuery, params...)
	assert.Nil(err)

	query := fmt.Sprintf(`SELECT
		p.domain 											AS domain,
		COALESCE(SUM(cd.imp),0) 							AS impression,
		COALESCE(SUM(cd.click),0) 							AS click,
		COALESCE(AVG(cd.cpc),0) 							AS ecpc,
		COALESCE(AVG(cd.cpm),0) 							AS ecpm,
		COALESCE((SUM(cd.click)/SUM(cd.imp))*10,0)  		AS ectr,
		COALESCE((SUM(cd.cpc)+SUM(cd.cpm)+SUM(cd.cpa)),0) 	AS spend,
		COALESCE(SUM(cd.conv),0) 							AS conversion,
		COALESCE((SUM(cd.conv)/SUM(cd.click))*100,0) 		AS conversion_rate,
		COALESCE(SUM(cd.cpa),0) 							AS cpa
		FROM %s AS cd
		INNER JOIN %s AS c ON c.id=cd.campaign_id
		INNER JOIN %s AS owner ON owner.id=c.user_id
		LEFT JOIN %s AS p ON p.id=cd.publisher_id %s `,
		CampaignDetailTableFull,
		CampaignTableFull,
		aaa.UserTableFull,
		"publishers",
		conds,
	)
	query += " GROUP BY cd.publisher_id"

	limit := c
	offset := (p - 1) * c
	if sort != "" {
		query += fmt.Sprintf(" ORDER BY %s %s ", sort, order)
	}
	query += fmt.Sprintf(" LIMIT %d OFFSET %d ", limit, offset)

	_, err = m.GetRDbMap().Select(&res, query, params...)
	assert.Nil(err)

	return res, count, nil
}
