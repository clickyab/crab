package orm

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"clickyab.com/crab/libs"
	"clickyab.com/crab/modules/user/aaa"
	"github.com/clickyab/services/assert"
	"github.com/clickyab/services/permission"
)

// CampaignGraphDaily is the campaign full data in data table
// @Graph {
//		url = /graph/daily/:id
//		entity = chartdaily
//		view = graph_daily_campaign:self
//		key = ID
//		controller = clickyab.com/crab/modules/campaign/controllers
//		fill = FillCampaignGraphDaily
// }
type CampaignGraphDaily struct {
	ID         int64   `json:"id" db:"id" type:"number"`
	AvgCPC     float64 `json:"avg_cpc" db:"avg_cpc" graph:"avg_cpc,Avg. CPC,line,false,4"`
	AvgCPM     float64 `json:"avg_cpm" db:"avg_cpm" graph:"avg_cpm,Avg. CPM,line,false,5"`
	Ctr        float64 `json:"ctr" db:"ctr" graph:"ctr,CTR,line,false,3"`
	TotalImp   int64   `json:"total_imp" db:"total_imp" graph:"imp,Total Impression,bar,true,2"`
	TotalClick int64   `json:"total_click" db:"total_click" graph:"click,Click,line,true,1"`
	TotalSpent int64   `json:"total_spent" db:"total_spent" graph:"total_spent,Total spent,line,false,6"`
}

// FillCampaignGraphDaily is the function to handle
func (m *Manager) FillCampaignGraphDaily(
	pc permission.InterfaceComplete,
	filters map[string]string,
	search map[string]string,
	contextparams map[string]string,
	from, to time.Time) []CampaignGraphDaily {
	res := make([]CampaignGraphDaily, 0)

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
	var params []interface{}
	params = append(params, pc.GetDomainID())
	//add campaign id
	val, ok := contextparams["id"]
	if !ok {
		return res
	}
	intVal, err := strconv.ParseInt(val, 10, 0)
	if err != nil {
		return res
	}

	where = append(where, "cp.id=?")
	params = append(params, intVal)

	where = append(where, fmt.Sprintf(`%s BETWEEN %d AND %d`, "cd.daily_id",
		libs.TimeToID(from),
		libs.TimeToID(to)))
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

	// find current user childes
	childes := pc.GetChildesPerm(permission.ScopeSelf, "graph_daily_campaign", pc.GetDomainID())
	childes = append(childes, pc.GetID())
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
	_, err = m.GetRDbMap().Select(&res, query, params...)
	assert.Nil(err)

	return res
}
