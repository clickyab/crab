// Code generated build with models DO NOT EDIT.

package orm

import (
	"fmt"

	"github.com/clickyab/services/mysql"
	gorp "gopkg.in/gorp.v2"
)

// AUTO GENERATED CODE. DO NOT EDIT!

const (
	// AssetTableFull is the Asset table name
	AssetTableFull = "assets"

	// CreativeDetailTableFull is the CreativeDetail table name
	CreativeDetailTableFull = "creative_detail"

	// CreativeRejectReasonsTableFull is the CreativeRejectReasons table name
	CreativeRejectReasonsTableFull = "creative_reject_reasons"

	// CreativeTableFull is the Creative table name
	CreativeTableFull = "creatives"
)

func GetSelectFields(tb string, alias string) string {
	if alias != "" {
		alias += "."
	}
	switch tb {

	case AssetTableFull:
		return fmt.Sprintf(`%[1]sid,%[1]screative_id,%[1]sasset_type,%[1]sproperty,%[1]sasset_key,%[1]sasset_value,%[1]screated_at,%[1]supdated_at`, alias)

	case CreativeDetailTableFull:
		return fmt.Sprintf(`%[1]scampaign_id,%[1]screative_id,%[1]sdaily_id,%[1]shour_id,%[1]sfake_imp,%[1]sfake_click,%[1]simp,%[1]sclick,%[1]scpc,%[1]scpa,%[1]scpm,%[1]sconv,%[1]screated_at,%[1]supdated_at`, alias)

	case CreativeRejectReasonsTableFull:
		return fmt.Sprintf(`%[1]sid,%[1]sreason,%[1]sstatus,%[1]screated_at,%[1]supdated_at`, alias)

	case CreativeTableFull:
		return fmt.Sprintf(`%[1]sid,%[1]suser_id,%[1]scampaign_id,%[1]sstatus,%[1]stype,%[1]surl,%[1]sname,%[1]smax_bid,%[1]sattributes,%[1]sreject_reasons_id,%[1]screated_at,%[1]supdated_at,%[1]sarchived_at`, alias)

	}
	return ""
}

// Manager is the model manager for orm package
type Manager struct {
	mysql.Manager
}

// NewOrmManager create and return a manager for this module
func NewOrmManager() *Manager {
	return &Manager{}
}

// NewOrmManagerFromTransaction create and return a manager for this module from a transaction
func NewOrmManagerFromTransaction(tx gorp.SqlExecutor) (*Manager, error) {
	m := &Manager{}
	err := m.Hijack(tx)
	if err != nil {
		return nil, err
	}

	return m, nil
}

// Initialize orm package
func (m *Manager) Initialize() {

	m.AddTableWithName(
		Asset{},
		AssetTableFull,
	).SetKeys(
		true,
		"ID",
	)

	m.AddTableWithName(
		CreativeDetail{},
		CreativeDetailTableFull,
	).SetKeys(
		false,
		"CreativeID",
		"HourID",
	)

	m.AddTableWithName(
		CreativeRejectReasons{},
		CreativeRejectReasonsTableFull,
	).SetKeys(
		true,
		"ID",
	)

	m.AddTableWithName(
		Creative{},
		CreativeTableFull,
	).SetKeys(
		true,
		"ID",
	)

}
func init() {
	mysql.Register(NewOrmManager())
}
