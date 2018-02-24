// Code generated build with models DO NOT EDIT.

package notif

import (
	"fmt"

	"github.com/clickyab/services/mysql"
	gorp "gopkg.in/gorp.v2"
)

// AUTO GENERATED CODE. DO NOT EDIT!

const (
	// NotificationTableFull is the Notification table name
	NotificationTableFull = "notification"
)

func getSelectFields(tb string, alias string) string {
	if alias != "" {
		alias += "."
	}
	switch tb {

	case NotificationTableFull:
		return fmt.Sprintf(`%[1]sid,%[1]suser_id,%[1]stitle,%[1]smessage,%[1]stype,%[1]screated_at`, alias)

	}
	return ""
}

// Manager is the model manager for notif package
type Manager struct {
	mysql.Manager
}

// NewNotifManager create and return a manager for this module
func NewNotifManager() *Manager {
	return &Manager{}
}

// NewNotifManagerFromTransaction create and return a manager for this module from a transaction
func NewNotifManagerFromTransaction(tx gorp.SqlExecutor) (*Manager, error) {
	m := &Manager{}
	err := m.Hijack(tx)
	if err != nil {
		return nil, err
	}

	return m, nil
}

// Initialize notif package
func (m *Manager) Initialize() {

	m.AddTableWithName(
		Notification{},
		NotificationTableFull,
	).SetKeys(
		true,
		"ID",
	)

}
func init() {
	mysql.Register(NewNotifManager())
}
