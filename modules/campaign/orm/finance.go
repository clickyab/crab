package orm

// UpdateCampaignBudget for updating campaign finance
func (m *Manager) UpdateCampaignBudget(c CampaignFinance, ca *Campaign) error {

	ca.TotalBudget = c.TotalBudget
	ca.DailyBudget = c.DailyBudget
	ca.Strategy = c.Strategy
	ca.MaxBid = c.MaxBid
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
