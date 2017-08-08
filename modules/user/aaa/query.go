package aaa

import (
	"fmt"

	"clickyab.com/crab/modules/domain/dmn"
)

// FindUserByEmailDomian return the User base on its email an domain
func (m *Manager) FindUserByEmailDomian(email string, domain *dmn.Domain) (*User, error) {
	var res User
	err := m.GetRDbMap().SelectOne(
		&res,
		fmt.Sprintf("SELECT * FROM %s AS u "+
			"INNER JOIN domain_user AS dm ON dm.user_id=u.id"+
			" WHERE u.email=? AND dm.domain_id=?", UserTableFull),
		email,
		domain.ID,
	)

	if err != nil {
		return nil, err
	}

	return &res, nil
}
