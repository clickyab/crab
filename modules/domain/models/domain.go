package models

import (
	"time"

	"fmt"

	"github.com/clickyab/services/mysql"
)

// Domains is slice of domain
type Domains []Domain

// Domain domain model in database
// @Model {
//		table = domains
//		primary = true, id
//		find_by = id,name
//		list = yes
// }
type Domain struct {
	ID          int64            `json:"id" db:"id"`
	Name        string           `json:"name" db:"name"`
	Description mysql.NullString `json:"description" db:"description"`
	Active      bool             `json:"active" db:"active"`
	CreatedAt   time.Time        `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time        `json:"updated_at" db:"updated_at"`
}

// DomainUser domain_user model in database
// @Model {
//		table = domain_user
//		primary = false, user_id, domain_id
// }
type DomainUser struct {
	DomainID int64 `json:"domain_id" db:"domain_id"`
	UserID   int64 `json:"user_id" db:"user_id"`
}

// FindActiveDomainByName find active domain by name
func (m *Manager) FindActiveDomainByName(name string) (*Domain, error) {
	var res Domain
	q := fmt.Sprintf("SELECT * FROM %s WHERE name=? AND active=?", DomainTableFull)
	err := m.GetRDbMap().SelectOne(&res, q, name, true)
	if err != nil {
		return nil, err
	}
	return &res, nil

}
