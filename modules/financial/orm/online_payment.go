package orm

import (
	"time"

	"github.com/clickyab/services/mysql"
)

// OnlinePaymentStatus is bill payment model
// @Enum{
// }
type OnlinePaymentStatus string

const (
	// Init pay transaction
	Init OnlinePaymentStatus = "init"
	// InGateway user go to gateway for payment
	InGateway OnlinePaymentStatus = "in_gateway"
	// BackToClickyab back to clickyab panel but not verified
	BackToClickyab OnlinePaymentStatus = "back_to_clickyab"
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
	ID          int64               `json:"id" db:"id"`
	DomainID    int64               `json:"domain_id" db:"domain_id"`
	UserID      int64               `json:"user_id" db:"user_id"`
	GatewayID   int64               `json:"gateway_id" db:"gateway_id"`
	Amount      int64               `json:"amount" db:"amount"`
	Status      OnlinePaymentStatus `json:"status" db:"status"`
	BankStatus  int64               `json:"bank_status" db:"bank_status"`
	RefNum      string              `json:"ref_num" db:"ref_num"`
	ResNum      mysql.NullString    `json:"res_num" db:"res_num"`
	CID         mysql.NullString    `json:"cid" db:"cid"`
	TraceNumber mysql.NullString    `json:"trace_number" db:"trace_number"`
	CreatedAt   time.Time           `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time           `json:"updated_at" db:"updated_at"`
}
