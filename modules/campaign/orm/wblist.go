package orm

import (
	"clickyab.com/crab/modules/inventory/orm"
	"github.com/clickyab/services/mysql"
)

// UpdateCampaignWhiteBlackList update white/black list
func (m *Manager) UpdateCampaignWhiteBlackList(w int64, exchange ExchangeType, white *bool, ca *Campaign) error {
	var domains mysql.StringMapJSONArray
	if w != 0 {
		l, err := orm.NewOrmManager().FindWhiteBlackListByID(w)
		if err != nil {
			return err
		}
		domains = l.Domains
	}
	ca.InventoryDomains = domains
	ca.InventoryID = mysql.NullInt64{
		Valid: w > 0,
		Int64: w,
	}
	ca.Exchange = exchange
	ca.InventoryType = NullInventoryState{
		Valid:          true,
		InventoryState: InventoryState("white_list"),
	}
	err := m.UpdateCampaign(ca)
	if err != nil {
		return err
	}

	err = m.attachSchedule(ca)
	if err != nil {
		return err
	}

	err = m.attachAttribute(ca)
	return err
}

// DeleteCampaignWhiteBlackList delete white/black list
func (m *Manager) DeleteCampaignWhiteBlackList(ca *Campaign) error {

	ca.InventoryID = mysql.NullInt64{
		Valid: false,
	}
	ca.InventoryType = NullInventoryState{
		Valid: false,
	}
	ca.InventoryDomains = mysql.StringMapJSONArray(make(map[string][]string))
	err := m.UpdateCampaign(ca)
	if err != nil {
		return err
	}

	err = m.attachSchedule(ca)
	if err != nil {
		return err
	}

	err = m.attachAttribute(ca)
	return err
}
