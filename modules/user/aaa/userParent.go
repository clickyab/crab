package aaa

import (
	"fmt"
	"time"
)

// ConsularCustomer CC model in database
// @Model {
//		table = consular_customer
//		primary = false, consulary_id,customer_id
//		list = yes
// }
type ConsularCustomer struct {
	ConsularyID int64     `db:"consulary_id"`
	CustomerID  int64     `db:"customer_id"`
	CreatedAt   time.Time `db:"created_at"`
}

func (m *Manager) ConsularCustomerExists(ConsularyID, CostumerID int64) bool {
	query := fmt.Sprintf(`SELECT * FROM %s where consulary_id=? AND costumer_id=?`, ConsularCustomerTableFull)
	err := m.GetRDbMap().SelectOne(nil, query)

	return err != nil
}
