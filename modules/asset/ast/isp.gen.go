// Code generated build with models DO NOT EDIT.

package ast

import (
	"fmt"
	"time"

	"github.com/clickyab/services/initializer"
)

// AUTO GENERATED CODE. DO NOT EDIT!

// CreateISP try to save a new ISP in database
func (m *Manager) CreateISP(isp *ISP) error {
	now := time.Now()
	isp.CreatedAt = now
	isp.UpdatedAt = now
	func(in interface{}) {
		if ii, ok := in.(initializer.Simple); ok {
			ii.Initialize()
		}
	}(isp)

	return m.GetWDbMap().Insert(isp)
}

// UpdateISP try to update ISP in database
func (m *Manager) UpdateISP(isp *ISP) error {
	now := time.Now()
	isp.UpdatedAt = now
	func(in interface{}) {
		if ii, ok := in.(initializer.Simple); ok {
			ii.Initialize()
		}
	}(isp)

	_, err := m.GetWDbMap().Update(isp)
	return err
}

// FindISPByID return the ISP base on its id
func (m *Manager) FindISPByID(id int64) (*ISP, error) {
	var res ISP
	err := m.GetRDbMap().SelectOne(
		&res,
		fmt.Sprintf("SELECT * FROM %s WHERE id=?", ISPTableFull),
		id,
	)

	if err != nil {
		return nil, err
	}

	return &res, nil
}

// FindISPByName return the ISP base on its name
func (m *Manager) FindISPByName(n string) (*ISP, error) {
	var res ISP
	err := m.GetRDbMap().SelectOne(
		&res,
		fmt.Sprintf("SELECT * FROM %s WHERE name=?", ISPTableFull),
		n,
	)

	if err != nil {
		return nil, err
	}

	return &res, nil
}
