// Code generated build with models DO NOT EDIT.

package orm

import (
	"time"

	"github.com/clickyab/services/initializer"
)

// AUTO GENERATED CODE. DO NOT EDIT!

// CreateCreativeDetail try to save a new CreativeDetail in database
func (m *Manager) CreateCreativeDetail(cd *CreativeDetail) error {
	now := time.Now()
	cd.CreatedAt = now
	cd.UpdatedAt = now
	func(in interface{}) {
		if ii, ok := in.(initializer.Simple); ok {
			ii.Initialize()
		}
	}(cd)

	return m.GetWDbMap().Insert(cd)
}

// UpdateCreativeDetail try to update CreativeDetail in database
func (m *Manager) UpdateCreativeDetail(cd *CreativeDetail) error {

	cd.UpdatedAt = time.Now()

	func(in interface{}) {
		if ii, ok := in.(initializer.Simple); ok {
			ii.Initialize()
		}
	}(cd)

	_, err := m.GetWDbMap().Update(cd)
	return err
}
