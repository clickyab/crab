package orm

import (
	"fmt"
	"strings"
	"time"

	"clickyab.com/crab/modules/ad/errors"
	"clickyab.com/crab/modules/user/aaa"
	"github.com/clickyab/services/assert"
	"github.com/clickyab/services/permission"
)

// BillingDataTable is the billing data table
// @DataTable {
//		url = /billing
//		entity = billingReport
//		checkable = false
//		datefilter = created_at
//		multiselect = false
//		view = get_billing:self
//		controller = clickyab.com/crab/modules/financial/controllers
//		fill = FillBilling
// }
type BillingDataTable struct {
	ID        int64     `type:"number" visible:"false" json:"id" db:"id"`
	UserID    int64     `type:"number" visible:"true" search:"true" json:"user_id" db:"user_id"`
	PayModel  PayModels `filter:"true" visible:"true" type:"enum" json:"pay_model" db:"pay_model"`
	FirstName string    `visible:"true" type:"string" json:"first_name" search:"true" db:"first_name"`
	LastName  string    `visible:"true" type:"string" json:"last_name" search:"true" db:"last_name"`
	Email     string    `visible:"true" type:"string" search:"true" json:"email" db:"email"`
	Amount    int64     `sort:"true" visible:"true" type:"number" json:"amount" db:"amount"`
	Balance   int64     `sort:"true" visible:"true" type:"number" json:"balance" db:"balance"`

	CreatedAt time.Time `sort:"true" visible:"true" type:"date" json:"created_at" db:"created_at"`

	OwnerID   int64   `type:"number" visible:"true" json:"-" db:"-"`
	DomainID  int64   `db:"-" json:"-"`
	ParentIDs []int64 `db:"-" json:"-" visible:"false"`
	Actions   string  `db:"-" json:"_actions" visible:"false"`
}

// FillBilling is the function to handle
func (m *Manager) FillBilling(
	pc permission.InterfaceComplete,
	filters map[string]string,
	from string,
	to string,
	search map[string]string,
	contextparams map[string]string,
	sort, order string, p, c int) (BillingDataTableArray, int64, error) {

	var params []interface{}
	var res BillingDataTableArray
	var where []string
	var whereLike []string

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
		where = append(where, fmt.Sprintf(`b.%s BETWEEN ? AND ?`, fromArr[0]))
		params = append(params, fromTime, toTime)
	}

	//check for domain
	where = append(where, fmt.Sprintf("%s=?", "b.domain_id"))
	params = append(params, pc.GetDomainID())

	highestScope := pc.GetCurrentScope()
	if highestScope == permission.ScopeSelf {
		// find current user childes
		childes := pc.GetChildesPerm(permission.ScopeSelf, "get_billing", pc.GetDomainID())
		childes = append(childes, pc.GetID())
		where = append(where, fmt.Sprintf("b.user_id IN (%s)",
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

	countQuery := fmt.Sprintf(`SELECT COUNT(b.id) FROM %s AS b
		INNER JOIN %s AS owner ON owner.id=b.user_id
		%s`, BillingTableFull, aaa.UserTableFull, conds)

	query := fmt.Sprintf(`SELECT
		b.id AS id,
		b.pay_model AS pay_model,
		b.user_id AS user_id,
		owner.first_name AS first_name,
		owner.last_name AS last_name,
		owner.email AS email,
		b.created_at AS created_at,
		b.amount AS amount,
		b.balance AS balance
		FROM %s AS b
		INNER JOIN %s AS owner ON owner.id=b.user_id
		%s GROUP BY b.id`,
		BillingTableFull,
		aaa.UserTableFull,
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
