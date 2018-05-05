// Code generated build with models DO NOT EDIT.

package orm

import (
	"fmt"
	"strings"

	"github.com/clickyab/services/assert"
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

// ListCategoriesWithFilter try to list all Categories without pagination
func (m *Manager) ListCategoriesWithFilter(filter string, params ...interface{}) []Category {
	filter = strings.Trim(filter, "\n\t ")
	if filter != "" {
		filter = "WHERE " + filter
	}
	var res []Category
	_, err := m.GetRDbMap().Select(
		&res,
		fmt.Sprintf("SELECT %s FROM %s %s", getSelectFields(CategoryTableFull, ""), CategoryTableFull, filter),
		params...,
	)
	assert.Nil(err)

	return res
}

// ListCategories try to list all Categories without pagination
func (m *Manager) ListCategories() []Category {
	return m.ListCategoriesWithFilter("")
}

// CountCategoriesWithFilter count entity in Categories table with valid where filter
func (m *Manager) CountCategoriesWithFilter(filter string, params ...interface{}) int64 {
	filter = strings.Trim(filter, "\n\t ")
	if filter != "" {
		filter = "WHERE " + filter
	}
	cnt, err := m.GetRDbMap().SelectInt(
		fmt.Sprintf("SELECT COUNT(*) FROM %s %s", CategoryTableFull, filter),
		params...,
	)
	assert.Nil(err)

	return cnt
}

// CountCategories count entity in Categories table
func (m *Manager) CountCategories() int64 {
	return m.CountCategoriesWithFilter("")
}

// ListCategoriesWithPaginationFilter try to list all Categories with pagination and filter
func (m *Manager) ListCategoriesWithPaginationFilter(
	offset, perPage int, filter string, params ...interface{}) []Category {
	var res []Category
	filter = strings.Trim(filter, "\n\t ")
	if filter != "" {
		filter = "WHERE " + filter
	}

	filter += " LIMIT ?, ? "
	params = append(params, offset, perPage)

	// TODO : better pagination without offset and limit
	_, err := m.GetRDbMap().Select(
		&res,
		fmt.Sprintf("SELECT %s FROM %s %s", getSelectFields(CategoryTableFull, ""), CategoryTableFull, filter),
		params...,
	)
	assert.Nil(err)

	return res
}

// ListCategoriesWithPagination try to list all Categories with pagination
func (m *Manager) ListCategoriesWithPagination(offset, perPage int) []Category {
	return m.ListCategoriesWithPaginationFilter(offset, perPage, "")
}

// FindCategoryByName return the Category base on its name
func (m *Manager) FindCategoryByName(n string) (*Category, error) {
	var res Category
	err := m.GetRDbMap().SelectOne(
		&res,
		fmt.Sprintf("SELECT %s FROM %s WHERE name=?", getSelectFields(CategoryTableFull, ""), CategoryTableFull),
		n,
	)

	if err != nil {
		return nil, err
	}

	return &res, nil
}

// CreateCategoryModel try to save a new CategoryModel in database
func (m *Manager) CreateCategoryModel(cm *CategoryModel) error {

	func(in interface{}) {
		if ii, ok := in.(initializer.Simple); ok {
			ii.Initialize()
		}
	}(cm)

	return m.GetWDbMap().Insert(cm)
}

// UpdateCategoryModel try to update CategoryModel in database
func (m *Manager) UpdateCategoryModel(cm *CategoryModel) error {

	func(in interface{}) {
		if ii, ok := in.(initializer.Simple); ok {
			ii.Initialize()
		}
	}(cm)

	_, err := m.GetWDbMap().Update(cm)
	return err
}

// ListCategoryModelsWithFilter try to list all CategoryModels without pagination
func (m *Manager) ListCategoryModelsWithFilter(filter string, params ...interface{}) []CategoryModel {
	filter = strings.Trim(filter, "\n\t ")
	if filter != "" {
		filter = "WHERE " + filter
	}
	var res []CategoryModel
	_, err := m.GetRDbMap().Select(
		&res,
		fmt.Sprintf("SELECT %s FROM %s %s", getSelectFields(CategoryModelTableFull, ""), CategoryModelTableFull, filter),
		params...,
	)
	assert.Nil(err)

	return res
}

// ListCategoryModels try to list all CategoryModels without pagination
func (m *Manager) ListCategoryModels() []CategoryModel {
	return m.ListCategoryModelsWithFilter("")
}

// CountCategoryModelsWithFilter count entity in CategoryModels table with valid where filter
func (m *Manager) CountCategoryModelsWithFilter(filter string, params ...interface{}) int64 {
	filter = strings.Trim(filter, "\n\t ")
	if filter != "" {
		filter = "WHERE " + filter
	}
	cnt, err := m.GetRDbMap().SelectInt(
		fmt.Sprintf("SELECT COUNT(*) FROM %s %s", CategoryModelTableFull, filter),
		params...,
	)
	assert.Nil(err)

	return cnt
}

// CountCategoryModels count entity in CategoryModels table
func (m *Manager) CountCategoryModels() int64 {
	return m.CountCategoryModelsWithFilter("")
}

// ListCategoryModelsWithPaginationFilter try to list all CategoryModels with pagination and filter
func (m *Manager) ListCategoryModelsWithPaginationFilter(
	offset, perPage int, filter string, params ...interface{}) []CategoryModel {
	var res []CategoryModel
	filter = strings.Trim(filter, "\n\t ")
	if filter != "" {
		filter = "WHERE " + filter
	}

	filter += " LIMIT ?, ? "
	params = append(params, offset, perPage)

	// TODO : better pagination without offset and limit
	_, err := m.GetRDbMap().Select(
		&res,
		fmt.Sprintf("SELECT %s FROM %s %s", getSelectFields(CategoryModelTableFull, ""), CategoryModelTableFull, filter),
		params...,
	)
	assert.Nil(err)

	return res
}

// ListCategoryModelsWithPagination try to list all CategoryModels with pagination
func (m *Manager) ListCategoryModelsWithPagination(offset, perPage int) []CategoryModel {
	return m.ListCategoryModelsWithPaginationFilter(offset, perPage, "")
}
