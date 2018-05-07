package orm

import (
	"time"

	"github.com/go-sql-driver/mysql"
)

// RejectReasonStatusTypes is the reject reason status type
type (
	// RejectReasonStatusTypes is the reject reason status type
	// @Enum{
	// }
	RejectReasonStatusTypes string
)

const (
	//EnableRejectReason RejectReasonStatusTypes enable
	EnableRejectReason RejectReasonStatusTypes = "enable"
	//DisableRejectReason RejectReasonStatusTypes disable
	DisableRejectReason RejectReasonStatusTypes = "disable"
)

// CreativeRejectReasons model in database
// @Model {
//		table = creative_reject_reasons
//		primary = true, id
//		find_by = id
//		list = yes
// }
type CreativeRejectReasons struct {
	ID        int64                   `json:"id" db:"id"`
	Reason    string                  `json:"reason" db:"reason"`
	Status    RejectReasonStatusTypes `json:"status" db:"status"`
	CreatedAt time.Time               `json:"created_at" db:"created_at"`
	UpdatedAt mysql.NullTime          `json:"updated_at" db:"updated_at"`
}
