package orm

import "time"

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

// ManualCashChangeStatus is bill payment model
// @Enum{
// }
type ManualCashChangeStatus string

const (
	// WaitToCheckCashChange with operator
	WaitToCheckCashChange ManualCashChangeStatus = "pending"
	// VerifiedCashChange with operator
	VerifiedCashChange ManualCashChangeStatus = "accepted"
	// RejectedCashChange with operator
	RejectedCashChange ManualCashChangeStatus = "rejected"
)

// ManualCashChange model in database
// @Model {
//		table = manual_cash_changes
//		primary = true, id
//		find_by = id
//		list = yes
// }
type ManualCashChange struct {
	ID          int64                  `json:"id" db:"id"`
	DomainID    int64                  `json:"domain_id" db:"domain_id"`
	UserID      int64                  `json:"user_id" db:"user_id"`
	OperatorID  int64                  `json:"operator_id" db:"operator_id"`
	Reason      ChangeCashReasons      `json:"reason" db:"reason"`
	Amount      int64                  `json:"amount" db:"amount"`
	Status      ManualCashChangeStatus `json:"status" db:"status"`
	Description string                 `json:"description" db:"description"`
	CreatedAt   time.Time              `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time              `json:"updated_at" db:"updated_at"`
}
