package dmn

import (
	"time"

	"fmt"

	"github.com/clickyab/services/mysql"
)

// DomainStatus is the user domain status
type (
	// DomainStatus is the user domain status
	// @Enum{
	// }
	DomainStatus string
)

const (
	// EnableDomainStatus enable domain
	EnableDomainStatus DomainStatus = "enable"
	// DisableDomainStatus disable domain
	DisableDomainStatus DomainStatus = "disable"
)

// Domain domain model in database
// @Model {
//		table = domains
//		primary = true, id
//		find_by = id,name
//		list = yes
// }
type Domain struct {
	ID          int64                  `json:"id" db:"id"`
	Name        string                 `json:"name" db:"name"`
	Description mysql.NullString       `json:"description" db:"description"`
	Attributes  mysql.GenericJSONField `json:"attributes" db:"attributes"`
	Status      DomainStatus           `json:"status" db:"status"`
	CreatedAt   time.Time              `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time              `json:"updated_at" db:"updated_at"`
}

// DomainUser domain_user model in database
// @Model {
//		table = users_domains
//		primary = false, user_id, domain_id
// }
type DomainUser struct {
	DomainID int64        `json:"domain_id" db:"domain_id"`
	Status   DomainStatus `json:"status" db:"status"`
	UserID   int64        `json:"user_id" db:"user_id"`
}

// FindActiveDomainByName find active domain by name
func (m *Manager) FindActiveDomainByName(name string) (*Domain, error) {
	var res Domain
	q := fmt.Sprintf("SELECT * FROM %s WHERE name=? AND status=?", DomainTableFull)
	err := m.GetRDbMap().SelectOne(&res, q, name, EnableDomainStatus)
	if err != nil {
		return nil, err
	}
	return &res, nil

}
