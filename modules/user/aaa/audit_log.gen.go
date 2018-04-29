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

// CreateAuditLog try to save a new AuditLog in database
func (m *Manager) CreateAuditLog(al *AuditLog) error {
	now := time.Now()
	al.CreatedAt = now

	func(in interface{}) {
		if ii, ok := in.(initializer.Simple); ok {
			ii.Initialize()
		}
	}(al)

	return m.GetWDbMap().Insert(al)
}

// UpdateAuditLog try to update AuditLog in database
func (m *Manager) UpdateAuditLog(al *AuditLog) error {

	al.UpdatedAt = mysql.NullTime{Valid: true, Time: time.Now()}

	func(in interface{}) {
		if ii, ok := in.(initializer.Simple); ok {
			ii.Initialize()
		}
	}(al)

	_, err := m.GetWDbMap().Update(al)
	return err
}

// ListAuditLogsWithFilter try to list all AuditLogs without pagination
func (m *Manager) ListAuditLogsWithFilter(filter string, params ...interface{}) []AuditLog {
	filter = strings.Trim(filter, "\n\t ")
	if filter != "" {
		filter = "WHERE " + filter
	}
	var res []AuditLog
	_, err := m.GetRDbMap().Select(
		&res,
		fmt.Sprintf("SELECT %s FROM %s %s", getSelectFields(AuditLogTableFull, ""), AuditLogTableFull, filter),
		params...,
	)
	assert.Nil(err)

	return res
}

// ListAuditLogs try to list all AuditLogs without pagination
func (m *Manager) ListAuditLogs() []AuditLog {
	return m.ListAuditLogsWithFilter("")
}

// CountAuditLogsWithFilter count entity in AuditLogs table with valid where filter
func (m *Manager) CountAuditLogsWithFilter(filter string, params ...interface{}) int64 {
	filter = strings.Trim(filter, "\n\t ")
	if filter != "" {
		filter = "WHERE " + filter
	}
	cnt, err := m.GetRDbMap().SelectInt(
		fmt.Sprintf("SELECT COUNT(*) FROM %s %s", AuditLogTableFull, filter),
		params...,
	)
	assert.Nil(err)

	return cnt
}

// CountAuditLogs count entity in AuditLogs table
func (m *Manager) CountAuditLogs() int64 {
	return m.CountAuditLogsWithFilter("")
}

// ListAuditLogsWithPaginationFilter try to list all AuditLogs with pagination and filter
func (m *Manager) ListAuditLogsWithPaginationFilter(
	offset, perPage int, filter string, params ...interface{}) []AuditLog {
	var res []AuditLog
	filter = strings.Trim(filter, "\n\t ")
	if filter != "" {
		filter = "WHERE " + filter
	}

	filter += " LIMIT ?, ? "
	params = append(params, offset, perPage)

	// TODO : better pagination without offset and limit
	_, err := m.GetRDbMap().Select(
		&res,
		fmt.Sprintf("SELECT %s FROM %s %s", getSelectFields(AuditLogTableFull, ""), AuditLogTableFull, filter),
		params...,
	)
	assert.Nil(err)

	return res
}

// ListAuditLogsWithPagination try to list all AuditLogs with pagination
func (m *Manager) ListAuditLogsWithPagination(offset, perPage int) []AuditLog {
	return m.ListAuditLogsWithPaginationFilter(offset, perPage, "")
}

// FindAuditLogByID return the AuditLog base on its id
func (m *Manager) FindAuditLogByID(id int64) (*AuditLog, error) {
	var res AuditLog
	err := m.GetRDbMap().SelectOne(
		&res,
		fmt.Sprintf("SELECT %s FROM %s WHERE id=?", getSelectFields(AuditLogTableFull, ""), AuditLogTableFull),
		id,
	)

	if err != nil {
		return nil, err
	}

	return &res, nil
}
