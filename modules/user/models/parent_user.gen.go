// Code generated build with models DO NOT EDIT.

package models

import (
	"time"

	"github.com/clickyab/services/initializer"
)

// AUTO GENERATED CODE. DO NOT EDIT!

// CreateParentUser try to save a new ParentUser in database
func (m *Manager) CreateParentUser(pu *ParentUser) error {
	now := time.Now()
	pu.CreatedAt = now

	func(in interface{}) {
		if ii, ok := in.(initializer.Simple); ok {
			ii.Initialize()
		}
	}(pu)

	return m.GetWDbMap().Insert(pu)
}

// UpdateParentUser try to update ParentUser in database
func (m *Manager) UpdateParentUser(pu *ParentUser) error {

	func(in interface{}) {
		if ii, ok := in.(initializer.Simple); ok {
			ii.Initialize()
		}
	}(pu)

	_, err := m.GetWDbMap().Update(pu)
	return err
}
