package aaa

import (
	"time"

	"github.com/clickyab/services/mysql"
)

// AuditLogDetail audit log detail model in database
// @Model {
//		table = audit_log_details
//		primary = true, id
//		find_by = id
//		list = yes
// }
type AuditLogDetail struct {
	ID         int64                  `json:"id" db:"id"`
	AuditLogID int64                  `json:"audit_log_id" db:"audit_log_id"`
	Data       mysql.GenericJSONField `json:"data" db:"data"`
	CreatedAt  time.Time              `json:"created_at" db:"created_at"`
	UpdatedAt  mysql.NullTime         `json:"updated_at" db:"updated_at"`
}
