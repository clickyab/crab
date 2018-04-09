package orm

import (
	"time"

	"fmt"
	"strings"

	"github.com/clickyab/services/assert"
	"github.com/clickyab/services/mysql"
	"github.com/clickyab/services/permission"
)

// Publisher is model for publishers table in database
// @Model {
//		table = publishers
//		primary =  true, id
//		find_by = id
//		list = yes
// }
type Publisher struct {
	ID         int64                 `json:"id" db:"id" type:"number"`
	Name       string                `json:"name" db:"name" type:"string" search:"true"`
	Domain     string                `json:"domain" db:"domain" type:"string" search:"true"`
	Categories mysql.StringJSONArray `json:"categories" db:"categories" type:"array"`
	Supplier   string                `json:"supplier" db:"supplier" type:"string" search:"true"`
	Kind       PublisherType         `json:"kind" db:"kind" type:"enum" filter:"true"`
	Status     Status                `json:"status" db:"status" type:"enum" filter:"true"`
	CreatedAt  time.Time             `json:"created_at" db:"created_at" type:"date" sort:"true"`
	UpdatedAt  time.Time             `json:"updated_at" db:"updated_at" type:"date"`
	DeletedAt  mysql.NullTime        `json:"deleted_at" db:"deleted_at" type:"date"`
}

// PublisherDataTable is the inventory full data in data table
// @DataTable {
//		url = /publisher/list
//		entity = publisher
//		view = publisher_list:self
//		searchkey = q
//		checkable = true
//		multiselect = true
//		datefilter = created_at
//		map_prefix = publishers
//		_edit = none
//		controller = clickyab.com/crab/modules/inventory/controllers
//		fill = FillPublisherDataTableArray
// }
type PublisherDataTable struct {
	Publisher
	OwnerID   int64   `json:"owner_id" db:"owner_id" visible:"false"`
	DomainID  int64   `json:"domain_id" db:"domain_id" visible:"false"`
	ParentIDs []int64 `json:"parent_ids" db:"parent_ids" visible:"false"`
	Actions   string  `db:"-" json:"_actions" visible:"false"`
}

// FillPublisherDataTableArray is the function to handle
func (m *Manager) FillPublisherDataTableArray(
	pc permission.InterfaceComplete,
	filters map[string]string,
	dateRange map[string]string,
	search map[string]string,
	contextparams map[string]string,
	sort, order string, p, c int) (PublisherDataTableArray, int64) {
	var params []interface{}
	var res PublisherDataTableArray
	var where []string
	var whereLike []string
	countQuery := fmt.Sprintf("SELECT COUNT(id) FROM %s",
		PublisherTableFull,
	)
	query := fmt.Sprintf("SELECT * FROM %s",
		PublisherTableFull,
	)
	for field, value := range filters {
		where = append(where, fmt.Sprintf("%s=?", field))
		params = append(params, value)
	}

	for column, val := range search {
		whereLike = append(whereLike, fmt.Sprintf("%s LIKE ?", column))
		params = append(params, "%"+val+"%")
	}
	//check for perm
	if len(where)+len(whereLike) > 0 {
		query += fmt.Sprintf("%s %s ", query, " WHERE ")
		countQuery += fmt.Sprintf("%s %s ", countQuery, " WHERE ")
	}
	query += strings.Join(where, " AND ")
	countQuery += strings.Join(where, " AND ")
	if len(where) > 0 && len(whereLike) > 0 {
		query += fmt.Sprintf("%s %s ", query, " AND ")
		countQuery += fmt.Sprintf("%s %s ", countQuery, " AND ")
	}
	query += strings.Join(whereLike, " OR ")
	countQuery += strings.Join(whereLike, " OR ")
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

	return res, count
}
