package orm

import (
	"database/sql/driver"
	"time"

	"strings"

	"strconv"

	"github.com/clickyab/services/trans"
)

type base struct {
	ID        int64     `json:"id" db:"id"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

type Int64Array []int64

func (sa *Int64Array) Value() (driver.Value, error) {
	return []int64(sa), nil
}

func (sa *Int64Array) Scan(src interface{}) error {
	var b []byte
	switch src.(type) {
	case []byte:
		b = src.([]byte)
	case string:
		b = []byte(src.(string))
	case nil:
		b = make([]byte, 0)
	default:
		return trans.E("unsupported type")
	}

	intArrayInString := strings.Split(string(b), ",")
	var ints = make([]int64, len(intArrayInString))
	for i := range intArrayInString {
		val, err := strconv.ParseInt(intArrayInString[i], 10, 0)
		if err != nil {
			return trans.E("unsupported value in int array")
		}

		ints = append(ints, val)
	}

	*sa = Int64Array(ints)

	return nil
}
