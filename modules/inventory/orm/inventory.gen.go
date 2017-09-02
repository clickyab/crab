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

// CreateInventory try to save a new Inventory in database
func (m *Manager) CreateInventory(i *Inventory) error {
	now := time.Now()
	i.CreatedAt = now
	i.UpdatedAt = now
	func(in interface{}) {
		if ii, ok := in.(initializer.Simple); ok {
			ii.Initialize()
		}
	}(i)

	return m.GetWDbMap().Insert(i)
}

// UpdateInventory try to update Inventory in database
func (m *Manager) UpdateInventory(i *Inventory) error {
	now := time.Now()
	i.UpdatedAt = now
	func(in interface{}) {
		if ii, ok := in.(initializer.Simple); ok {
			ii.Initialize()
		}
	}(i)

	_, err := m.GetWDbMap().Update(i)
	return err
}

// ListInventoriesWithFilter try to list all Inventories without pagination
func (m *Manager) ListInventoriesWithFilter(filter string, params ...interface{}) []Inventory {
	filter = strings.Trim(filter, "\n\t ")
	if filter != "" {
		filter = "WHERE " + filter
	}
	var res []Inventory
	_, err := m.GetRDbMap().Select(
		&res,
		fmt.Sprintf("SELECT * FROM %s %s", InventoryTableFull, filter),
		params...,
	)
	assert.Nil(err)

	return res
}

// ListInventories try to list all Inventories without pagination
func (m *Manager) ListInventories() []Inventory {
	return m.ListInventoriesWithFilter("")
}

// CountInventoriesWithFilter count entity in Inventories table with valid where filter
func (m *Manager) CountInventoriesWithFilter(filter string, params ...interface{}) int64 {
	filter = strings.Trim(filter, "\n\t ")
	if filter != "" {
		filter = "WHERE " + filter
	}
	cnt, err := m.GetRDbMap().SelectInt(
		fmt.Sprintf("SELECT COUNT(*) FROM %s %s", InventoryTableFull, filter),
		params...,
	)
	assert.Nil(err)

	return cnt
}

// CountInventories count entity in Inventories table
func (m *Manager) CountInventories() int64 {
	return m.CountInventoriesWithFilter("")
}

// ListInventoriesWithPaginationFilter try to list all Inventories with pagination and filter
func (m *Manager) ListInventoriesWithPaginationFilter(
	offset, perPage int, filter string, params ...interface{}) []Inventory {
	var res []Inventory
	filter = strings.Trim(filter, "\n\t ")
	if filter != "" {
		filter = "WHERE " + filter
	}

	filter += " LIMIT ?, ? "
	params = append(params, offset, perPage)

	// TODO : better pagination without offset and limit
	_, err := m.GetRDbMap().Select(
		&res,
		fmt.Sprintf("SELECT * FROM %s %s", InventoryTableFull, filter),
		params...,
	)
	assert.Nil(err)

	return res
}

// ListInventoriesWithPagination try to list all Inventories with pagination
func (m *Manager) ListInventoriesWithPagination(offset, perPage int) []Inventory {
	return m.ListInventoriesWithPaginationFilter(offset, perPage, "")
}

// FindInventoryByID return the Inventory base on its id
func (m *Manager) FindInventoryByID(id int64) (*Inventory, error) {
	var res Inventory
	err := m.GetRDbMap().SelectOne(
		&res,
		fmt.Sprintf("SELECT * FROM %s WHERE id=?", InventoryTableFull),
		id,
	)

	if err != nil {
		return nil, err
	}

	return &res, nil
}

// FindInventoryByUserID return the Inventory base on its user_id
func (m *Manager) FindInventoryByUserID(ui int64) (*Inventory, error) {
	var res Inventory
	err := m.GetRDbMap().SelectOne(
		&res,
		fmt.Sprintf("SELECT * FROM %s WHERE user_id=?", InventoryTableFull),
		ui,
	)

	if err != nil {
		return nil, err
	}

	return &res, nil
}
