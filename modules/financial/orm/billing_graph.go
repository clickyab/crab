package orm

import (
	"fmt"
	"strings"
	"time"

	"clickyab.com/crab/libs"
	"clickyab.com/crab/modules/campaign/orm"
	"clickyab.com/crab/modules/user/aaa"
	"github.com/clickyab/services/assert"
	"github.com/clickyab/services/permission"
)

// BillingSpendGraph is the billing graph
// @Graph {
//		url = /graph/spend
//		entity = billingGraphReport
//		key = ID
//		view = get_billing:self
//		controller = clickyab.com/crab/modules/financial/controllers
//		fill = FillGraph
// }
type BillingSpendGraph struct {
	ID    int64 `type:"number" visible:"false" json:"id" db:"id"`
	Spend int64 `json:"spend" db:"spend" graph:"spend,Spend,line,false,1"`
}

// FillGraph is the function to handle
func (m *Manager) FillGraph(
	pc permission.InterfaceComplete,
	filters map[string]string,
	search map[string]string,
	contextparams map[string]string,
	from, to time.Time) []BillingSpendGraph {

	var params []interface{}
	var res = make([]BillingSpendGraph, 0)
	var where []string
	where = append(where, fmt.Sprintf(`%s BETWEEN %d AND %d`, "cd.daily_id",
		libs.TimeToID(from),
		libs.TimeToID(to)))
	var whereLike []string

	for field, value := range filters {
		where = append(where, fmt.Sprintf("%s=?", field))
		params = append(params, value)
	}

	//check for domain
	where = append(where, fmt.Sprintf("%s=?", "cp.domain_id"))
	params = append(params, pc.GetDomainID())

	highestScope := pc.GetCurrentScope()
	if highestScope == permission.ScopeSelf {
		// find current user childes
		childes := pc.GetChildesPerm(permission.ScopeSelf, "get_billing", pc.GetDomainID())
		childes = append(childes, pc.GetID())
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

	query := fmt.Sprintf(`SELECT
		cd.daily_id as id,
		COALESCE( SUM(cd.cpc) + SUM(cd.cpm) + SUM(cd.cpa),0) as spend
		FROM %s AS cp
		JOIN %s as cd ON cp.id = cd.campaign_id
		INNER JOIN %s AS owner ON owner.id=cp.user_id
		%s GROUP BY cd.daily_id`,
		orm.CampaignTableFull,
		orm.CampaignDetailTableFull,
		aaa.UserTableFull,
		conds,
	)

	_, err := m.GetRDbMap().Select(&res, query, params...)
	assert.Nil(err)
	return res
}
