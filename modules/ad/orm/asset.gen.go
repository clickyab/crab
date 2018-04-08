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

// CreateAsset try to save a new Asset in database
func (m *Manager) CreateAsset(a *Asset) error {
	now := time.Now()
	a.CreatedAt = now
	a.UpdatedAt = now
	func(in interface{}) {
		if ii, ok := in.(initializer.Simple); ok {
			ii.Initialize()
		}
	}(a)

	return m.GetWDbMap().Insert(a)
}

// UpdateAsset try to update Asset in database
func (m *Manager) UpdateAsset(a *Asset) error {
	now := time.Now()
	a.UpdatedAt = now
	func(in interface{}) {
		if ii, ok := in.(initializer.Simple); ok {
			ii.Initialize()
		}
	}(a)

	_, err := m.GetWDbMap().Update(a)
	return err
}

// ListAssetsWithFilter try to list all Assets without pagination
func (m *Manager) ListAssetsWithFilter(filter string, params ...interface{}) []Asset {
	filter = strings.Trim(filter, "\n\t ")
	if filter != "" {
		filter = "WHERE " + filter
	}
	var res []Asset
	_, err := m.GetRDbMap().Select(
		&res,
		fmt.Sprintf("SELECT %s FROM %s %s", getSelectFields(AssetTableFull, ""), AssetTableFull, filter),
		params...,
	)
	assert.Nil(err)

	return res
}

// ListAssets try to list all Assets without pagination
func (m *Manager) ListAssets() []Asset {
	return m.ListAssetsWithFilter("")
}

// CountAssetsWithFilter count entity in Assets table with valid where filter
func (m *Manager) CountAssetsWithFilter(filter string, params ...interface{}) int64 {
	filter = strings.Trim(filter, "\n\t ")
	if filter != "" {
		filter = "WHERE " + filter
	}
	cnt, err := m.GetRDbMap().SelectInt(
		fmt.Sprintf("SELECT COUNT(*) FROM %s %s", AssetTableFull, filter),
		params...,
	)
	assert.Nil(err)

	return cnt
}

// CountAssets count entity in Assets table
func (m *Manager) CountAssets() int64 {
	return m.CountAssetsWithFilter("")
}

// ListAssetsWithPaginationFilter try to list all Assets with pagination and filter
func (m *Manager) ListAssetsWithPaginationFilter(
	offset, perPage int, filter string, params ...interface{}) []Asset {
	var res []Asset
	filter = strings.Trim(filter, "\n\t ")
	if filter != "" {
		filter = "WHERE " + filter
	}

	filter += " LIMIT ?, ? "
	params = append(params, offset, perPage)

	// TODO : better pagination without offset and limit
	_, err := m.GetRDbMap().Select(
		&res,
		fmt.Sprintf("SELECT %s FROM %s %s", getSelectFields(AssetTableFull, ""), AssetTableFull, filter),
		params...,
	)
	assert.Nil(err)

	return res
}

// ListAssetsWithPagination try to list all Assets with pagination
func (m *Manager) ListAssetsWithPagination(offset, perPage int) []Asset {
	return m.ListAssetsWithPaginationFilter(offset, perPage, "")
}

// FindAssetByID return the Asset base on its id
func (m *Manager) FindAssetByID(id int64) (*Asset, error) {
	var res Asset
	err := m.GetRDbMap().SelectOne(
		&res,
		fmt.Sprintf("SELECT %s FROM %s WHERE id=?", getSelectFields(AssetTableFull, ""), AssetTableFull),
		id,
	)

	if err != nil {
		return nil, err
	}

	return &res, nil
}

// FindAssetByCreativeID return the Asset base on its creative_id
func (m *Manager) FindAssetByCreativeID(ci int64) (*Asset, error) {
	var res Asset
	err := m.GetRDbMap().SelectOne(
		&res,
		fmt.Sprintf("SELECT %s FROM %s WHERE creative_id=?", getSelectFields(AssetTableFull, ""), AssetTableFull),
		ci,
	)

	if err != nil {
		return nil, err
	}

	return &res, nil
}
