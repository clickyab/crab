package errors

import (
	"github.com/clickyab/services/gettext/t9e"
)

var (
	// KeyFieldsRequired error
	KeyFieldsRequired error = t9e.G("key fields is required. you can set key fields in command like: keyfields=x or keyfields=x,y,z, min is 1 and max is 4")
	// TargetFieldsRequired error
	TargetFieldsRequired error = t9e.G("target fields is required. you can set key fields in command like: targetfields=x or targetfields=x,y,z, min is 1 and max is 4")
	// SelectFieldMin query with no selected field!
	SelectFieldMin error = t9e.G("invalid query! at least query shoul select 2 fields.")
	// SelectedFieldsCount when selected field != count of needed fields
	SelectedFieldsCount error = t9e.G("invalid query! count of selected fields should be equal of sum number of key fiedls and target fields")
	// QueryConditionRequired error
	QueryConditionRequired error = t9e.G("query where condition is required.")
)

// NotFoundFieldError maker
func NotFoundFieldError(key string) error {
	if key != "" {
		return t9e.G("required field %s not found. at least required to select as data1 and keyfield1", key)
	}

	return t9e.G("required field %s not found. at least required to select as data1 and keyfield1.")
}
