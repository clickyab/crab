// Code generated build with models DO NOT EDIT.

package orm

import (
	"fmt"
	"time"

	"github.com/clickyab/services/initializer"
)

// AUTO GENERATED CODE. DO NOT EDIT!

// CreateCampaignReportReceivers try to save a new CampaignReportReceivers in database
func (m *Manager) CreateCampaignReportReceivers(crr *CampaignReportReceivers) error {
	now := time.Now()
	crr.CreatedAt = now

	func(in interface{}) {
		if ii, ok := in.(initializer.Simple); ok {
			ii.Initialize()
		}
	}(crr)

	return m.GetWDbMap().Insert(crr)
}

// UpdateCampaignReportReceivers try to update CampaignReportReceivers in database
func (m *Manager) UpdateCampaignReportReceivers(crr *CampaignReportReceivers) error {

	func(in interface{}) {
		if ii, ok := in.(initializer.Simple); ok {
			ii.Initialize()
		}
	}(crr)

	_, err := m.GetWDbMap().Update(crr)
	return err
}

// FindCampaignReportReceiversByCampaignID return the CampaignReportReceivers base on its campaign_id
func (m *Manager) FindCampaignReportReceiversByCampaignID(ci int64) (*CampaignReportReceivers, error) {
	var res CampaignReportReceivers
	err := m.GetRDbMap().SelectOne(
		&res,
		fmt.Sprintf("SELECT %s FROM %s WHERE campaign_id=?", getSelectFields(CampaignReportReceiversTableFull, ""), CampaignReportReceiversTableFull),
		ci,
	)

	if err != nil {
		return nil, err
	}

	return &res, nil
}

// FindCampaignReportReceiversByUserID return the CampaignReportReceivers base on its user_id
func (m *Manager) FindCampaignReportReceiversByUserID(ui int64) (*CampaignReportReceivers, error) {
	var res CampaignReportReceivers
	err := m.GetRDbMap().SelectOne(
		&res,
		fmt.Sprintf("SELECT %s FROM %s WHERE user_id=?", getSelectFields(CampaignReportReceiversTableFull, ""), CampaignReportReceiversTableFull),
		ui,
	)

	if err != nil {
		return nil, err
	}

	return &res, nil
}
