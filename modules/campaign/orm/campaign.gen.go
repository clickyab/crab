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

// CreateCampaignAttributes try to save a new CampaignAttributes in database
func (m *Manager) CreateCampaignAttributes(ca *CampaignAttributes) error {

	func(in interface{}) {
		if ii, ok := in.(initializer.Simple); ok {
			ii.Initialize()
		}
	}(ca)

	return m.GetWDbMap().Insert(ca)
}

// UpdateCampaignAttributes try to update CampaignAttributes in database
func (m *Manager) UpdateCampaignAttributes(ca *CampaignAttributes) error {

	func(in interface{}) {
		if ii, ok := in.(initializer.Simple); ok {
			ii.Initialize()
		}
	}(ca)

	_, err := m.GetWDbMap().Update(ca)
	return err
}

// ListCampaignAttributesWithFilter try to list all CampaignAttributes without pagination
func (m *Manager) ListCampaignAttributesWithFilter(filter string, params ...interface{}) []CampaignAttributes {
	filter = strings.Trim(filter, "\n\t ")
	if filter != "" {
		filter = "WHERE " + filter
	}
	var res []CampaignAttributes
	_, err := m.GetRDbMap().Select(
		&res,
		fmt.Sprintf("SELECT * FROM %s %s", CampaignAttributesTableFull, filter),
		params...,
	)
	assert.Nil(err)

	return res
}

// ListCampaignAttributes try to list all CampaignAttributes without pagination
func (m *Manager) ListCampaignAttributes() []CampaignAttributes {
	return m.ListCampaignAttributesWithFilter("")
}

// CountCampaignAttributesWithFilter count entity in CampaignAttributes table with valid where filter
func (m *Manager) CountCampaignAttributesWithFilter(filter string, params ...interface{}) int64 {
	filter = strings.Trim(filter, "\n\t ")
	if filter != "" {
		filter = "WHERE " + filter
	}
	cnt, err := m.GetRDbMap().SelectInt(
		fmt.Sprintf("SELECT COUNT(*) FROM %s %s", CampaignAttributesTableFull, filter),
		params...,
	)
	assert.Nil(err)

	return cnt
}

// CountCampaignAttributes count entity in CampaignAttributes table
func (m *Manager) CountCampaignAttributes() int64 {
	return m.CountCampaignAttributesWithFilter("")
}

// ListCampaignAttributesWithPaginationFilter try to list all CampaignAttributes with pagination and filter
func (m *Manager) ListCampaignAttributesWithPaginationFilter(
	offset, perPage int, filter string, params ...interface{}) []CampaignAttributes {
	var res []CampaignAttributes
	filter = strings.Trim(filter, "\n\t ")
	if filter != "" {
		filter = "WHERE " + filter
	}

	filter += " LIMIT ?, ? "
	params = append(params, offset, perPage)

	// TODO : better pagination without offset and limit
	_, err := m.GetRDbMap().Select(
		&res,
		fmt.Sprintf("SELECT * FROM %s %s", CampaignAttributesTableFull, filter),
		params...,
	)
	assert.Nil(err)

	return res
}

// ListCampaignAttributesWithPagination try to list all CampaignAttributes with pagination
func (m *Manager) ListCampaignAttributesWithPagination(offset, perPage int) []CampaignAttributes {
	return m.ListCampaignAttributesWithPaginationFilter(offset, perPage, "")
}

// FindCampaignAttributesByID return the CampaignAttributes base on its id
func (m *Manager) FindCampaignAttributesByID(id int64) (*CampaignAttributes, error) {
	var res CampaignAttributes
	err := m.GetRDbMap().SelectOne(
		&res,
		fmt.Sprintf("SELECT * FROM %s WHERE id=?", CampaignAttributesTableFull),
		id,
	)

	if err != nil {
		return nil, err
	}

	return &res, nil
}

// FindCampaignAttributesByCampaignID return the CampaignAttributes base on its campaign_id
func (m *Manager) FindCampaignAttributesByCampaignID(ci int64) (*CampaignAttributes, error) {
	var res CampaignAttributes
	err := m.GetRDbMap().SelectOne(
		&res,
		fmt.Sprintf("SELECT * FROM %s WHERE campaign_id=?", CampaignAttributesTableFull),
		ci,
	)

	if err != nil {
		return nil, err
	}

	return &res, nil
}

// CreateCampaign try to save a new Campaign in database
func (m *Manager) CreateCampaign(c *Campaign) error {
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

// UpdateCampaign try to update Campaign in database
func (m *Manager) UpdateCampaign(c *Campaign) error {
	now := time.Now()
	c.UpdatedAt = now
	func(in interface{}) {
		if ii, ok := in.(initializer.Simple); ok {
			ii.Initialize()
		}
	}(c)

	_, err := m.GetWDbMap().Update(c)
	return err
}

// ListCampaignsWithFilter try to list all Campaigns without pagination
func (m *Manager) ListCampaignsWithFilter(filter string, params ...interface{}) []Campaign {
	filter = strings.Trim(filter, "\n\t ")
	if filter != "" {
		filter = "WHERE " + filter
	}
	var res []Campaign
	_, err := m.GetRDbMap().Select(
		&res,
		fmt.Sprintf("SELECT * FROM %s %s", CampaignTableFull, filter),
		params...,
	)
	assert.Nil(err)

	return res
}

// ListCampaigns try to list all Campaigns without pagination
func (m *Manager) ListCampaigns() []Campaign {
	return m.ListCampaignsWithFilter("")
}

// CountCampaignsWithFilter count entity in Campaigns table with valid where filter
func (m *Manager) CountCampaignsWithFilter(filter string, params ...interface{}) int64 {
	filter = strings.Trim(filter, "\n\t ")
	if filter != "" {
		filter = "WHERE " + filter
	}
	cnt, err := m.GetRDbMap().SelectInt(
		fmt.Sprintf("SELECT COUNT(*) FROM %s %s", CampaignTableFull, filter),
		params...,
	)
	assert.Nil(err)

	return cnt
}

// CountCampaigns count entity in Campaigns table
func (m *Manager) CountCampaigns() int64 {
	return m.CountCampaignsWithFilter("")
}

// ListCampaignsWithPaginationFilter try to list all Campaigns with pagination and filter
func (m *Manager) ListCampaignsWithPaginationFilter(
	offset, perPage int, filter string, params ...interface{}) []Campaign {
	var res []Campaign
	filter = strings.Trim(filter, "\n\t ")
	if filter != "" {
		filter = "WHERE " + filter
	}

	filter += " LIMIT ?, ? "
	params = append(params, offset, perPage)

	// TODO : better pagination without offset and limit
	_, err := m.GetRDbMap().Select(
		&res,
		fmt.Sprintf("SELECT * FROM %s %s", CampaignTableFull, filter),
		params...,
	)
	assert.Nil(err)

	return res
}

// ListCampaignsWithPagination try to list all Campaigns with pagination
func (m *Manager) ListCampaignsWithPagination(offset, perPage int) []Campaign {
	return m.ListCampaignsWithPaginationFilter(offset, perPage, "")
}

// FindCampaignByID return the Campaign base on its id
func (m *Manager) FindCampaignByID(id int64) (*Campaign, error) {
	var res Campaign
	err := m.GetRDbMap().SelectOne(
		&res,
		fmt.Sprintf("SELECT * FROM %s WHERE id=?", CampaignTableFull),
		id,
	)

	if err != nil {
		return nil, err
	}

	return &res, nil
}
