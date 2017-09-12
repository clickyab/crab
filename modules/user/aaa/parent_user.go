package aaa

import "time"

// ParentUser ParentUser model in database
// @Model {
//		table = parent_user
//		primary = false, user_id,parent_id,domain_id
// }
type ParentUser struct {
	UserID    int64     `json:"user_id" db:"user_id"`
	ParentID  int64     `json:"parent_id" db:"parent_id"`
	DomainID  int64     `json:"domain_id" db:"domain_id"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
}
