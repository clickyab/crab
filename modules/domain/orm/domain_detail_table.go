package orm

import (
	"fmt"
	"strings"

	"github.com/clickyab/services/assert"
	"github.com/clickyab/services/mysql"
	"github.com/clickyab/services/permission"
)

// FIXME : added because of import cycle error
const (
	usersTable        = "users"
	corporationsTable = "corporations"
	rolesTable        = "roles"
	DomainOwnerRole   = "Owner"
)

// DomainDetails is the domain daily data in data table
// @DataTable {
//		url = /list
//		entity = domains_data_table
//		checkable = false
//		multiselect = false
//		view = list_domain:superGlobal
//		controller = clickyab.com/crab/modules/domain/controllers
//		fill = FillDomainDetails
//		_detail = get_domain:superGlobal
//		_edit = edit_domain:superGlobal
// }
type DomainDetails struct {
	ID              int64            `sort:"true" json:"id" db:"id" visible:"true" type:"number"`
	Title           string           `sort:"true" json:"title" db:"title" visible:"true" type:"string" search:"true"`
	Status          DomainStatus     `sort:"true" filter:"true" json:"status" db:"status" visible:"true" type:"enum"`
	CorporationName mysql.NullString `sort:"true" json:"corporation_name" db:"corporation_name" visible:"true" type:"string"`
	DomainBase      string           `sort:"true" json:"domain_base" db:"domain_base" visible:"true" type:"string" search:"true" search:"true"`
	OwnerEmail      string           `sort:"true" json:"owner_email" map:"u.email" db:"owner_email" visible:"true" type:"string" search:"true"`
	Balance         int64            `sort:"true" json:"balance" db:"balance" visible:"true" type:"number"`

	OwnerID   int64   `db:"owner_id" json:"-" visible:"false"`
	DomainID  int64   `db:"domain_id" json:"-" visible:"false"`
	ParentIDs []int64 `db:"-" json:"-" visible:"false"`
	Actions   string  `db:"-" json:"_actions" visible:"false"`
}

// FillDomainDetails is the function to handle
func (m *Manager) FillDomainDetails(
	pc permission.InterfaceComplete,
	filters map[string]string,
	from string,
	to string,
	search map[string]string,
	contextparams map[string]string,
	sort, order string, p, c int) (DomainDetailsArray, int64, error) {

	// ORDER MATTER
	var params = []interface{}{
		DomainOwnerRole,
	}
	var countParams = []interface{}{
		DomainOwnerRole,
	}

	var where []string

	for field, value := range filters {
		where = append(where, fmt.Sprintf("d.%s=?", field))
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

	var conds string
	if len(where) > 0 {
		conds += " WHERE "
	}
	conds += strings.Join(where, " AND ")

	q := fmt.Sprintf(`
		SELECT
		  d.id          AS id,
		  d.id          AS domain_id,
		  u.id          AS owner_id,
		  d.title		AS title,
		  d.status		AS status,
		  c.legal_name	AS corporation_name,
		  d.domain_base AS domain_base,
		  u.email		AS owner_email,
		  u.balance		AS balance
		FROM %s AS d
		INNER JOIN %s AS du ON (d.id = du.domain_id)
		INNER JOIN %s AS u ON (du.user_id = u.id)
		INNER JOIN %s AS r ON (du.role_id = r.id AND r.name = ?)
		LEFT JOIN %s AS c ON u.id = c.user_id
		%s
		`,
		DomainTableFull,
		DomainUserTableFull,
		usersTable,
		rolesTable,
		corporationsTable,
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
		FROM %s AS d
		INNER JOIN %s AS du ON (d.id = du.domain_id)
		INNER JOIN %s AS u ON (du.user_id = u.id)
		INNER JOIN %s AS r ON (du.role_id = r.id AND r.name = ?)
		LEFT JOIN %s AS c ON u.id = c.user_id
		%s`,

		DomainTableFull,
		DomainUserTableFull,
		usersTable,
		rolesTable,
		corporationsTable,
		conds,
	)

	count, err := m.GetRDbMap().SelectInt(countQuery, countParams...)
	assert.Nil(err)

	var res DomainDetailsArray
	_, err = m.GetRDbMap().Select(&res, q, params...)
	assert.Nil(err)

	return res, count, nil
}
