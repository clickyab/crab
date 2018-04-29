package aaa

import (
	"fmt"
	"time"

	"github.com/clickyab/services/assert"
	"github.com/clickyab/services/permission"
)

// Advisor Advisor model in database
// @Model {
//		table = advisors
//		primary = false, user_id,advisor_id,domain_id
// }
type Advisor struct {
	UserID    int64     `json:"user_id" db:"user_id"`
	AdvisorID int64     `json:"advisor_id" db:"advisor_id"`
	DomainID  int64     `json:"domain_id" db:"domain_id"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
}

// GetUserParentsIDDomain get user parent by id and domain
func (m *Manager) GetUserParentsIDDomain(id, d int64) []Advisor {
	var res []Advisor
	q := fmt.Sprintf("SELECT * FROM %s AS pu WHERE user_id=? AND domain_id=?", AdvisorTableFull)
	_, err := m.GetRDbMap().Select(&res, q, id, d)
	assert.Nil(err)
	return res
}

// GetUserChildesIDDomain get user child ids
func (m *Manager) GetUserChildesIDDomain(id, d int64) []int64 {
	var res []Advisor
	var final []int64
	q := fmt.Sprintf(`SELECT user_id FROM %s AS pu WHERE pu.advisor_id=? AND pu.domain_id=?`, AdvisorTableFull)
	_, err := m.GetRDbMap().Select(&res, q, id, d)
	assert.Nil(err)
	for i := range res {
		final = append(final, res[i].UserID)
	}
	return final
}

// getUserChildesIDDomainPerm get user child ids perm considered
func (m *Manager) getUserChildesIDDomainPerm(id, d int64, scope permission.UserScope, perm string) []int64 {
	var res []Advisor
	var final []int64
	q := fmt.Sprintf("SELECT %s FROM %s AS a "+
		"INNER JOIN %s AS u ON u.id=a.user_id "+
		"INNER JOIN %s AS ru ON ru.user_id=u.id "+
		"INNER JOIN %s AS rp ON rp.role_id=ru.role_id "+
		"WHERE a.advisor_id=? AND a.domain_id=? AND rp.perm=? AND rp.scope=? GROUP BY u.id",
		getSelectFields(AdvisorTableFull, "a"),
		AdvisorTableFull,
		UserTableFull,
		RoleUserTableFull,
		RolePermissionTableFull,
	)
	_, err := m.GetRDbMap().Select(&res, q, id, d, perm, scope)
	assert.Nil(err)
	for i := range res {
		final = append(final, res[i].UserID)
	}
	return final
}
