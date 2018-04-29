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

// CreateGateway try to save a new Gateway in database
func (m *Manager) CreateGateway(g *Gateway) error {
	now := time.Now()
	g.CreatedAt = now
	g.UpdatedAt = now
	func(in interface{}) {
		if ii, ok := in.(initializer.Simple); ok {
			ii.Initialize()
		}
	}(g)

	return m.GetWDbMap().Insert(g)
}

// UpdateGateway try to update Gateway in database
func (m *Manager) UpdateGateway(g *Gateway) error {

	g.UpdatedAt = time.Now()

	func(in interface{}) {
		if ii, ok := in.(initializer.Simple); ok {
			ii.Initialize()
		}
	}(g)

	_, err := m.GetWDbMap().Update(g)
	return err
}

// ListGatewaysWithFilter try to list all Gateways without pagination
func (m *Manager) ListGatewaysWithFilter(filter string, params ...interface{}) []Gateway {
	filter = strings.Trim(filter, "\n\t ")
	if filter != "" {
		filter = "WHERE " + filter
	}
	var res []Gateway
	_, err := m.GetRDbMap().Select(
		&res,
		fmt.Sprintf("SELECT %s FROM %s %s", getSelectFields(GatewayTableFull, ""), GatewayTableFull, filter),
		params...,
	)
	assert.Nil(err)

	return res
}

// ListGateways try to list all Gateways without pagination
func (m *Manager) ListGateways() []Gateway {
	return m.ListGatewaysWithFilter("")
}

// CountGatewaysWithFilter count entity in Gateways table with valid where filter
func (m *Manager) CountGatewaysWithFilter(filter string, params ...interface{}) int64 {
	filter = strings.Trim(filter, "\n\t ")
	if filter != "" {
		filter = "WHERE " + filter
	}
	cnt, err := m.GetRDbMap().SelectInt(
		fmt.Sprintf("SELECT COUNT(*) FROM %s %s", GatewayTableFull, filter),
		params...,
	)
	assert.Nil(err)

	return cnt
}

// CountGateways count entity in Gateways table
func (m *Manager) CountGateways() int64 {
	return m.CountGatewaysWithFilter("")
}

// ListGatewaysWithPaginationFilter try to list all Gateways with pagination and filter
func (m *Manager) ListGatewaysWithPaginationFilter(
	offset, perPage int, filter string, params ...interface{}) []Gateway {
	var res []Gateway
	filter = strings.Trim(filter, "\n\t ")
	if filter != "" {
		filter = "WHERE " + filter
	}

	filter += " LIMIT ?, ? "
	params = append(params, offset, perPage)

	// TODO : better pagination without offset and limit
	_, err := m.GetRDbMap().Select(
		&res,
		fmt.Sprintf("SELECT %s FROM %s %s", getSelectFields(GatewayTableFull, ""), GatewayTableFull, filter),
		params...,
	)
	assert.Nil(err)

	return res
}

// ListGatewaysWithPagination try to list all Gateways with pagination
func (m *Manager) ListGatewaysWithPagination(offset, perPage int) []Gateway {
	return m.ListGatewaysWithPaginationFilter(offset, perPage, "")
}

// FindGatewayByID return the Gateway base on its id
func (m *Manager) FindGatewayByID(id int64) (*Gateway, error) {
	var res Gateway
	err := m.GetRDbMap().SelectOne(
		&res,
		fmt.Sprintf("SELECT %s FROM %s WHERE id=?", getSelectFields(GatewayTableFull, ""), GatewayTableFull),
		id,
	)

	if err != nil {
		return nil, err
	}

	return &res, nil
}
