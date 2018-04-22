package workers

import (
	"strings"

	"github.com/clickyab/services/array"
	"github.com/sirupsen/logrus"

	"github.com/clickyab/services/assert"

	"clickyab.com/crab/modules/dataworker/errors"
	"clickyab.com/crab/modules/dataworker/orm"
)

// Run command main process
func Run(findQuery, targetTable string, keyFields, targetFields []string) {
	db := orm.NewOrmManager()

	err := validate(findQuery, keyFields, targetFields)
	assert.Nil(err)

	logrus.Debugf("try to find data and update fields %s and %s \n", targetFields[0], targetFields[1])
	info := db.FindInformation(findQuery)

	logrus.Debug("\n______________ we found _________________\n")
	logrus.Debugf("ok, we found %d result. try to update target fields", len(info))

	queries := orm.GenerateQueries(targetTable, keyFields, targetFields, info)

	logrus.Debug("\n______________ we generate _________________\n")
	logrus.Debugf("@@@@ check query: %s \n", queries.CheckQuery)
	logrus.Debugf("@@@@ insert query: %s \n", queries.InsertQuery)
	logrus.Debugf("@@@@ update query: %s \n", queries.UpdateQuery)

	logrus.Debug("\n______________ try to insert/update _________________\n")
	ucnt, icnt := 0, 0
	for i := range info {
		if db.RowExists(queries.CheckQuery, queries.Params[i]) {
			_, err := db.RunQueryWithParams(queries.UpdateQuery, queries.Params[i])
			assert.Nil(err)
			ucnt++
		} else {
			_, err := db.RunQueryWithParams(queries.InsertQuery, queries.Params[i])
			assert.Nil(err)
			icnt++
		}
	}

	logrus.Debug("\n______________ finally we run _________________\n")
	logrus.Debugf("we update %d rows and insert %d new row data. \n", ucnt, icnt)
}

func validate(findQuery string, keyFields, targetFields []string) error {
	if len(keyFields) == 0 {
		return errors.KeyFieldsRequired
	}

	if len(targetFields) == 0 {
		return errors.TargetFieldsRequired
	}

	qparts := getQueryParts(findQuery)
	if len(qparts.SelectFields) < 2 {
		return errors.SelectFieldMin
	}

	if len(keyFields)+len(targetFields) != len(qparts.SelectFields) {
		return errors.SelectedFieldsCount
	}

	if !array.StringInArray("data1", qparts.SelectFields...) {
		return errors.NotFoundFieldError("data1")
	}

	if !array.StringInArray("keyfield1", qparts.SelectFields...) {
		return errors.NotFoundFieldError("keyfield1")
	}

	if !qparts.HasWhere {
		return errors.QueryConditionRequired
	}
	return nil
}

type queryParts struct {
	SelectFields []string
	HasWhere     bool
}

func getQueryParts(query string) queryParts {
	query = strings.Replace(query, ",", " ", -1)
	parts := strings.Split(query, " ")

	res := queryParts{}
	for i, val := range parts {
		if strings.ToUpper(val) == "AS" {
			res.SelectFields = append(res.SelectFields, parts[i+1])
		}

		if strings.ToUpper(val) == "WHERE" {
			res.HasWhere = true
		}
	}

	return res
}
