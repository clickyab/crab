// Code generated build with models DO NOT EDIT.

package location

import (
	"fmt"
	"strings"

	"github.com/clickyab/services/assert"
	"github.com/clickyab/services/initializer"
)

// AUTO GENERATED CODE. DO NOT EDIT!

// CreateCity try to save a new City in database
func (m *Manager) CreateCity(c *City) error {

	func(in interface{}) {
		if ii, ok := in.(initializer.Simple); ok {
			ii.Initialize()
		}
	}(c)

	return m.GetWDbMap().Insert(c)
}

// UpdateCity try to update City in database
func (m *Manager) UpdateCity(c *City) error {

	func(in interface{}) {
		if ii, ok := in.(initializer.Simple); ok {
			ii.Initialize()
		}
	}(c)

	_, err := m.GetWDbMap().Update(c)
	return err
}

// FindCityByID return the City base on its id
func (m *Manager) FindCityByID(id int64) (*City, error) {
	var res City
	err := m.GetRDbMap().SelectOne(
		&res,
		fmt.Sprintf("SELECT %s FROM %s WHERE id=?", getSelectFields(CityTableFull, ""), CityTableFull),
		id,
	)

	if err != nil {
		return nil, err
	}

	return &res, nil
}

// FindCityByName return the City base on its name
func (m *Manager) FindCityByName(n string) (*City, error) {
	var res City
	err := m.GetRDbMap().SelectOne(
		&res,
		fmt.Sprintf("SELECT %s FROM %s WHERE name=?", getSelectFields(CityTableFull, ""), CityTableFull),
		n,
	)

	if err != nil {
		return nil, err
	}

	return &res, nil
}

// GetProvinceCities return all Cities belong to Province
func (m *Manager) GetProvinceCities(p *Province) []City {
	var res []City
	_, err := m.GetRDbMap().Select(
		&res,
		fmt.Sprintf(
			"SELECT %s FROM %s WHERE province_id=?",
			getSelectFields(CityTableFull, ""),
			CityTableFull,
		),
		p.ID,
	)

	assert.Nil(err)
	return res
}

// CountProvinceCities return count Cities belong to Province
func (m *Manager) CountProvinceCities(p *Province) int64 {
	res, err := m.GetRDbMap().SelectInt(
		fmt.Sprintf(
			"SELECT COUNT(*) FROM %s WHERE province_id=?",
			CityTableFull,
		),
		p.ID,
	)

	assert.Nil(err)
	return res
}

// CreateCountry try to save a new Country in database
func (m *Manager) CreateCountry(c *Country) error {

	func(in interface{}) {
		if ii, ok := in.(initializer.Simple); ok {
			ii.Initialize()
		}
	}(c)

	return m.GetWDbMap().Insert(c)
}

// UpdateCountry try to update Country in database
func (m *Manager) UpdateCountry(c *Country) error {

	func(in interface{}) {
		if ii, ok := in.(initializer.Simple); ok {
			ii.Initialize()
		}
	}(c)

	_, err := m.GetWDbMap().Update(c)
	return err
}

// ListCountriesWithFilter try to list all Countries without pagination
func (m *Manager) ListCountriesWithFilter(filter string, params ...interface{}) []Country {
	filter = strings.Trim(filter, "\n\t ")
	if filter != "" {
		filter = "WHERE " + filter
	}
	var res []Country
	_, err := m.GetRDbMap().Select(
		&res,
		fmt.Sprintf("SELECT %s FROM %s %s", getSelectFields(CountryTableFull, ""), CountryTableFull, filter),
		params...,
	)
	assert.Nil(err)

	return res
}

// ListCountries try to list all Countries without pagination
func (m *Manager) ListCountries() []Country {
	return m.ListCountriesWithFilter("")
}

// CountCountriesWithFilter count entity in Countries table with valid where filter
func (m *Manager) CountCountriesWithFilter(filter string, params ...interface{}) int64 {
	filter = strings.Trim(filter, "\n\t ")
	if filter != "" {
		filter = "WHERE " + filter
	}
	cnt, err := m.GetRDbMap().SelectInt(
		fmt.Sprintf("SELECT COUNT(*) FROM %s %s", CountryTableFull, filter),
		params...,
	)
	assert.Nil(err)

	return cnt
}

