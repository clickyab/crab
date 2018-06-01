package aaa

import (
	"fmt"
	"strings"
	"time"

	"clickyab.com/crab/modules/domain/orm"
	upload "clickyab.com/crab/modules/upload/model"
	"github.com/clickyab/services/assert"
	"github.com/clickyab/services/mysql"
	"github.com/clickyab/services/permission"
)

// UserList is user list data
// @DataTable {
//		url = /list
//		entity = users_list
//		checkable = false
//		checklevel = true
//		preventself = true
//		searchkey = q
//		multiselect = false
//		view = list_user:self
//		datefilter = created_at
//		map_prefix = users
//		controller = clickyab.com/crab/modules/user/controllers
//		fill = FillUsers
//		_edit = edit_user:global
//		_impersonate = impersonate_user:self
//		_change_pass = edit_user:global
// }
type UserList struct {
	ID          int64            `json:"id" db:"id" visible:"true" type:"number"`
	FullName    string           `json:"full_name" db:"full_name" type:"string" search:"true" visible:"true"`
	Status      UserValidStatus  `json:"status" db:"status" type:"enum" filter:"true" visible:"true"`
	Balance     int64            `json:"balance" db:"balance" type:"string" sort:"true" visible:"true"`
	Email       string           `json:"email" db:"email" type:"string" search:"true" visible:"true"`
	CellPhone   mysql.NullString `json:"cellphone" db:"cellphone" type:"string" search:"true" visible:"true"`
	LandLine    mysql.NullString `json:"land_line" db:"land_line" type:"string" search:"true" visible:"true"`
	AccountType AccountType      `json:"account_type" db:"account_type" type:"string" visible:"true"`
	SSN         mysql.NullString `json:"ssn" db:"ssn" type:"string" search:"true" visible:"true"`
	Avatar      mysql.NullString `json:"avatar" db:"avatar" type:"string" visible:"true"`
	CreatedAt   time.Time        `json:"created_at" db:"created_at" type:"string" sort:"true" visible:"true"`

	OwnerID   int64   `db:"owner_id" json:"-" visible:"false"`
	DomainID  int64   `db:"domain_id" json:"-" visible:"false"`
	ParentIDs []int64 `db:"-" json:"-" visible:"false"`
	Actions   string  `db:"-" json:"_actions" visible:"false"`
}

// FillUsers is the function to handle
func (m *Manager) FillUsers(
	pc permission.InterfaceComplete,
	filters map[string]string,
	from string,
	to string,
	search map[string]string,
	contextparams map[string]string,
	sort, order string, p, c int) (UserListArray, int64, error) {

	// ORDER MATTER
	var params = []interface{}{
		pc.GetDomainID(),
	}
	var countParams = []interface{}{
		pc.GetDomainID(),
	}

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
	childes := pc.GetChildesPerm(permission.ScopeSelf, "list_user", pc.GetDomainID())
	childes = append(childes, pc.GetID())
	// self or parent
	if pc.GetCurrentScope() == permission.ScopeSelf {
		//check if parent or owner
		where = append(where, fmt.Sprintf("users.id IN (%s)",
			func() string {
				return strings.TrimRight(strings.Repeat("?,", len(childes)), ",")
			}(),
		),
		)

		for i := range childes {
			params = append(params, childes[i])
			countParams = append(countParams, childes[i])
		}

	}

	var conds string
	if len(where) > 0 {
		conds += " WHERE "
	}

	conds += strings.Join(where, " AND ")

	q := fmt.Sprintf(`
		SELECT
		users.id,
		duser.domain_id AS domain_id,
		full_name,
		users.status,
		balance,
		email,
		cellphone,
		land_line,
		(
			CASE WHEN EXISTS(SELECT NULL FROM %s AS corp WHERE corp.user_id=users.id)
				THEN 'corporation'
				ELSE 'personal'
			END 
		)AS account_type,
		ssn,
		file.id AS avatar,
		users.created_at,
		users.id AS owner_id
		FROM %s AS users
		INNER JOIN %s AS duser ON (duser.user_id = users.id AND duser.domain_id=?)
		LEFT JOIN %s AS file ON file.user_id=users.id AND section='user-avatar'
		%s`,
		CorporationTableFull,
		UserTableFull,
		orm.DomainUserTableFull,
		upload.UploadTableFull,
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
		FROM %s AS users
		INNER JOIN %s AS duser ON (duser.user_id = users.id AND duser.domain_id=?)
		%s`,
		UserTableFull,
		orm.DomainUserTableFull,
		conds,
	)

	count, err := m.GetRDbMap().SelectInt(countQuery, countParams...)
	assert.Nil(err)

	var res UserListArray
	_, err = m.GetRDbMap().Select(&res, q, params...)
	assert.Nil(err)

	return res, count, nil
}
