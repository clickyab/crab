package orm

import (
	"fmt"
	"strings"
	"time"

	"clickyab.com/crab/libs"
	"clickyab.com/crab/modules/ad/errors"
	caOrm "clickyab.com/crab/modules/campaign/orm"
	"clickyab.com/crab/modules/user/aaa"
	"github.com/clickyab/services/assert"
	"github.com/clickyab/services/permission"
)

// PublishersBaseStatistics is the campaign full data in data table
// @DataTable {
//		url = /base-publishers/statistics
//		entity = base_publisher_statistics
//		view = publisher_base_statistics:self
//		searchkey = q
//		checkable = false
//		multiselect = false
//		datefilter = created_at
//		map_prefix = pub
//		controller = clickyab.com/crab/modules/inventory/controllers
//		fill = FillPublishersBaseStatistics
// }
type PublishersBaseStatistics struct {
	Count         int64   `json:"count" db:"count" type:"number"`
	AvgImp        float64 `json:"avg_imp" db:"avg_imp"`
	ExchangeCount int64   `json:"exchange_count" db:"exchange_count" type:"number"`

	Name     string        `json:"name" db:"-" type:"string" search:"true" visible:"false"`
	Domain   string        `json:"domain" db:"-" type:"string" search:"true" visible:"false"`
	Supplier string        `json:"supplier" db:"-" type:"string" search:"true" visible:"false"`
	Kind     PublisherType `json:"kind" db:"-" type:"enum" filter:"true" visible:"false"`
	Status   Status        `json:"status" db:"-" type:"enum" filter:"true" visible:"false"`

	Actions string `db:"-" json:"-" visible:"false"`
}

// FillPublishersBaseStatistics is the function to handle
func (m *Manager) FillPublishersBaseStatistics(
	pc permission.InterfaceComplete,
	filters map[string]string,
	from string,
	to string,
	search map[string]string,
	contextparams map[string]string,
	sort, order string, p, c int) (PublishersBaseStatisticsArray, int64, error) {

	var res PublishersBaseStatisticsArray
	query := fmt.Sprintf(`
		SELECT 
		COUNT(DISTINCT(pub.id))						AS count,
		COALESCE((SUM(cd.imp)/(COUNT(1)/30)),0) 	AS avg_imp,
		COALESCE(COUNT(DISTINCT(camp.exchange)),0)	AS exchange_count

		FROM %s AS pub
		
		INNER JOIN %s AS cd ON cd.publisher_id=pub.id AND  cd.daily_id BETWEEN ? AND ?
		INNER JOIN %s AS camp ON cd.campaign_id=camp.id AND camp.domain_id = ?
		`,
		PublisherTableFull,
		caOrm.CampaignDetailTableFull,
		caOrm.CampaignTableFull,
	)

	var where []string
	var params []interface{}

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
		params = append(params, libs.TimeToID(fromTime), libs.TimeToID(toTime))
	}

	params = append(params, pc.GetDomainID())

	for field, value := range filters {
		where = append(where, fmt.Sprintf("%s=?", field))
		params = append(params, value)
	}

	var whereLike []string
	for column, val := range search {
		whereLike = append(whereLike, fmt.Sprintf("%s LIKE ?", column))
		params = append(params, "%"+val+"%")
	}
	if len(whereLike) > 0 {
		wl := "(" + strings.Join(whereLike, " OR ") + ")"
		where = append(where, wl)
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
		where = append(where, fmt.Sprintf("camp.user_id IN (%s)",
			func() string {
				return strings.TrimRight(strings.Repeat("?,", len(childes)), ",")
			}(),
		),
		)
		for i := range childes {
			params = append(params, childes[i])
		}

	}

	query += fmt.Sprintf("%s %s", " WHERE ", strings.Join(where, " AND "))

	_, err := m.GetRDbMap().Select(&res, query, params...)
	assert.Nil(err)

	return res, 1, nil
}
