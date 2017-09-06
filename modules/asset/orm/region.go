package orm

// Region model in database
// @Model {
//		table = regions
//		primary = true, id
//		list = yes
// }
type Region struct {
	base
	Name string `json:"name" db:"name"`
}
