package orm

import (
	"strconv"
	"time"

	"fmt"
	"strings"

	"clickyab.com/crab/modules/ad/errors"
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
	ID         int64                 `json:"id" db:"id" type:"number" sort:"true"`
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
		fromArr := strings.Split(from, ":")
		toArr := strings.Split(to, ":")
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

// ListPublisherDomainsByIDs list publisher domain based on their ids
func (m *Manager) ListPublisherDomainsByIDs(IDs []int64) []string {
	var final []string
	publishers := m.ListPublishersWithFilter(
		fmt.Sprintf("id IN (%s)", strings.TrimRight(strings.Repeat("?,", len(IDs)), ",")),
		func() []interface{} {
			var res = make([]interface{}, len(IDs))
			for i := range IDs {
				res[i] = IDs[i]
			}
			return res
		}()...,
	)
	for j := range publishers {
		final = append(final, publishers[j].Domain)
	}

	return final
}

// GetValidPubs get valid publisher base on pub ids
func (m *Manager) GetValidPubs(pubIDs []int64) []Publisher {
	var validPublishers []Publisher
	bind := strings.TrimRight(strings.Repeat("?,", len(pubIDs)), ",")
	validPublishers = m.ListPublishersWithFilter(fmt.Sprintf("id IN (%s)", bind),
		func() []interface{} {
			var res = make([]interface{}, len(pubIDs))
			for i := range pubIDs {
				res[i] = pubIDs[i]
			}
			return res
		}()...,
	)
	return validPublishers
}

// SinglePublisherDataTable is the single publisher
// @DataTable {
//		url = /publisher/list/single/:id
//		entity = invpublisher
//		view = list_inventory:self
//		searchkey = q
//		checkable = true
//		multiselect = true
//		datefilter = created_at
//		map_prefix = p
//		controller = clickyab.com/crab/modules/inventory/controllers
//		fill = FillSinglePublisherDataTableArray
// }
type SinglePublisherDataTable struct {
	Publisher
	OwnerID   int64   `json:"-" db:"-" visible:"false"`
	DomainID  int64   `json:"-" db:"-" visible:"false"`
	ParentIDs []int64 `json:"-" db:"-" visible:"false"`
	Actions   string  `db:"-" json:"_actions" visible:"false"`
}

// FillSinglePublisherDataTableArray is the function to handle
func (m *Manager) FillSinglePublisherDataTableArray(
	pc permission.InterfaceComplete,
	filters map[string]string,
	from string,
	to string,
	search map[string]string,
	contextparams map[string]string,
	sort, order string, p, c int) (SinglePublisherDataTableArray, int64, error) {
	var params []interface{}
	var res SinglePublisherDataTableArray
	var where []string
	var whereLike []string
	countQuery := fmt.Sprintf(`SELECT COUNT(p.id) FROM %s AS p
		INNER JOIN %s AS ip ON ip.publisher_id=p.id
		INNER JOIN %s AS i ON i.id=ip.inventory_id`,
		PublisherTableFull,
		InventoryPublisherTableFull,
		InventoryTableFull,
	)
	query := fmt.Sprintf(`SELECT p.id,
		p.name,
		p.domain,
		p.created_at,
		p.updated_at,
		p.categories,
		p.supplier,
		p.kind,
		p.status,
		p.deleted_at FROM %s AS p
		INNER JOIN %s AS ip ON ip.publisher_id=p.id
		INNER JOIN %s AS i ON i.id=ip.inventory_id`,
		PublisherTableFull,
		InventoryPublisherTableFull,
		InventoryTableFull,
	)

	//add inventory
	val, ok := contextparams["id"]
	if !ok {
		return nil, 0, errors.DBError
	}
	intVal, err := strconv.ParseInt(val, 10, 0)
	if err != nil {
		return nil, 0, errors.DBError
	}

	where = append(where, "i.id=?")
	params = append(params, intVal)

	for field, value := range filters {
		where = append(where, fmt.Sprintf("%s=?", field))
		params = append(params, value)
	}

	//check for date filter
	if from != "" && to != "" {
		fromArr := strings.Split(from, ":")
		toArr := strings.Split(to, ":")
		where = append(where, fmt.Sprintf(`p.%s BETWEEN ? AND ?`, fromArr[0]))
		params = append(params, fromArr[1], toArr[1])
	}

	//check for domain
	where = append(where, fmt.Sprintf("%s=?", "i.domain_id"))
	params = append(params, pc.GetDomainID())

	highestScope := pc.GetCurrentScope()
	if highestScope == permission.ScopeSelf {
		// find current user childes
		childes := pc.GetChildesPerm(permission.ScopeSelf, "list_inventory", pc.GetDomainID())
		childes = append(childes, pc.GetID())
		where = append(where, fmt.Sprintf("i.user_id IN (%s)",
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
		query = fmt.Sprintf("%s %s ", query, "WHERE")
		countQuery = fmt.Sprintf("%s %s ", countQuery, "WHERE")
	}
	query += strings.Join(where, " AND ")
	countQuery += strings.Join(where, " AND ")
	if len(where) > 0 && len(whereLike) > 0 {
		query = fmt.Sprintf("%s %s ", query, "AND")
		countQuery = fmt.Sprintf("%s %s ", countQuery, "AND")
	}
	query += strings.Join(whereLike, " OR ")
	countQuery += strings.Join(whereLike, " OR ")
	query += fmt.Sprintf(" GROUP BY p.id ")
	countQuery += fmt.Sprintf(" GROUP BY p.id ")
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
