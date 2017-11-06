package orm

import (
	"strings"
	"time"

	"fmt"

	"github.com/clickyab/services/assert"
	"github.com/clickyab/services/mysql"
	"github.com/clickyab/services/permission"
)

// PublisherType is the pub type
type (
	// PublisherType is the pub type
	// @Enum{
	// }
	PublisherType string
)

const (
	// PublisherTypeWeb web pub type
	PublisherTypeWeb PublisherType = "web"
	// PublisherTypeAPP web pub type
	PublisherTypeAPP PublisherType = "app"
)

// Status is the publisher status
type (
	// Status is the publisher status
	// @Enum{
	// }
	Status string
)

const (
	// ActiveStatus active publisher
	ActiveStatus Status = "accepted"
	// PendingStatus pending publisher
	PendingStatus Status = "pending"
	// BlockedStatus blocked publisher
	BlockedStatus Status = "blocked"
)

// WhiteBlackList user_wlbl_presets model in database
// @Model {
//		table = user_wlbl_presets
//		primary = true, id
//		find_by = id
//		list = yes
// }
type WhiteBlackList struct {
	ID        int64                    `json:"id" db:"id"`
	CreatedAt time.Time                `json:"created_at" db:"created_at"`
	UpdatedAt time.Time                `json:"updated_at" db:"updated_at"`
	Active    bool                     `json:"active" db:"active"`
	UserID    int64                    `json:"user_id" db:"user_id"`
	DomainID  int64                    `json:"domain_id" db:"domain_id"`
	Label     string                   `json:"label" db:"label"`
	Domains   mysql.StringMapJSONArray `json:"domains" db:"domains"`
	// Kind shows if it's a white list (true) or blacklist (false)
	Kind          bool          `json:"kind" db:"kind"`
	PublisherType PublisherType `json:"publisher_type" db:"publisher_type"`
}

// Inventory is model for inventories table in database
// @Model {
//		table = inventories
//		primary =  false, id
//		find_by = id, domain
//		list = yes
// }
type Inventory struct {
	ID        int64                 `json:"id" db:"id" type:"number"`
	CreatedAt time.Time             `json:"created_at" db:"created_at" type:"date" sort:"true"`
	UpdatedAt time.Time             `json:"updated_at" db:"updated_at" type:"date"`
	Active    bool                  `json:"active" db:"active" type:"bool"`
	Name      string                `json:"name" db:"name" type:"string" search:"true"`
	Domain    string                `json:"domain" db:"domain" type:"string" search:"true"`
	Cat       mysql.StringJSONArray `json:"cat" db:"cat" type:"array"`
	Publisher string                `json:"publisher" db:"publisher" type:"string" sort:"true"`
	Kind      PublisherType         `json:"kind" db:"kind" type:"enum" filter:"true"`
	Status    Status                `json:"status" db:"status" type:"enum"`
}

// InventoryDataTable is the inventory full data in data table
// @DataTable {
//		url = /list
//		entity = inventory
//		view = inventory_list:self
//		checkable = true
//		multiselect = true
//		map_prefix = inventories
//		controller = clickyab.com/crab/modules/inventory/controllers
//		fill = FillInventoryDataTableArray
// }
type InventoryDataTable struct {
	Inventory
	Actions string `db:"-" json:"_actions" visible:"false"`
}

// FillInventoryDataTableArray is the function to handle
func (m *Manager) FillInventoryDataTableArray(
	pc permission.InterfaceComplete,
	filters map[string]string,
	search map[string]string,
	contextparams map[string]string,
	sort, order string, p, c int) (InventoryDataTableArray, int64) {
	var params []interface{}
	var res InventoryDataTableArray
	var where []string
	countQuery := fmt.Sprintf("SELECT COUNT(id) FROM %s",
		InventoryTableFull,
	)
	query := fmt.Sprintf("SELECT * FROM %s",
		InventoryTableFull,
	)
	for field, value := range filters {
		where = append(where, fmt.Sprintf("%s=?", field))
		params = append(params, value)
	}

	for column, val := range search {
		where = append(where, fmt.Sprintf("%s LIKE ?", column))
		params = append(params, "%"+val+"%")
	}
	//check for perm
	if len(where) > 0 {
		query += " WHERE "
		countQuery += " WHERE "
	}
	query += strings.Join(where, " AND ")
	countQuery += strings.Join(where, " AND ")
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

// GetDomainPublishers try to get all Inventory with ids
func (m *Manager) GetDomainPublishers(ids []int64) []Inventory {
	var res []Inventory
	g := strings.Repeat("?,", len(ids))
	gg := strings.TrimRight(g, ",")
	_, err := m.GetRDbMap().Select(
		&res,
		fmt.Sprintf("SELECT * FROM %s WHERE id IN (%s)",
			InventoryTableFull,
			gg,
		),
		func(i []int64) []interface{} {
			x := []interface{}{}
			for _, v := range ids {

				x = append(x, v)
			}
			return x
		}(ids)...,
	)
	assert.Nil(err)

	return res
}
