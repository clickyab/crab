package aaa

import (
	"fmt"
	"time"

	"clickyab.com/crab/modules/domain/dmn"
	"clickyab.com/crab/modules/user/ucfg"
	"github.com/clickyab/services/assert"
	"github.com/clickyab/services/kv"
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
	// ActiveUserStatus user active
	ActiveUserStatus UserValidStatus = "active"
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

// GenderType is the user gender
type (
	// GenderType is the user gender
	// @Enum{
	// }
	GenderType string
)

const (
	// MaleGender male
	MaleGender GenderType = "male"
	// FemaleGender female
	FemaleGender GenderType = "female"
	// NotSpecifiedGender not specified
	NotSpecifiedGender GenderType = "not_specified"
)

// User user model in database
// @Model {
//		table = users
//		primary = true, id
//		find_by = id,email,access_token
//		list = yes
// }
type User struct {
	ID          int64                 `json:"id" db:"id"`
	Email       string                `json:"email" db:"email"`
	Password    string                `json:"password" db:"password"`
	AccessToken string                `json:"-" db:"access_token"`
	Avatar      mysql.NullString      `json:"avatar" db:"avatar"`
	UserType    UserTyp               `json:"user_type" db:"user_type"`
	Status      UserValidStatus       `json:"status" db:"status"`
	CreatedAt   time.Time             `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time             `json:"updated_at" db:"updated_at"`
	OldPassword mysql.StringJSONArray `json:"-"  db:"old_password"`

	profile interface{} `db:"-"`
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

// UserPersonal user personal model
// @Model {
//		table = user_personal
//		primary = false, user_id
//		find_by = user_id
// }
type UserPersonal struct {
	UserID    int64            `json:"user_id" db:"user_id"`
	FirstName string           `json:"first_name" db:"first_name"`
	LastName  string           `json:"last_name" db:"last_name"`
	Gender    GenderType       `json:"gender" db:"gender"`
	Cellphone mysql.NullString `json:"cellphone" db:"cellphone"`
	Phone     mysql.NullString `json:"phone" db:"phone"`
	Address   mysql.NullString `json:"address" db:"address"`
	CityID    mysql.NullInt64  `json:"city_id" db:"city_id"`
	CreatedAt time.Time        `json:"created_at" db:"created_at"`
	UpdatedAt time.Time        `json:"updated_at" db:"updated_at"`
}

// UserCorporation user corporation model
// @Model {
//		table = user_corporation
//		primary = false, user_id
//		find_by = user_id
// }
type UserCorporation struct {
	UserID       int64            `json:"user_id" db:"user_id"`
	FirstName    mysql.NullString `json:"first_name" db:"first_name"`
	LastName     mysql.NullString `json:"last_name" db:"last_name"`
	Name         mysql.NullString `json:"name" db:"name"`
	Cellphone    mysql.NullString `json:"cellphone" db:"cellphone"`
	Phone        mysql.NullString `json:"phone" db:"phone"`
	Address      mysql.NullString `json:"address" db:"address"`
	EconomicCode mysql.NullString `json:"economic_code" db:"economic_code"`
	RegisterCode mysql.NullString `json:"register_code" db:"register_code"`
	CityID       mysql.NullInt64  `json:"city_id" db:"city_id"`
	CreatedAt    time.Time        `json:"created_at" db:"created_at"`
	UpdatedAt    time.Time        `json:"updated_at" db:"updated_at"`
}

// RegisterUserPayload register
type RegisterUserPayload struct {
	Email       string
	Password    string
	FirstName   string
	Mobile      string
	LastName    string
	CompanyName string
	UserType    UserTyp
}

// RegisterUser try to register user
func (m *Manager) RegisterUser(pl RegisterUserPayload, domainID int64) (*User, error) {
	password, err := bcrypt.GenerateFromPassword([]byte(pl.Password), bcrypt.DefaultCost)
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
		Email:       pl.Email,
		Password:    string(password),
		UserType:    pl.UserType,
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
	var up *UserPersonal
	var uc *UserCorporation
	if pl.UserType == PersonalUserTyp {
		up = &UserPersonal{
			UserID:    u.ID,
			FirstName: pl.FirstName,
			LastName:  pl.LastName,
			Cellphone: mysql.NullString{String: pl.Mobile, Valid: pl.Mobile != ""},
			Gender:    NotSpecifiedGender,
		}
		err = m.CreateUserPersonal(up)
		if err != nil {
			return nil, err
		}
		u.profile = up
	} else {
		uc = &UserCorporation{
			UserID:    u.ID,
			FirstName: mysql.NullString{String: pl.FirstName, Valid: true},
			LastName:  mysql.NullString{String: pl.LastName, Valid: true},
			Name:      mysql.NullString{String: pl.CompanyName, Valid: true},
			Cellphone: mysql.NullString{String: pl.Mobile, Valid: pl.Mobile != ""},
		}
		err = m.CreateUserCorporation(uc)
		if err != nil {
			return nil, err
		}
		u.profile = uc
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
	generated := kv.NewEavStore(t).SetSubKey("token", user.AccessToken)
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

// FindUserByAccessTokenDomain return the User base on its access_token and domain
func (m *Manager) FindUserByAccessTokenDomain(at string, domainID int64) (*User, error) {
	var res User
	err := m.GetRDbMap().SelectOne(
		&res,
		fmt.Sprintf("SELECT u.* FROM %s AS u "+
			"INNER JOIN %s AS dm ON dm.user_id=u.id "+
			"WHERE u.access_token=? AND dm.domain_id=?", UserTableFull, dmn.DomainUserTableFull),
		at,
		domainID,
	)

	if err != nil {
		return nil, err
	}

	return &res, nil
}

// GetUserPersonal get personal profile
func (u *User) GetUserPersonal() *UserPersonal {
	if u.profile == nil {
		m := NewAaaManager()
		up, err := m.FindUserPersonalByUserID(u.ID)
		assert.Nil(err)
		u.profile = up
	}
	return u.profile.(*UserPersonal)
}

// GetUserCorporation get corporation profile
func (u *User) GetUserCorporation() *UserCorporation {
	if u.profile == nil {
		m := NewAaaManager()
		uc, err := m.FindUserCorporationByUserID(u.ID)
		assert.Nil(err)
		u.profile = uc
	}
	return u.profile.(*UserCorporation)
}

// UpdateOldPassword get user id and password and will update oldpassword column
func (m *Manager) UpdateOldPassword(d int64, p string) error {
	u, e := m.FindUserByID(d)
	if e != nil {
		return e
	}
	np, e := bcrypt.GenerateFromPassword([]byte(p), bcrypt.DefaultCost)
	if e != nil {
		return e
	}
	if len(u.OldPassword) < 3 {
		u.OldPassword = append(u.OldPassword, string(np))
	} else {
		u.OldPassword = append([]string{string(np)}, u.OldPassword[:2]...)
	}
	return m.UpdateUser(u)
}

// IsOldPassword will return true if user used this password
func (m *Manager) IsOldPassword(d int64, p string) (bool, error) {
	u, e := m.FindUserByID(d)
	if e != nil {
		return false, e
	}
	for _, o := range u.OldPassword {
		if bcrypt.CompareHashAndPassword([]byte(o), []byte(p)) != nil {
			return true, nil
		}
	}
	return false, nil
}

// FindUserByEmailDomain return the User base on its email an domain
func (m *Manager) FindUserByEmailDomain(email string, domain *dmn.Domain) (*User, error) {
	var res User
	err := m.GetRDbMap().SelectOne(
		&res,
		fmt.Sprintf("SELECT u.* FROM %s AS u "+
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