// CountCountries count entity in Countries table
func (m *Manager) CountCountries() int64 {
	return m.CountCountriesWithFilter("")
}

// ListCountriesWithPaginationFilter try to list all Countries with pagination and filter
func (m *Manager) ListCountriesWithPaginationFilter(
	offset, perPage int, filter string, params ...interface{}) []Country {
	var res []Country
	filter = strings.Trim(filter, "\n\t ")
	if filter != "" {
		filter = "WHERE " + filter
	}

	filter += " LIMIT ?, ? "
	params = append(params, offset, perPage)

	// TODO : better pagination without offset and limit
	_, err := m.GetRDbMap().Select(
		&res,
		fmt.Sprintf("SELECT %s FROM %s %s", getSelectFields(CountryTableFull, ""), CountryTableFull, filter),
		params...,
	)
	assert.Nil(err)

	return res
}

// ListCountriesWithPagination try to list all Countries with pagination
func (m *Manager) ListCountriesWithPagination(offset, perPage int) []Country {
	return m.ListCountriesWithPaginationFilter(offset, perPage, "")
}

// FindCountryByID return the Country base on its id
func (m *Manager) FindCountryByID(id int64) (*Country, error) {
	var res Country
	err := m.GetRDbMap().SelectOne(
		&res,
		fmt.Sprintf("SELECT %s FROM %s WHERE id=?", getSelectFields(CountryTableFull, ""), CountryTableFull),
		id,
	)

	if err != nil {
		return nil, err
	}

	return &res, nil
}

// FindCountryByName return the Country base on its name
func (m *Manager) FindCountryByName(n string) (*Country, error) {
	var res Country
	err := m.GetRDbMap().SelectOne(
		&res,
		fmt.Sprintf("SELECT %s FROM %s WHERE name=?", getSelectFields(CountryTableFull, ""), CountryTableFull),
		n,
	)

	if err != nil {
		return nil, err
	}

	return &res, nil
}

// CreateProvince try to save a new Province in database
func (m *Manager) CreateProvince(p *Province) error {

	func(in interface{}) {
		if ii, ok := in.(initializer.Simple); ok {
			ii.Initialize()
		}
	}(p)

	return m.GetWDbMap().Insert(p)
}

// UpdateProvince try to update Province in database
func (m *Manager) UpdateProvince(p *Province) error {

	func(in interface{}) {
		if ii, ok := in.(initializer.Simple); ok {
			ii.Initialize()
		}
	}(p)

	_, err := m.GetWDbMap().Update(p)
	return err
}

// FindProvinceByID return the Province base on its id
func (m *Manager) FindProvinceByID(id int64) (*Province, error) {
	var res Province
	err := m.GetRDbMap().SelectOne(
		&res,
		fmt.Sprintf("SELECT %s FROM %s WHERE id=?", getSelectFields(ProvinceTableFull, ""), ProvinceTableFull),
		id,
	)

	if err != nil {
		return nil, err
	}

	return &res, nil
}

// FindProvinceByName return the Province base on its name
func (m *Manager) FindProvinceByName(n string) (*Province, error) {
	var res Province
	err := m.GetRDbMap().SelectOne(
		&res,
		fmt.Sprintf("SELECT %s FROM %s WHERE name=?", getSelectFields(ProvinceTableFull, ""), ProvinceTableFull),
		n,
	)

	if err != nil {
		return nil, err
	}

	return &res, nil
}

// GetCountryProvinces return all Provinces belong to Country
func (m *Manager) GetCountryProvinces(c *Country) []Province {
	var res []Province
	_, err := m.GetRDbMap().Select(
		&res,
		fmt.Sprintf(
			"SELECT %s FROM %s WHERE country_id=?",
			getSelectFields(ProvinceTableFull, ""),
			ProvinceTableFull,
		),
		c.ID,
	)

	assert.Nil(err)
	return res
}

// CountCountryProvinces return count Provinces belong to Country
func (m *Manager) CountCountryProvinces(c *Country) int64 {
	res, err := m.GetRDbMap().SelectInt(
		fmt.Sprintf(
			"SELECT COUNT(*) FROM %s WHERE country_id=?",
			ProvinceTableFull,
		),
		c.ID,
	)

	assert.Nil(err)
	return res
}
