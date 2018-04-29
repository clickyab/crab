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

// TODO: declared here because of import cycle error
var creativeTableFull = "creatives"

// CampaignDetails is the campaign daily data in data table
// @DataTable {
//		url = /list
//		entity = campaigns
//		checkable = false
//		multiselect = false
//		view = campaign_list:self
//		controller = clickyab.com/crab/modules/campaign/controllers
//		fill = FillCampaigns
//		_detail = campaign_detail:self
//		_edit = campaign_edit:self
//		_copy = campaign_copy:self
//		_archive = campaign_archive:self
// }
type CampaignDetails struct {
	ID             int64        `json:"id" db:"id" type:"number"`
	Title          string       `sort:"true" type:"string" search:"true"  json:"title" db:"title"`
	Status         Status       `sort:"true" type:"enum" filter:"true"  json:"status" db:"status"`
	Kind           CampaignKind `sort:"true" type:"enum" filter:"true"  json:"kind" db:"kind"`
	TotalImp       int64        `sort:"true" type:"number" json:"total_imp" db:"total_imp"`
	TotalClick     int64        `sort:"true" type:"number" json:"total_click" db:"total_click"`
	ECTR           float64      `sort:"true" type:"number" json:"ectr" db:"ectr"`
	ECPC           int64        `sort:"true" type:"number" json:"ecpc" db:"ecpc"`
	ECPM           int64        `sort:"true" type:"number" json:"ecpm" db:"ecpm"`
	TotalSpend     int64        `sort:"true" type:"number" json:"total_spend" db:"total_spend"`
	MaxBid         int64        `sort:"true" type:"number" json:"max_bid" db:"max_bid"`
	Conversion     int64        `sort:"true" type:"number" json:"conversion" db:"conversion"`
	TotalBudget    int64        `sort:"true" type:"number" json:"total_budget" db:"total_budget"`
	TodaySpend     int64        `sort:"true" type:"number" json:"today_spend" db:"today_spend"`
	CreatedAt      time.Time    `sort:"true" type:"date" json:"created_at" db:"created_at"`
	StartAt        time.Time    `sort:"true" type:"date" json:"start_at" db:"start_at"`
	EndAt          time.Time    `sort:"true" type:"date" json:"end_at" db:"end_at"`
	TodayCTR       float64      `sort:"true" type:"number" json:"today_ctr" db:"today_ctr"`
	TodayImp       int64        `sort:"true" type:"number" json:"today_imp" db:"today_imp"`
	TodayClick     int64        `sort:"true" type:"number" json:"today_click" db:"today_click"`
	Creative       int64        `sort:"true" type:"number" json:"creative" db:"creative"`
	OwnerEmail     string       `sort:"true" type:"number" search:"true" json:"owner_email" db:"owner_email"`
	ConversionRate int64        `sort:"true" type:"number" json:"conversion_rate" db:"conversion_rate"`
	CPA            int64        `sort:"true" type:"number" json:"cpa" db:"cpa"`
	Strategy       Strategy     `sort:"true" type:"enum" filter:"true" json:"strategy" db:"strategy"`
	Exchange       ExchangeType `sort:"true" type:"enum" filter:"true" json:"exchange" db:"exchange"`

	OwnerID   int64   `db:"-" json:"owner_id" visible:"false"`
	DomainID  int64   `db:"-" json:"domain_id"`
	ParentIDs []int64 `db:"-" json:"-" visible:"false"`
	Actions   string  `db:"-" json:"_actions" visible:"false"`
}

