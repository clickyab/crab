// Code generated build with models DO NOT EDIT.

package aaa

import (
	"time"

	"github.com/clickyab/services/initializer"
)

// AUTO GENERATED CODE. DO NOT EDIT!

// CreateAdvisor try to save a new Advisor in database
func (m *Manager) CreateAdvisor(a *Advisor) error {
	now := time.Now()
	a.CreatedAt = now

	func(in interface{}) {
		if ii, ok := in.(initializer.Simple); ok {
			ii.Initialize()
		}
	}(a)

	return m.GetWDbMap().Insert(a)
}

// UpdateAdvisor try to update Advisor in database
func (m *Manager) UpdateAdvisor(a *Advisor) error {

	func(in interface{}) {
		if ii, ok := in.(initializer.Simple); ok {
			ii.Initialize()
		}
	}(a)

	_, err := m.GetWDbMap().Update(a)
	return err
}
