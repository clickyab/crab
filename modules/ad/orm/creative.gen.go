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

// CreateCreative try to save a new Creative in database
func (m *Manager) CreateCreative(c *Creative) error {
	now := time.Now()
	c.CreatedAt = now
	c.UpdatedAt = now
	func(in interface{}) {
		if ii, ok := in.(initializer.Simple); ok {
			ii.Initialize()
		}
	}(c)

	return m.GetWDbMap().Insert(c)
}

// UpdateCreative try to update Creative in database
func (m *Manager) UpdateCreative(c *Creative) error {

	c.UpdatedAt = time.Now()

	func(in interface{}) {
		if ii, ok := in.(initializer.Simple); ok {
			ii.Initialize()
		}
	}(c)

	_, err := m.GetWDbMap().Update(c)
	return err
}

// ListCreativesWithFilter try to list all Creatives without pagination
func (m *Manager) ListCreativesWithFilter(filter string, params ...interface{}) []Creative {
	filter = strings.Trim(filter, "\n\t ")
	if filter != "" {
		filter = "WHERE " + filter
	}
	var res []Creative
	_, err := m.GetRDbMap().Select(
		&res,
		fmt.Sprintf("SELECT %s FROM %s %s", getSelectFields(CreativeTableFull, ""), CreativeTableFull, filter),
		params...,
	)
	assert.Nil(err)

	return res
}

// ListCreatives try to list all Creatives without pagination
func (m *Manager) ListCreatives() []Creative {
	return m.ListCreativesWithFilter("")
}

// CountCreativesWithFilter count entity in Creatives table with valid where filter
func (m *Manager) CountCreativesWithFilter(filter string, params ...interface{}) int64 {
	filter = strings.Trim(filter, "\n\t ")
	if filter != "" {
		filter = "WHERE " + filter
	}
	cnt, err := m.GetRDbMap().SelectInt(
		fmt.Sprintf("SELECT COUNT(*) FROM %s %s", CreativeTableFull, filter),
		params...,
	)
	assert.Nil(err)

	return cnt
}

// CountCreatives count entity in Creatives table
func (m *Manager) CountCreatives() int64 {
	return m.CountCreativesWithFilter("")
}

// ListCreativesWithPaginationFilter try to list all Creatives with pagination and filter
func (m *Manager) ListCreativesWithPaginationFilter(
	offset, perPage int, filter string, params ...interface{}) []Creative {
	var res []Creative
	filter = strings.Trim(filter, "\n\t ")
	if filter != "" {
		filter = "WHERE " + filter
	}

	filter += " LIMIT ?, ? "
	params = append(params, offset, perPage)

	// TODO : better pagination without offset and limit
	_, err := m.GetRDbMap().Select(
		&res,
		fmt.Sprintf("SELECT %s FROM %s %s", getSelectFields(CreativeTableFull, ""), CreativeTableFull, filter),
		params...,
	)
	assert.Nil(err)

	return res
}

// ListCreativesWithPagination try to list all Creatives with pagination
func (m *Manager) ListCreativesWithPagination(offset, perPage int) []Creative {
	return m.ListCreativesWithPaginationFilter(offset, perPage, "")
}

// FindCreativeByID return the Creative base on its id
func (m *Manager) FindCreativeByID(id int64) (*Creative, error) {
	var res Creative
	err := m.GetRDbMap().SelectOne(
		&res,
		fmt.Sprintf("SELECT %s FROM %s WHERE id=?", getSelectFields(CreativeTableFull, ""), CreativeTableFull),
		id,
	)

	if err != nil {
		return nil, err
	}

	return &res, nil
}

// FindCreativeByCampaignID return the Creative base on its campaign_id
func (m *Manager) FindCreativeByCampaignID(ci int64) (*Creative, error) {
	var res Creative
	err := m.GetRDbMap().SelectOne(
		&res,
		fmt.Sprintf("SELECT %s FROM %s WHERE campaign_id=?", getSelectFields(CreativeTableFull, ""), CreativeTableFull),
		ci,
	)

	if err != nil {
		return nil, err
	}

	return &res, nil
}
