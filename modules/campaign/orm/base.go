package orm

import (
	"database/sql/driver"
	"time"

	"encoding/json"

	"github.com/clickyab/services/trans"
)

type base struct {
	ID        int64     `json:"id" db:"id"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

type stringArray []string

func (sa *stringArray) Value() (driver.Value, error) {
	return []string(sa), nil
}

func (sa *stringArray) Scan(src interface{}) error {
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

	err := json.Unmarshal(b, sa)
	if err != nil {
		return trans.EE(err)
	}

	return nil
}
