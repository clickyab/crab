// Code generated build with models DO NOT EDIT.

package models

import (
	"fmt"
	"strings"
	"time"

	"github.com/clickyab/services/assert"
	"github.com/clickyab/services/initializer"
)

// AUTO GENERATED CODE. DO NOT EDIT!

// CreateSchedule try to save a new Schedule in database
func (m *Manager) CreateSchedule(s *Schedule) error {
	now := time.Now()

	s.UpdatedAt = now
	func(in interface{}) {
		if ii, ok := in.(initializer.Simple); ok {
			ii.Initialize()
		}
	}(s)

	return m.GetWDbMap().Insert(s)
}

// UpdateSchedule try to update Schedule in database
func (m *Manager) UpdateSchedule(s *Schedule) error {
	now := time.Now()
	s.UpdatedAt = now
	func(in interface{}) {
		if ii, ok := in.(initializer.Simple); ok {
			ii.Initialize()
		}
	}(s)

	_, err := m.GetWDbMap().Update(s)
	return err
}

// ListSchedulesWithFilter try to list all Schedules without pagination
func (m *Manager) ListSchedulesWithFilter(filter string, params ...interface{}) []Schedule {
	filter = strings.Trim(filter, "\n\t ")
	if filter != "" {
		filter = "WHERE " + filter
	}
	var res []Schedule
	_, err := m.GetRDbMap().Select(
		&res,
		fmt.Sprintf("SELECT * FROM %s %s", ScheduleTableFull, filter),
		params...,
	)
	assert.Nil(err)

	return res
}

// ListSchedules try to list all Schedules without pagination
func (m *Manager) ListSchedules() []Schedule {
	return m.ListSchedulesWithFilter("")
}

// CountSchedulesWithFilter count entity in Schedules table with valid where filter
func (m *Manager) CountSchedulesWithFilter(filter string, params ...interface{}) int64 {
	filter = strings.Trim(filter, "\n\t ")
	if filter != "" {
		filter = "WHERE " + filter
	}
	cnt, err := m.GetRDbMap().SelectInt(
		fmt.Sprintf("SELECT COUNT(*) FROM %s %s", ScheduleTableFull, filter),
		params...,
	)
	assert.Nil(err)

	return cnt
}

// CountSchedules count entity in Schedules table
func (m *Manager) CountSchedules() int64 {
	return m.CountSchedulesWithFilter("")
}

// ListSchedulesWithPaginationFilter try to list all Schedules with pagination and filter
func (m *Manager) ListSchedulesWithPaginationFilter(
	offset, perPage int, filter string, params ...interface{}) []Schedule {
	var res []Schedule
	filter = strings.Trim(filter, "\n\t ")
	if filter != "" {
		filter = "WHERE " + filter
	}

	filter += " LIMIT ?, ? "
	params = append(params, offset, perPage)

	// TODO : better pagination without offset and limit
	_, err := m.GetRDbMap().Select(
		&res,
		fmt.Sprintf("SELECT * FROM %s %s", ScheduleTableFull, filter),
		params...,
	)
	assert.Nil(err)

	return res
}

// ListSchedulesWithPagination try to list all Schedules with pagination
func (m *Manager) ListSchedulesWithPagination(offset, perPage int) []Schedule {
	return m.ListSchedulesWithPaginationFilter(offset, perPage, "")
}

// FindScheduleByID return the Schedule base on its id
func (m *Manager) FindScheduleByID(id int64) (*Schedule, error) {
	var res Schedule
	err := m.GetRDbMap().SelectOne(
		&res,
		fmt.Sprintf("SELECT * FROM %s WHERE id=?", ScheduleTableFull),
		id,
	)

	if err != nil {
		return nil, err
	}

	return &res, nil
}

// FindScheduleByCampaignID return the Schedule base on its campaign_id
func (m *Manager) FindScheduleByCampaignID(ci int64) (*Schedule, error) {
	var res Schedule
	err := m.GetRDbMap().SelectOne(
		&res,
		fmt.Sprintf("SELECT * FROM %s WHERE campaign_id=?", ScheduleTableFull),
		ci,
	)

	if err != nil {
		return nil, err
	}

	return &res, nil
}
