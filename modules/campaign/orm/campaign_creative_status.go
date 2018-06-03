package orm

import (
	"fmt"
	"strings"
	"time"

	"clickyab.com/crab/modules/user/aaa"
	"github.com/clickyab/services/assert"
	"github.com/clickyab/services/permission"
)

// TODO added because of cycle error
const creativePendingStatus = "pending"

// CampaignCreativeStatus is the campaign creative status in data table
// @DataTable {
//		url = /status-list
//		entity = campaigns_creative
//		checkable = false
//		preventself = true
//		multiselect = false
//		view = list_campaign:superGlobal
//		controller = clickyab.com/crab/modules/campaign/controllers
//		fill = FillCampaignCreativeStatus
//		_bulk_accept = change_creative_status:superGlobal
//		_accept_reject = change_creatives_status:superGlobal
// }
type CampaignCreativeStatus struct {
	ID            int64        `sort:"true" map:"c.id" json:"id" db:"id" search:"true" visible:"true" type:"number"`
	Title         string       `sort:"true" type:"string" search:"true" visible:"true" json:"title" db:"title"`
	Kind          CampaignKind `sort:"true" type:"enum" filter:"true" visible:"true" json:"kind" db:"kind"`
	CreativeCount int64        `sort:"true" type:"number" visible:"true" json:"creative_count" db:"creative_count"`
	OwnerEmail    string       `sort:"true" type:"number" visible:"true" map:"u.email" json:"owner_email" db:"owner_email"`
	OwnerMobile   string       `sort:"true" type:"number" visible:"true" map:"u.cellphone" json:"owner_mobile" db:"owner_mobile"`
	CreatedAt     time.Time    `sort:"true" type:"date" visible:"true" json:"created_at" db:"created_at"`

	PendingCount int64   `db:"pending_count" json:"-" visible:"false"`
	OwnerID      int64   `db:"owner_id" json:"-" visible:"false"`
	DomainID     int64   `db:"domain_id" json:"-" visible:"false"`
	ParentIDs    []int64 `db:"-" json:"-" visible:"false"`
	Actions      string  `db:"-" json:"_actions" visible:"false"`
}

// FillCampaignCreativeStatus is the function to handle
func (m *Manager) FillCampaignCreativeStatus(
	pc permission.InterfaceComplete,
	filters map[string]string,
	from string,
	to string,
	search map[string]string,
	contextparams map[string]string,
	sort, order string, p, c int) (CampaignCreativeStatusArray, int64, error) {

	// ORDER MATTER
	var params = []interface{}{
		creativePendingStatus,
	}
	var countParams = []interface{}{
		creativePendingStatus,
	}

	var where []string

	for field, value := range filters {
		where = append(where, fmt.Sprintf("c.%s=?", field))
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

	var conds = wh + " c.archived_at IS NULL "
	if len(where) > 0 {
		conds += " AND "
	}
	conds += strings.Join(where, " AND ")
	q := fmt.Sprintf(`
		SELECT
		c.user_id                                        AS owner_id,
		c.domain_id                                      AS domain_id,
		c.id AS id,
    	c.title                                          AS title,
    	c.kind                                           AS kind,
    	COALESCE(COUNT(cr.id), 0)		                 AS creative_count,
    	u.email                                          AS owner_email,
    	u.cellphone										 AS owner_mobile,
    	c.created_at                     				 AS created_at,
    	COALESCE(COUNT(cr.id), 0)					     AS pending_count
    	FROM %s AS c
    	JOIN %s u ON u.id = c.user_id
    	INNER JOIN %s cr ON (c.id = cr.campaign_id AND cr.status = ?)
    	%s
    	GROUP BY c.id`,
		CampaignTableFull,
		aaa.UserTableFull,
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
		`SELECT COUNT(distinct(c.id))
				FROM %s AS c
				INNER JOIN %s cr ON (c.id = cr.campaign_id AND cr.status = ?)
				%s`,
		CampaignTableFull,
		creativeTableFull,
		conds,
	)
	count, err := m.GetRDbMap().SelectInt(countQuery, countParams...)
	assert.Nil(err)

	var res CampaignCreativeStatusArray
	_, err = m.GetRDbMap().Select(&res, q, params...)
	assert.Nil(err)

	return res, count, nil
}
