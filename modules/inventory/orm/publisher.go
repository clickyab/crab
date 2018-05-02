package orm

import (
	"time"

	"fmt"
	"strings"

	"github.com/clickyab/services/mysql"
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

// GetPublishersStatistics to get publishers count and statistics
func (m *Manager) GetPublishersStatistics() {

}
