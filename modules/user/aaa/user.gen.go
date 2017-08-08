// Code generated build with models DO NOT EDIT.

package aaa

import (
	"fmt"
	"strings"

	"github.com/clickyab/services/assert"
	"github.com/clickyab/services/initializer"
)

// AUTO GENERATED CODE. DO NOT EDIT!

// CreateRolePermission try to save a new RolePermission in database
func (m *Manager) CreateRolePermission(rp *RolePermission) error {

	func(in interface{}) {
		if ii, ok := in.(initializer.Simple); ok {
			ii.Initialize()
		}
	}(rp)

	return m.GetWDbMap().Insert(rp)
}

// UpdateRolePermission try to update RolePermission in database
func (m *Manager) UpdateRolePermission(rp *RolePermission) error {

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
		fmt.Sprintf("SELECT * FROM %s WHERE id=?", RolePermissionTableFull),
		id,
	)

	if err != nil {
		return nil, err
	}

	return &res, nil
}

// CreateRole try to save a new Role in database
func (m *Manager) CreateRole(r *Role) error {

	func(in interface{}) {
		if ii, ok := in.(initializer.Simple); ok {
			ii.Initialize()
		}
	}(r)

	return m.GetWDbMap().Insert(r)
}

// UpdateRole try to update Role in database
func (m *Manager) UpdateRole(r *Role) error {

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
		fmt.Sprintf("SELECT * FROM %s %s", RoleTableFull, filter),
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
		fmt.Sprintf("SELECT * FROM %s %s", RoleTableFull, filter),
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
		fmt.Sprintf("SELECT * FROM %s WHERE id=?", RoleTableFull),
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
		fmt.Sprintf("SELECT * FROM %s WHERE name=?", RoleTableFull),
		n,
	)

	if err != nil {
		return nil, err
	}

	return &res, nil
}

// CreateUser try to save a new User in database
func (m *Manager) CreateUser(u *User) error {

	func(in interface{}) {
		if ii, ok := in.(initializer.Simple); ok {
			ii.Initialize()
		}
	}(u)

	return m.GetWDbMap().Insert(u)
}

// UpdateUser try to update User in database
func (m *Manager) UpdateUser(u *User) error {

	func(in interface{}) {
		if ii, ok := in.(initializer.Simple); ok {
			ii.Initialize()
		}
	}(u)

	_, err := m.GetWDbMap().Update(u)
	return err
}

// ListUsersWithFilter try to list all Users without pagination
func (m *Manager) ListUsersWithFilter(filter string, params ...interface{}) []User {
	filter = strings.Trim(filter, "\n\t ")
	if filter != "" {
		filter = "WHERE " + filter
	}
	var res []User
	_, err := m.GetRDbMap().Select(
		&res,
		fmt.Sprintf("SELECT * FROM %s %s", UserTableFull, filter),
		params...,
	)
	assert.Nil(err)

	return res
}

// ListUsers try to list all Users without pagination
func (m *Manager) ListUsers() []User {
	return m.ListUsersWithFilter("")
}

// CountUsersWithFilter count entity in Users table with valid where filter
func (m *Manager) CountUsersWithFilter(filter string, params ...interface{}) int64 {
	filter = strings.Trim(filter, "\n\t ")
	if filter != "" {
		filter = "WHERE " + filter
	}
	cnt, err := m.GetRDbMap().SelectInt(
		fmt.Sprintf("SELECT COUNT(*) FROM %s %s", UserTableFull, filter),
		params...,
	)
	assert.Nil(err)

	return cnt
}

// CountUsers count entity in Users table
func (m *Manager) CountUsers() int64 {
	return m.CountUsersWithFilter("")
}

// ListUsersWithPaginationFilter try to list all Users with pagination and filter
func (m *Manager) ListUsersWithPaginationFilter(
	offset, perPage int, filter string, params ...interface{}) []User {
	var res []User
	filter = strings.Trim(filter, "\n\t ")
	if filter != "" {
		filter = "WHERE " + filter
	}

	filter += " LIMIT ?, ? "
	params = append(params, offset, perPage)

	// TODO : better pagination without offset and limit
	_, err := m.GetRDbMap().Select(
		&res,
		fmt.Sprintf("SELECT * FROM %s %s", UserTableFull, filter),
		params...,
	)
	assert.Nil(err)

	return res
}

// ListUsersWithPagination try to list all Users with pagination
func (m *Manager) ListUsersWithPagination(offset, perPage int) []User {
	return m.ListUsersWithPaginationFilter(offset, perPage, "")
}

// FindUserByID return the User base on its id
func (m *Manager) FindUserByID(id int64) (*User, error) {
	var res User
	err := m.GetRDbMap().SelectOne(
		&res,
		fmt.Sprintf("SELECT * FROM %s WHERE id=?", UserTableFull),
		id,
	)

	if err != nil {
		return nil, err
	}

	return &res, nil
}

// FindUserByEmail return the User base on its email
func (m *Manager) FindUserByEmail(e string) (*User, error) {
	var res User
	err := m.GetRDbMap().SelectOne(
		&res,
		fmt.Sprintf("SELECT * FROM %s WHERE email=?", UserTableFull),
		e,
	)

	if err != nil {
		return nil, err
	}

	return &res, nil
}

// FindUserByAccessToken return the User base on its access_token
func (m *Manager) FindUserByAccessToken(at string) (*User, error) {
	var res User
	err := m.GetRDbMap().SelectOne(
		&res,
		fmt.Sprintf("SELECT * FROM %s WHERE access_token=?", UserTableFull),
		at,
	)

	if err != nil {
		return nil, err
	}

	return &res, nil
}
