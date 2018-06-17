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

// CreateOSBrowser try to save a new OSBrowser in database
func (m *Manager) CreateOSBrowser(osb *OSBrowser) error {

	func(in interface{}) {
		if ii, ok := in.(initializer.Simple); ok {
			ii.Initialize()
		}
	}(osb)

	return m.GetWDbMap().Insert(osb)
}

// UpdateOSBrowser try to update OSBrowser in database
func (m *Manager) UpdateOSBrowser(osb *OSBrowser) error {

	func(in interface{}) {
		if ii, ok := in.(initializer.Simple); ok {
			ii.Initialize()
		}
	}(osb)

	_, err := m.GetWDbMap().Update(osb)
	return err
}

// ListOSBrowsersWithFilter try to list all OSBrowsers without pagination
func (m *Manager) ListOSBrowsersWithFilter(filter string, params ...interface{}) []OSBrowser {
	filter = strings.Trim(filter, "\n\t ")
	if filter != "" {
		filter = "WHERE " + filter
	}
	var res []OSBrowser
	_, err := m.GetRDbMap().Select(
		&res,
		fmt.Sprintf("SELECT %s FROM %s %s", GetSelectFields(OSBrowserTableFull, ""), OSBrowserTableFull, filter),
		params...,
	)
	assert.Nil(err)

	return res
}

// ListOSBrowsers try to list all OSBrowsers without pagination
func (m *Manager) ListOSBrowsers() []OSBrowser {
	return m.ListOSBrowsersWithFilter("")
}

// CountOSBrowsersWithFilter count entity in OSBrowsers table with valid where filter
func (m *Manager) CountOSBrowsersWithFilter(filter string, params ...interface{}) int64 {
	filter = strings.Trim(filter, "\n\t ")
	if filter != "" {
		filter = "WHERE " + filter
	}
	cnt, err := m.GetRDbMap().SelectInt(
		fmt.Sprintf("SELECT COUNT(*) FROM %s %s", OSBrowserTableFull, filter),
		params...,
	)
	assert.Nil(err)

	return cnt
}

// CountOSBrowsers count entity in OSBrowsers table
func (m *Manager) CountOSBrowsers() int64 {
	return m.CountOSBrowsersWithFilter("")
}

// ListOSBrowsersWithPaginationFilter try to list all OSBrowsers with pagination and filter
func (m *Manager) ListOSBrowsersWithPaginationFilter(
	offset, perPage int, filter string, params ...interface{}) []OSBrowser {
	var res []OSBrowser
	filter = strings.Trim(filter, "\n\t ")
	if filter != "" {
		filter = "WHERE " + filter
	}

	filter += " LIMIT ?, ? "
	params = append(params, offset, perPage)

	// TODO : better pagination without offset and limit
	_, err := m.GetRDbMap().Select(
		&res,
		fmt.Sprintf("SELECT %s FROM %s %s", GetSelectFields(OSBrowserTableFull, ""), OSBrowserTableFull, filter),
		params...,
	)
	assert.Nil(err)

	return res
}

// ListOSBrowsersWithPagination try to list all OSBrowsers with pagination
func (m *Manager) ListOSBrowsersWithPagination(offset, perPage int) []OSBrowser {
	return m.ListOSBrowsersWithPaginationFilter(offset, perPage, "")
}

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

	os.UpdatedAt = time.Now()

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
		fmt.Sprintf("SELECT %s FROM %s %s", GetSelectFields(OSTableFull, ""), OSTableFull, filter),
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
		fmt.Sprintf("SELECT %s FROM %s %s", GetSelectFields(OSTableFull, ""), OSTableFull, filter),
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
		fmt.Sprintf("SELECT %s FROM %s WHERE name=?", GetSelectFields(OSTableFull, ""), OSTableFull),
		n,
	)

	if err != nil {
		return nil, err
	}

	return &res, nil
}

// FindOSByID return the OS base on its id
func (m *Manager) FindOSByID(id int64) (*OS, error) {
	var res OS
	err := m.GetRDbMap().SelectOne(
		&res,
		fmt.Sprintf("SELECT %s FROM %s WHERE id=?", GetSelectFields(OSTableFull, ""), OSTableFull),
		id,
	)

	if err != nil {
		return nil, err
	}

	return &res, nil
}
