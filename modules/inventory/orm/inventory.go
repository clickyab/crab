package orm

import (
	"errors"
	"strings"
	"time"

	"fmt"

	"database/sql"

	"clickyab.com/crab/modules/campaign/orm"
	"clickyab.com/crab/modules/user/aaa"
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
	ID        int64           `json:"id" db:"id" type:"number" sort:"true"`
	CreatedAt time.Time       `json:"created_at" db:"created_at" type:"date" sort:"true"`
	UpdatedAt time.Time       `json:"updated_at" db:"updated_at" type:"date"`
	UserID    int64           `json:"user_id" db:"user_id" type:"number"`
	DomainID  int64           `json:"domain_id" db:"domain_id" type:"number" visible:"false"`
	Label     string          `json:"label" db:"label" type:"string" search:"true"`
	Status    InventoryStatus `json:"status" db:"status" type:"enum" filter:"true"`
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

// InventoryWithPubCount type
type InventoryWithPubCount struct {
	Inventory
	PublisherCount int64 `json:"publisher_count" db:"publisher_count"`
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

// AddInventoryPub add inventory with its pivot table
func (m *Manager) AddInventoryPub(currentInventory *Inventory, invPublishers []InventoryPublisher, pubIDs []int64) (*Inventory, error) {
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

	var resDomains []string
	if len(invPubIDsRes) != 0 {
		//find all domains to be updated in campaigns table
		resDomains = m.ListPublisherDomainsByIDs(invPubIDsRes)
	}

	// update campaigns domains related to this inventory
	q := fmt.Sprintf("UPDATE %s SET inventory_domains=? WHERE inventory_id=?", orm.CampaignTableFull)
	newManager, err := orm.NewOrmManagerFromTransaction(m.GetWDbMap())
	_, err = newManager.GetWDbMap().Exec(q, mysql.StringJSONArray(resDomains), currentInventory.ID)
	if err != nil {
		return nil, err
	}
	if err != nil {
		return nil, err
	}
	return currentInventory, nil
}

// RemoveInventoryPub remove inventory with its pivot table
func (m *Manager) RemoveInventoryPub(currentInventory *Inventory, pubIDs []int64, oldPubIDs []int64) (*Inventory, error) {
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

	invPubIDsRes := subtractsTwoArray(oldPubIDs, pubIDs)
	var resDomains []string
	if len(invPubIDsRes) != 0 {
		//find all domains to be updated in campaigns table
		resDomains = m.ListPublisherDomainsByIDs(invPubIDsRes)
	}

	// update campaigns domains related to this inventory
	q := fmt.Sprintf("UPDATE %s SET inventory_domains=? WHERE inventory_id=?", orm.CampaignTableFull)
	newManager, err := orm.NewOrmManagerFromTransaction(m.GetWDbMap())
	_, err = newManager.GetWDbMap().Exec(q, mysql.StringJSONArray(resDomains), currentInventory.ID)
	if err != nil {
		return nil, err
	}
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

// Duplicate inventory. the id argument is the target inventory id
func (m *Manager) Duplicate(id int64) (*Inventory, error) {
	o, err := m.FindInventoryByID(id)
	if err != nil {
		if err == sql.ErrNoRows {
			assert.Nil(err)
		}
		return nil, err
	}
	oid := o.ID

	err = m.Begin()
	defer func() {
		if err != nil {
			m.Rollback()
			return
		}
		m.Commit()
	}()

	err = m.CreateInventory(o)

	if err != nil {
		return nil, err
	}

	_, err = m.GetWDbMap().Exec(fmt.Sprintf(`INSERT into %[2]s
	SELECT %[1]d as inventory_id,publisher_id  FROM %[2]s
	WHERE inventory_id=?`, oid, InventoryPublisherTableFull), o.ID)

	return o, err
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
//		map_prefix = i
//		controller = clickyab.com/crab/modules/inventory/controllers
//		fill = FillInventoryDataTableArray
// }
type InventoryDataTable struct {
	Inventory
	OwnerID           int64   `json:"-" db:"owner_id" visible:"false"`
	AttachedCampaigns int64   `json:"attached" db:"attached" visible:"false"`
	DomainID          int64   `json:"-" db:"domain_id" visible:"false"`
	ParentIDs         []int64 `json:"-" db:"-" visible:"false"`
	Actions           string  `db:"-" json:"_actions" visible:"false"`
}

// FillInventoryDataTableArray is the function to handle
func (m *Manager) FillInventoryDataTableArray(
	pc permission.InterfaceComplete,
	filters map[string]string,
	from string,
	to string,
	search map[string]string,
	contextparams map[string]string,
	sort, order string, p, c int) (InventoryDataTableArray, int64, error) {
	var params []interface{}
	var res InventoryDataTableArray
	var where []string
	var whereLike []string
	countQuery := fmt.Sprintf("SELECT COUNT(i.id) FROM %s AS i "+
		"INNER JOIN %s AS owner ON owner.id=i.user_id LEFT JOIN campaigns AS c ON (c.inventory_id=i.id AND c.status=?)",
		InventoryTableFull,
		aaa.UserTableFull,
	)
	query := fmt.Sprintf("SELECT %s,owner.id AS owner_id,COUNT(c.id) AS attached FROM %s AS i "+
		"INNER JOIN %s AS owner ON owner.id=i.user_id LEFT JOIN campaigns AS c ON (c.inventory_id=i.id AND c.status=?)",
		"i.id,i.label,i.status,i.created_at,i.updated_at,i.domain_id AS domain_id",
		InventoryTableFull,
		aaa.UserTableFull,
	)
	params = append(params, orm.StartStatus)
	for field, value := range filters {
		where = append(where, fmt.Sprintf("%s=?", field))
		params = append(params, value)
	}

	//check for date filter
	if from != "" && to != "" {
		fromArr := strings.Split(from, ":")
		toArr := strings.Split(to, ":")
		where = append(where, fmt.Sprintf(`i.%s BETWEEN ? AND ?`, fromArr[0]))
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

	wl, lp := generateSearchQuery(search)
	whereLike = append(whereLike, wl...)
	params = append(params, lp...)

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
	query += fmt.Sprintf(" GROUP BY i.id ")
	countQuery += fmt.Sprintf(" GROUP BY i.id ")
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

// FindInventoryByIDDomain find inventory by id and domain
func (m *Manager) FindInventoryByIDDomain(ID, domainID int64) (*Inventory, error) {
	inventories := m.ListInventoriesWithFilter("id=? AND domain_id=?", ID, domainID)
	if len(inventories) != 1 {
		return nil, errors.New("inventory not found")
	}
	return &inventories[0], nil
}

// FindInventoryDomainsByInvID find publisher
func (m *Manager) FindInventoryDomainsByInvID(ID int64) []string {
	var res []string
	q := fmt.Sprintf("SELECT p.domain FROM %s AS i "+
		"INNER JOIN %s AS ip ON ip.inventory_id=i.id "+
		"INNER JOIN %s AS p ON p.id=ip.publisher_id WHERE i.id=?",
		InventoryTableFull,
		InventoryPublisherTableFull,
		PublisherTableFull,
	)
	_, err := m.GetRDbMap().Select(&res, q, ID)
	assert.Nil(err)
	return res
}

// subtractsTwoArray similar to (a-b)
func subtractsTwoArray(a, b []int64) []int64 {
	x := make(map[int64]bool)
	for _, i := range a {
		x[i] = true
	}
	for _, i := range b {
		if _, ok := x[i]; ok {
			delete(x, i)
		}
	}
	var res []int64
	for i := range x {
		res = append(res, i)
	}
	return res
}

// FindInventoryAndPubCount find inventory info + publishers count
func (m *Manager) FindInventoryAndPubCount(id int64) (InventoryWithPubCount, error) {
	var res InventoryWithPubCount

	err := m.GetRDbMap().SelectOne(
		&res,
		fmt.Sprintf("SELECT COUNT(ip.publisher_id) AS publisher_count,i.* FROM %s AS i "+
			"INNER JOIN inventories_publishers AS ip ON ip.inventory_id=i.id "+
			"WHERE id=?",
			InventoryTableFull,
		),
		id,
	)

	return res, err
}
