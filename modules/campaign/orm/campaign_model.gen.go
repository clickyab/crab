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

	c.UpdatedAt = time.Now()

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
		fmt.Sprintf("SELECT %s FROM %s %s", getSelectFields(CampaignTableFull, ""), CampaignTableFull, filter),
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
		fmt.Sprintf("SELECT %s FROM %s %s", getSelectFields(CampaignTableFull, ""), CampaignTableFull, filter),
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
		fmt.Sprintf("SELECT %s FROM %s WHERE id=?", getSelectFields(CampaignTableFull, ""), CampaignTableFull),
		id,
	)

	if err != nil {
		return nil, err
	}

	return &res, nil
}
