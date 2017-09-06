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

// CreateRegion try to save a new Region in database
func (m *Manager) CreateRegion(r *Region) error {
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

// UpdateRegion try to update Region in database
func (m *Manager) UpdateRegion(r *Region) error {
	now := time.Now()
	r.UpdatedAt = now
	func(in interface{}) {
		if ii, ok := in.(initializer.Simple); ok {
			ii.Initialize()
		}
	}(r)

	_, err := m.GetWDbMap().Update(r)
	return err
}

// ListRegionsWithFilter try to list all Regions without pagination
func (m *Manager) ListRegionsWithFilter(filter string, params ...interface{}) []Region {
	filter = strings.Trim(filter, "\n\t ")
	if filter != "" {
		filter = "WHERE " + filter
	}
	var res []Region
	_, err := m.GetRDbMap().Select(
		&res,
		fmt.Sprintf("SELECT * FROM %s %s", RegionTableFull, filter),
		params...,
	)
	assert.Nil(err)

	return res
}

// ListRegions try to list all Regions without pagination
func (m *Manager) ListRegions() []Region {
	return m.ListRegionsWithFilter("")
}

// CountRegionsWithFilter count entity in Regions table with valid where filter
func (m *Manager) CountRegionsWithFilter(filter string, params ...interface{}) int64 {
	filter = strings.Trim(filter, "\n\t ")
	if filter != "" {
		filter = "WHERE " + filter
	}
	cnt, err := m.GetRDbMap().SelectInt(
		fmt.Sprintf("SELECT COUNT(*) FROM %s %s", RegionTableFull, filter),
		params...,
	)
	assert.Nil(err)

	return cnt
}

// CountRegions count entity in Regions table
func (m *Manager) CountRegions() int64 {
	return m.CountRegionsWithFilter("")
}

// ListRegionsWithPaginationFilter try to list all Regions with pagination and filter
func (m *Manager) ListRegionsWithPaginationFilter(
	offset, perPage int, filter string, params ...interface{}) []Region {
	var res []Region
	filter = strings.Trim(filter, "\n\t ")
	if filter != "" {
		filter = "WHERE " + filter
	}

	filter += " LIMIT ?, ? "
	params = append(params, offset, perPage)

	// TODO : better pagination without offset and limit
	_, err := m.GetRDbMap().Select(
		&res,
		fmt.Sprintf("SELECT * FROM %s %s", RegionTableFull, filter),
		params...,
	)
	assert.Nil(err)

	return res
}

// ListRegionsWithPagination try to list all Regions with pagination
func (m *Manager) ListRegionsWithPagination(offset, perPage int) []Region {
	return m.ListRegionsWithPaginationFilter(offset, perPage, "")
}
