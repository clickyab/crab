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

// CreateDomainUser try to save a new DomainUser in database
func (m *Manager) CreateDomainUser(du *DomainUser) error {

	func(in interface{}) {
		if ii, ok := in.(initializer.Simple); ok {
			ii.Initialize()
		}
	}(du)

	return m.GetWDbMap().Insert(du)
}

// UpdateDomainUser try to update DomainUser in database
func (m *Manager) UpdateDomainUser(du *DomainUser) error {

	func(in interface{}) {
		if ii, ok := in.(initializer.Simple); ok {
			ii.Initialize()
		}
	}(du)

	_, err := m.GetWDbMap().Update(du)
	return err
}

// CreateDomain try to save a new Domain in database
func (m *Manager) CreateDomain(d *Domain) error {
	now := time.Now()
	d.CreatedAt = now
	d.UpdatedAt = now
	func(in interface{}) {
		if ii, ok := in.(initializer.Simple); ok {
			ii.Initialize()
		}
	}(d)

	return m.GetWDbMap().Insert(d)
}

// UpdateDomain try to update Domain in database
func (m *Manager) UpdateDomain(d *Domain) error {
	now := time.Now()
	d.UpdatedAt = now
	func(in interface{}) {
		if ii, ok := in.(initializer.Simple); ok {
			ii.Initialize()
		}
	}(d)

	_, err := m.GetWDbMap().Update(d)
	return err
}

// ListDomainsWithFilter try to list all Domains without pagination
func (m *Manager) ListDomainsWithFilter(filter string, params ...interface{}) []Domain {
	filter = strings.Trim(filter, "\n\t ")
	if filter != "" {
		filter = "WHERE " + filter
	}
	var res []Domain
	_, err := m.GetRDbMap().Select(
		&res,
		fmt.Sprintf("SELECT * FROM %s %s", DomainTableFull, filter),
		params...,
	)
	assert.Nil(err)

	return res
}

// ListDomains try to list all Domains without pagination
func (m *Manager) ListDomains() []Domain {
	return m.ListDomainsWithFilter("")
}

// CountDomainsWithFilter count entity in Domains table with valid where filter
func (m *Manager) CountDomainsWithFilter(filter string, params ...interface{}) int64 {
	filter = strings.Trim(filter, "\n\t ")
	if filter != "" {
		filter = "WHERE " + filter
	}
	cnt, err := m.GetRDbMap().SelectInt(
		fmt.Sprintf("SELECT COUNT(*) FROM %s %s", DomainTableFull, filter),
		params...,
	)
	assert.Nil(err)

	return cnt
}

// CountDomains count entity in Domains table
func (m *Manager) CountDomains() int64 {
	return m.CountDomainsWithFilter("")
}

// ListDomainsWithPaginationFilter try to list all Domains with pagination and filter
func (m *Manager) ListDomainsWithPaginationFilter(
	offset, perPage int, filter string, params ...interface{}) []Domain {
	var res []Domain
	filter = strings.Trim(filter, "\n\t ")
	if filter != "" {
		filter = "WHERE " + filter
	}

	filter += " LIMIT ?, ? "
	params = append(params, offset, perPage)

	// TODO : better pagination without offset and limit
	_, err := m.GetRDbMap().Select(
		&res,
		fmt.Sprintf("SELECT * FROM %s %s", DomainTableFull, filter),
		params...,
	)
	assert.Nil(err)

	return res
}

// ListDomainsWithPagination try to list all Domains with pagination
func (m *Manager) ListDomainsWithPagination(offset, perPage int) []Domain {
	return m.ListDomainsWithPaginationFilter(offset, perPage, "")
}

// FindDomainByID return the Domain base on its id
func (m *Manager) FindDomainByID(id int64) (*Domain, error) {
	var res Domain
	err := m.GetRDbMap().SelectOne(
		&res,
		fmt.Sprintf("SELECT * FROM %s WHERE id=?", DomainTableFull),
		id,
	)

	if err != nil {
		return nil, err
	}

	return &res, nil
}

// FindDomainByName return the Domain base on its name
func (m *Manager) FindDomainByName(n string) (*Domain, error) {
	var res Domain
	err := m.GetRDbMap().SelectOne(
		&res,
		fmt.Sprintf("SELECT * FROM %s WHERE name=?", DomainTableFull),
		n,
	)

	if err != nil {
		return nil, err
	}

	return &res, nil
}
