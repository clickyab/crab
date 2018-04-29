// Code generated build with models DO NOT EDIT.

package aaa

import (
	"fmt"
	"strings"
	"time"

	"github.com/clickyab/services/assert"
	"github.com/clickyab/services/initializer"
	"github.com/clickyab/services/mysql"
)

// AUTO GENERATED CODE. DO NOT EDIT!

// CreateAuditLogDetail try to save a new AuditLogDetail in database
func (m *Manager) CreateAuditLogDetail(ald *AuditLogDetail) error {
	now := time.Now()
	ald.CreatedAt = now

	func(in interface{}) {
		if ii, ok := in.(initializer.Simple); ok {
			ii.Initialize()
		}
	}(ald)

	return m.GetWDbMap().Insert(ald)
}

// UpdateAuditLogDetail try to update AuditLogDetail in database
func (m *Manager) UpdateAuditLogDetail(ald *AuditLogDetail) error {

	ald.UpdatedAt = mysql.NullTime{Valid: true, Time: time.Now()}

	func(in interface{}) {
		if ii, ok := in.(initializer.Simple); ok {
			ii.Initialize()
		}
	}(ald)

	_, err := m.GetWDbMap().Update(ald)
	return err
}

// ListAuditLogDetailsWithFilter try to list all AuditLogDetails without pagination
func (m *Manager) ListAuditLogDetailsWithFilter(filter string, params ...interface{}) []AuditLogDetail {
	filter = strings.Trim(filter, "\n\t ")
	if filter != "" {
		filter = "WHERE " + filter
	}
	var res []AuditLogDetail
	_, err := m.GetRDbMap().Select(
		&res,
		fmt.Sprintf("SELECT %s FROM %s %s", getSelectFields(AuditLogDetailTableFull, ""), AuditLogDetailTableFull, filter),
		params...,
	)
	assert.Nil(err)

	return res
}

// ListAuditLogDetails try to list all AuditLogDetails without pagination
func (m *Manager) ListAuditLogDetails() []AuditLogDetail {
	return m.ListAuditLogDetailsWithFilter("")
}

// CountAuditLogDetailsWithFilter count entity in AuditLogDetails table with valid where filter
func (m *Manager) CountAuditLogDetailsWithFilter(filter string, params ...interface{}) int64 {
	filter = strings.Trim(filter, "\n\t ")
	if filter != "" {
		filter = "WHERE " + filter
	}
	cnt, err := m.GetRDbMap().SelectInt(
		fmt.Sprintf("SELECT COUNT(*) FROM %s %s", AuditLogDetailTableFull, filter),
		params...,
	)
	assert.Nil(err)

	return cnt
}

// CountAuditLogDetails count entity in AuditLogDetails table
func (m *Manager) CountAuditLogDetails() int64 {
	return m.CountAuditLogDetailsWithFilter("")
}

// ListAuditLogDetailsWithPaginationFilter try to list all AuditLogDetails with pagination and filter
func (m *Manager) ListAuditLogDetailsWithPaginationFilter(
	offset, perPage int, filter string, params ...interface{}) []AuditLogDetail {
	var res []AuditLogDetail
	filter = strings.Trim(filter, "\n\t ")
	if filter != "" {
		filter = "WHERE " + filter
	}

	filter += " LIMIT ?, ? "
	params = append(params, offset, perPage)

	// TODO : better pagination without offset and limit
	_, err := m.GetRDbMap().Select(
		&res,
		fmt.Sprintf("SELECT %s FROM %s %s", getSelectFields(AuditLogDetailTableFull, ""), AuditLogDetailTableFull, filter),
		params...,
	)
	assert.Nil(err)

	return res
}

// ListAuditLogDetailsWithPagination try to list all AuditLogDetails with pagination
func (m *Manager) ListAuditLogDetailsWithPagination(offset, perPage int) []AuditLogDetail {
	return m.ListAuditLogDetailsWithPaginationFilter(offset, perPage, "")
}

// FindAuditLogDetailByID return the AuditLogDetail base on its id
func (m *Manager) FindAuditLogDetailByID(id int64) (*AuditLogDetail, error) {
	var res AuditLogDetail
	err := m.GetRDbMap().SelectOne(
		&res,
		fmt.Sprintf("SELECT %s FROM %s WHERE id=?", getSelectFields(AuditLogDetailTableFull, ""), AuditLogDetailTableFull),
		id,
	)

	if err != nil {
		return nil, err
	}

	return &res, nil
}
