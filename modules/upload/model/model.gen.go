// Code generated build with models DO NOT EDIT.

package model

import (
	"fmt"

	"github.com/clickyab/services/mysql"
	gorp "gopkg.in/gorp.v2"
)

// AUTO GENERATED CODE. DO NOT EDIT!

const (
	// UploadTableFull is the Upload table name
	UploadTableFull = "uploads"
)

func getSelectFields(tb string, alias string) string {
	if alias != "" {
		alias += "."
	}
	switch tb {

	case UploadTableFull:
		return fmt.Sprintf(`%[1]s&#34;id&#34;,%[1]s&#34;created_at&#34;,%[1]s&#34;mime&#34;,%[1]s&#34;size&#34;,%[1]s&#34;user_id&#34;,%[1]s&#34;section&#34;,%[1]s&#34;Attr&#34;`, alias)

	}
	return ""
}

// Manager is the model manager for model package
type Manager struct {
	mysql.Manager
}

// NewModelManager create and return a manager for this module
func NewModelManager() *Manager {
	return &Manager{}
}

// NewModelManagerFromTransaction create and return a manager for this module from a transaction
func NewModelManagerFromTransaction(tx gorp.SqlExecutor) (*Manager, error) {
	m := &Manager{}
	err := m.Hijack(tx)
	if err != nil {
		return nil, err
	}

	return m, nil
}

// Initialize model package
func (m *Manager) Initialize() {

	m.AddTableWithName(
		Upload{},
		UploadTableFull,
	).SetKeys(
		false,
		"ID",
	)

}
func init() {
	mysql.Register(NewModelManager())
}
