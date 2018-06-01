// Code generated build with models DO NOT EDIT.

package aaa

import (
	"fmt"
	"strings"
	"time"

	"github.com/clickyab/services/assert"
	"github.com/clickyab/services/initializer"
)

// AUTO GENERATED CODE. DO NOT EDIT!

// CreateRolePermission try to save a new RolePermission in database
func (m *Manager) CreateRolePermission(rp *RolePermission) error {
	now := time.Now()
	rp.CreatedAt = now
	rp.UpdatedAt = now
	func(in interface{}) {
		if ii, ok := in.(initializer.Simple); ok {
			ii.Initialize()
		}
	}(rp)

	return m.GetWDbMap().Insert(rp)
}

// UpdateRolePermission try to update RolePermission in database
func (m *Manager) UpdateRolePermission(rp *RolePermission) error {

	rp.UpdatedAt = time.Now()

	func(in interface{}) {
		if ii, ok := in.(initializer.Simple); ok {
			ii.Initialize()
		}
	}(rp)

	_, err := m.GetWDbMap().Update(rp)
	return err
}

// FindRolePermissionByID return the RolePermission base on its id
func (m *Manager) FindRolePermissionByID(id int64) (*RolePermission, error) {
	var res RolePermission
	err := m.GetRDbMap().SelectOne(
		&res,
		fmt.Sprintf("SELECT %s FROM %s WHERE id=?", GetSelectFields(RolePermissionTableFull, ""), RolePermissionTableFull),
		id,
	)

	if err != nil {
		return nil, err
	}

	return &res, nil
}

// CreateRole try to save a new Role in database
func (m *Manager) CreateRole(r *Role) error {
	now := time.Now()
	r.CreatedAt = now
	r.UpdatedAt = now
	func(in interface{}) {
		if ii, ok := in.(initializer.Simple); ok {
			ii.Initialize()
		}
	}(r)

	return m.GetWDbMap().Insert(r)
}

// UpdateRole try to update Role in database
func (m *Manager) UpdateRole(r *Role) error {

	r.UpdatedAt = time.Now()

	func(in interface{}) {
		if ii, ok := in.(initializer.Simple); ok {
			ii.Initialize()
		}
	}(r)

	_, err := m.GetWDbMap().Update(r)
	return err
}

// ListRolesWithFilter try to list all Roles without pagination
func (m *Manager) ListRolesWithFilter(filter string, params ...interface{}) []Role {
	filter = strings.Trim(filter, "\n\t ")
	if filter != "" {
		filter = "WHERE " + filter
	}
	var res []Role
	_, err := m.GetRDbMap().Select(
		&res,
		fmt.Sprintf("SELECT %s FROM %s %s", GetSelectFields(RoleTableFull, ""), RoleTableFull, filter),
		params...,
	)
	assert.Nil(err)

	return res
}

// ListRoles try to list all Roles without pagination
func (m *Manager) ListRoles() []Role {
	return m.ListRolesWithFilter("")
}

// CountRolesWithFilter count entity in Roles table with valid where filter
func (m *Manager) CountRolesWithFilter(filter string, params ...interface{}) int64 {
	filter = strings.Trim(filter, "\n\t ")
	if filter != "" {
		filter = "WHERE " + filter
	}
	cnt, err := m.GetRDbMap().SelectInt(
		fmt.Sprintf("SELECT COUNT(*) FROM %s %s", RoleTableFull, filter),
		params...,
	)
	assert.Nil(err)

	return cnt
}

// CountRoles count entity in Roles table
func (m *Manager) CountRoles() int64 {
	return m.CountRolesWithFilter("")
}

// ListRolesWithPaginationFilter try to list all Roles with pagination and filter
func (m *Manager) ListRolesWithPaginationFilter(
	offset, perPage int, filter string, params ...interface{}) []Role {
	var res []Role
	filter = strings.Trim(filter, "\n\t ")
	if filter != "" {
		filter = "WHERE " + filter
	}

	filter += " LIMIT ?, ? "
	params = append(params, offset, perPage)

	// TODO : better pagination without offset and limit
	_, err := m.GetRDbMap().Select(
		&res,
		fmt.Sprintf("SELECT %s FROM %s %s", GetSelectFields(RoleTableFull, ""), RoleTableFull, filter),
		params...,
	)
	assert.Nil(err)

	return res
}

// ListRolesWithPagination try to list all Roles with pagination
func (m *Manager) ListRolesWithPagination(offset, perPage int) []Role {
	return m.ListRolesWithPaginationFilter(offset, perPage, "")
}

// FindRoleByID return the Role base on its id
func (m *Manager) FindRoleByID(id int64) (*Role, error) {
	var res Role
	err := m.GetRDbMap().SelectOne(
		&res,
		fmt.Sprintf("SELECT %s FROM %s WHERE id=?", GetSelectFields(RoleTableFull, ""), RoleTableFull),
		id,
	)

	if err != nil {
		return nil, err
	}

	return &res, nil
}

// FindRoleByName return the Role base on its name
func (m *Manager) FindRoleByName(n string) (*Role, error) {
	var res Role
	err := m.GetRDbMap().SelectOne(
		&res,
		fmt.Sprintf("SELECT %s FROM %s WHERE name=?", GetSelectFields(RoleTableFull, ""), RoleTableFull),
		n,
	)

	if err != nil {
		return nil, err
	}

	return &res, nil
}
