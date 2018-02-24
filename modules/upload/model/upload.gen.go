// Code generated build with models DO NOT EDIT.

package model

import (
	"fmt"
	"strings"
	"time"

	"github.com/clickyab/services/assert"
	"github.com/clickyab/services/initializer"
)

// AUTO GENERATED CODE. DO NOT EDIT!

// CreateUpload try to save a new Upload in database
func (m *Manager) CreateUpload(u *Upload) error {
	now := time.Now()
	u.CreatedAt = now

	func(in interface{}) {
		if ii, ok := in.(initializer.Simple); ok {
			ii.Initialize()
		}
	}(u)

	return m.GetWDbMap().Insert(u)
}

// UpdateUpload try to update Upload in database
func (m *Manager) UpdateUpload(u *Upload) error {

	func(in interface{}) {
		if ii, ok := in.(initializer.Simple); ok {
			ii.Initialize()
		}
	}(u)

	_, err := m.GetWDbMap().Update(u)
	return err
}

// ListUploadsWithFilter try to list all Uploads without pagination
func (m *Manager) ListUploadsWithFilter(filter string, params ...interface{}) []Upload {
	filter = strings.Trim(filter, "\n\t ")
	if filter != "" {
		filter = "WHERE " + filter
	}
	var res []Upload
	_, err := m.GetRDbMap().Select(
		&res,
		fmt.Sprintf("SELECT %s FROM %s %s", getSelectFields(UploadTableFull, ""), UploadTableFull, filter),
		params...,
	)
	assert.Nil(err)

	return res
}

// ListUploads try to list all Uploads without pagination
func (m *Manager) ListUploads() []Upload {
	return m.ListUploadsWithFilter("")
}

// CountUploadsWithFilter count entity in Uploads table with valid where filter
func (m *Manager) CountUploadsWithFilter(filter string, params ...interface{}) int64 {
	filter = strings.Trim(filter, "\n\t ")
	if filter != "" {
		filter = "WHERE " + filter
	}
	cnt, err := m.GetRDbMap().SelectInt(
		fmt.Sprintf("SELECT COUNT(*) FROM %s %s", UploadTableFull, filter),
		params...,
	)
	assert.Nil(err)

	return cnt
}

// CountUploads count entity in Uploads table
func (m *Manager) CountUploads() int64 {
	return m.CountUploadsWithFilter("")
}

// ListUploadsWithPaginationFilter try to list all Uploads with pagination and filter
func (m *Manager) ListUploadsWithPaginationFilter(
	offset, perPage int, filter string, params ...interface{}) []Upload {
	var res []Upload
	filter = strings.Trim(filter, "\n\t ")
	if filter != "" {
		filter = "WHERE " + filter
	}

	filter += " LIMIT ?, ? "
	params = append(params, offset, perPage)

	// TODO : better pagination without offset and limit
	_, err := m.GetRDbMap().Select(
		&res,
		fmt.Sprintf("SELECT %s FROM %s %s", getSelectFields(UploadTableFull, ""), UploadTableFull, filter),
		params...,
	)
	assert.Nil(err)

	return res
}

// ListUploadsWithPagination try to list all Uploads with pagination
func (m *Manager) ListUploadsWithPagination(offset, perPage int) []Upload {
	return m.ListUploadsWithPaginationFilter(offset, perPage, "")
}

// FindUploadByID return the Upload base on its id
func (m *Manager) FindUploadByID(id string) (*Upload, error) {
	var res Upload
	err := m.GetRDbMap().SelectOne(
		&res,
		fmt.Sprintf("SELECT %s FROM %s WHERE id=?", getSelectFields(UploadTableFull, ""), UploadTableFull),
		id,
	)

	if err != nil {
		return nil, err
	}

	return &res, nil
}
