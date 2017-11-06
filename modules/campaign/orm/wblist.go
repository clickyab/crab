package orm

import (
	"clickyab.com/crab/modules/inventory/orm"
	"github.com/clickyab/services/mysql"
)

// UpdateCampaignWhiteBlackList update white/black list
func (m *Manager) UpdateCampaignWhiteBlackList(w int64, ca *Campaign) error {

	l, err := orm.NewOrmManager().FindWhiteBlackListByID(w)
	if err != nil {
		return err
	}

	ca.WhiteBlackID = mysql.NullInt64{
		Valid: l.ID > 0,
		Int64: l.ID,
	}

	ca.WhiteBlackType = func() BlackWhiteTyp {
		if l.Kind {
			return WhiteTyp
		}
		return BlackTyp
	}()

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
	ca.WhiteBlackType = ClickyabTyp
	ca.WhiteBlackValue = []string{}
	err := m.UpdateCampaign(ca)
	if err != nil {
		return err
	}
	m.attachSchedule(ca)
	m.attachAttribute(ca)
	return nil
}
