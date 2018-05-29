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

// CreateManufacturer try to save a new Manufacturer in database
func (m *Manager) CreateManufacturer(mm *Manufacturer) error {
	now := time.Now()
	mm.CreatedAt = now
	mm.UpdatedAt = now
	func(in interface{}) {
		if ii, ok := in.(initializer.Simple); ok {
			ii.Initialize()
		}
	}(mm)

	return m.GetWDbMap().Insert(mm)
}

// UpdateManufacturer try to update Manufacturer in database
func (m *Manager) UpdateManufacturer(mm *Manufacturer) error {

	mm.UpdatedAt = time.Now()

	func(in interface{}) {
		if ii, ok := in.(initializer.Simple); ok {
			ii.Initialize()
		}
	}(mm)

	_, err := m.GetWDbMap().Update(mm)
	return err
}

// ListManufacturersWithFilter try to list all Manufacturers without pagination
func (m *Manager) ListManufacturersWithFilter(filter string, params ...interface{}) []Manufacturer {
	filter = strings.Trim(filter, "\n\t ")
	if filter != "" {
		filter = "WHERE " + filter
	}
	var res []Manufacturer
	_, err := m.GetRDbMap().Select(
		&res,
		fmt.Sprintf("SELECT %s FROM %s %s", GetSelectFields(ManufacturerTableFull, ""), ManufacturerTableFull, filter),
		params...,
	)
	assert.Nil(err)

	return res
}

// ListManufacturers try to list all Manufacturers without pagination
func (m *Manager) ListManufacturers() []Manufacturer {
	return m.ListManufacturersWithFilter("")
}

// CountManufacturersWithFilter count entity in Manufacturers table with valid where filter
func (m *Manager) CountManufacturersWithFilter(filter string, params ...interface{}) int64 {
	filter = strings.Trim(filter, "\n\t ")
	if filter != "" {
		filter = "WHERE " + filter
	}
	cnt, err := m.GetRDbMap().SelectInt(
		fmt.Sprintf("SELECT COUNT(*) FROM %s %s", ManufacturerTableFull, filter),
		params...,
	)
	assert.Nil(err)

	return cnt
}

// CountManufacturers count entity in Manufacturers table
func (m *Manager) CountManufacturers() int64 {
	return m.CountManufacturersWithFilter("")
}

// ListManufacturersWithPaginationFilter try to list all Manufacturers with pagination and filter
func (m *Manager) ListManufacturersWithPaginationFilter(
	offset, perPage int, filter string, params ...interface{}) []Manufacturer {
	var res []Manufacturer
	filter = strings.Trim(filter, "\n\t ")
	if filter != "" {
		filter = "WHERE " + filter
	}

	filter += " LIMIT ?, ? "
	params = append(params, offset, perPage)

	// TODO : better pagination without offset and limit
	_, err := m.GetRDbMap().Select(
		&res,
		fmt.Sprintf("SELECT %s FROM %s %s", GetSelectFields(ManufacturerTableFull, ""), ManufacturerTableFull, filter),
		params...,
	)
	assert.Nil(err)

	return res
}

// ListManufacturersWithPagination try to list all Manufacturers with pagination
func (m *Manager) ListManufacturersWithPagination(offset, perPage int) []Manufacturer {
	return m.ListManufacturersWithPaginationFilter(offset, perPage, "")
}

// FindManufacturerByName return the Manufacturer base on its name
func (m *Manager) FindManufacturerByName(n string) (*Manufacturer, error) {
	var res Manufacturer
	err := m.GetRDbMap().SelectOne(
		&res,
		fmt.Sprintf("SELECT %s FROM %s WHERE name=?", GetSelectFields(ManufacturerTableFull, ""), ManufacturerTableFull),
		n,
	)

	if err != nil {
		return nil, err
	}

	return &res, nil
}
