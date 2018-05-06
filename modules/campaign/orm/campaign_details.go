package orm

import (
	"fmt"
	"strings"
	"time"

	"clickyab.com/crab/libs"
	"clickyab.com/crab/modules/user/aaa"
	"github.com/clickyab/services/assert"
	"github.com/clickyab/services/mysql"
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
//		_detail = get_campaign:self
//		_edit = edit_campaign:self
//		_copy = copy_campaign:self
//		_archive = archive_campaign:self
// }
type CampaignDetails struct {
	ID             int64          `json:"id" db:"id" visible:"true" type:"number"`
	Title          string         `sort:"true" type:"string" search:"true" visible:"true" json:"title" db:"title"`
	Status         Status         `sort:"true" type:"enum" filter:"true" visible:"true"  json:"status" db:"status"`
	Kind           CampaignKind   `sort:"true" type:"enum" filter:"true" visible:"true" json:"kind" db:"kind"`
	TotalImp       int64          `sort:"true" type:"number" visible:"true" json:"total_imp" db:"total_imp"`
	TotalClick     int64          `sort:"true" type:"number" json:"total_click" db:"total_click"`
	ECTR           float64        `sort:"true" type:"number" visible:"true" visible:"true" json:"ectr" db:"ectr"`
	ECPC           float64        `sort:"true" type:"number" visible:"true" json:"ecpc" db:"ecpc"`
	ECPM           float64        `sort:"true" type:"number" visible:"true" json:"ecpm" db:"ecpm"`
	TotalSpend     int64          `sort:"true" type:"number" visible:"true" json:"total_spend" db:"total_spend"`
	MaxBid         int64          `sort:"true" type:"number" visible:"false" json:"max_bid" db:"max_bid"`
	Conversion     float64        `sort:"true" type:"number" visible:"false" json:"conversion" db:"conversion"`
	TotalBudget    int64          `sort:"true" type:"number" visible:"false" json:"total_budget" db:"total_budget"`
	TodaySpend     int64          `sort:"true" type:"number" visible:"false" json:"today_spend" db:"today_spend"`
	CreatedAt      time.Time      `sort:"true" type:"date"  visible:"false" json:"created_at" db:"created_at"`
	StartAt        time.Time      `sort:"true" type:"date"  visible:"false" json:"start_at" db:"start_at"`
	EndAt          mysql.NullTime `sort:"true" type:"date"  visible:"false" json:"end_at" db:"end_at"`
	TodayCTR       float64        `sort:"true" type:"number"  visible:"false" json:"today_ctr" db:"today_ctr"`
	TodayImp       int64          `sort:"true" type:"number"  visible:"false" json:"today_imp" db:"today_imp"`
	TodayClick     int64          `sort:"true" type:"number"  visible:"false" json:"today_click" db:"today_click"`
	Creative       int64          `sort:"true" type:"number"  visible:"false" json:"creative" db:"creative"`
	OwnerEmail     string         `sort:"true" type:"number"  visible:"false" search:"true" map:"u.email" json:"owner_email" db:"owner_email"`
	ConversionRate float64        `sort:"true" type:"number"  visible:"false" json:"conversion_rate" db:"conversion_rate"`
	CPA            int64          `sort:"true" type:"number"  visible:"false" json:"cpa" db:"cpa"`
	Strategy       Strategy       `sort:"true" type:"enum" visible:"false" filter:"true" json:"strategy" db:"strategy"`
	Exchange       ExchangeType   `sort:"true" type:"enum" visible:"false" filter:"true" json:"exchange" db:"exchange"`

	OwnerID   int64   `db:"owner_id" json:"-" visible:"false"`
	DomainID  int64   `db:"domain_id" json:"-" visible:"false"`
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
	queryParams map[string]string,
	sort, order string, p, c int) (CampaignDetailsArray, int64, error) {

	// ORDER MATTER
	var params = []interface{}{
		pc.GetDomainID(),
		libs.TimeToID(time.Now()),
	}
	var countParams []interface{}

	var where []string

	for field, value := range filters {
		where = append(where, fmt.Sprintf("%s=?", field))
		params = append(params, value)
		countParams = append(countParams, value)
	}

	var whereLike []string
	for column, val := range search {
		whereLike = append(whereLike, fmt.Sprintf("%s LIKE ?", column))
		params = append(params, "%"+val+"%")
		countParams = append(countParams, "%"+val+"%")
	}
	if len(whereLike) > 0 {
		wl := "(" + strings.Join(whereLike, " OR ") + ")"
		where = append(where, wl)
	}

	//check for perm
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
		params = append(params, strings.Trim(strings.Join(strings.Fields(fmt.Sprint(childes)), ","), "[]"))
		countParams = append(countParams, strings.Trim(strings.Join(strings.Fields(fmt.Sprint(childes)), ","), "[]"))
	}

	var conds = wh + " c.archived_at IS NULL "
	if len(where) > 0 {
		conds += " AND "
	}
	conds += strings.Join(where, " AND ")

	q := fmt.Sprintf(`
		SELECT
		c.user_id                                             AS owner_id,
		c.id                                             AS id,
		c.title                                          AS title,
		c.status                                         AS status,
		c.kind                                           AS kind,
		COALESCE(SUM(c2.imp), 0)                         AS total_imp,
		COALESCE(SUM(c2.click), 0)                       AS total_click,
		COALESCE((SUM(c2.click)/SUM(c2.imp))*10,0) 	     AS ectr,
		COALESCE(SUM(c2.cpc)/SUM(c2.click),0)            AS ecpc,
		COALESCE(SUM(c2.cpm)/SUM(c2.imp),0)              AS ecpm,
		COALESCE(c.total_spend, 0)                       AS total_spend,
		COALESCE(c.max_bid, 0)                           AS max_bid,
		COALESCE(SUM(c2.cpa) / SUM(c2.click),0)          AS conversion,
		c.total_budget			                         AS total_budget,
		COALESCE(c.today_spend, 0)                       AS today_spend,
		c.created_at                     				 AS created_at,
		c.start_at                         				 AS start_at,
		c.end_at                            			 AS end_at,
		COALESCE(SUM(cd.click)/SUM(cd.imp)*10, 0)    	 AS today_ctr,
		COALESCE(SUM(cd.imp), 0)                         AS today_imp,
		COALESCE(SUM(cd.click), 0)                       AS today_click,
		COALESCE(COUNT(c3.id), 0)                        AS creative,
		u.email                                          AS owner_email,
		COALESCE((SUM(cd.cpa)*100)/SUM(cd.click), 0) 	 AS conversion_rate,
		COALESCE(SUM(cd.cpa), 0)                         AS cpa,
		c.strategy                                       AS strategy,
		c.exchange                                       AS exchange
		FROM %s AS c
		JOIN %s u ON u.id = c.user_id
		LEFT JOIN %s AS pu ON (pu.user_id = u.id AND c.domain_id = ?)
		LEFT JOIN %s cd ON (c.id = cd.campaign_id  AND cd.daily_id = ?)
		LEFT JOIN %s c2 ON (c.id = c2.campaign_id)
		LEFT JOIN %s c3 ON (c.id = c3.campaign_id AND c3.status = 'accepted')
			%s
		GROUP BY c.id`,
		CampaignTableFull,
		aaa.UserTableFull,
		aaa.AdvisorTableFull,
		CampaignDetailTableFull,
		CampaignDetailTableFull,
		creativeTableFull,
		conds,
	)

	if sort != "" {
		q += fmt.Sprintf(" ORDER BY %s %s ", sort, order)
	}

	limit := c
	offset := (p - 1) * c
	q += fmt.Sprintf(" LIMIT %d OFFSET %d ", limit, offset)

	countQuery := fmt.Sprintf(
		`
		SELECT COUNT(1)
		FROM %s AS c
		JOIN %s u ON u.id = c.user_id
		%s`,
		CampaignTableFull,
		aaa.UserTableFull,
		conds,
	)

	count, err := m.GetRDbMap().SelectInt(countQuery, countParams...)
	assert.Nil(err)

	var res CampaignDetailsArray
	_, err = m.GetRDbMap().Select(&res, q, params...)
	assert.Nil(err)

	return res, count, nil
}
