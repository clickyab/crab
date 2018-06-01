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

// CreatePlatform try to save a new Platform in database
func (m *Manager) CreatePlatform(p *Platform) error {
	now := time.Now()
	p.CreatedAt = now
	p.UpdatedAt = now
	func(in interface{}) {
		if ii, ok := in.(initializer.Simple); ok {
			ii.Initialize()
		}
	}(p)

	return m.GetWDbMap().Insert(p)
}

// UpdatePlatform try to update Platform in database
func (m *Manager) UpdatePlatform(p *Platform) error {

	p.UpdatedAt = time.Now()

	func(in interface{}) {
		if ii, ok := in.(initializer.Simple); ok {
			ii.Initialize()
		}
	}(p)

	_, err := m.GetWDbMap().Update(p)
	return err
}

// ListPlatformsWithFilter try to list all Platforms without pagination
func (m *Manager) ListPlatformsWithFilter(filter string, params ...interface{}) []Platform {
	filter = strings.Trim(filter, "\n\t ")
	if filter != "" {
		filter = "WHERE " + filter
	}
	var res []Platform
	_, err := m.GetRDbMap().Select(
		&res,
		fmt.Sprintf("SELECT %s FROM %s %s", GetSelectFields(PlatformTableFull, ""), PlatformTableFull, filter),
		params...,
	)
	assert.Nil(err)

	return res
}

// ListPlatforms try to list all Platforms without pagination
func (m *Manager) ListPlatforms() []Platform {
	return m.ListPlatformsWithFilter("")
}

// CountPlatformsWithFilter count entity in Platforms table with valid where filter
func (m *Manager) CountPlatformsWithFilter(filter string, params ...interface{}) int64 {
	filter = strings.Trim(filter, "\n\t ")
	if filter != "" {
		filter = "WHERE " + filter
	}
	cnt, err := m.GetRDbMap().SelectInt(
		fmt.Sprintf("SELECT COUNT(*) FROM %s %s", PlatformTableFull, filter),
		params...,
	)
	assert.Nil(err)

	return cnt
}

// CountPlatforms count entity in Platforms table
func (m *Manager) CountPlatforms() int64 {
	return m.CountPlatformsWithFilter("")
}

// ListPlatformsWithPaginationFilter try to list all Platforms with pagination and filter
func (m *Manager) ListPlatformsWithPaginationFilter(
	offset, perPage int, filter string, params ...interface{}) []Platform {
	var res []Platform
	filter = strings.Trim(filter, "\n\t ")
	if filter != "" {
		filter = "WHERE " + filter
	}

	filter += " LIMIT ?, ? "
	params = append(params, offset, perPage)

	// TODO : better pagination without offset and limit
	_, err := m.GetRDbMap().Select(
		&res,
		fmt.Sprintf("SELECT %s FROM %s %s", GetSelectFields(PlatformTableFull, ""), PlatformTableFull, filter),
		params...,
	)
	assert.Nil(err)

	return res
}

// ListPlatformsWithPagination try to list all Platforms with pagination
func (m *Manager) ListPlatformsWithPagination(offset, perPage int) []Platform {
	return m.ListPlatformsWithPaginationFilter(offset, perPage, "")
}

// FindPlatformByName return the Platform base on its name
func (m *Manager) FindPlatformByName(n string) (*Platform, error) {
	var res Platform
	err := m.GetRDbMap().SelectOne(
		&res,
		fmt.Sprintf("SELECT %s FROM %s WHERE name=?", GetSelectFields(PlatformTableFull, ""), PlatformTableFull),
		n,
	)

	if err != nil {
		return nil, err
	}

	return &res, nil
}
