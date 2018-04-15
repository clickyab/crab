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

// CreateBilling try to save a new Billing in database
func (m *Manager) CreateBilling(b *Billing) error {
	now := time.Now()
	b.CreatedAt = now

	func(in interface{}) {
		if ii, ok := in.(initializer.Simple); ok {
			ii.Initialize()
		}
	}(b)

	return m.GetWDbMap().Insert(b)
}

// UpdateBilling try to update Billing in database
func (m *Manager) UpdateBilling(b *Billing) error {

	func(in interface{}) {
		if ii, ok := in.(initializer.Simple); ok {
			ii.Initialize()
		}
	}(b)

	_, err := m.GetWDbMap().Update(b)
	return err
}

// ListBillingsWithFilter try to list all Billings without pagination
func (m *Manager) ListBillingsWithFilter(filter string, params ...interface{}) []Billing {
	filter = strings.Trim(filter, "\n\t ")
	if filter != "" {
		filter = "WHERE " + filter
	}
	var res []Billing
	_, err := m.GetRDbMap().Select(
		&res,
		fmt.Sprintf("SELECT %s FROM %s %s", getSelectFields(BillingTableFull, ""), BillingTableFull, filter),
		params...,
	)
	assert.Nil(err)

	return res
}

// ListBillings try to list all Billings without pagination
func (m *Manager) ListBillings() []Billing {
	return m.ListBillingsWithFilter("")
}

// CountBillingsWithFilter count entity in Billings table with valid where filter
func (m *Manager) CountBillingsWithFilter(filter string, params ...interface{}) int64 {
	filter = strings.Trim(filter, "\n\t ")
	if filter != "" {
		filter = "WHERE " + filter
	}
	cnt, err := m.GetRDbMap().SelectInt(
		fmt.Sprintf("SELECT COUNT(*) FROM %s %s", BillingTableFull, filter),
		params...,
	)
	assert.Nil(err)

	return cnt
}

// CountBillings count entity in Billings table
func (m *Manager) CountBillings() int64 {
	return m.CountBillingsWithFilter("")
}

// ListBillingsWithPaginationFilter try to list all Billings with pagination and filter
func (m *Manager) ListBillingsWithPaginationFilter(
	offset, perPage int, filter string, params ...interface{}) []Billing {
	var res []Billing
	filter = strings.Trim(filter, "\n\t ")
	if filter != "" {
		filter = "WHERE " + filter
	}

	filter += " LIMIT ?, ? "
	params = append(params, offset, perPage)

	// TODO : better pagination without offset and limit
	_, err := m.GetRDbMap().Select(
		&res,
		fmt.Sprintf("SELECT %s FROM %s %s", getSelectFields(BillingTableFull, ""), BillingTableFull, filter),
		params...,
	)
	assert.Nil(err)

	return res
}

// ListBillingsWithPagination try to list all Billings with pagination
func (m *Manager) ListBillingsWithPagination(offset, perPage int) []Billing {
	return m.ListBillingsWithPaginationFilter(offset, perPage, "")
}

// FindBillingByID return the Billing base on its id
func (m *Manager) FindBillingByID(id int64) (*Billing, error) {
	var res Billing
	err := m.GetRDbMap().SelectOne(
		&res,
		fmt.Sprintf("SELECT %s FROM %s WHERE id=?", getSelectFields(BillingTableFull, ""), BillingTableFull),
		id,
	)

	if err != nil {
		return nil, err
	}

	return &res, nil
}
