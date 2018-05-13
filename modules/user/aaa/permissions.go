package aaa

import (
	"fmt"

	"strings"

	"strconv"

	"github.com/clickyab/services/assert"
	"github.com/clickyab/services/permission"
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
		if u.resource[permission.ScopeGlobal][perm] || u.resource[scope][string(permission.God)] {
			rScope = permission.ScopeGlobal
			permGranted = true
		}
	}

	return rScope, permGranted
}

// HasOn check for entity perm
func (u *User) HasOn(perm permission.Token, ownerID int64, parentIDs []int64, DomainID int64, scopes ...permission.UserScope) (permission.UserScope, bool) {
	permission.Registered(perm)
	u.setUserPermissions(DomainID)
	var (
		self, global bool
	)
	if len(scopes) == 0 {
		self = true
		global = true
	} else {
		for i := range scopes {
			if scopes[i] == permission.ScopeSelf {
				self = true
			} else if scopes[i] == permission.ScopeGlobal {
				global = true
			}
		}
	}
	if self {
		if ownerID == u.ID {
			if u.resource[permission.ScopeSelf][string(perm)] {
				return permission.ScopeSelf, true
			}
		} else { //check parents
			for i := range parentIDs {
				if parentIDs[i] == u.ID {
					return permission.ScopeSelf, true
				}
			}
		}
	}

	if global {
		if u.resource[permission.ScopeGlobal][string(perm)] || u.resource[permission.ScopeGlobal]["god"] {
			return permission.ScopeGlobal, true
		}
	}
	return permission.ScopeSelf, false
}

func (u *User) getUserRoles(DomainID int64) []Role {
	var roles []Role
	query := fmt.Sprintf("SELECT roles.* FROM %[1]s INNER JOIN %[2]s ON %[2]s.role_id=roles.id WHERE %[2]s.user_id=? AND %[1]s.domain_id=?", RoleTableFull, RoleUserTableFull)

	_, err := NewAaaManager().GetRDbMap().Select(&roles, query, u.ID, DomainID)
	assert.Nil(err)
	return roles
}

func (u *User) setUserRoles(DomainID int64) {
	if len(u.roles) == 0 {
		u.roles = u.getUserRoles(DomainID)
	}
}

func (u *User) getUserPermissions(DomainID int64) map[permission.UserScope]map[string]bool {
	var roleIDs []string
	var rolePerm []RolePermission
	var resp = make(map[permission.UserScope]map[string]bool)
	resp[permission.ScopeGlobal] = make(map[string]bool)
	resp[permission.ScopeSelf] = make(map[string]bool)

	if len(u.roles) == 0 {
		u.setUserRoles(DomainID)
	}
	roles := u.roles

	for i := range roles {
		roleIDs = append(roleIDs, fmt.Sprintf("%d", roles[i].ID))
	}

	var ids string
	if len(roleIDs) > 1 {
		ids = strings.Join(roleIDs, ",")
	} else {
		ids = strconv.FormatInt(roles[0].ID, 10)
	}

	query := fmt.Sprintf(
		`SELECT %s from %s WHERE role_id IN (%s)`,
		getSelectFields(RolePermissionTableFull, ""),
		RolePermissionTableFull,
		ids,
	)
	_, err := NewAaaManager().GetRDbMap().Select(&rolePerm, query)
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
		u.childes = NewAaaManager().getUserChildesIDDomainPerm(u.ID, DomainID, scope, perm)
	}
	return u.childes
}

// GetAllUserPerms return a slice of user permissions with scope
func (u *User) GetAllUserPerms(domainID int64) (*[]string, error) {
	var perms []*RolePermission
	q := fmt.Sprintf(
		"select rp.* from %s rp "+
			"JOIN %s ro on rp.role_id = ro.id "+
			"JOIN %s ru on ro.id = ru.role_id "+
			"WHERE ru.user_id=? and ro.domain_id=?",
		RolePermissionTableFull,
		RoleTableFull,
		RoleUserTableFull,
	)
	_, err := NewAaaManager().GetRDbMap().Select(&perms, q, u.ID, domainID)
	if err != nil {
		return nil, err
	}
	var res []string
	for _, perm := range perms {
		res = append(res, fmt.Sprintf("%s:%s", perm.Perm, perm.Scope))
	}
	return &res, nil
}
