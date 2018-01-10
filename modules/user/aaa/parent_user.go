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

// GetUserChildesIDDomain get user child ids
func (m *Manager) GetUserChildesIDDomain(id, d int64) []int64 {
	var res []ParentUser
	var final []int64
	q := fmt.Sprintf(`SELECT user_id FROM %s AS pu WHERE pu.parent_id=? AND pu.domain_id=?`, ParentUserTableFull)
	_, err := m.GetRDbMap().Select(&res, q, id, d)
	assert.Nil(err)
	for i := range res {
		final = append(final, res[i].UserID)
	}
	return final
}
