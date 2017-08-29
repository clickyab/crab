package orm

// Manufacturer model in database
// @Model {
//		table = manufacturers
//		primary = true, id
//		find_by = id, brand
//		list = yes
// }
type Manufacturer struct {
	base
	Brand string `json:"brand" db:"brand"`
}
