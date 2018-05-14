package orm

import (
	"fmt"
	"time"
)

// PayModels is bill payment model
// @Enum{
// }
type PayModels string

const (
	// OnlinePaymentModel pay type
	OnlinePaymentModel PayModels = "online_payment"
	// BankSnapModel pay with bank snap
	BankSnapModel PayModels = "bank_snap"
	// ManualCashChangeModel payment type
	ManualCashChangeModel PayModels = "manual_cash_change"
)

// Billing model in database
// @Model {
//		table = billings
//		primary = true, id
//		find_by = id
//		list = yes
// }
type Billing struct {
	ID        int64     `json:"id" db:"id"`
	DomainID  int64     `json:"domain_id" db:"domain_id"`
	UserID    int64     `json:"user_id" db:"user_id"`
	PayModel  PayModels `json:"pay_model" db:"pay_model"`
	IncomeID  int64     `json:"income_id" db:"income_id"`
	VAT       int64     `json:"vat" db:"vat"`
	Amount    int64     `json:"amount" db:"amount"`
	PayAmount int64     `json:"pay_amount" db:"pay_amount"`
	Balance   int64     `json:"balance" db:"balance"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
}

// GetUserBalance try to get user balance
func (m *Manager) GetUserBalance(userID int64) (int64, error) {
	billQ := fmt.Sprintf("SELECT COALESCE(SUM(amount),0) AS balance FROM %s WHERE user_id=?", BillingTableFull)
	userBalance, err := m.GetRDbMap().SelectInt(billQ, userID)
	if err != nil {
		return 0, err
	}
	return userBalance, err
}

// GetUserDomainBalance try to get user balance in a domain
func (m *Manager) GetUserDomainBalance(userID int64, domainID int64) (int64, error) {
	billQ := fmt.Sprintf("SELECT COALESCE(SUM(amount),0) AS balance FROM %s WHERE user_id=? AND domain_id=?", BillingTableFull)
	userBalance, err := m.GetRDbMap().SelectInt(billQ, userID, domainID)
	if err != nil {
		return 0, err
	}
	return userBalance, err
}
