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

// CreateWhiteBlackList try to save a new WhiteBlackList in database
func (m *Manager) CreateWhiteBlackList(wbl *WhiteBlackList) error {
	now := time.Now()
	wbl.CreatedAt = now
	wbl.UpdatedAt = now
	func(in interface{}) {
		if ii, ok := in.(initializer.Simple); ok {
			ii.Initialize()
		}
	}(wbl)

	return m.GetWDbMap().Insert(wbl)
}

// UpdateWhiteBlackList try to update WhiteBlackList in database
func (m *Manager) UpdateWhiteBlackList(wbl *WhiteBlackList) error {
	now := time.Now()
	wbl.UpdatedAt = now
	func(in interface{}) {
		if ii, ok := in.(initializer.Simple); ok {
			ii.Initialize()
		}
	}(wbl)

	_, err := m.GetWDbMap().Update(wbl)
	return err
}

// ListWhiteBlackListsWithFilter try to list all WhiteBlackLists without pagination
func (m *Manager) ListWhiteBlackListsWithFilter(filter string, params ...interface{}) []WhiteBlackList {
	filter = strings.Trim(filter, "\n\t ")
	if filter != "" {
		filter = "WHERE " + filter
	}
	var res []WhiteBlackList
	_, err := m.GetRDbMap().Select(
		&res,
		fmt.Sprintf("SELECT * FROM %s %s", WhiteBlackListTableFull, filter),
		params...,
	)
	assert.Nil(err)

	return res
}

// ListWhiteBlackLists try to list all WhiteBlackLists without pagination
func (m *Manager) ListWhiteBlackLists() []WhiteBlackList {
	return m.ListWhiteBlackListsWithFilter("")
}

// CountWhiteBlackListsWithFilter count entity in WhiteBlackLists table with valid where filter
func (m *Manager) CountWhiteBlackListsWithFilter(filter string, params ...interface{}) int64 {
	filter = strings.Trim(filter, "\n\t ")
	if filter != "" {
		filter = "WHERE " + filter
	}
	cnt, err := m.GetRDbMap().SelectInt(
		fmt.Sprintf("SELECT COUNT(*) FROM %s %s", WhiteBlackListTableFull, filter),
		params...,
	)
	assert.Nil(err)

	return cnt
}

// CountWhiteBlackLists count entity in WhiteBlackLists table
func (m *Manager) CountWhiteBlackLists() int64 {
	return m.CountWhiteBlackListsWithFilter("")
}

// ListWhiteBlackListsWithPaginationFilter try to list all WhiteBlackLists with pagination and filter
func (m *Manager) ListWhiteBlackListsWithPaginationFilter(
	offset, perPage int, filter string, params ...interface{}) []WhiteBlackList {
	var res []WhiteBlackList
	filter = strings.Trim(filter, "\n\t ")
	if filter != "" {
		filter = "WHERE " + filter
	}

	filter += " LIMIT ?, ? "
	params = append(params, offset, perPage)

	// TODO : better pagination without offset and limit
	_, err := m.GetRDbMap().Select(
		&res,
		fmt.Sprintf("SELECT * FROM %s %s", WhiteBlackListTableFull, filter),
		params...,
	)
	assert.Nil(err)

	return res
}

// ListWhiteBlackListsWithPagination try to list all WhiteBlackLists with pagination
func (m *Manager) ListWhiteBlackListsWithPagination(offset, perPage int) []WhiteBlackList {
	return m.ListWhiteBlackListsWithPaginationFilter(offset, perPage, "")
}

// FindWhiteBlackListByID return the WhiteBlackList base on its id
func (m *Manager) FindWhiteBlackListByID(id int64) (*WhiteBlackList, error) {
	var res WhiteBlackList
	err := m.GetRDbMap().SelectOne(
		&res,
		fmt.Sprintf("SELECT * FROM %s WHERE id=?", WhiteBlackListTableFull),
		id,
	)

	if err != nil {
		return nil, err
	}

	return &res, nil
}
