package orm

import (
	"strings"
	"time"

	"fmt"

	"github.com/clickyab/services/assert"
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

// InventoryStatus is the inventory status
type (
	// InventoryStatus is the inventory status
	// @Enum{
	// }
	InventoryStatus string
)

const (
	// EnableInventoryStatus enable status
	EnableInventoryStatus InventoryStatus = "enable"
	// DisableInventoryStatus disable status
	DisableInventoryStatus InventoryStatus = "disable"
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

// Inventory is model for inventories table in database
// @Model {
//		table = inventories
//		primary =  true, id
//		find_by = id
//		list = yes
// }
type Inventory struct {
	ID        int64           `json:"id" db:"id" type:"number"`
	CreatedAt time.Time       `json:"created_at" db:"created_at" type:"date" sort:"true"`
	UpdatedAt time.Time       `json:"updated_at" db:"updated_at" type:"date"`
	UserID    int64           `json:"user_id" db:"user_id" type:"number"`
	DomainID  int64           `json:"domain_id" db:"domain_id" type:"number" visible:"false"`
	Label     string          `json:"label" db:"label" type:"string" search:"true"`
	Status    InventoryStatus `json:"status" db:"status" type:"enum"`
}

// InventoryPublisher is model for inventories_publishers table in database
// @Model {
//		table = inventories_publishers
//		primary =  false, publisher_id,inventory_id
//		list = yes
// }
type InventoryPublisher struct {
	PublisherID int64 `json:"publisher_id" db:"publisher_id"`
	InventoryID int64 `json:"inventory_id" db:"inventory_id"`
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

// CreateInventoryComplete create inventory with its pivot table
func (m *Manager) CreateInventoryComplete(label string, pubIDs []int64, domainID, userID int64) (*Inventory, error) {
	err := m.Begin()
	assert.Nil(err)
	defer func() {
		if err != nil {
			assert.Nil(m.Rollback())
		} else {
			assert.Nil(m.Commit())
		}
	}()
	newInventory := &Inventory{
		Label:    label,
		Status:   EnableInventoryStatus,
		DomainID: domainID,
		UserID:   userID,
	}
	err = m.CreateInventory(newInventory)
	if err != nil {
		return nil, err
	}
	// create inventory publishers stuff
	for i := range pubIDs {
		inventoryPublisher := &InventoryPublisher{
			InventoryID: newInventory.ID,
			PublisherID: pubIDs[i],
		}
		err = m.CreateInventoryPublisher(inventoryPublisher)
		if err != nil {
			return nil, err
		}
	}

	return newInventory, nil
}

// AddInventoryComplete add inventory with its pivot table
func (m *Manager) AddInventoryComplete(currentInventory *Inventory, invPublishers []InventoryPublisher, pubIDs []int64) (*Inventory, error) {
	err := m.Begin()
	assert.Nil(err)
	defer func() {
		if err != nil {
			assert.Nil(m.Rollback())
		} else {
			assert.Nil(m.Commit())
		}
	}()

	//Remove inventory publisher
	_, err = m.GetWDbMap().Exec(fmt.Sprintf("DELETE FROM %s WHERE inventory_id=?", InventoryPublisherTableFull), currentInventory.ID)
	if err != nil {
		return nil, err
	}

	var invPubIDs = make([]int64, 0)
	for j := range invPublishers {
		invPubIDs = append(invPubIDs, invPublishers[j].PublisherID)
	}
	for i := range pubIDs {
		invPubIDs = append(invPubIDs, pubIDs[i])
	}
	invPubIDsRes := removeDuplicate(invPubIDs)

	//create inventory publishers
	for k := range invPubIDsRes {
		inventoryPublisher := &InventoryPublisher{
			InventoryID: currentInventory.ID,
			PublisherID: invPubIDsRes[k],
		}
		err = m.CreateInventoryPublisher(inventoryPublisher)
		if err != nil {
			return nil, err
		}
	}

	return currentInventory, nil
}

// RemoveInventoryPub remove inventory with its pivot table
func (m *Manager) RemoveInventoryPub(currentInventory *Inventory, invPublishers []InventoryPublisher, pubIDs []int64) (*Inventory, error) {
	err := m.Begin()
	assert.Nil(err)
	defer func() {
		if err != nil {
			assert.Nil(m.Rollback())
		} else {
			assert.Nil(m.Commit())
		}
	}()

	//Remove inventory publisher
	_, err = m.GetWDbMap().Exec(fmt.Sprintf("DELETE FROM %s WHERE inventory_id=? AND publisher_id IN (%s)",
		InventoryPublisherTableFull,
		strings.TrimRight(strings.Repeat("?,", len(pubIDs)), ","),
	),
		func(pubIDs []int64, ID int64) []interface{} {
			var pubStringIDs = make([]interface{}, len(pubIDs)+1)
			pubStringIDs[0] = ID
			for i := range pubIDs {
				pubStringIDs[i+1] = pubIDs[i]
			}
			return pubStringIDs
		}(pubIDs, currentInventory.ID)...,
	)
	if err != nil {
		return nil, err
	}

	return currentInventory, nil
}

// ChangeLabel update inventory label
func (m *Manager) ChangeLabel(currentInventory *Inventory, label string) (*Inventory, error) {
	currentInventory.Label = label
	err := m.UpdateInventory(currentInventory)
	return currentInventory, err
}

func removeDuplicate(elements []int64) []int64 {
	encountered := map[int64]bool{}
	var result []int64

	for v := range elements {
		if encountered[elements[v]] {
		} else {
			encountered[elements[v]] = true
			result = append(result, elements[v])
		}
	}
	return result
}

// InventoryDataTable is the inventory full data in data table
// @DataTable {
//		url = /inventory/list
//		entity = inventory
//		view = list_inventory:self
//		searchkey = q
//		datefilter = created_at
//		checkable = false
//		multiselect = false
//		map_prefix = inventories
//		controller = clickyab.com/crab/modules/inventory/controllers
//		fill = FillInventoryDataTableArray
// }
type InventoryDataTable struct {
	Inventory
	OwnerID   int64   `json:"owner_id" db:"owner_id" visible:"false"`
	DomainID  int64   `json:"domain_id" db:"domain_id" visible:"false"`
	ParentIDs []int64 `json:"parent_ids" db:"parent_ids" visible:"false"`
	Actions   string  `db:"-" json:"_actions" visible:"false"`
}

// FillInventoryDataTableArray is the function to handle
func (m *Manager) FillInventoryDataTableArray(
	pc permission.InterfaceComplete,
	filters map[string]string,
	dateRange map[string]string,
	search map[string]string,
	contextparams map[string]string,
	sort, order string, p, c int) (InventoryDataTableArray, int64) {
	var params []interface{}
	var res InventoryDataTableArray
	var where []string
	var whereLike []string
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
		whereLike = append(whereLike, fmt.Sprintf("%s LIKE ?", column))
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