// FillCampaigns is the function to handle
func (m *Manager) FillCampaigns(
	pc permission.InterfaceComplete,
	filters map[string]string,
	from string,
	to string,
	search map[string]string,
	contextparams map[string]string,
	sort, order string, p, c int) (CampaignDetailsArray, int64, error) {

	// ORDER MATTER
	var params = []interface{}{
		pc.GetID(),
		pc.GetDomainID(),
		libs.TimeToID(time.Now()),
	}
	var where []string

	params = append(params, pc.GetDomainID())
	for field, value := range filters {
		where = append(where, fmt.Sprintf("%s=?", field))
		params = append(params, value)
	}

	for column, val := range search {
		where = append(where, fmt.Sprintf("%s LIKE ?", column))
		params = append(params, "%"+val+"%")
	}

	// find current user childes
	userManager := aaa.NewAaaManager()
	childes := userManager.GetUserChildesIDDomain(pc.GetID(), pc.GetDomainID())
	childes = append(childes, pc.GetID())
	// self or parent
	if pc.GetCurrentScope() == permission.ScopeSelf {
		//check if parent or owner
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
	//check for perm

	var conds string
	if len(where) > 0 {
		conds += wh
	}
	conds += strings.Join(where, " AND ")

	q := fmt.Sprintf(`FROM %s AS c
  JOIN %s u ON c.user_id = ?
  LEFT JOIN %s AS pu ON (pu.user_id = owner.id AND cp.domain_id = ?)
  LEFT JOIN %s AS parent ON parent.id = pu.advisor_id
  JOIN %s cd ON (c.id = cd.campaign_id AND cd.daily_id = ?)
  JOIN %s c2 ON c.id = c2.campaign_id
  LEFT JOIN %s c3 ON (c.id = c3.campaign_id AND c3.status = 'accepted')
	%s
GROUP BY c.id, cd.daily_id, c3.campaign_id `,
		CampaignTableFull, aaa.UserTableFull, aaa.AdvisorTableFull, aaa.UserTableFull,
		aaa.AdvisorTableFull, aaa.AdvisorTableFull, creativeTableFull, conds)

	if sort != "" {
		q += fmt.Sprintf(" ORDER BY %s %s ", sort, order)
	}
	limit := c
	offset := (p - 1) * c
	q += fmt.Sprintf(" LIMIT %d OFFSET %d ", limit, offset)

	fields := fmt.Sprintf(`SELECT
  c.id                                             AS id,
  c.title                                          AS title,
  c.status                                         AS status,
  c.kind                                           AS kind,
  COALESCE(sum(c2.imp), 0)                         AS total_imp,
  COALESCE(sum(c2.click), 0)                       AS total_click,
  COALESCE(avg(c2.imp) / avg(c2.click) * 10, 0)    AS ectr,
  COALESCE(avg(c2.cpc), 0)                         AS ecpc,
  COALESCE(avg(c2.cpm), 0)                         AS ecpm,
  COALESCE(c.total_spend, 0)                       AS total_spend,
  COALESCE(c.max_bid, 0)                           AS max_bid,
  COALESCE(avg(c2.cpa) / avg(c2.click)             AS conversion,
  COALESCE(c.total_budget, 0)                      AS total_budget,
  COALESCE(c.today_spend, 0)                       AS today_spend,
  COALESCE(c.created_at, 0)                        AS created_at,
  COALESCE(c.start_at, 0)                          AS start_at,
  COALESCE(c.end_at, 0)                            AS end_at,
  COALESCE(sum(cd.imp) / sum(cd.click) * 10, 0)    AS today_ctr,
  COALESCE(sum(cd.imp), 0)                         AS today_imp,
  COALESCE(sum(cd.click), 0)                       AS today_click,
  COALESCE(count(c3.id), 0)                        AS creative,
  u.email                                          AS owner_email,
  COALESCE((sum(cs.cpa) * 100) / sum(cd.click), 0) AS conversion_rate,
  COALESCE(sum(cd.cpa), 0)                         AS cpa,
  c.strategy                                       AS strategy,
  c.exchange                                       AS exchange  %s`, q)

	count, err := m.GetRDbMap().SelectInt(fmt.Sprintf(`SELECT COUNT(c.id) %s`, q), params...)
	assert.Nil(err)

	var res CampaignDetailsArray
	_, err = m.GetRDbMap().Select(&res, fields, params...)
	assert.Nil(err)

	return res, count, nil
}
