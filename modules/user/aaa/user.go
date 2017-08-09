package aaa

import (
	"fmt"
	"time"

	"clickyab.com/crab/modules/domain/dmn"
	ucfg "clickyab.com/crab/modules/user/config"
	"github.com/clickyab/services/assert"
	"github.com/clickyab/services/eav"
	"github.com/clickyab/services/mysql"
	"github.com/clickyab/services/random"
	"golang.org/x/crypto/bcrypt"
)

// ActiveStatus is the user active status
type (
	// ActiveStatus is the user active status
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

// UserTyp is the user type status
type (
	// UserTyp is the user type status
	// @Enum{
	// }
	UserTyp string
)

const (
	// PersonalUserTyp user personal
	PersonalUserTyp UserTyp = "personal"
	// CorporationUserTyp user corporation
	CorporationUserTyp UserTyp = "corporation"
)

// UserValidStatus is the user status
type (
	// UserValidStatus is the user status
	// @Enum{
	// }
	UserValidStatus string
)

const (
	// RegisteredUserStatus user registered
	RegisteredUserStatus UserValidStatus = "registered"
	// BlockedUserStatus user blocked
	BlockedUserStatus UserValidStatus = "blocked"
)

// UserScope is the user perm
type (
	// UserScope is the user perm
	// @Enum{
	// }
	UserScope string
)

const (
	// SelfPerm user self
	SelfPerm UserScope = "self"
	// ParentPerm user parent
	ParentPerm UserScope = "parent"
	// GlobalPerm user global
	GlobalPerm UserScope = "global"
)

// User user model in database
// @Model {
//		table = users
//		primary = true, id
//		find_by = id,email,access_token
//		list = yes
// }
type User struct {
	ID          int64            `json:"id" db:"id"`
	Email       string           `json:"email" db:"email"`
	Password    string           `json:"password" db:"password"`
	AccessToken string           `json:"-" db:"access_token"`
	Avatar      mysql.NullString `json:"avatar" db:"avatar"`
	UserType    UserTyp          `json:"user_type" db:"user_type"`
	Status      UserValidStatus  `json:"status" db:"status"`
	CreatedAt   time.Time        `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time        `json:"updated_at" db:"updated_at"`
}

// Role role model in database
// @Model {
//		table = roles
//		primary = true, id
//		find_by = id
//		list = yes
// }
type Role struct {
	ID          int64            `json:"id" db:"id"`
	Name        string           `json:"name" db:"name"`
	Description mysql.NullString `json:"description" db:"description"`
	DomainID    int64            `json:"domain_id" db:"domain_id"`
	CreatedAt   time.Time        `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time        `json:"updated_at" db:"updated_at"`
}

// RoleUser RoleUser model in database
// @Model {
//		table = role_user
//		primary = false, user_id, role_id
// }
type RoleUser struct {
	UserID    int64     `json:"user_id" db:"user_id"`
	RoleID    int64     `json:"role_id" db:"role_id"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
}

// RolePermission RolePermission model in database
// @Model {
//		table = role_permission
//		primary = true, id
//		find_by = id
// }
type RolePermission struct {
	ID        int64     `json:"id" db:"id"`
	RoleID    int64     `json:"role_id" db:"role_id"`
	Scope     UserScope `json:"scope" db:"scope"`
	Perm      string    `json:"perm" db:"perm"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

// RegisterUser try to register user
func (m *Manager) RegisterUser(email, pass string, typ UserTyp, domainID int64) (*User, error) {
	password, err := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)
	assert.Nil(err)
	err = m.Begin()
	assert.Nil(err)
	defer func() {
		if err != nil {
			assert.Nil(m.Rollback())
		} else {
			assert.Nil(m.Commit())
		}
	}()
	u := &User{
		Email:       email,
		Password:    string(password),
		UserType:    typ,
		Status:      RegisteredUserStatus,
		AccessToken: <-random.ID,
	}
	err = m.CreateUser(u)
	if err != nil {
		return nil, err
	}
	role, err := m.FindRoleByNameDomain(ucfg.DefaultRole.String(), domainID)
	if err != nil {
		return nil, err
	}
	ur := &RoleUser{RoleID: role.ID, UserID: u.ID}
	err = m.CreateRoleUser(ur)
	if err != nil {
		return nil, err
	}
	dManager, err := dmn.NewDmnManagerFromTransaction(m.GetRDbMap())
	if err != nil {
		return nil, err
	}
	du := &dmn.DomainUser{UserID: u.ID, DomainID: domainID}
	err = dManager.CreateDomainUser(du)
	if err != nil {
		return nil, err
	}
	return u, nil
}

// FindRoleByNameDomain return the Role base on its name and domain
func (m *Manager) FindRoleByNameDomain(n string, domainID int64) (*Role, error) {
	var res Role
	err := m.GetRDbMap().SelectOne(
		&res,
		fmt.Sprintf("SELECT * FROM %s WHERE name=? AND domain_id=?", RoleTableFull),
		n,
		domainID,
	)

	if err != nil {
		return nil, err
	}

	return &res, nil
}

// GetNewToken save new token
func GetNewToken(user *User) string {
	t := fmt.Sprintf("%d:%s", user.ID, <-random.ID)
	generated := eav.NewEavStore(t).SetSubKey("token", user.AccessToken)
	assert.Nil(generated.Save(ucfg.TokenTimeout.Duration()))
	return t
}

// FindUserDomainsByEmail find active user domain based on its email
func (m *Manager) FindUserDomainsByEmail(e string) []dmn.Domain {
	var res []dmn.Domain
	q := "SELECT d.* FROM domains AS d " +
		"INNER JOIN domain_user AS dm ON dm.domain_id=d.id " +
		"INNER JOIN users AS u ON u.id=dm.user_id " +
		"WHERE u.email=? AND d.active=?"
	_, err := m.GetRDbMap().Select(&res, q, e, dmn.ActiveStatusYes)
	assert.Nil(err)
	return res
}
