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

// CreatePreset try to save a new Preset in database
func (m *Manager) CreatePreset(p *Preset) error {
	now := time.Now()
	p.CreatedAt = now
	p.UpdatedAt = now
	func(in interface{}) {
		if ii, ok := in.(initializer.Simple); ok {
			ii.Initialize()
		}
	}(p)

	return m.GetWDbMap().Insert(p)
}

// UpdatePreset try to update Preset in database
func (m *Manager) UpdatePreset(p *Preset) error {
	now := time.Now()
	p.UpdatedAt = now
	func(in interface{}) {
		if ii, ok := in.(initializer.Simple); ok {
			ii.Initialize()
		}
	}(p)

	_, err := m.GetWDbMap().Update(p)
	return err
}

// ListPresetsWithFilter try to list all Presets without pagination
func (m *Manager) ListPresetsWithFilter(filter string, params ...interface{}) []Preset {
	filter = strings.Trim(filter, "\n\t ")
	if filter != "" {
		filter = "WHERE " + filter
	}
	var res []Preset
	_, err := m.GetRDbMap().Select(
		&res,
		fmt.Sprintf("SELECT * FROM %s %s", PresetTableFull, filter),
		params...,
	)
	assert.Nil(err)

	return res
}

// ListPresets try to list all Presets without pagination
func (m *Manager) ListPresets() []Preset {
	return m.ListPresetsWithFilter("")
}

// CountPresetsWithFilter count entity in Presets table with valid where filter
func (m *Manager) CountPresetsWithFilter(filter string, params ...interface{}) int64 {
	filter = strings.Trim(filter, "\n\t ")
	if filter != "" {
		filter = "WHERE " + filter
	}
	cnt, err := m.GetRDbMap().SelectInt(
		fmt.Sprintf("SELECT COUNT(*) FROM %s %s", PresetTableFull, filter),
		params...,
	)
	assert.Nil(err)

	return cnt
}

// CountPresets count entity in Presets table
func (m *Manager) CountPresets() int64 {
	return m.CountPresetsWithFilter("")
}

// ListPresetsWithPaginationFilter try to list all Presets with pagination and filter
func (m *Manager) ListPresetsWithPaginationFilter(
	offset, perPage int, filter string, params ...interface{}) []Preset {
	var res []Preset
	filter = strings.Trim(filter, "\n\t ")
	if filter != "" {
		filter = "WHERE " + filter
	}

	filter += " LIMIT ?, ? "
	params = append(params, offset, perPage)

	// TODO : better pagination without offset and limit
	_, err := m.GetRDbMap().Select(
		&res,
		fmt.Sprintf("SELECT * FROM %s %s", PresetTableFull, filter),
		params...,
	)
	assert.Nil(err)

	return res
}

// ListPresetsWithPagination try to list all Presets with pagination
func (m *Manager) ListPresetsWithPagination(offset, perPage int) []Preset {
	return m.ListPresetsWithPaginationFilter(offset, perPage, "")
}

// FindPresetByID return the Preset base on its id
func (m *Manager) FindPresetByID(id int64) (*Preset, error) {
	var res Preset
	err := m.GetRDbMap().SelectOne(
		&res,
		fmt.Sprintf("SELECT * FROM %s WHERE id=?", PresetTableFull),
		id,
	)

	if err != nil {
		return nil, err
	}

	return &res, nil
}
