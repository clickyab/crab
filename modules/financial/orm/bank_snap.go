package orm

import (
	"time"

	"github.com/clickyab/services/mysql"
)

// BankSnapCheckStatus is bill payment model
// @Enum{
// }
type BankSnapCheckStatus string

const (
	// PendingStatus snap with operator
	PendingStatus BankSnapCheckStatus = "pending"
	// AcceptedStatus snap with operator
	AcceptedStatus BankSnapCheckStatus = "accepted"
	// RejectedStatus snap with operator
	RejectedStatus BankSnapCheckStatus = "rejected"
)

// BankSnap model in database
// @Model {
//		table = bank_snaps
//		primary = true, id
//		find_by = id
//		list = yes
// }
type BankSnap struct {
	ID          int64               `json:"id" db:"id"`
	DomainID    int64               `json:"domain_id" db:"domain_id"`
	UserID      int64               `json:"user_id" db:"user_id"`
	TraceNumber int64               `json:"trace_number" db:"trace_number"`
	VAT         int64               `json:"vat" db:"vat"`
	Amount      int64               `json:"amount" db:"amount"`
	PayAmount   int64               `json:"pay_amount" db:"pay_amount"`
	Status      BankSnapCheckStatus `json:"status" db:"status"`
	CheckedBy   mysql.NullInt64     `json:"checked_by" db:"checked_by"`
	CreatedAt   time.Time           `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time           `json:"updated_at" db:"updated_at"`
}
