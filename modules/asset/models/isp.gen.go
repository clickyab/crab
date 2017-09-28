// Code generated build with models DO NOT EDIT.

package models

import (
	"fmt"
	"strings"
	"time"

	"github.com/clickyab/services/assert"
	"github.com/clickyab/services/initializer"
)

// AUTO GENERATED CODE. DO NOT EDIT!

// CreateISP try to save a new ISP in database
func (m *Manager) CreateISP(isp *ISP) error {
	now := time.Now()
	isp.CreatedAt = now
	isp.UpdatedAt = now
	func(in interface{}) {
		if ii, ok := in.(initializer.Simple); ok {
			ii.Initialize()
		}
	}(isp)

	return m.GetWDbMap().Insert(isp)
}

// UpdateISP try to update ISP in database
func (m *Manager) UpdateISP(isp *ISP) error {
	now := time.Now()
	isp.UpdatedAt = now
	func(in interface{}) {
		if ii, ok := in.(initializer.Simple); ok {
			ii.Initialize()
		}
	}(isp)

	_, err := m.GetWDbMap().Update(isp)
	return err
}

// ListISPSWithFilter try to list all ISPS without pagination
func (m *Manager) ListISPSWithFilter(filter string, params ...interface{}) []ISP {
	filter = strings.Trim(filter, "\n\t ")
	if filter != "" {
		filter = "WHERE " + filter
	}
	var res []ISP
	_, err := m.GetRDbMap().Select(
		&res,
		fmt.Sprintf("SELECT * FROM %s %s", ISPTableFull, filter),
		params...,
	)
	assert.Nil(err)

	return res
}

// ListISPS try to list all ISPS without pagination
func (m *Manager) ListISPS() []ISP {
	return m.ListISPSWithFilter("")
}

// CountISPSWithFilter count entity in ISPS table with valid where filter
func (m *Manager) CountISPSWithFilter(filter string, params ...interface{}) int64 {
	filter = strings.Trim(filter, "\n\t ")
	if filter != "" {
		filter = "WHERE " + filter
	}
	cnt, err := m.GetRDbMap().SelectInt(
		fmt.Sprintf("SELECT COUNT(*) FROM %s %s", ISPTableFull, filter),
		params...,
	)
	assert.Nil(err)

	return cnt
}

// CountISPS count entity in ISPS table
func (m *Manager) CountISPS() int64 {
	return m.CountISPSWithFilter("")
}

// ListISPSWithPaginationFilter try to list all ISPS with pagination and filter
func (m *Manager) ListISPSWithPaginationFilter(
	offset, perPage int, filter string, params ...interface{}) []ISP {
	var res []ISP
	filter = strings.Trim(filter, "\n\t ")
	if filter != "" {
		filter = "WHERE " + filter
	}

	filter += " LIMIT ?, ? "
	params = append(params, offset, perPage)

	// TODO : better pagination without offset and limit
	_, err := m.GetRDbMap().Select(
		&res,
		fmt.Sprintf("SELECT * FROM %s %s", ISPTableFull, filter),
		params...,
	)
	assert.Nil(err)

	return res
}

// ListISPSWithPagination try to list all ISPS with pagination
func (m *Manager) ListISPSWithPagination(offset, perPage int) []ISP {
	return m.ListISPSWithPaginationFilter(offset, perPage, "")
}

// FindISPByName return the ISP base on its name
func (m *Manager) FindISPByName(n string) (*ISP, error) {
	var res ISP
	err := m.GetRDbMap().SelectOne(
		&res,
		fmt.Sprintf("SELECT * FROM %s WHERE name=?", ISPTableFull),
		n,
	)

	if err != nil {
		return nil, err
	}

	return &res, nil
}
