package models

import (
	"database/sql/driver"
	"errors"
)

type SpecialEffectType int

const (
	Variable SpecialEffectType = iota
)

func (s SpecialEffectType) String() string {
	return [...]string{"Variable"}[s]
}

func (s SpecialEffectType) Value() (driver.Value, error) {
	return int(s), nil
}

func (s *SpecialEffectType) Scan(value interface{}) error {
	intValue, ok := value.(int64)
	if !ok {
		return errors.New("type assertion to int64 failed")
	}
	*s = SpecialEffectType(intValue)
	return nil
}
