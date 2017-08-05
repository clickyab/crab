package dmn

import (
	"time"

	"github.com/clickyab/services/mysql"
)

// ActiveStatus is the domain active status
type (
	// ActiveStatus is the domain active status
	// @Enum{
	// }
	ActiveStatus string
)

const (
	// ActiveStatusYes domain active
	ActiveStatusYes ActiveStatus = "yes"
	// ActiveStatusNo for inactive domain
	ActiveStatusNo ActiveStatus = "no"
)

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
	Active      int64            `json:"active" db:"active"`
	CreatedAt   *time.Time       `json:"created_at" db:"created_at"`
	UpdatedAt   *time.Time       `json:"updated_at" db:"uop"`
}

// DomainUser domain_user model in database
// @Model {
//		table = domain_user
// }
type DomainUser struct {
	DomainID int64  `json:"domain_id" db:"domain_id"`
	UserID   string `json:"user_id" db:"user_id"`
}

// FindActiveDomainByName find active domain by name
func (m *Manager) FindActiveDomainByName(name string) (*Domain, error) {
	var res *Domain
	q := "SELECT * FROM domains WHERE name=? AND active=?"
	err := m.GetRDbMap().SelectOne(res, q, name, ActiveStatusYes)
	if err != nil {
		return nil, err
	}
	return res, nil

}
