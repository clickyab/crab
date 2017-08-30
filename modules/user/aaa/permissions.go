package aaa

import (
	"fmt"

	"strings"

	"strconv"

	"github.com/Sirupsen/logrus"
	"github.com/clickyab/services/assert"
	"github.com/clickyab/services/permission"
)

func (u *User) Has(scope permission.UserScope, p permission.Token) (permission.UserScope, bool) {
	perm := string(p)
	if !scope.IsValid() {
		return permission.ScopeSelf, false
	}

	u.setUserRoles()
	u.setUserPermissions()

	var (
		rScope      permission.UserScope
		permGranted bool
	)

	switch scope {
	case permission.ScopeSelf:
		if u.resource[permission.ScopeSelf][perm] {
			rScope = scope
			permGranted = true
		}
		fallthrough
	case permission.ScopeParent:
		if u.resource[permission.ScopeParent][perm] {
			rScope = scope
			permGranted = true
		}
		fallthrough
	case permission.ScopeGlobal:
		if u.resource[permission.ScopeGlobal][perm] || u.resource[scope][string(permission.God)] {
			rScope = scope
			permGranted = true
		}
	}

	return rScope, permGranted
}

func (u *User) HasOn(perm permission.Token, ownerID, parentID int64, scopes ...permission.UserScope) (permission.UserScope, bool) {
	if len(scopes) == 0 {
		return permission.ScopeSelf, false
	}

	u.setUserRoles()
	u.setUserPermissions()

	if !NewAaaManager().ConsularCustomerExists(parentID, ownerID) {
		return permission.ScopeSelf, false
	}

	var self, parent, global bool

	if len(scopes) == 0 {
		self = true
		parent = true
		global = true
	} else {
		for i := range scopes {
			if scopes[i] == permission.ScopeSelf {
				self = true
			} else if scopes[i] == permission.ScopeParent {
				parent = true
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
		}
	}
	if parent {
		if parentID == u.ID {
			if u.resource[permission.ScopeParent][string(perm)] {
				return permission.ScopeParent, true
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

func (u *User) getUserRoles() []Role {
	var roles []Role
	query := fmt.Sprintf("SELECT roles.* FROM %[1]s INNER JOIN %[2]s ON %[2]s.role_id=roles.id WHERE %[2]s.user_id=?", RoleTableFull, RoleUserTableFull)

	_, err := NewAaaManager().GetRDbMap().Select(&roles, query, u.ID)
	assert.Nil(err)
	return roles
}

func (u *User) setUserRoles() {
	if len(u.roles) == 0 {
		u.roles = u.getUserRoles()
	}
}

func (u *User) getUserPermissions() map[permission.UserScope]map[string]bool {
	var roleIDs []string
	var rolePerm []RolePermission
	var resp = make(map[permission.UserScope]map[string]bool, 0)

	if len(u.roles) == 0 {
		u.setUserRoles()
	}
	roles := u.roles

	for i := range roles {
		roleIDs = append(roleIDs, string(roles[i].ID))
	}

	var ids string
	if len(roleIDs) > 1 {
		ids = strings.Join(roleIDs, ",")
	} else {
		ids = strconv.FormatInt(roles[0].ID, 10)
	}

	query := fmt.Sprintf(`SELECT * from %s WHERE role_id IN (%s)`, RolePermissionTableFull, ids)
	_, err := NewAaaManager().GetRDbMap().Select(&rolePerm, query)
	assert.Nil(err)

	for i := range rolePerm {
		g := make(map[string]bool, 0)
		scope := rolePerm[i].Scope
		perm := rolePerm[i].Perm
		g[perm] = true
		logrus.Warn(scope, perm)
		resp[scope] = g
	}

	return resp
}

func (u *User) setUserPermissions() {
	if u.resource == nil {
		u.resource = u.getUserPermissions()
	}
}
