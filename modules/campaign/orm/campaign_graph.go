package orm

import (
	"fmt"
	"strings"
	"time"

	"clickyab.com/crab/libs"
	"clickyab.com/crab/modules/user/aaa"
	"github.com/clickyab/services/assert"
	"github.com/clickyab/services/permission"
)

// CampaignGraph is the campaign full data in data table
// @Graph {
//		url = /graph/all
//		entity = chartall
//		view = graph_campaign:self
//		key = ID
//		controller = clickyab.com/crab/modules/campaign/controllers
//		fill = FillCampaignGraph
// }
type CampaignGraph struct {
	OwnerEmail string       `db:"owner_email" json:"owner_email" type:"string" search:"true" map:"owner.email"`
	Kind       CampaignKind `json:"kind" db:"kind" type:"enum" filter:"true" map:"cp.kind"`
	Title      string       `json:"title" db:"title" type:"string" search:"true" map:"cp.title"`

	ID         int64   `json:"id" db:"id" type:"number"`
	AvgCPC     float64 `json:"avg_cpc" db:"avg_cpc" graph:"avg_cpc,Avg. CPC,line,true,4"`
	AvgCPM     float64 `json:"avg_cpm" db:"avg_cpm" graph:"avg_cpm,Avg. CPM,line,true,5"`
	Ctr        float64 `json:"ctr" db:"ctr" graph:"ctr,CTR,line,true,3"`
	TotalImp   int64   `json:"total_imp" db:"total_imp" graph:"imp,Total Impression,bar,false,2"`
	TotalClick int64   `json:"total_click" db:"total_click" graph:"click,Click,line,false,1"`
	TotalSpent int64   `json:"total_spent" db:"total_spent" graph:"total_spent,Total spent,line,true,6"`
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
	COALESCE(SUM(cd.cpc)/SUM(cd.click),0) AS avg_cpc,
	COALESCE((SUM(cd.cpm)/SUM(cd.imp))*1000,0) AS avg_cpm,
	COALESCE(SUM(cd.click),0) AS total_click,
	COALESCE(SUM(cd.imp),0) AS total_imp,
	COALESCE((SUM(cd.click)/SUM(cd.imp))*10,0) AS ctr,
	COALESCE(SUM(cd.cpc)+SUM(cd.cpm)+SUM(cd.cpa),0) AS total_spent
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
	//check for domain
	where = append(where, fmt.Sprintf("%s=?", "cp.domain_id"))
	params = append(params, pc.GetDomainID())

	for field, value := range filters {
		where = append(where, fmt.Sprintf("%s=?", field))
		params = append(params, value)
	}
	for column, val := range search {
		where = append(where, fmt.Sprintf("%s LIKE ?", column))
		params = append(params, "%"+val+"%")
	}
	highestScope := pc.GetCurrentScope()

	// self or parent
	if highestScope == permission.ScopeSelf {
		//check if parent or owner
		// find current user childes
		childes := pc.GetChildesPerm(permission.ScopeSelf, "campaign_graph", pc.GetDomainID())
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
