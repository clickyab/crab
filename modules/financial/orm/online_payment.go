package orm

import (
	"time"

	"fmt"

	"clickyab.com/crab/modules/user/aaa"
	"github.com/clickyab/services/assert"
	"github.com/clickyab/services/config"
	"github.com/clickyab/services/mysql"
)

// VatPercent global payment vat percent
var VatPercent = config.RegisterInt("crab.modules.financial.vat", 9, "vat percent")

// OnlinePaymentStatus is bill payment model
// @Enum{
// }
type OnlinePaymentStatus string

const (
	// Init pay transaction
	Init OnlinePaymentStatus = "init"
	// BackToSite back to clickyab panel but not verified
	BackToSite OnlinePaymentStatus = "back_to_site"
	// Finalized verified and finalized payment successfully
	Finalized OnlinePaymentStatus = "finalized"
)

// OnlinePayment model in database
// @Model {
//		table = online_payments
//		primary = true, id
//		find_by = id
//		list = yes
// }
type OnlinePayment struct {
	ID          int64                  `json:"id" db:"id"`
	DomainID    int64                  `json:"domain_id" db:"domain_id"`
	UserID      int64                  `json:"user_id" db:"user_id"`
	GatewayID   int64                  `json:"gateway_id" db:"gateway_id"`
	Amount      int64                  `json:"amount" db:"amount"`
	Status      OnlinePaymentStatus    `json:"status" db:"status"`
	BankStatus  mysql.NullInt64        `json:"bank_status" db:"bank_status"`
	RefNum      mysql.NullString       `json:"ref_num" db:"ref_num"`
	ResNum      string                 `json:"res_num" db:"res_num"`
	CID         mysql.NullString       `json:"cid" db:"cid"`
	TraceNumber mysql.NullString       `json:"trace_number" db:"trace_number"`
	Attr        mysql.GenericJSONField `json:"attr" db:"attr"`
	Reason      mysql.NullString       `json:"reason" db:"reason"`
	CreatedAt   time.Time              `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time              `json:"updated_at" db:"updated_at"`
}

// FindInitPaymentByResNum return the OnlinePayment base on its res num
func (m *Manager) FindInitPaymentByResNum(id string) (*OnlinePayment, error) {
	var res OnlinePayment
	err := m.GetRDbMap().SelectOne(
		&res,
		fmt.Sprintf("SELECT %s FROM %s WHERE res_num=? AND status=?", getSelectFields(OnlinePaymentTableFull, ""), OnlinePaymentTableFull),
		id,
		Init,
	)

	if err != nil {
		return nil, err
	}

	return &res, nil
}

// ChargeUser charge user transaction
func (m *Manager) ChargeUser(payment *OnlinePayment, domainID, chargeAmount int64) error {
	err := m.Begin()
	assert.Nil(err)
	defer func() {
		if err == nil {
			assert.Nil(m.Commit())
			return
		}
		assert.Nil(m.Rollback())
	}()
	err = m.UpdateOnlinePayment(payment)
	if err != nil {
		return err
	}

	billQ := fmt.Sprintf("SELECT COALESCE(SUM(amount),0) AS balance FROM %s WHERE user_id=?", BillingTableFull)
	oldBalance, err := m.GetRDbMap().SelectInt(billQ, payment.UserID)
	if err != nil {
		return err
	}

	newBalance := oldBalance + chargeAmount

	// create billing
	bill := &Billing{
		Amount:    chargeAmount,
		PayAmount: payment.Amount,
		VAT:       payment.Amount - chargeAmount,
		UserID:    payment.UserID,
		IncomeID:  payment.ID,
		DomainID:  domainID,
		PayModel:  OnlinePaymentModel,
		Balance:   newBalance,
	}
	err = m.CreateBilling(bill)
	if err != nil {
		return err
	}

	// update user balance
	userManager, err := aaa.NewAaaManagerFromTransaction(m.GetWDbMap())
	if err != nil {
		return err
	}

	userQ := fmt.Sprintf("UPDATE %s SET balance=? WHERE id=?", aaa.UserTableFull)
	_, err = userManager.GetWDbMap().Exec(userQ, newBalance, payment.UserID)
	return err
}
