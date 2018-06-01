package aaa

import (
	"fmt"
	"time"

	"clickyab.com/crab/modules/domain/orm"
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
	q := fmt.Sprintf("SELECT %s FROM %s AS pu WHERE user_id=? AND domain_id=?", GetSelectFields(AdvisorTableFull, ""), AdvisorTableFull)
	_, err := m.GetRDbMap().Select(&res, q, id, d)
	assert.Nil(err)
	return res
}

// getUserChildesIDPerAdviser get user child ids perm considered
func (m *Manager) getUserChildesIDPerAdviser(id, d int64, scope permission.UserScope, perm string) []int64 {
	var res []Advisor
	var final []int64
	q := fmt.Sprintf(`SELECT %s FROM %s AS a 
		INNER JOIN %s AS du ON (du.user_id=a.user_id AND du.domain_id=?)
		INNER JOIN %s AS rp ON (rp.role_id=du.role_id) WHERE a.advisor_id=? AND rp.perm=? AND rp.scope=?`,
		GetSelectFields(AdvisorTableFull, "a"),
		AdvisorTableFull,
		orm.DomainUserTableFull,
		RolePermissionTableFull,
	)
	_, err := m.GetRDbMap().Select(&res, q, d, id, perm, scope)
	assert.Nil(err)
	for i := range res {
		final = append(final, res[i].UserID)
	}
	return final
}
