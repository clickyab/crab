// Code generated build with models DO NOT EDIT.

package orm

import (
	"time"

	"github.com/clickyab/services/initializer"
)

// AUTO GENERATED CODE. DO NOT EDIT!

// CreateCampaignDetail try to save a new CampaignDetail in database
func (m *Manager) CreateCampaignDetail(cd *CampaignDetail) error {
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

// UpdateCampaignDetail try to update CampaignDetail in database
func (m *Manager) UpdateCampaignDetail(cd *CampaignDetail) error {

	cd.UpdatedAt = time.Now()

	func(in interface{}) {
		if ii, ok := in.(initializer.Simple); ok {
			ii.Initialize()
		}
	}(cd)

	_, err := m.GetWDbMap().Update(cd)
	return err
}
