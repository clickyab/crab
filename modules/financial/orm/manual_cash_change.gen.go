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

// CreateManualCashChange try to save a new ManualCashChange in database
func (m *Manager) CreateManualCashChange(mcc *ManualCashChange) error {
	now := time.Now()
	mcc.CreatedAt = now
	mcc.UpdatedAt = now
	func(in interface{}) {
		if ii, ok := in.(initializer.Simple); ok {
			ii.Initialize()
		}
	}(mcc)

	return m.GetWDbMap().Insert(mcc)
}

// UpdateManualCashChange try to update ManualCashChange in database
func (m *Manager) UpdateManualCashChange(mcc *ManualCashChange) error {
	now := time.Now()
	mcc.UpdatedAt = now
	func(in interface{}) {
		if ii, ok := in.(initializer.Simple); ok {
			ii.Initialize()
		}
	}(mcc)

	_, err := m.GetWDbMap().Update(mcc)
	return err
}

// ListManualCashChangesWithFilter try to list all ManualCashChanges without pagination
func (m *Manager) ListManualCashChangesWithFilter(filter string, params ...interface{}) []ManualCashChange {
	filter = strings.Trim(filter, "\n\t ")
	if filter != "" {
		filter = "WHERE " + filter
	}
	var res []ManualCashChange
	_, err := m.GetRDbMap().Select(
		&res,
		fmt.Sprintf("SELECT %s FROM %s %s", getSelectFields(ManualCashChangeTableFull, ""), ManualCashChangeTableFull, filter),
		params...,
	)
	assert.Nil(err)

	return res
}

// ListManualCashChanges try to list all ManualCashChanges without pagination
func (m *Manager) ListManualCashChanges() []ManualCashChange {
	return m.ListManualCashChangesWithFilter("")
}

// CountManualCashChangesWithFilter count entity in ManualCashChanges table with valid where filter
func (m *Manager) CountManualCashChangesWithFilter(filter string, params ...interface{}) int64 {
	filter = strings.Trim(filter, "\n\t ")
	if filter != "" {
		filter = "WHERE " + filter
	}
	cnt, err := m.GetRDbMap().SelectInt(
		fmt.Sprintf("SELECT COUNT(*) FROM %s %s", ManualCashChangeTableFull, filter),
		params...,
	)
	assert.Nil(err)

	return cnt
}

// CountManualCashChanges count entity in ManualCashChanges table
func (m *Manager) CountManualCashChanges() int64 {
	return m.CountManualCashChangesWithFilter("")
}

// ListManualCashChangesWithPaginationFilter try to list all ManualCashChanges with pagination and filter
func (m *Manager) ListManualCashChangesWithPaginationFilter(
	offset, perPage int, filter string, params ...interface{}) []ManualCashChange {
	var res []ManualCashChange
	filter = strings.Trim(filter, "\n\t ")
	if filter != "" {
		filter = "WHERE " + filter
	}

	filter += " LIMIT ?, ? "
	params = append(params, offset, perPage)

	// TODO : better pagination without offset and limit
	_, err := m.GetRDbMap().Select(
		&res,
		fmt.Sprintf("SELECT %s FROM %s %s", getSelectFields(ManualCashChangeTableFull, ""), ManualCashChangeTableFull, filter),
		params...,
	)
	assert.Nil(err)

	return res
}

// ListManualCashChangesWithPagination try to list all ManualCashChanges with pagination
func (m *Manager) ListManualCashChangesWithPagination(offset, perPage int) []ManualCashChange {
	return m.ListManualCashChangesWithPaginationFilter(offset, perPage, "")
}

// FindManualCashChangeByID return the ManualCashChange base on its id
func (m *Manager) FindManualCashChangeByID(id int64) (*ManualCashChange, error) {
	var res ManualCashChange
	err := m.GetRDbMap().SelectOne(
		&res,
		fmt.Sprintf("SELECT %s FROM %s WHERE id=?", getSelectFields(ManualCashChangeTableFull, ""), ManualCashChangeTableFull),
		id,
	)

	if err != nil {
		return nil, err
	}

	return &res, nil
}
