package orm

import "time"

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
	H00 bool `json:"h00" db:"h00"`
	H01 bool `json:"h01" db:"h01"`
	H02 bool `json:"h02" db:"h02"`
	H03 bool `json:"h03" db:"h03"`
	H04 bool `json:"h04" db:"h04"`
	H05 bool `json:"h05" db:"h05"`
	H06 bool `json:"h06" db:"h06"`
	H07 bool `json:"h07" db:"h07"`
	H08 bool `json:"h08" db:"h08"`
	H09 bool `json:"h09" db:"h09"`
	H10 bool `json:"h10" db:"h10"`
	H11 bool `json:"h11" db:"h11"`
	H12 bool `json:"h12" db:"h12"`
	H13 bool `json:"h13" db:"h13"`
	H14 bool `json:"h14" db:"h14"`
	H15 bool `json:"h15" db:"h15"`
	H16 bool `json:"h16" db:"h16"`
	H17 bool `json:"h17" db:"h17"`
	H18 bool `json:"h18" db:"h18"`
	H19 bool `json:"h19" db:"h19"`
	H20 bool `json:"h20" db:"h20"`
	H21 bool `json:"h21" db:"h21"`
	H22 bool `json:"h22" db:"h22"`
	H23 bool `json:"h23" db:"h23"`
}
