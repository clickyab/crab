package models

import (
	"fmt"

	"strings"

	"strconv"

	"github.com/clickyab/services/assert"
	"github.com/clickyab/services/permission"
	"github.com/sirupsen/logrus"
)

// Has check for pern
func (u *User) Has(scope permission.UserScope, p permission.Token, d int64) (permission.UserScope, bool) {
	perm := string(p)
	if !scope.IsValid() {
		return permission.ScopeSelf, false
	}

	u.setUserRoles(d)
	u.setUserPermissions(d)

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
	case permission.ScopeGlobal:
		if u.resource[permission.ScopeGlobal][perm] || u.resource[scope][string(permission.God)] {
			rScope = scope
			permGranted = true
		}
	}

	return rScope, permGranted
}

// HasOn check for entity perm
func (u *User) HasOn(perm permission.Token, ownerID, parentID int64, DomainID int64, scopes ...permission.UserScope) (permission.UserScope, bool) {
	if len(scopes) == 0 {
		return permission.ScopeSelf, false
	}

	u.setUserRoles(DomainID)
	u.setUserPermissions(DomainID)

	var self, global bool

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

	_, err := NewModelsManager().GetRDbMap().Select(&roles, query, u.ID, DomainID)
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
	var resp = make(map[permission.UserScope]map[string]bool, 0)

	if len(u.roles) == 0 {
		u.setUserRoles(DomainID)
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
	_, err := NewModelsManager().GetRDbMap().Select(&rolePerm, query)
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

func (u *User) setUserPermissions(DomainID int64) {
	if u.resource == nil {
		u.resource = u.getUserPermissions(DomainID)
	}
}
