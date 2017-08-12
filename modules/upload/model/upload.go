package model

import "time"

// Upload model in database
// @Model {
//		table = uploads
//		primary = true, id
//		find_by = id
//		list = yes
// }
type Upload struct {
	ID        int64     `json:"id" db:"id"`
	Path      string    `json:"path"  db:"path"`
	CreatedAt time.Time `json:"created_at"  db:"created_at"`
	MIME      string    `json:"mime"  db:"mime"`
	Size      int64     `json:"size"  db:"size"`
	UserID    int64     `json:"user_id"  db:"user_id"`
	Section   string    `json:"section" db:"section"`
}
