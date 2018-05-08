package aaa

import (
	"fmt"
	"time"

	"strings"

	domainOrm "clickyab.com/crab/modules/domain/orm"
	"clickyab.com/crab/modules/user/ucfg"
	"github.com/clickyab/services/assert"
	"github.com/clickyab/services/config"
	"github.com/clickyab/services/gettext/t9e"
	"github.com/clickyab/services/kv"
	"github.com/clickyab/services/mysql"
	"github.com/clickyab/services/permission"
	"github.com/clickyab/services/random"
	"golang.org/x/crypto/bcrypt"
)

const manageablePerm = "can_manage"

// AccountType is the user account type
type (
	// AccountType is the user account type
	// @Enum{
	// }
	AccountType string
)

const (
	// PersonalUser male
	PersonalUser AccountType = "personal"
	// CorporationUser corporation
	CorporationUser AccountType = "corporation"
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
	ID          int64                                    `json:"id" db:"id"`
	Email       string                                   `json:"email" db:"email"`
	Password    string                                   `json:"password" db:"password"`
	AccessToken string                                   `json:"-" db:"access_token"`
	Avatar      mysql.NullString                         `json:"avatar" db:"avatar"`
	Status      UserValidStatus                          `json:"status" db:"status"`
	OldPassword mysql.StringJSONArray                    `json:"-"  db:"old_password"`
	CityID      mysql.NullInt64                          `json:"city_id" db:"city_id"`
	LandLine    mysql.NullString                         `json:"land_line" db:"land_line"`
	Cellphone   mysql.NullString                         `json:"cellphone" db:"cellphone"`
	PostalCode  mysql.NullString                         `json:"postal_code" db:"postal_code"`
	FirstName   string                                   `json:"first_name" db:"first_name"`
	LastName    string                                   `json:"last_name" db:"last_name"`
	Address     mysql.NullString                         `json:"address" db:"address"`
	Gender      GenderType                               `json:"gender" db:"gender"`
	SSN         mysql.NullString                         `json:"ssn" db:"ssn"`
	Balance     int64                                    `json:"balance" db:"balance"`
	Attributes  mysql.GenericJSONField                   `json:"attributes" db:"attributes"`
	Advantage   int                                      `json:"advantage" db:"advantage"`
	Corporation *Corporation                             `json:"corporation, omitempty" db:"-"`
	parents     []int64                                  `json:"-" db:"-"`
	childes     []int64                                  `json:"-" db:"-"`
	roles       []Role                                   `db:"-"`
	resource    map[permission.UserScope]map[string]bool `db:"-"`
	CreatedAt   time.Time                                `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time                                `json:"updated_at" db:"updated_at"`
}

// Corporation user corporation model
// @Model {
//		table = corporations
//		primary = false, user_id
//		find_by = user_id
// }
type Corporation struct {
	UserID        int64            `json:"user_id" db:"user_id"`
	LegalName     string           `json:"legal_name" db:"legal_name"`
	LegalRegister mysql.NullString `json:"legal_register" db:"legal_register"`
	EconomicCode  mysql.NullString `json:"economic_code" db:"economic_code"`
}

// RegisterUser try to register user
func (m *Manager) RegisterUser(user *User, corp *Corporation, domainID, roleID int64) error {
	password, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
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

	user.AccessToken = <-random.ID
	user.Password = string(password)

	if !user.Gender.IsValid() {
		user.Gender = NotSpecifiedGender
	}

	if !user.Status.IsValid() {
		user.Status = RegisteredUserStatus
	}

	err = m.CreateUser(user)
	if err != nil {
		return err
	}

	ur := &RoleUser{RoleID: roleID, UserID: user.ID}
	err = m.CreateRoleUser(ur)
	if err != nil {
		return err
	}

	if corp != nil && corp.LegalName != "" {
		corp.UserID = user.ID
		err = m.CreateCorporation(corp)
		if err != nil {
			return err
		}

		user.Corporation = corp
	}

	dManager, err := domainOrm.NewOrmManagerFromTransaction(m.GetWDbMap())
	if err != nil {
		return err
	}
	du := &domainOrm.DomainUser{
		UserID:   user.ID,
		DomainID: domainID,
		Status:   domainOrm.EnableDomainStatus,
	}

	err = dManager.CreateDomainUser(du)

	return err
}

// FindRoleByNameDomain return the Role base on its name and domain
func (m *Manager) FindRoleByNameDomain(n string, domainID int64) (*Role, error) {
	var res Role
	err := m.GetRDbMap().SelectOne(
		&res,
		fmt.Sprintf(
			"SELECT %s FROM %s WHERE name=? AND domain_id=?",
			getSelectFields(RoleTableFull, ""),
			RoleTableFull,
		),
		n,
		domainID,
	)

	if err != nil {
		return nil, err
	}

	return &res, nil
}

// FindRoleByIDAndDomain return the Role base on its name and domain
func (m *Manager) FindRoleByIDAndDomain(rID, domainID int64) (*Role, error) {
	var res Role
	err := m.GetRDbMap().SelectOne(
		&res,
		fmt.Sprintf(
			"SELECT %s FROM %s WHERE id=? AND domain_id=?",
			getSelectFields(RoleTableFull, ""),
			RoleTableFull,
		),
		rID,
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
func (m *Manager) FindUserDomainsByEmail(e string) []domainOrm.Domain {
	var res []domainOrm.Domain
	q := fmt.Sprintf("SELECT d.* FROM %s AS d "+
		"INNER JOIN %s AS dm ON dm.domain_id=d.id "+
		"INNER JOIN %s AS u ON u.id=dm.user_id "+
		"WHERE u.email=? AND d.status=?", domainOrm.DomainTableFull, domainOrm.DomainUserTableFull, UserTableFull)
	_, err := m.GetRDbMap().Select(&res, q, e, domainOrm.EnableDomainStatus)
	assert.Nil(err)
	return res
}

// FindUserByAccessTokenDomain return the User base on its access_token and domain
func (m *Manager) FindUserByAccessTokenDomain(at string, domainID int64) (*User, error) {
	var res User
	err := m.GetRDbMap().SelectOne(
		&res,
		fmt.Sprintf(
			`SELECT %s FROM %s AS u
			INNER JOIN %s AS dm ON dm.user_id=u.id
			WHERE u.access_token=? AND dm.domain_id=?`,
			getSelectFields(UserTableFull, "u"),
			UserTableFull,
			domainOrm.DomainUserTableFull,
		),
		at,
		domainID,
	)

	if err != nil {
		return nil, err
	}

	return &res, nil
}

// FindUserByEmailDomain return the User base on its email an domain
func (m *Manager) FindUserByEmailDomain(email string, domain *domainOrm.Domain) (*User, error) {
	var res User
	err := m.GetRDbMap().SelectOne(
		&res,
		fmt.Sprintf(
			`SELECT %s FROM %s AS u 
			INNER JOIN %s AS dm ON dm.user_id=u.id
			WHERE u.email=? AND dm.domain_id=?`,
			getSelectFields(UserTableFull, "u"),
			UserTableFull,
			domainOrm.DomainUserTableFull,
		),
		email,
		domain.ID,
	)

	if err != nil {
		return nil, err
	}

	return &res, nil
}

var allowOldPassword = config.RegisterBoolean("crab.user.allow_old_pass", true,
	"determine if the user can change it's password to the one of old ones")

var (
	// ErrorOldPass when The password was used before
	ErrorOldPass error = t9e.G("this password was used before, please try another one")
	// ErrorWrongPassword when The current user password and claimed password doesn't match
	ErrorWrongPassword error = t9e.G("current password is not correct")
)

// UpdatePassword will change password (p param) if the current given password (c param) be correct.
func (u *User) UpdatePassword(c, p string) error {
	if !u.ValidatePassword(c) {
		return ErrorWrongPassword
	}
	return u.ChangePassword(p)

}

// ValidatePassword return true if the given password is the user current password.
func (u *User) ValidatePassword(p string) bool {
	return bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(p)) == nil
}

// ChangePassword get user and password and will update oldpassword column
func (u *User) ChangePassword(p string) error {

	b := u.isOldPassword(p)

	if b && !allowOldPassword.Bool() {
		return ErrorOldPass
	}

	np, e := bcrypt.GenerateFromPassword([]byte(p), bcrypt.DefaultCost)
	u.Password = string(np)
	if e != nil {
		return e
	}
	if len(u.OldPassword) < 3 {
		u.OldPassword = append(u.OldPassword, string(np))
	} else {
		u.OldPassword = append([]string{string(np)}, u.OldPassword[:2]...)
	}
	m := NewAaaManager()

	return m.UpdateUser(u)
}

// isOldPassword will return true if user used this password
func (u *User) isOldPassword(p string) bool {
	for _, o := range u.OldPassword {
		if bcrypt.CompareHashAndPassword([]byte(o), []byte(p)) == nil {
			return true
		}
	}
	return false
}

// setUserParents try to get user parent for the specified domain
func (u *User) setUserParents(d int64) {
	if len(u.parents) == 0 {
		u.parents = u.getUserParents(d)
	}
}

// getUserParents try to get user parent ids
func (u *User) getUserParents(d int64) []int64 {
	var res []int64
	m := NewAaaManager()
	parents := m.GetUserParentsIDDomain(u.ID, d)
	for i := range parents {
		res = append(res, parents[i].AdvisorID)
	}
	return res
}

// FindUserWithParentsByID return the User with parent base on its id
func (m *Manager) FindUserWithParentsByID(id, d int64) (*User, error) {
	var res User
	err := m.GetRDbMap().SelectOne(
		&res,
		fmt.Sprintf("SELECT %s FROM %s AS u WHERE u.id=?", getSelectFields(UserTableFull, "u"), UserTableFull),
		id,
	)

	if err != nil {
		return nil, err
	}
	//fetch parents
	res.setUserParents(d)
	return &res, nil
}

// CheckPermOn has perm on wrapper to use inside controllers
func CheckPermOn(owner *User, currentUser *User, perm permission.Token, domainID int64, scopes ...permission.UserScope) (permission.UserScope, bool) {
	if currentUser.ID == owner.ID { //user is the owner
		return owner.HasOn(perm, owner.ID, owner.parents, domainID, permission.ScopeSelf)
	}
	// user is not owner (check parent or global)
	ownerResources := owner.resource
	s, ok := currentUser.HasOn(perm, owner.ID, owner.parents, domainID)
	if s == permission.ScopeSelf && ok {
		return s, ownerResources[permission.ScopeSelf][string(perm)]
	}
	return s, ok
}

// UserSearchResult search result
type UserSearchResult struct {
	Email string `json:"email"`
	ID    int64  `json:"id"`
}

// ListUserByEmail find user by email and domain
func (m *Manager) ListUserByEmail(email string, domainID int64) []UserSearchResult {
	var res []UserSearchResult
	q := fmt.Sprintf(
		`SELECT u.id,u.email FROM %s AS u
				INNER JOIN %s AS du ON (du.user_id=u.id AND du.domain_id=? AND du.status=?)
				WHERE u.email LIKE ?`,
		UserTableFull,
		domainOrm.DomainUserTableFull,
	)
	_, err := m.GetRDbMap().Select(&res, q, domainID, domainOrm.EnableDomainStatus, "%"+email+"%")
	assert.Nil(err)
	return res
}

// FindUserByIDDomain return the User base on its id and domain
func (m *Manager) FindUserByIDDomain(id, domainID int64) (*User, error) {
	var res User
	err := m.GetRDbMap().SelectOne(
		&res,
		fmt.Sprintf(`SELECT %s FROM %s WHERE id=? 
			INNER JOIN %s AS du ON (du.users_id=users.id AND du.domain_id=? AND du.status=?) 
			LIMIT 1`,
			getSelectFields(UserTableFull, ""),
			UserTableFull,
			domainOrm.DomainUserTableFull,
		),
		id,
		domainID,
		domainOrm.EnableDomainStatus,
	)

	if err != nil {
		return nil, err
	}

	return &res, nil
}

// SearchAdvisorsByMailDomain search advisors by mail and domain
func (m *Manager) SearchAdvisorsByMailDomain(mail string, d int64) []UserSearchResult {
	var res []UserSearchResult
	var params []interface{}
	q := fmt.Sprintf(`SELECT u.id,u.email FROM %s AS u 
	INNER JOIN %s AS du ON (du.user_id=u.id AND du.domain_id=? AND du.status=?)
	INNER JOIN %s AS ru ON ru.user_id=u.id 
	INNER JOIN %s AS rp ON rp.role_id=ru.role_id
	WHERE u.email LIKE ? AND rp.perm=? GROUP BY u.id`,
		UserTableFull,
		domainOrm.DomainUserTableFull,
		RoleUserTableFull,
		RolePermissionTableFull,
	)
	params = append(params, d, domainOrm.EnableDomainStatus, "%"+mail+"%", manageablePerm)

	_, err := m.GetRDbMap().Select(&res, q, params...)
	assert.Nil(err)
	return res
}

// FindManagersByIDsDomain find managers by ids and domains
func (m *Manager) FindManagersByIDsDomain(ids []int64, d int64) []int64 {
	var res []int64
	var params []interface{}
	q := fmt.Sprintf(`SELECT u.id FROM %s AS u 
	INNER JOIN %s AS du ON (du.user_id=u.id AND du.domain_id=? AND du.status=?)
	INNER JOIN %s AS ru ON ru.user_id=u.id 
	INNER JOIN %s AS rp ON rp.role_id=ru.role_id
	WHERE u.id IN (%s) AND rp.perm=? GROUP BY u.id`,
		UserTableFull,
		domainOrm.DomainUserTableFull,
		RoleUserTableFull,
		RolePermissionTableFull,
		func() string {
			return strings.TrimRight(strings.Repeat("?,", len(ids)), ",")
		}(),
	)
	params = append(params, d, domainOrm.EnableDomainStatus)
	for i := range ids {
		params = append(params, ids[i])
	}
	params = append(params, manageablePerm)

	_, err := m.GetRDbMap().Select(&res, q, params...)
	assert.Nil(err)
	return res
}

// deleteAdvisorsByUserID delete advisor by user id
func (m *Manager) deleteAdvisorsByUserID(id, d int64) error {
	// delete old managers
	var err error
	q := fmt.Sprintf(`DELETE FROM %s WHERE user_id=? AND domain_id=?`, AdvisorTableFull)
	_, err = m.GetWDbMap().Exec(q, id, d)
	return err
}

// AssignManagers assign managers to user
func (m *Manager) AssignManagers(userID, d int64, managerIDs []int64) ([]*Advisor, error) {
	err := m.Begin()
	assert.Nil(err)
	defer func() {
		if err != nil {
			assert.Nil(m.Rollback())
		} else {
			assert.Nil(m.Commit())
		}
	}()
	err = m.deleteAdvisorsByUserID(userID, d)
	if err != nil {
		return nil, err
	}

	res, err := m.assignManagers(managerIDs, d, userID)
	if err != nil {
		return nil, err
	}

	return res, nil
}

// assignManagers assign advisor to user
func (m *Manager) assignManagers(managerIDs []int64, d, userID int64) ([]*Advisor, error) {
	var res []*Advisor
	var err error
	// insert new managers
	for i := range managerIDs {
		advisor := &Advisor{
			AdvisorID: managerIDs[i],
			DomainID:  d,
			UserID:    userID,
		}
		err = m.CreateAdvisor(advisor)
		if err != nil {
			return res, err
		}
		res = append(res, advisor)
	}
	return res, nil
}
