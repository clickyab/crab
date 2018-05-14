package orm

import (
	"fmt"
	"time"

	"clickyab.com/crab/modules/user/aaa"
	"github.com/clickyab/services/assert"
	"github.com/clickyab/services/mysql"
)

// ChangeCashReasons is bill payment model
// @Enum{
// }
type ChangeCashReasons string

const (
	// Gift to user
	Gift ChangeCashReasons = "gift"
	// ManualPay pay manually
	ManualPay ChangeCashReasons = "manual_pay"
	// Refund for some reasons
	Refund ChangeCashReasons = "refund"
)

// ManualCashChange model in database
// @Model {
//		table = manual_cash_changes
//		primary = true, id
//		find_by = id
//		list = yes
// }
type ManualCashChange struct {
	ID          int64             `json:"id" db:"id"`
	DomainID    int64             `json:"domain_id" db:"domain_id"`
	UserID      int64             `json:"user_id" db:"user_id"`
	OperatorID  int64             `json:"operator_id" db:"operator_id"`
	Reason      ChangeCashReasons `json:"reason" db:"reason"`
	Amount      int64             `json:"amount" db:"amount"`
	Description string            `json:"description" db:"description"`
	CreatedAt   time.Time         `json:"created_at" db:"created_at"`
	UpdatedAt   mysql.NullTime    `json:"updated_at" db:"updated_at"`
}

func addChangeToBilling(m *Manager, user *aaa.User, incomeID int64, domainID int64, balance int64, amount int64) error {
	billingDetail := &Billing{
		DomainID:  domainID,
		UserID:    user.ID,
		PayModel:  ManualCashChangeModel,
		IncomeID:  incomeID,
		VAT:       0,
		Amount:    amount,
		PayAmount: 0,
		Balance:   balance,
	}
	createBillErr := m.CreateBilling(billingDetail)
	return createBillErr
}

//ApplyManualCash apply change cash
func (m *Manager) ApplyManualCash(fromUser *aaa.User, toUser *aaa.User, ch *ManualCashChange) error {
	err := m.Begin()
	assert.Nil(err)
	defer func() {
		if err != nil {
			assert.Nil(m.Rollback())
		} else {
			assert.Nil(m.Commit())
		}
	}()
	// dec/inc amount from/to creator
	operatorAmount := ch.Amount * -1
	var newCreatorBalance int64
	newCreatorBalance, err = m.changeUserBalance(fromUser, ch.DomainID, operatorAmount)
	if err != nil {
		return err
	}
	// add amount to manual change cash
	err = m.CreateManualCashChange(ch)
	if err != nil {
		return err
	}
	// insert creator billing
	err = addChangeToBilling(m, fromUser, ch.ID, ch.DomainID, newCreatorBalance, operatorAmount)
	if err != nil {
		return err
	}
	// inc/dec user balance
	var newUserBalance int64
	newUserBalance, err = m.changeUserBalance(toUser, ch.DomainID, ch.Amount)
	if err != nil {
		return err
	}
	// insert user billing
	err = addChangeToBilling(m, toUser, ch.ID, ch.DomainID, newUserBalance, ch.Amount)
	return err
}

//changeUserBalance change user balance by user id, amount can be positive or negative
func (m *Manager) changeUserBalance(user *aaa.User, domainID int64, amount int64) (int64, error) {
	oldBalance, err := m.GetUserDomainBalance(user.ID, domainID)
	if err != nil {
		return 0, err
	}
	newBalance := oldBalance + amount
	// update user balance
	userQ := fmt.Sprintf("UPDATE %s SET balance=? WHERE id=?", aaa.UserTableFull)
	_, err = m.GetWDbMap().Exec(userQ, newBalance, user.ID)
	return newBalance, err
}
