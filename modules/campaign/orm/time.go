package orm

import (
	"time"

	"github.com/clickyab/services/mysql"
)

// Schedule model in database
// @Model {
//		table = schedules
//		primary = true, id
//		find_by = id, campaign_id
//		list = yes
// }
type Schedule struct {
	ID         int64     `json:"id" db:"id"`
	CampaignID int64     `json:"campaign_id" db:"campaign_id"`
	UpdatedAt  time.Time `json:"updated_at" db:"updated_at"`
	ScheduleSheet
}

// ScheduleSheet is a hour/view chart
type ScheduleSheet struct {
	H00 mysql.NullString `json:"h00" db:"h00"`
	H01 mysql.NullString `json:"h01" db:"h01"`
	H02 mysql.NullString `json:"h02" db:"h02"`
	H03 mysql.NullString `json:"h03" db:"h03"`
	H04 mysql.NullString `json:"h04" db:"h04"`
	H05 mysql.NullString `json:"h05" db:"h05"`
	H06 mysql.NullString `json:"h06" db:"h06"`
	H07 mysql.NullString `json:"h07" db:"h07"`
	H08 mysql.NullString `json:"h08" db:"h08"`
	H09 mysql.NullString `json:"h09" db:"h09"`
	H10 mysql.NullString `json:"h10" db:"h10"`
	H11 mysql.NullString `json:"h11" db:"h11"`
	H12 mysql.NullString `json:"h12" db:"h12"`
	H13 mysql.NullString `json:"h13" db:"h13"`
	H14 mysql.NullString `json:"h14" db:"h14"`
	H15 mysql.NullString `json:"h15" db:"h15"`
	H16 mysql.NullString `json:"h16" db:"h16"`
	H17 mysql.NullString `json:"h17" db:"h17"`
	H18 mysql.NullString `json:"h18" db:"h18"`
	H19 mysql.NullString `json:"h19" db:"h19"`
	H20 mysql.NullString `json:"h20" db:"h20"`
	H21 mysql.NullString `json:"h21" db:"h21"`
	H22 mysql.NullString `json:"h22" db:"h22"`
	H23 mysql.NullString `json:"h23" db:"h23"`
}
