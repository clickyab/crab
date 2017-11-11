package orm

import (
	"clickyab.com/crab/modules/inventory/orm"
	"github.com/clickyab/services/mysql"
)

// UpdateCampaignWhiteBlackList update white/black list
func (m *Manager) UpdateCampaignWhiteBlackList(w int64, exchange *bool, white *bool, ca *Campaign) error {
	if w != 0 {
		l, err := orm.NewOrmManager().FindWhiteBlackListByID(w)
		if err != nil {
			return err
		}
		ca.WhiteBlackID = mysql.NullInt64{
			Valid: w > 0,
			Int64: w,
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
