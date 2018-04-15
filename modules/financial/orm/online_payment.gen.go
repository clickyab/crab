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

// CreateOnlinePayment try to save a new OnlinePayment in database
func (m *Manager) CreateOnlinePayment(op *OnlinePayment) error {
	now := time.Now()
	op.CreatedAt = now
	op.UpdatedAt = now
	func(in interface{}) {
		if ii, ok := in.(initializer.Simple); ok {
			ii.Initialize()
		}
	}(op)

	return m.GetWDbMap().Insert(op)
}

// UpdateOnlinePayment try to update OnlinePayment in database
func (m *Manager) UpdateOnlinePayment(op *OnlinePayment) error {
	now := time.Now()
	op.UpdatedAt = now
	func(in interface{}) {
		if ii, ok := in.(initializer.Simple); ok {
			ii.Initialize()
		}
	}(op)

	_, err := m.GetWDbMap().Update(op)
	return err
}

// ListOnlinePaymentsWithFilter try to list all OnlinePayments without pagination
func (m *Manager) ListOnlinePaymentsWithFilter(filter string, params ...interface{}) []OnlinePayment {
	filter = strings.Trim(filter, "\n\t ")
	if filter != "" {
		filter = "WHERE " + filter
	}
	var res []OnlinePayment
	_, err := m.GetRDbMap().Select(
		&res,
		fmt.Sprintf("SELECT %s FROM %s %s", getSelectFields(OnlinePaymentTableFull, ""), OnlinePaymentTableFull, filter),
		params...,
	)
	assert.Nil(err)

	return res
}

// ListOnlinePayments try to list all OnlinePayments without pagination
func (m *Manager) ListOnlinePayments() []OnlinePayment {
	return m.ListOnlinePaymentsWithFilter("")
}

// CountOnlinePaymentsWithFilter count entity in OnlinePayments table with valid where filter
func (m *Manager) CountOnlinePaymentsWithFilter(filter string, params ...interface{}) int64 {
	filter = strings.Trim(filter, "\n\t ")
	if filter != "" {
		filter = "WHERE " + filter
	}
	cnt, err := m.GetRDbMap().SelectInt(
		fmt.Sprintf("SELECT COUNT(*) FROM %s %s", OnlinePaymentTableFull, filter),
		params...,
	)
	assert.Nil(err)

	return cnt
}

// CountOnlinePayments count entity in OnlinePayments table
func (m *Manager) CountOnlinePayments() int64 {
	return m.CountOnlinePaymentsWithFilter("")
}

// ListOnlinePaymentsWithPaginationFilter try to list all OnlinePayments with pagination and filter
func (m *Manager) ListOnlinePaymentsWithPaginationFilter(
	offset, perPage int, filter string, params ...interface{}) []OnlinePayment {
	var res []OnlinePayment
	filter = strings.Trim(filter, "\n\t ")
	if filter != "" {
		filter = "WHERE " + filter
	}

	filter += " LIMIT ?, ? "
	params = append(params, offset, perPage)

	// TODO : better pagination without offset and limit
	_, err := m.GetRDbMap().Select(
		&res,
		fmt.Sprintf("SELECT %s FROM %s %s", getSelectFields(OnlinePaymentTableFull, ""), OnlinePaymentTableFull, filter),
		params...,
	)
	assert.Nil(err)

	return res
}

// ListOnlinePaymentsWithPagination try to list all OnlinePayments with pagination
func (m *Manager) ListOnlinePaymentsWithPagination(offset, perPage int) []OnlinePayment {
	return m.ListOnlinePaymentsWithPaginationFilter(offset, perPage, "")
}

// FindOnlinePaymentByID return the OnlinePayment base on its id
func (m *Manager) FindOnlinePaymentByID(id int64) (*OnlinePayment, error) {
	var res OnlinePayment
	err := m.GetRDbMap().SelectOne(
		&res,
		fmt.Sprintf("SELECT %s FROM %s WHERE id=?", getSelectFields(OnlinePaymentTableFull, ""), OnlinePaymentTableFull),
		id,
	)

	if err != nil {
		return nil, err
	}

	return &res, nil
}
