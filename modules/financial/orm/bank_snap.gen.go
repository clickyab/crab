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

// CreateBankSnap try to save a new BankSnap in database
func (m *Manager) CreateBankSnap(bs *BankSnap) error {
	now := time.Now()
	bs.CreatedAt = now
	bs.UpdatedAt = now
	func(in interface{}) {
		if ii, ok := in.(initializer.Simple); ok {
			ii.Initialize()
		}
	}(bs)

	return m.GetWDbMap().Insert(bs)
}

// UpdateBankSnap try to update BankSnap in database
func (m *Manager) UpdateBankSnap(bs *BankSnap) error {

	bs.UpdatedAt = time.Now()

	func(in interface{}) {
		if ii, ok := in.(initializer.Simple); ok {
			ii.Initialize()
		}
	}(bs)

	_, err := m.GetWDbMap().Update(bs)
	return err
}

// ListBankSnapsWithFilter try to list all BankSnaps without pagination
func (m *Manager) ListBankSnapsWithFilter(filter string, params ...interface{}) []BankSnap {
	filter = strings.Trim(filter, "\n\t ")
	if filter != "" {
		filter = "WHERE " + filter
	}
	var res []BankSnap
	_, err := m.GetRDbMap().Select(
		&res,
		fmt.Sprintf("SELECT %s FROM %s %s", getSelectFields(BankSnapTableFull, ""), BankSnapTableFull, filter),
		params...,
	)
	assert.Nil(err)

	return res
}

// ListBankSnaps try to list all BankSnaps without pagination
func (m *Manager) ListBankSnaps() []BankSnap {
	return m.ListBankSnapsWithFilter("")
}

// CountBankSnapsWithFilter count entity in BankSnaps table with valid where filter
func (m *Manager) CountBankSnapsWithFilter(filter string, params ...interface{}) int64 {
	filter = strings.Trim(filter, "\n\t ")
	if filter != "" {
		filter = "WHERE " + filter
	}
	cnt, err := m.GetRDbMap().SelectInt(
		fmt.Sprintf("SELECT COUNT(*) FROM %s %s", BankSnapTableFull, filter),
		params...,
	)
	assert.Nil(err)

	return cnt
}

// CountBankSnaps count entity in BankSnaps table
func (m *Manager) CountBankSnaps() int64 {
	return m.CountBankSnapsWithFilter("")
}

// ListBankSnapsWithPaginationFilter try to list all BankSnaps with pagination and filter
func (m *Manager) ListBankSnapsWithPaginationFilter(
	offset, perPage int, filter string, params ...interface{}) []BankSnap {
	var res []BankSnap
	filter = strings.Trim(filter, "\n\t ")
	if filter != "" {
		filter = "WHERE " + filter
	}

	filter += " LIMIT ?, ? "
	params = append(params, offset, perPage)

	// TODO : better pagination without offset and limit
	_, err := m.GetRDbMap().Select(
		&res,
		fmt.Sprintf("SELECT %s FROM %s %s", getSelectFields(BankSnapTableFull, ""), BankSnapTableFull, filter),
		params...,
	)
	assert.Nil(err)

	return res
}

// ListBankSnapsWithPagination try to list all BankSnaps with pagination
func (m *Manager) ListBankSnapsWithPagination(offset, perPage int) []BankSnap {
	return m.ListBankSnapsWithPaginationFilter(offset, perPage, "")
}

// FindBankSnapByID return the BankSnap base on its id
func (m *Manager) FindBankSnapByID(id int64) (*BankSnap, error) {
	var res BankSnap
	err := m.GetRDbMap().SelectOne(
		&res,
		fmt.Sprintf("SELECT %s FROM %s WHERE id=?", getSelectFields(BankSnapTableFull, ""), BankSnapTableFull),
		id,
	)

	if err != nil {
		return nil, err
	}

	return &res, nil
}
