package aaa

import (
	"fmt"

	"clickyab.com/crab/modules/domain/orm"
	"github.com/clickyab/services/assert"
	"github.com/clickyab/services/permission"
	"github.com/sirupsen/logrus"
)

// Has check for pern
func (u *User) Has(scope permission.UserScope, p permission.Token, d int64) (permission.UserScope, bool) {
	permission.Registered(p)
	perm := string(p)
	if !scope.IsValid() {
		return permission.ScopeSelf, false
	}
	u.setUserPermissions(d)
	var (
		rScope      = permission.ScopeSelf
		permGranted bool
	)
	switch scope {
	case permission.ScopeSelf:
		if u.resource[permission.ScopeSelf][perm] {
			rScope = permission.ScopeSelf
			permGranted = true
		}
		fallthrough
	case permission.ScopeGlobal:
		if u.resource[permission.ScopeGlobal][perm] {
			rScope = permission.ScopeGlobal
			permGranted = true
		}
		fallthrough
	case permission.ScopeSuperGlobal:
		if u.resource[permission.ScopeSuperGlobal][perm] {
			rScope = permission.ScopeSuperGlobal
			permGranted = true
		}

	}

	return rScope, permGranted
}

// HasOn check for entity perm
func (u *User) HasOn(
	perm permission.Token,
	ownerID int64,
	DomainID int64,
	checkLevel bool,
	preventSelf bool,
	scopes ...permission.UserScope) (permission.UserScope, bool) {
	if preventSelf && ownerID == u.ID {
		return permission.ScopeSelf, false
	}
	permission.Registered(perm)
	u.setUserPermissions(DomainID)
	var (
		self, global, superGlobal bool
	)
	if len(scopes) == 0 {
		self = true
		global = true
		superGlobal = true
	} else {
		for i := range scopes {
			if scopes[i] == permission.ScopeSelf {
				self = true
			} else if scopes[i] == permission.ScopeGlobal {
				global = true
			} else if scopes[i] == permission.ScopeSuperGlobal {
				superGlobal = true
			}
		}
	}
	owner, err := NewAaaManager().FindUserByID(ownerID)
	assert.Nil(err)
	owner.setUserParents(DomainID)
	owner.setUserPermissions(DomainID)
	if self {
		if ownerID == u.ID {
			if u.resource[permission.ScopeSelf][string(perm)] {
				return permission.ScopeSelf, true
			}
		} else { //check Parents
			for i := range owner.Parents {
				if owner.Parents[i] == u.ID {
					var l = true
					if checkLevel {
						l = u.Role.Level > owner.Role.Level
					}
					if owner.resource[permission.ScopeSelf][string(perm)] && l {
						return permission.ScopeSelf, true
					}

				}
			}
		}
	}

	if global {
		var l = true
		if checkLevel {
			l = u.Role.Level > owner.Role.Level
		}
		if u.resource[permission.ScopeGlobal][string(perm)] && NewAaaManager().CheckUserDomain(ownerID, DomainID) && l {

			return permission.ScopeGlobal, true
		}

	}
	if superGlobal {
		if u.resource[permission.ScopeSuperGlobal][string(perm)] {
			return permission.ScopeSuperGlobal, true
		}
	}
	return permission.ScopeSelf, false
}

func (u *User) getUserRole(DomainID int64) *Role {
	logrus.Warn(DomainID)
	var role *Role
	var where = "WHERE domain_id IS NULL"
	var params []interface{}
	params = append(params, u.ID)
	if !u.DomainLess {
		where = "WHERE domain_id=?"
		params = append(params, DomainID)
	}
	logrus.Warn(params)
	query := fmt.Sprintf("SELECT %s FROM %s AS r INNER JOIN %s AS du ON (du.role_id=r.id AND du.user_id=?) %s",
		GetSelectFields(RoleTableFull, "r"),
		RoleTableFull,
		orm.DomainUserTableFull,
		where,
	)

	err := NewAaaManager().GetRDbMap().SelectOne(&role, query, params...)
	assert.Nil(err)
	return role
}

// SetUserRole set user roles
func (u *User) SetUserRole(DomainID int64) {
	if u.Role == nil {
		u.Role = u.getUserRole(DomainID)
	}
}

func (u *User) getUserPermissions(DomainID int64) map[permission.UserScope]map[string]bool {
	var rolePerm []RolePermission
	var resp = make(map[permission.UserScope]map[string]bool)
	resp[permission.ScopeGlobal] = make(map[string]bool)
	resp[permission.ScopeSelf] = make(map[string]bool)
	resp[permission.ScopeSuperGlobal] = make(map[string]bool)
	u.SetUserRole(DomainID)
	var where = "AND rp.scope=?"
	if !u.DomainLess {
		where = "AND rp.scope!=?"
	}
	query := fmt.Sprintf(
		`SELECT %s from %s AS rp WHERE rp.role_id=? %s`,
		GetSelectFields(RolePermissionTableFull, "rp"),
		RolePermissionTableFull,
		where,
	)
	_, err := NewAaaManager().GetRDbMap().Select(&rolePerm, query, u.Role.ID, permission.ScopeSuperGlobal)
	assert.Nil(err)
	for i := range rolePerm {
		resp[rolePerm[i].Scope][rolePerm[i].Perm] = true
	}
	return resp
}

func (u *User) setUserPermissions(DomainID int64) {
	if u.resource == nil {
		u.resource = u.getUserPermissions(DomainID)
	}
}

// GetChildesPerm get user childes based on perm
func (u *User) GetChildesPerm(scope permission.UserScope, perm string, DomainID int64) []int64 {
	if u.childes == nil {
		u.childes = NewAaaManager().getUserChildesIDPerAdviser(u.ID, DomainID, scope, perm)
	}
	return u.childes
}

// GetAllUserPerms return a slice of user permissions with scope
func (u *User) GetAllUserPerms(domainID int64) (*[]string, error) {
	userPerms := u.getUserPermissions(domainID)
	var res []string
	for scope, perm := range userPerms {
		for permName := range perm {
			if perm[permName] {
				res = append(res, fmt.Sprintf("%s:%s", permName, scope))
			}
		}
	}
	return &res, nil
}
