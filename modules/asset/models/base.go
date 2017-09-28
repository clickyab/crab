package models

import (
	"fmt"
	"time"
)

type base struct {
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
	Active    bool      `json:"active" db:"active"`
}

type baseName struct {
	base
	Name string `json:"name" db:"name"`
}

func (m *Manager) allActive(o interface{}, t string) error {
	q := fmt.Sprintf("SELECT * FROM %s WHERE active=?", t)
	return m.GetRDbMap().SelectOne(&o, q, true)
}
