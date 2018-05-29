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

// CreateCorporation try to save a new Corporation in database
func (m *Manager) CreateCorporation(c *Corporation) error {

	func(in interface{}) {
		if ii, ok := in.(initializer.Simple); ok {
			ii.Initialize()
		}
	}(c)

	return m.GetWDbMap().Insert(c)
}

// UpdateCorporation try to update Corporation in database
func (m *Manager) UpdateCorporation(c *Corporation) error {

	func(in interface{}) {
		if ii, ok := in.(initializer.Simple); ok {
			ii.Initialize()
		}
	}(c)

	_, err := m.GetWDbMap().Update(c)
	return err
}

// FindCorporationByUserID return the Corporation base on its user_id
func (m *Manager) FindCorporationByUserID(ui int64) (*Corporation, error) {
	var res Corporation
	err := m.GetRDbMap().SelectOne(
		&res,
		fmt.Sprintf("SELECT %s FROM %s WHERE user_id=?", GetSelectFields(CorporationTableFull, ""), CorporationTableFull),
		ui,
	)

	if err != nil {
		return nil, err
	}

	return &res, nil
}

// CreateUser try to save a new User in database
func (m *Manager) CreateUser(u *User) error {
	now := time.Now()
	u.CreatedAt = now
	u.UpdatedAt = now
	func(in interface{}) {
		if ii, ok := in.(initializer.Simple); ok {
			ii.Initialize()
		}
	}(u)

	return m.GetWDbMap().Insert(u)
}

// UpdateUser try to update User in database
func (m *Manager) UpdateUser(u *User) error {

	u.UpdatedAt = time.Now()

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
		fmt.Sprintf("SELECT %s FROM %s %s", GetSelectFields(UserTableFull, ""), UserTableFull, filter),
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
		fmt.Sprintf("SELECT %s FROM %s %s", GetSelectFields(UserTableFull, ""), UserTableFull, filter),
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
		fmt.Sprintf("SELECT %s FROM %s WHERE id=?", GetSelectFields(UserTableFull, ""), UserTableFull),
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
		fmt.Sprintf("SELECT %s FROM %s WHERE email=?", GetSelectFields(UserTableFull, ""), UserTableFull),
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
		fmt.Sprintf("SELECT %s FROM %s WHERE access_token=?", GetSelectFields(UserTableFull, ""), UserTableFull),
		at,
	)

	if err != nil {
		return nil, err
	}

	return &res, nil
}
