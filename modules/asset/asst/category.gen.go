// Code generated build with models DO NOT EDIT.

package asst

import (
	"fmt"

	"github.com/clickyab/services/initializer"
)

// AUTO GENERATED CODE. DO NOT EDIT!

// CreateCategory try to save a new Category in database
func (m *Manager) CreateCategory(c *Category) error {

	func(in interface{}) {
		if ii, ok := in.(initializer.Simple); ok {
			ii.Initialize()
		}
	}(c)

	return m.GetWDbMap().Insert(c)
}

// UpdateCategory try to update Category in database
func (m *Manager) UpdateCategory(c *Category) error {

	func(in interface{}) {
		if ii, ok := in.(initializer.Simple); ok {
			ii.Initialize()
		}
	}(c)

	_, err := m.GetWDbMap().Update(c)
	return err
}

// FindCategoryByID return the Category base on its id
func (m *Manager) FindCategoryByID(id int64) (*Category, error) {
	var res Category
	err := m.GetRDbMap().SelectOne(
		&res,
		fmt.Sprintf("SELECT * FROM %s WHERE id=?", CategoryTableFull),
		id,
	)

	if err != nil {
		return nil, err
	}

	return &res, nil
}

// FindCategoryByName return the Category base on its name
func (m *Manager) FindCategoryByName(n string) (*Category, error) {
	var res Category
	err := m.GetRDbMap().SelectOne(
		&res,
		fmt.Sprintf("SELECT * FROM %s WHERE name=?", CategoryTableFull),
		n,
	)

	if err != nil {
		return nil, err
	}

	return &res, nil
}
