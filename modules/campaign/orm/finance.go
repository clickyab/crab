package orm

// UpdateCampaignBudget for updating campaign finance
func (m *Manager) UpdateCampaignBudget(c CampaignFinance, ca *Campaign) error {

	ca.Budget = c.Budget
	ca.DailyLimit = c.DailyLimit
	ca.CostType = c.CostType
	ca.MaxBid = c.MaxBid
	ca.NotifyEmail = c.NotifyEmail
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
