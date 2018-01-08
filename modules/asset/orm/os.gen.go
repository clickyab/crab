// Code generated build with models DO NOT EDIT.

package orm

import (
	"fmt"
	"strings"
	"time"

	"github.com/clickyab/services/assert"
	"github.com/clickyab/services/initializer"
)

// AUTO GENERATED CODE. DO NOT EDIT!

// CreateOS try to save a new OS in database
func (m *Manager) CreateOS(os *OS) error {
	now := time.Now()
	os.CreatedAt = now
	os.UpdatedAt = now
	func(in interface{}) {
		if ii, ok := in.(initializer.Simple); ok {
			ii.Initialize()
		}
	}(os)

	return m.GetWDbMap().Insert(os)
}

// UpdateOS try to update OS in database
func (m *Manager) UpdateOS(os *OS) error {
	now := time.Now()
	os.UpdatedAt = now
	func(in interface{}) {
		if ii, ok := in.(initializer.Simple); ok {
			ii.Initialize()
		}
	}(os)

	_, err := m.GetWDbMap().Update(os)
	return err
}

// ListOSWithFilter try to list all OS without pagination
func (m *Manager) ListOSWithFilter(filter string, params ...interface{}) []OS {
	filter = strings.Trim(filter, "\n\t ")
	if filter != "" {
		filter = "WHERE " + filter
	}
	var res []OS
	_, err := m.GetRDbMap().Select(
		&res,
		fmt.Sprintf("SELECT %s FROM %s %s", getSelectFields(OSTableFull, ""), OSTableFull, filter),
		params...,
	)
	assert.Nil(err)

	return res
}

// ListOS try to list all OS without pagination
func (m *Manager) ListOS() []OS {
	return m.ListOSWithFilter("")
}

// CountOSWithFilter count entity in OS table with valid where filter
func (m *Manager) CountOSWithFilter(filter string, params ...interface{}) int64 {
	filter = strings.Trim(filter, "\n\t ")
	if filter != "" {
		filter = "WHERE " + filter
	}
	cnt, err := m.GetRDbMap().SelectInt(
		fmt.Sprintf("SELECT COUNT(*) FROM %s %s", OSTableFull, filter),
		params...,
	)
	assert.Nil(err)

	return cnt
}

// CountOS count entity in OS table
func (m *Manager) CountOS() int64 {
	return m.CountOSWithFilter("")
}

// ListOSWithPaginationFilter try to list all OS with pagination and filter
func (m *Manager) ListOSWithPaginationFilter(
	offset, perPage int, filter string, params ...interface{}) []OS {
	var res []OS
	filter = strings.Trim(filter, "\n\t ")
	if filter != "" {
		filter = "WHERE " + filter
	}

	filter += " LIMIT ?, ? "
	params = append(params, offset, perPage)

	// TODO : better pagination without offset and limit
	_, err := m.GetRDbMap().Select(
		&res,
		fmt.Sprintf("SELECT %s FROM %s %s", getSelectFields(OSTableFull, ""), OSTableFull, filter),
		params...,
	)
	assert.Nil(err)

	return res
}

// ListOSWithPagination try to list all OS with pagination
func (m *Manager) ListOSWithPagination(offset, perPage int) []OS {
	return m.ListOSWithPaginationFilter(offset, perPage, "")
}

// FindOSByName return the OS base on its name
func (m *Manager) FindOSByName(n string) (*OS, error) {
	var res OS
	err := m.GetRDbMap().SelectOne(
		&res,
		fmt.Sprintf("SELECT %s FROM %s WHERE name=?", getSelectFields(OSTableFull, ""), OSTableFull),
		n,
	)

	if err != nil {
		return nil, err
	}

	return &res, nil
}
