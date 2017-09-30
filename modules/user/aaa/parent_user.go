package aaa

import (
	"fmt"
	"time"

	"github.com/clickyab/services/assert"
)

// ParentUser ParentUser model in database
// @Model {
//		table = parent_user
//		primary = false, user_id,parent_id,domain_id
// }
type ParentUser struct {
	UserID    int64     `json:"user_id" db:"user_id"`
	ParentID  int64     `json:"parent_id" db:"parent_id"`
	DomainID  int64     `json:"domain_id" db:"domain_id"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
}

// GetUserParentsIDDomain get user parent by id and domain
func (m *Manager) GetUserParentsIDDomain(id, d int64) []ParentUser {
	var res []ParentUser
	q := fmt.Sprintf("SELECT * FROM %s AS pu WHERE user_id=? AND domain_id=?", ParentUserTableFull)
	_, err := m.GetRDbMap().Select(&res, q, id, d)
	assert.Nil(err)
	return res
}
