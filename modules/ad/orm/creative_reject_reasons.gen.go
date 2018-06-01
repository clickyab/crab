// Code generated build with models DO NOT EDIT.

package orm

import (
	"fmt"
	"strings"
	"time"

	"github.com/clickyab/services/assert"
	"github.com/clickyab/services/initializer"
	"github.com/go-sql-driver/mysql"
)

// AUTO GENERATED CODE. DO NOT EDIT!

// CreateCreativeRejectReasons try to save a new CreativeRejectReasons in database
func (m *Manager) CreateCreativeRejectReasons(crr *CreativeRejectReasons) error {
	now := time.Now()
	crr.CreatedAt = now

	func(in interface{}) {
		if ii, ok := in.(initializer.Simple); ok {
			ii.Initialize()
		}
	}(crr)

	return m.GetWDbMap().Insert(crr)
}

// UpdateCreativeRejectReasons try to update CreativeRejectReasons in database
func (m *Manager) UpdateCreativeRejectReasons(crr *CreativeRejectReasons) error {

	crr.UpdatedAt = mysql.NullTime{Valid: true, Time: time.Now()}

	func(in interface{}) {
		if ii, ok := in.(initializer.Simple); ok {
			ii.Initialize()
		}
	}(crr)

	_, err := m.GetWDbMap().Update(crr)
	return err
}

// ListCreativeRejectReasonsWithFilter try to list all CreativeRejectReasons without pagination
func (m *Manager) ListCreativeRejectReasonsWithFilter(filter string, params ...interface{}) []CreativeRejectReasons {
	filter = strings.Trim(filter, "\n\t ")
	if filter != "" {
		filter = "WHERE " + filter
	}
	var res []CreativeRejectReasons
	_, err := m.GetRDbMap().Select(
		&res,
		fmt.Sprintf("SELECT %s FROM %s %s", GetSelectFields(CreativeRejectReasonsTableFull, ""), CreativeRejectReasonsTableFull, filter),
		params...,
	)
	assert.Nil(err)

	return res
}

// ListCreativeRejectReasons try to list all CreativeRejectReasons without pagination
func (m *Manager) ListCreativeRejectReasons() []CreativeRejectReasons {
	return m.ListCreativeRejectReasonsWithFilter("")
}

// CountCreativeRejectReasonsWithFilter count entity in CreativeRejectReasons table with valid where filter
func (m *Manager) CountCreativeRejectReasonsWithFilter(filter string, params ...interface{}) int64 {
	filter = strings.Trim(filter, "\n\t ")
	if filter != "" {
		filter = "WHERE " + filter
	}
	cnt, err := m.GetRDbMap().SelectInt(
		fmt.Sprintf("SELECT COUNT(*) FROM %s %s", CreativeRejectReasonsTableFull, filter),
		params...,
	)
	assert.Nil(err)

	return cnt
}

// CountCreativeRejectReasons count entity in CreativeRejectReasons table
func (m *Manager) CountCreativeRejectReasons() int64 {
	return m.CountCreativeRejectReasonsWithFilter("")
}

// ListCreativeRejectReasonsWithPaginationFilter try to list all CreativeRejectReasons with pagination and filter
func (m *Manager) ListCreativeRejectReasonsWithPaginationFilter(
	offset, perPage int, filter string, params ...interface{}) []CreativeRejectReasons {
	var res []CreativeRejectReasons
	filter = strings.Trim(filter, "\n\t ")
	if filter != "" {
		filter = "WHERE " + filter
	}

	filter += " LIMIT ?, ? "
	params = append(params, offset, perPage)

	// TODO : better pagination without offset and limit
	_, err := m.GetRDbMap().Select(
		&res,
		fmt.Sprintf("SELECT %s FROM %s %s", GetSelectFields(CreativeRejectReasonsTableFull, ""), CreativeRejectReasonsTableFull, filter),
		params...,
	)
	assert.Nil(err)

	return res
}

// ListCreativeRejectReasonsWithPagination try to list all CreativeRejectReasons with pagination
func (m *Manager) ListCreativeRejectReasonsWithPagination(offset, perPage int) []CreativeRejectReasons {
	return m.ListCreativeRejectReasonsWithPaginationFilter(offset, perPage, "")
}

// FindCreativeRejectReasonsByID return the CreativeRejectReasons base on its id
func (m *Manager) FindCreativeRejectReasonsByID(id int64) (*CreativeRejectReasons, error) {
	var res CreativeRejectReasons
	err := m.GetRDbMap().SelectOne(
		&res,
		fmt.Sprintf("SELECT %s FROM %s WHERE id=?", GetSelectFields(CreativeRejectReasonsTableFull, ""), CreativeRejectReasonsTableFull),
		id,
	)

	if err != nil {
		return nil, err
	}

	return &res, nil
}
