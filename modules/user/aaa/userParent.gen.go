// Code generated build with models DO NOT EDIT.

package aaa

import (
	"fmt"
	"strings"
	"time"

	"github.com/clickyab/services/assert"
	"github.com/clickyab/services/initializer"
)

// AUTO GENERATED CODE. DO NOT EDIT!

// CreateConsularCustomer try to save a new ConsularCustomer in database
func (m *Manager) CreateConsularCustomer(cc *ConsularCustomer) error {
	now := time.Now()
	cc.CreatedAt = now

	func(in interface{}) {
		if ii, ok := in.(initializer.Simple); ok {
			ii.Initialize()
		}
	}(cc)

	return m.GetWDbMap().Insert(cc)
}

// UpdateConsularCustomer try to update ConsularCustomer in database
func (m *Manager) UpdateConsularCustomer(cc *ConsularCustomer) error {

	func(in interface{}) {
		if ii, ok := in.(initializer.Simple); ok {
			ii.Initialize()
		}
	}(cc)

	_, err := m.GetWDbMap().Update(cc)
	return err
}

// ListConsularCustomersWithFilter try to list all ConsularCustomers without pagination
func (m *Manager) ListConsularCustomersWithFilter(filter string, params ...interface{}) []ConsularCustomer {
	filter = strings.Trim(filter, "\n\t ")
	if filter != "" {
		filter = "WHERE " + filter
	}
	var res []ConsularCustomer
	_, err := m.GetRDbMap().Select(
		&res,
		fmt.Sprintf("SELECT * FROM %s %s", ConsularCustomerTableFull, filter),
		params...,
	)
	assert.Nil(err)

	return res
}

// ListConsularCustomers try to list all ConsularCustomers without pagination
func (m *Manager) ListConsularCustomers() []ConsularCustomer {
	return m.ListConsularCustomersWithFilter("")
}

// CountConsularCustomersWithFilter count entity in ConsularCustomers table with valid where filter
func (m *Manager) CountConsularCustomersWithFilter(filter string, params ...interface{}) int64 {
	filter = strings.Trim(filter, "\n\t ")
	if filter != "" {
		filter = "WHERE " + filter
	}
	cnt, err := m.GetRDbMap().SelectInt(
		fmt.Sprintf("SELECT COUNT(*) FROM %s %s", ConsularCustomerTableFull, filter),
		params...,
	)
	assert.Nil(err)

	return cnt
}

// CountConsularCustomers count entity in ConsularCustomers table
func (m *Manager) CountConsularCustomers() int64 {
	return m.CountConsularCustomersWithFilter("")
}

// ListConsularCustomersWithPaginationFilter try to list all ConsularCustomers with pagination and filter
func (m *Manager) ListConsularCustomersWithPaginationFilter(
	offset, perPage int, filter string, params ...interface{}) []ConsularCustomer {
	var res []ConsularCustomer
	filter = strings.Trim(filter, "\n\t ")
	if filter != "" {
		filter = "WHERE " + filter
	}

	filter += " LIMIT ?, ? "
	params = append(params, offset, perPage)

	// TODO : better pagination without offset and limit
	_, err := m.GetRDbMap().Select(
		&res,
		fmt.Sprintf("SELECT * FROM %s %s", ConsularCustomerTableFull, filter),
		params...,
	)
	assert.Nil(err)

	return res
}

// ListConsularCustomersWithPagination try to list all ConsularCustomers with pagination
func (m *Manager) ListConsularCustomersWithPagination(offset, perPage int) []ConsularCustomer {
	return m.ListConsularCustomersWithPaginationFilter(offset, perPage, "")
}
