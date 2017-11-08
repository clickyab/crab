package orm

import (
	"clickyab.com/crab/modules/inventory/orm"
	"github.com/clickyab/services/mysql"
)

// UpdateCampaignWhiteBlackList update white/black list
func (m *Manager) UpdateCampaignWhiteBlackList(w mysql.NullInt64, exchange *bool, white *bool, ca *Campaign) error {
	if w.Int64 != 0 {
		l, err := orm.NewOrmManager().FindWhiteBlackListByID(w.Int64)
		if err != nil {
			return err
		}
		ca.WhiteBlackID = mysql.NullInt64{
			Valid: l.ID > 0,
			Int64: l.ID,
		}
		ca.WhiteBlackValue = l.Domains
	}

	ca.Exchange = *exchange
	ca.WhiteBlackType = mysql.NullBool{
		Valid: true,
		Bool:  *white,
	}
	err := m.UpdateCampaign(ca)
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
	ca.WhiteBlackValue = mysql.StringMapJSONArray(make(map[string][]string))
	err := m.UpdateCampaign(ca)
	if err != nil {
		return err
	}
	m.attachSchedule(ca)
	m.attachAttribute(ca)
	return nil
}
