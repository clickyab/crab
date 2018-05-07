package aaa

import (
	"time"

	"github.com/clickyab/services/mysql"
	"github.com/clickyab/services/permission"
)

// Role role model in database
// @Model {
//		table = roles
//		primary = true, id
//		find_by = id
//		list = yes
// }
type Role struct {
	ID          int64            `json:"id" db:"id"`
	Name        string           `json:"name" db:"name"`
	Description mysql.NullString `json:"description" db:"description"`
	DomainID    int64            `json:"domain_id" db:"domain_id"`
	CreatedAt   time.Time        `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time        `json:"updated_at" db:"updated_at"`
}

// RoleUser RoleUser model in database
// @Model {
//		table = role_user
//		primary = false, user_id, role_id
// }
type RoleUser struct {
	UserID    int64     `json:"user_id" db:"user_id"`
	RoleID    int64     `json:"role_id" db:"role_id"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
}

// RolePermission RolePermission model in database
// @Model {
//		table = role_permission
//		primary = true, id
//		find_by = id
// }
type RolePermission struct {
	ID        int64                `json:"id" db:"id"`
	RoleID    int64                `json:"role_id" db:"role_id"`
	Scope     permission.UserScope `json:"scope" db:"scope"`
	Perm      string               `json:"perm" db:"perm"`
	CreatedAt time.Time            `json:"created_at" db:"created_at"`
	UpdatedAt time.Time            `json:"updated_at" db:"updated_at"`
}
