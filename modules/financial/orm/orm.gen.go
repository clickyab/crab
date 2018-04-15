// Code generated build with models DO NOT EDIT.

package orm

import (
	"fmt"

	"github.com/clickyab/services/mysql"
	gorp "gopkg.in/gorp.v2"
)

// AUTO GENERATED CODE. DO NOT EDIT!

const (
	// BankSnapTableFull is the BankSnap table name
	BankSnapTableFull = "bank_snaps"

	// BillingTableFull is the Billing table name
	BillingTableFull = "billings"

	// GatewayTableFull is the Gateway table name
	GatewayTableFull = "gateways"

	// ManualCashChangeTableFull is the ManualCashChange table name
	ManualCashChangeTableFull = "manual_cash_changes"

	// OnlinePaymentTableFull is the OnlinePayment table name
	OnlinePaymentTableFull = "online_payments"
)

func getSelectFields(tb string, alias string) string {
	if alias != "" {
		alias += "."
	}
	switch tb {

	case BankSnapTableFull:
		return fmt.Sprintf(`%[1]sid,%[1]sdomain_id,%[1]suser_id,%[1]strace_number,%[1]samount,%[1]stype,%[1]schecked_by,%[1]screated_at,%[1]supdated_at`, alias)

	case BillingTableFull:
		return fmt.Sprintf(`%[1]sid,%[1]sdomain_id,%[1]suser_id,%[1]spay_model,%[1]sincome_id,%[1]svat,%[1]samount,%[1]spay_amount,%[1]sbalance,%[1]screated_at`, alias)

	case GatewayTableFull:
		return fmt.Sprintf(`%[1]sid,%[1]sname,%[1]sstatus,%[1]sdefault,%[1]screated_at,%[1]screated_at`, alias)

	case ManualCashChangeTableFull:
		return fmt.Sprintf(`%[1]sid,%[1]sdomain_id,%[1]suser_id,%[1]soperator_id,%[1]sreason,%[1]samount,%[1]sstatus,%[1]sdescription,%[1]screated_at,%[1]supdated_at`, alias)

	case OnlinePaymentTableFull:
		return fmt.Sprintf(`%[1]sid,%[1]sdomain_id,%[1]suser_id,%[1]sgateway_id,%[1]samount,%[1]sstatus,%[1]sbank_status,%[1]sref_num,%[1]sres_num,%[1]scid,%[1]strace_number,%[1]screated_at,%[1]supdated_at`, alias)

	}
	return ""
}

// Manager is the model manager for orm package
type Manager struct {
	mysql.Manager
}

// NewOrmManager create and return a manager for this module
func NewOrmManager() *Manager {
	return &Manager{}
}

// NewOrmManagerFromTransaction create and return a manager for this module from a transaction
func NewOrmManagerFromTransaction(tx gorp.SqlExecutor) (*Manager, error) {
	m := &Manager{}
	err := m.Hijack(tx)
	if err != nil {
		return nil, err
	}

	return m, nil
}

// Initialize orm package
func (m *Manager) Initialize() {

	m.AddTableWithName(
		BankSnap{},
		BankSnapTableFull,
	).SetKeys(
		true,
		"ID",
	)

	m.AddTableWithName(
		Billing{},
		BillingTableFull,
	).SetKeys(
		true,
		"ID",
	)

	m.AddTableWithName(
		Gateway{},
		GatewayTableFull,
	).SetKeys(
		true,
		"ID",
	)

	m.AddTableWithName(
		ManualCashChange{},
		ManualCashChangeTableFull,
	).SetKeys(
		true,
		"ID",
	)

	m.AddTableWithName(
		OnlinePayment{},
		OnlinePaymentTableFull,
	).SetKeys(
		true,
		"ID",
	)

}
func init() {
	mysql.Register(NewOrmManager())
}
