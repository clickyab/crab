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
		return fmt.Sprintf(`%[1]sid,%[1]stitle,%[1]scampaign_id,%[1]ssrc,%[1]smime,%[1]starget,%[1]swidth,%[1]sheight,%[1]sstatus,%[1]stype,%[1]sattr,%[1]screated_at,%[1]supdated_at`, alias)

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
