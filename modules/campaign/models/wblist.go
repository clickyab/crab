package models

import (
	"clickyab.com/crab/modules/inventory/models"
	"github.com/clickyab/services/mysql"
)

// UpdateCampaignWhiteBlackList update white/black list
func (m *Manager) UpdateCampaignWhiteBlackList(w int64, ca *Campaign) error {

	l, err := models.NewModelsManager().FindPresetByID(w)
	if err != nil {
		return err
	}

	ca.WhiteBlackID = mysql.NullInt64{
		Valid: l.ID > 0,
		Int64: l.ID,
	}
	ca.WhiteBlackType = mysql.NullBool{
		Valid: l.ID > 0,
		Bool:  l.Kind,
	}
	ca.WhiteBlackValue = l.Domains
	err = m.UpdateCampaign(ca)
	if err != nil {
		return err
	}
	m.attachSchedule(ca)
	m.attachAttribute(ca)
	return nil
}

// DeleteCampaignWhiteBlackList delete white/black list
func (m *Manager) DeleteCampaignWhiteBlackList(ca *Campaign) error {

	ca.WhiteBlackID = mysql.NullInt64{
		Valid: false,
	}
	ca.WhiteBlackType = mysql.NullBool{
		Valid: false,
	}
	ca.WhiteBlackValue = []string{}
	err := m.UpdateCampaign(ca)
	if err != nil {
		return err
	}
	m.attachSchedule(ca)
	m.attachAttribute(ca)
	return nil
}
