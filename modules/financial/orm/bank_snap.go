package orm

import "time"

// BankSnapCheckStatus is bill payment model
// @Enum{
// }
type BankSnapCheckStatus string

const (
	// WaitToCheck snap with operator
	WaitToCheck BankSnapCheckStatus = "wait_to_check"
	// Verified snap with operator
	Verified BankSnapCheckStatus = "verified"
	// Rejected snap with operator
	Rejected BankSnapCheckStatus = "rejected"
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
	Amount      int64               `json:"amount" db:"amount"`
	Status      BankSnapCheckStatus `json:"type" db:"type"`
	CheckedBy   int64               `json:"checked_by" db:"checked_by"`
	CreatedAt   time.Time           `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time           `json:"updated_at" db:"updated_at"`
}
