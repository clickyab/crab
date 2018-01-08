// Code generated build with models DO NOT EDIT.

package add

import (
	"fmt"

	"github.com/clickyab/services/mysql"
	gorp "gopkg.in/gorp.v2"
)

// AUTO GENERATED CODE. DO NOT EDIT!

const (
	// AdTableFull is the Ad table name
	AdTableFull = "ads"
)

func getSelectFields(tb string, alias string) string {
	if alias != "" {
		alias += "."
	}
	switch tb {

	case AdTableFull:
		return fmt.Sprintf(`%[1]s&#34;id&#34;,%[1]s&#34;campaign_id&#34;,%[1]s&#34;src&#34;,%[1]s&#34;mime&#34;,%[1]s&#34;target&#34;,%[1]s&#34;width&#34;,%[1]s&#34;height&#34;,%[1]s&#34;status&#34;,%[1]s&#34;type&#34;,%[1]s&#34;attr&#34;,%[1]s&#34;created_at&#34;,%[1]s&#34;updated_at&#34;`, alias)

	}
	return ""
}

// Manager is the model manager for add package
type Manager struct {
	mysql.Manager
}

// NewAddManager create and return a manager for this module
func NewAddManager() *Manager {
	return &Manager{}
}

// NewAddManagerFromTransaction create and return a manager for this module from a transaction
func NewAddManagerFromTransaction(tx gorp.SqlExecutor) (*Manager, error) {
	m := &Manager{}
	err := m.Hijack(tx)
	if err != nil {
		return nil, err
	}

	return m, nil
}

// Initialize add package
func (m *Manager) Initialize() {

	m.AddTableWithName(
		Ad{},
		AdTableFull,
	).SetKeys(
		true,
		"ID",
	)

}
func init() {
	mysql.Register(NewAddManager())
}
