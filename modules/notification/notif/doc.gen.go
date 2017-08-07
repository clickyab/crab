// Code generated build with models DO NOT EDIT.

package notif

import (
	"fmt"
	"strings"
	"time"

	"github.com/clickyab/services/assert"
	"github.com/clickyab/services/initializer"
)

// AUTO GENERATED CODE. DO NOT EDIT!

// CreateNotification try to save a new Notification in database
func (m *Manager) CreateNotification(n *Notification) error {
	now := time.Now()
	n.CreatedAt = &now

	func(in interface{}) {
		if ii, ok := in.(initializer.Simple); ok {
			ii.Initialize()
		}
	}(n)

	return m.GetWDbMap().Insert(n)
}

// UpdateNotification try to update Notification in database
func (m *Manager) UpdateNotification(n *Notification) error {

	func(in interface{}) {
		if ii, ok := in.(initializer.Simple); ok {
			ii.Initialize()
		}
	}(n)

	_, err := m.GetWDbMap().Update(n)
	return err
}

// ListNotificationsWithFilter try to list all Notifications without pagination
func (m *Manager) ListNotificationsWithFilter(filter string, params ...interface{}) []Notification {
	filter = strings.Trim(filter, "\n\t ")
	if filter != "" {
		filter = "WHERE " + filter
	}
	var res []Notification
	_, err := m.GetRDbMap().Select(
		&res,
		fmt.Sprintf("SELECT * FROM %s %s", NotificationTableFull, filter),
		params...,
	)
	assert.Nil(err)

	return res
}

// ListNotifications try to list all Notifications without pagination
func (m *Manager) ListNotifications() []Notification {
	return m.ListNotificationsWithFilter("")
}

// CountNotificationsWithFilter count entity in Notifications table with valid where filter
func (m *Manager) CountNotificationsWithFilter(filter string, params ...interface{}) int64 {
	filter = strings.Trim(filter, "\n\t ")
	if filter != "" {
		filter = "WHERE " + filter
	}
	cnt, err := m.GetRDbMap().SelectInt(
		fmt.Sprintf("SELECT COUNT(*) FROM %s %s", NotificationTableFull, filter),
		params...,
	)
	assert.Nil(err)

	return cnt
}

// CountNotifications count entity in Notifications table
func (m *Manager) CountNotifications() int64 {
	return m.CountNotificationsWithFilter("")
}

// ListNotificationsWithPaginationFilter try to list all Notifications with pagination and filter
func (m *Manager) ListNotificationsWithPaginationFilter(
	offset, perPage int, filter string, params ...interface{}) []Notification {
	var res []Notification
	filter = strings.Trim(filter, "\n\t ")
	if filter != "" {
		filter = "WHERE " + filter
	}

	filter += " LIMIT ?, ? "
	params = append(params, offset, perPage)

	// TODO : better pagination without offset and limit
	_, err := m.GetRDbMap().Select(
		&res,
		fmt.Sprintf("SELECT * FROM %s %s", NotificationTableFull, filter),
		params...,
	)
	assert.Nil(err)

	return res
}

// ListNotificationsWithPagination try to list all Notifications with pagination
func (m *Manager) ListNotificationsWithPagination(offset, perPage int) []Notification {
	return m.ListNotificationsWithPaginationFilter(offset, perPage, "")
}

// FindNotificationByID return the Notification base on its id
func (m *Manager) FindNotificationByID(id int64) (*Notification, error) {
	var res Notification
	err := m.GetRDbMap().SelectOne(
		&res,
		fmt.Sprintf("SELECT * FROM %s WHERE id=?", NotificationTableFull),
		id,
	)

	if err != nil {
		return nil, err
	}

	return &res, nil
}
