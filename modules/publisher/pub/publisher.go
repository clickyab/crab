package pub

import (
	"time"

	"fmt"
	"strings"

	"github.com/clickyab/services/assert"
	"github.com/clickyab/services/permission"
)

// PType is the publisher type
type (
	// PType is the publisher type
	// @Enum{
	// }
	PType string
)

const (
	// AppPubType app publisher
	AppPubType PType = "app"
	// WebPubType web publisher
	WebPubType PType = "web"
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

// Publisher publisher model in database
// @Model {
//		table = publishers
//		primary = true, id
//		find_by = id
//		list = yes
// }
type Publisher struct {
	ID            int64     `json:"id" db:"id" type:"number"`
	UserID        int64     `json:"user_id" db:"user_id" type:"number"`
	Name          string    `json:"name" db:"name" type:"string"`
	Supplier      string    `json:"supplier" db:"supplier" type:"string"`
	Domain        string    `json:"domain" db:"domain" type:"string" search:"true"`
	PublisherType PType     `json:"pub_type" db:"pub_type" type:"enum" filter:"true"`
	PubStatus     Status    `json:"status" db:"status" type:"enum" filter:"true"`
	CreatedAt     time.Time `json:"created_at" db:"created_at" type:"date"`
	UpdatedAt     time.Time `json:"updated_at" db:"updated_at" type:"date"`
}

//PublisherDataTable is the publisher full data in data table
// @DataTable {
//		url = /list
//		entity = publisher
//		view = pub_list:self
//		checkable = true
//		multiselect = true
//		map_prefix = publishers
//		controller = clickyab.com/crab/modules/publisher/controllers
//		fill = FillPublisherDataTableArray
// }
type PublisherDataTable struct {
	Publisher
	Actions string `db:"-" json:"_actions" visible:"false"`
}

// FillPublisherDataTableArray is the function to handle
func (m *Manager) FillPublisherDataTableArray(
	pc permission.InterfaceComplete,
	filters map[string]string,
	search map[string]string,
	contextparams map[string]string,
	sort, order string, p, c int) (PublisherDataTableArray, int64) {
	var params []interface{}
	var res PublisherDataTableArray
	var where []string
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
