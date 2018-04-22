package orm

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/clickyab/services/assert"

	"github.com/clickyab/services/mysql"
)

// Info main data of select query
// @Model {
//		table = -
//		primary = false, keyfield1
// }
type Info struct {
	Data1 mysql.NullInt64 `db:"data1"`
	Data2 mysql.NullInt64 `db:"data2"`
	Data3 mysql.NullInt64 `db:"data3"`
	Data4 mysql.NullInt64 `db:"data4"`

	KeyField1 mysql.NullInt64 `db:"keyfield1"`
	KeyField2 mysql.NullInt64 `db:"keyfield2"`
	KeyField3 mysql.NullInt64 `db:"keyfield3"`
	KeyField4 mysql.NullInt64 `db:"keyfield4"`
}

// Queries is all need query to check and isert or update data
type Queries struct {
	CheckQuery  string
	InsertQuery string
	UpdateQuery string
	Params      []QueryParams
}

// QueryParams all params for queries
type QueryParams struct {
	DataFields []interface{}
	KeyFields  []interface{}
	All        []interface{}
}

type checkedQueryRes struct {
	Ok int64 `db:"ok"`
}

// FindInformation to run select query and find information
func (m *Manager) FindInformation(query string) []Info {
	var res []Info

	_, err := m.GetRDbMap().Select(
		&res,
		query,
	)
	assert.Nil(err)

	return res
}

// GenerateQueries generate check, insert and update queries base on key fiedls, target fields and target table
func GenerateQueries(targetTable string, keyFields, targetFields []string, foundData []Info) Queries {
	var keysE, targetsE, args []string

	for _, k := range keyFields {
		keysE = append(keysE, k+"=?")
		args = append(args, "?")
	}

	for _, k := range targetFields {
		targetsE = append(targetsE, k+"=?")
		args = append(args, "?")
	}

	cq := fmt.Sprintf("select 1 AS ok from %s where %s", targetTable, strings.Join(keysE, " and "))
	iq := fmt.Sprintf("insert into %s (%s,%s) values (%s)", targetTable, strings.Join(targetFields, ","), strings.Join(keyFields, ","), strings.Join(args, ","))
	uq := fmt.Sprintf("update %s set %s where %s", targetTable, strings.Join(targetsE, ","), strings.Join(keysE, " and "))

	var df, kf []interface{}
	params := make([]QueryParams, len(foundData))
	for k, data := range foundData {
		df = GetDataFields(data)
		kf = GetKeyFields(data)

		params[k] = QueryParams{
			DataFields: df,
			KeyFields:  kf,
			All:        append(kf, df...),
		}
	}

	qs := Queries{
		CheckQuery:  cq,
		InsertQuery: iq,
		UpdateQuery: uq,
		Params:      params,
	}

	return qs
}

// RowExists check if row exists base on key fields
func (m *Manager) RowExists(query string, args QueryParams) bool {
	var res checkedQueryRes

	err := m.GetWDbMap().SelectOne(&res, query, args.KeyFields...)

	if err == sql.ErrNoRows {
		return false
	}
	assert.Nil(err)

	return true
}

// RunQueryWithParams run update or insert query with params
func (m *Manager) RunQueryWithParams(query string, args QueryParams) (sql.Result, error) {
	res, err := m.GetWDbMap().Exec(query, args.All...)

	return res, err
}

// GetKeyFields get specified key fields found data
func GetKeyFields(args Info) []interface{} {
	var argsList []interface{}

	if args.KeyField1.Valid {
		argsList = append(argsList, args.KeyField1.Int64)
	}
	if args.KeyField2.Valid {
		argsList = append(argsList, args.KeyField2.Int64)
	}
	if args.KeyField3.Valid {
		argsList = append(argsList, args.KeyField3.Int64)
	}
	if args.KeyField4.Valid {
		argsList = append(argsList, args.KeyField4.Int64)
	}

	return argsList
}

// GetDataFields get specified data fields found data
func GetDataFields(args Info) []interface{} {
	var argsList []interface{}

	if args.Data1.Valid {
		argsList = append(argsList, args.Data1.Int64)
	}
	if args.Data2.Valid {
		argsList = append(argsList, args.Data2.Int64)
	}
	if args.Data3.Valid {
		argsList = append(argsList, args.Data3.Int64)
	}
	if args.Data4.Valid {
		argsList = append(argsList, args.Data4.Int64)
	}

	return argsList
}
