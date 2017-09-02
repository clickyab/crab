package orm

// @Model {
//		table = campaign_manufacturer
//		primary = true, id
//		find_by = id
//		list = yes
// }
type manufacturer struct {
	base
	Brands Int64Array `json:"brands" db:"brands"`
}

// @Model {
//		table = campaign_isp
//		primary = true, id
//		find_by = id
//		list = yes
// }
type isp struct {
	base
	ISPs Int64Array `json:"isps" db:"isps"`
}

// @Model {
//		table = campaign_os
//		primary = true, id
//		find_by = id
//		list = yes
// }
type os struct {
	base
	OSs Int64Array `json:"oss" db:"oss"`
}

// @Model {
//		table = campaign_category
//		primary = true, id
//		find_by = id
//		list = yes
// }
type categories struct {
	base
	Categories Int64Array `json:"categories" db:"categories"`
}
