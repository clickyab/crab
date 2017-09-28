// Code generated build with models DO NOT EDIT.

package models

import (
	"fmt"
	"strings"
	"time"

	"github.com/clickyab/services/assert"
	"github.com/clickyab/services/initializer"
)

// AUTO GENERATED CODE. DO NOT EDIT!

// CreateBrowser try to save a new Browser in database
func (m *Manager) CreateBrowser(b *Browser) error {
	now := time.Now()
	b.CreatedAt = now
	b.UpdatedAt = now
	func(in interface{}) {
		if ii, ok := in.(initializer.Simple); ok {
			ii.Initialize()
		}
	}(b)

	return m.GetWDbMap().Insert(b)
}

// UpdateBrowser try to update Browser in database
func (m *Manager) UpdateBrowser(b *Browser) error {
	now := time.Now()
	b.UpdatedAt = now
	func(in interface{}) {
		if ii, ok := in.(initializer.Simple); ok {
			ii.Initialize()
		}
	}(b)

	_, err := m.GetWDbMap().Update(b)
	return err
}

// ListBrowsersWithFilter try to list all Browsers without pagination
func (m *Manager) ListBrowsersWithFilter(filter string, params ...interface{}) []Browser {
	filter = strings.Trim(filter, "\n\t ")
	if filter != "" {
		filter = "WHERE " + filter
	}
	var res []Browser
	_, err := m.GetRDbMap().Select(
		&res,
		fmt.Sprintf("SELECT * FROM %s %s", BrowserTableFull, filter),
		params...,
	)
	assert.Nil(err)

	return res
}

// ListBrowsers try to list all Browsers without pagination
func (m *Manager) ListBrowsers() []Browser {
	return m.ListBrowsersWithFilter("")
}

// CountBrowsersWithFilter count entity in Browsers table with valid where filter
func (m *Manager) CountBrowsersWithFilter(filter string, params ...interface{}) int64 {
	filter = strings.Trim(filter, "\n\t ")
	if filter != "" {
		filter = "WHERE " + filter
	}
	cnt, err := m.GetRDbMap().SelectInt(
		fmt.Sprintf("SELECT COUNT(*) FROM %s %s", BrowserTableFull, filter),
		params...,
	)
	assert.Nil(err)

	return cnt
}

// CountBrowsers count entity in Browsers table
func (m *Manager) CountBrowsers() int64 {
	return m.CountBrowsersWithFilter("")
}

// ListBrowsersWithPaginationFilter try to list all Browsers with pagination and filter
func (m *Manager) ListBrowsersWithPaginationFilter(
	offset, perPage int, filter string, params ...interface{}) []Browser {
	var res []Browser
	filter = strings.Trim(filter, "\n\t ")
	if filter != "" {
		filter = "WHERE " + filter
	}

	filter += " LIMIT ?, ? "
	params = append(params, offset, perPage)

	// TODO : better pagination without offset and limit
	_, err := m.GetRDbMap().Select(
		&res,
		fmt.Sprintf("SELECT * FROM %s %s", BrowserTableFull, filter),
		params...,
	)
	assert.Nil(err)

	return res
}

// ListBrowsersWithPagination try to list all Browsers with pagination
func (m *Manager) ListBrowsersWithPagination(offset, perPage int) []Browser {
	return m.ListBrowsersWithPaginationFilter(offset, perPage, "")
}

// FindBrowserByName return the Browser base on its name
func (m *Manager) FindBrowserByName(n string) (*Browser, error) {
	var res Browser
	err := m.GetRDbMap().SelectOne(
		&res,
		fmt.Sprintf("SELECT * FROM %s WHERE name=?", BrowserTableFull),
		n,
	)

	if err != nil {
		return nil, err
	}

	return &res, nil
}
