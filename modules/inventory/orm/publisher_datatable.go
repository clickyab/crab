package orm

import (
	"fmt"
	"strings"

	"github.com/clickyab/services/assert"
	"github.com/clickyab/services/permission"
)

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
//		controller = clickyab.com/crab/modules/inventory/controllers
//		fill = FillPublisherDataTableArray
// }
type PublisherDataTable struct {
	Publisher
	OwnerID   int64   `json:"-" db:"-" visible:"false"`
	DomainID  int64   `json:"-" db:"-" visible:"false"`
	ParentIDs []int64 `json:"-" db:"-" visible:"false"`
	Actions   string  `db:"-" json:"_actions" visible:"false"`
}

// FillPublisherDataTableArray is the function to handle
func (m *Manager) FillPublisherDataTableArray(
	pc permission.InterfaceComplete,
	filters map[string]string,
	from string,
	to string,
	search map[string]string,
	contextparams map[string]string,
	sort, order string, p, c int) (PublisherDataTableArray, int64, error) {
	var params []interface{}
	var res PublisherDataTableArray
	var where []string
	var whereLike []string
	countQuery := fmt.Sprintf("SELECT COUNT(id) FROM %s",
		PublisherTableFull,
	)
	query := fmt.Sprintf("SELECT %s FROM %s",
		getSelectFields(PublisherTableFull, ""),
		PublisherTableFull,
	)
	for field, value := range filters {
		where = append(where, fmt.Sprintf("%s=?", field))
		params = append(params, value)
	}

	//check for date filter
	if from != "" && to != "" {
		fromArr := strings.Split(from, "*")
		toArr := strings.Split(to, "*")
		where = append(where, fmt.Sprintf(`%s BETWEEN ? AND ?`, fromArr[0]))
		params = append(params, fromArr[1], toArr[1])
	}

	for column, val := range search {
		if len(whereLike) == 0 {
			if len(search) == 1 {
				whereLike = append(whereLike, fmt.Sprintf("(%s LIKE ?)", column))
			} else {
				whereLike = append(whereLike, fmt.Sprintf("(%s LIKE ?", column))
			}
		} else if len(whereLike) == len(search)-1 {
			whereLike = append(whereLike, fmt.Sprintf("%s LIKE ?)", column))
		} else {
			whereLike = append(whereLike, fmt.Sprintf("%s LIKE ?", column))
		}
		params = append(params, "%"+val+"%")
	}

	//check for perm
	if len(where)+len(whereLike) > 0 {
		query = fmt.Sprintf("%s %s ", query, " WHERE ")
		countQuery = fmt.Sprintf("%s %s ", countQuery, " WHERE ")
	}
	query += strings.Join(where, " AND ")
	countQuery += strings.Join(where, " AND ")
	if len(where) > 0 && len(whereLike) > 0 {
		query = fmt.Sprintf("%s %s ", query, " AND ")
		countQuery = fmt.Sprintf("%s %s ", countQuery, " AND ")
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

	return res, count, nil
}
