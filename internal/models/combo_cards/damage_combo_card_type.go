package models

import (
	"database/sql/driver"
	"errors"
)

type DamageComboCardType int

const (
	PhysicalDamage DamageComboCardType = iota
)

func (d DamageComboCardType) String() string {
	return [...]string{"Physical damage"}[d]
}

func (d DamageComboCardType) Value() (driver.Value, error) {
	return int(d), nil
}

func (d *DamageComboCardType) Scan(value interface{}) error {
	intValue, ok := value.(int64)
	if !ok {
		return errors.New("type assertion to int64 failed")
	}
	*d = DamageComboCardType(intValue)
	return nil
}
