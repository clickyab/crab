package orm

// @Model {
//		table = campaign_document
//		primary = true, id
//		find_by = id
//		list = yes
// }
type assets struct {
	base
	OSs      stringArray `json:"os" db:"os"`
	Browsers  stringArray `json:"browser" db:"browser"`
	Brands    stringArray `json:"brand" db:"brand"`
	Categories stringArray `json:"category" db:"category"`
	ISPs      stringArray `json:"isp" db:"isp"`
}
