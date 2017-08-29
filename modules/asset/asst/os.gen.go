// Code generated build with models DO NOT EDIT.

package asst

import (
	"fmt"
	"time"

	"github.com/clickyab/services/initializer"
)

// AUTO GENERATED CODE. DO NOT EDIT!

// CreateOS try to save a new OS in database
func (m *Manager) CreateOS(os *OS) error {
	now := time.Now()
	os.CreatedAt = now
	os.UpdatedAt = now
	func(in interface{}) {
		if ii, ok := in.(initializer.Simple); ok {
			ii.Initialize()
		}
	}(os)

	return m.GetWDbMap().Insert(os)
}

// UpdateOS try to update OS in database
func (m *Manager) UpdateOS(os *OS) error {
	now := time.Now()
	os.UpdatedAt = now
	func(in interface{}) {
		if ii, ok := in.(initializer.Simple); ok {
			ii.Initialize()
		}
	}(os)

	_, err := m.GetWDbMap().Update(os)
	return err
}

// FindOSByID return the OS base on its id
func (m *Manager) FindOSByID(id int64) (*OS, error) {
	var res OS
	err := m.GetRDbMap().SelectOne(
		&res,
		fmt.Sprintf("SELECT * FROM %s WHERE id=?", OSTableFull),
		id,
	)

	if err != nil {
		return nil, err
	}

	return &res, nil
}

// FindOSByName return the OS base on its name
func (m *Manager) FindOSByName(n string) (*OS, error) {
	var res OS
	err := m.GetRDbMap().SelectOne(
		&res,
		fmt.Sprintf("SELECT * FROM %s WHERE name=?", OSTableFull),
		n,
	)

	if err != nil {
		return nil, err
	}

	return &res, nil
}
