package location

// Country model in database
// @Model {
//		table = countries
//		primary = true, id
//		find_by = id,name
//		list = yes
// }
type Country struct {
	ID   int64  `json:"id" db:"id"`
	Name string `json:"name" db:"name"`
}

// Province model in database
// @Model {
//		table = provinces
//		primary = true, id
//		find_by = id,name
//		belong_to = Country:country_id
// }
type Province struct {
	ID        int64  `json:"id" db:"id"`
	Name      string `json:"name" db:"name"`
	CountryID int64  `json:"country_id" db:"country_id"`
}

// City model in database
// @Model {
//		table = cities
//		primary = true, id
//		find_by = id,name
//		belong_to = Province:province_id
// }
type City struct {
	ID         int64  `json:"id" db:"id"`
	Name       string `json:"name" db:"name"`
	ProvinceID int64  `json:"province_id" db:"province_id"`
}
