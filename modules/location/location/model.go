package location

import (
	"fmt"

	"github.com/clickyab/services/assert"
)

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
//		primary = false, name
//		find_by = name,code
//		belong_to = Country:country_id
// }
type Province struct {
	Code      string `json:"code" db:"code"`
	Name      string `json:"name" db:"name"`
	FAName    string `json:"fa_name" db:"fa_name"`
	CountryID int64  `json:"country_id" db:"country_id"`
}

// City model in database
// @Model {
//		table = cities
//		primary = true, id
//		find_by = id,name
// }
type City struct {
	ID       int64  `json:"id" db:"id"`
	Name     string `json:"name" db:"name"`
	Province string `json:"province" db:"province"`
}

// CityInfo is city info
type CityInfo struct {
	CityName     string `json:"city_name" db:"city_name"`
	CityID       int64  `json:"city_id" db:"city_id"`
	ProvinceName string `json:"province_name" db:"province_name"`
	Province     string `json:"province" db:"province"`
	CountryName  string `json:"country_name" db:"country_name"`
	CountryID    int64  `json:"country_id" db:"country_id"`
}

// FindAllByCityID find city by id
func (m *Manager) FindAllByCityID(id int64) CityInfo {
	c := CityInfo{}
	e := m.GetRDbMap().SelectOne(&c,
		fmt.Sprintf(
			`SELECT 
			c.name 			AS city_name,
			c.id 			AS city_id,
			p.name 			AS province_name,
			p.code 			AS province,
			cu.id 			AS country_id,
			cu.name 		AS country_name

			FROM %s AS c
			JOIN %s AS p ON c.province = p.name
			JOIN %s AS cu ON p.country_id = cu.id
			where c.id=?`,
			CityTableFull,
			ProvinceTableFull,
			CountryTableFull,
		),
		id,
	)
	assert.Nil(e)
	return c
}

// GetProvinceCities get city by province id
func (m *Manager) GetProvinceCities(province string) []City {
	var res []City
	q := fmt.Sprintf("SELECT %s FROM %s AS c WHERE province=?",
		getSelectFields(CityTableFull, "c"),
		CityTableFull,
	)
	_, err := m.GetRDbMap().Select(&res, q, province)
	assert.Nil(err)
	return res
}
