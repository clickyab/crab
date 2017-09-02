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

// Createassets try to save a new assets in database
func (m *Manager) Createassets(a *assets) error {
	now := time.Now()
	a.CreatedAt = now
	a.UpdatedAt = now
	func(in interface{}) {
		if ii, ok := in.(initializer.Simple); ok {
			ii.Initialize()
		}
	}(a)

	return m.GetWDbMap().Insert(a)
}

// Updateassets try to update assets in database
func (m *Manager) Updateassets(a *assets) error {
	now := time.Now()
	a.UpdatedAt = now
	func(in interface{}) {
		if ii, ok := in.(initializer.Simple); ok {
			ii.Initialize()
		}
	}(a)

	_, err := m.GetWDbMap().Update(a)
	return err
}

// ListassetsWithFilter try to list all assets without pagination
func (m *Manager) ListassetsWithFilter(filter string, params ...interface{}) []assets {
	filter = strings.Trim(filter, "\n\t ")
	if filter != "" {
		filter = "WHERE " + filter
	}
	var res []assets
	_, err := m.GetRDbMap().Select(
		&res,
		fmt.Sprintf("SELECT * FROM %s %s", assetsTableFull, filter),
		params...,
	)
	assert.Nil(err)

	return res
}

// Listassets try to list all assets without pagination
func (m *Manager) Listassets() []assets {
	return m.ListassetsWithFilter("")
}

// CountassetsWithFilter count entity in assets table with valid where filter
func (m *Manager) CountassetsWithFilter(filter string, params ...interface{}) int64 {
	filter = strings.Trim(filter, "\n\t ")
	if filter != "" {
		filter = "WHERE " + filter
	}
	cnt, err := m.GetRDbMap().SelectInt(
		fmt.Sprintf("SELECT COUNT(*) FROM %s %s", assetsTableFull, filter),
		params...,
	)
	assert.Nil(err)

	return cnt
}

// Countassets count entity in assets table
func (m *Manager) Countassets() int64 {
	return m.CountassetsWithFilter("")
}

// ListassetsWithPaginationFilter try to list all assets with pagination and filter
func (m *Manager) ListassetsWithPaginationFilter(
	offset, perPage int, filter string, params ...interface{}) []assets {
	var res []assets
	filter = strings.Trim(filter, "\n\t ")
	if filter != "" {
		filter = "WHERE " + filter
	}

	filter += " LIMIT ?, ? "
	params = append(params, offset, perPage)

	// TODO : better pagination without offset and limit
	_, err := m.GetRDbMap().Select(
		&res,
		fmt.Sprintf("SELECT * FROM %s %s", assetsTableFull, filter),
		params...,
	)
	assert.Nil(err)

	return res
}

// ListassetsWithPagination try to list all assets with pagination
func (m *Manager) ListassetsWithPagination(offset, perPage int) []assets {
	return m.ListassetsWithPaginationFilter(offset, perPage, "")
}

// FindassetsByID return the assets base on its id
func (m *Manager) FindassetsByID(id int64) (*assets, error) {
	var res assets
	err := m.GetRDbMap().SelectOne(
		&res,
		fmt.Sprintf("SELECT * FROM %s WHERE id=?", assetsTableFull),
		id,
	)

	if err != nil {
		return nil, err
	}

	return &res, nil
}
