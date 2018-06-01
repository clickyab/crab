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

// CreateCampaignReportReceivers try to save a new CampaignReportReceivers in database
func (m *Manager) CreateCampaignReportReceivers(crr *CampaignReportReceivers) error {
	now := time.Now()
	crr.CreatedAt = now

	func(in interface{}) {
		if ii, ok := in.(initializer.Simple); ok {
			ii.Initialize()
		}
	}(crr)

	return m.GetWDbMap().Insert(crr)
}

// UpdateCampaignReportReceivers try to update CampaignReportReceivers in database
func (m *Manager) UpdateCampaignReportReceivers(crr *CampaignReportReceivers) error {

	func(in interface{}) {
		if ii, ok := in.(initializer.Simple); ok {
			ii.Initialize()
		}
	}(crr)

	_, err := m.GetWDbMap().Update(crr)
	return err
}

// ListCampaignReportReceiversWithFilter try to list all CampaignReportReceivers without pagination
func (m *Manager) ListCampaignReportReceiversWithFilter(filter string, params ...interface{}) []CampaignReportReceivers {
	filter = strings.Trim(filter, "\n\t ")
	if filter != "" {
		filter = "WHERE " + filter
	}
	var res []CampaignReportReceivers
	_, err := m.GetRDbMap().Select(
		&res,
		fmt.Sprintf("SELECT %s FROM %s %s", GetSelectFields(CampaignReportReceiversTableFull, ""), CampaignReportReceiversTableFull, filter),
		params...,
	)
	assert.Nil(err)

	return res
}

// ListCampaignReportReceivers try to list all CampaignReportReceivers without pagination
func (m *Manager) ListCampaignReportReceivers() []CampaignReportReceivers {
	return m.ListCampaignReportReceiversWithFilter("")
}

// CountCampaignReportReceiversWithFilter count entity in CampaignReportReceivers table with valid where filter
func (m *Manager) CountCampaignReportReceiversWithFilter(filter string, params ...interface{}) int64 {
	filter = strings.Trim(filter, "\n\t ")
	if filter != "" {
		filter = "WHERE " + filter
	}
	cnt, err := m.GetRDbMap().SelectInt(
		fmt.Sprintf("SELECT COUNT(*) FROM %s %s", CampaignReportReceiversTableFull, filter),
		params...,
	)
	assert.Nil(err)

	return cnt
}

// CountCampaignReportReceivers count entity in CampaignReportReceivers table
func (m *Manager) CountCampaignReportReceivers() int64 {
	return m.CountCampaignReportReceiversWithFilter("")
}

// ListCampaignReportReceiversWithPaginationFilter try to list all CampaignReportReceivers with pagination and filter
func (m *Manager) ListCampaignReportReceiversWithPaginationFilter(
	offset, perPage int, filter string, params ...interface{}) []CampaignReportReceivers {
	var res []CampaignReportReceivers
	filter = strings.Trim(filter, "\n\t ")
	if filter != "" {
		filter = "WHERE " + filter
	}

	filter += " LIMIT ?, ? "
	params = append(params, offset, perPage)

	// TODO : better pagination without offset and limit
	_, err := m.GetRDbMap().Select(
		&res,
		fmt.Sprintf("SELECT %s FROM %s %s", GetSelectFields(CampaignReportReceiversTableFull, ""), CampaignReportReceiversTableFull, filter),
		params...,
	)
	assert.Nil(err)

	return res
}

// ListCampaignReportReceiversWithPagination try to list all CampaignReportReceivers with pagination
func (m *Manager) ListCampaignReportReceiversWithPagination(offset, perPage int) []CampaignReportReceivers {
	return m.ListCampaignReportReceiversWithPaginationFilter(offset, perPage, "")
}
