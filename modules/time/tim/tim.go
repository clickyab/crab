package tim

import "time"

// Date model in database
// @Model {
//		table = date_table
//		primary = true, id
//		find_by = id
// }
type Date struct {
	ID     int64     `json:"id" db:"id"`
	Year   int64     `json:"year" db:"year"`
	Month  int64     `json:"month" db:"month"`
	Day    int64     `json:"day_int_64" db:"day_int_64"`
	JYear  int64     `json:"j_year" db:"j_year"`
	JMonth int64     `json:"j_month" db:"j_month"`
	JDay   int64     `json:"j_day" db:"j_day"`
	Extra  string    `json:"extra" db:"extra"`
	Basis  time.Time `json:"basis" db:"basis"`
}

// Hour model in database
// @Model {
//		table = hour_table
//		primary = true, id
//		find_by = id
// }
type Hour struct {
	ID     int64     `json:"id" db:"id"`
	Year   int64     `json:"year" db:"year"`
	Month  int64     `json:"month" db:"month"`
	Day    int64     `json:"day_int_64" db:"day_int_64"`
	Hour   int64     `json:"hour" db:"hour"`
	JYear  int64     `json:"j_year" db:"j_year"`
	JMonth int64     `json:"j_month" db:"j_month"`
	JDay   int64     `json:"j_day" db:"j_day"`
	DateID int64     `json:"date_id" db:"date_id"`
	Extra  string    `json:"extra" db:"extra"`
	Basis  time.Time `json:"basis" db:"basis"`
}
