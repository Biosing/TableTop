package models

import (
	"database/sql/driver"
	"errors"
)

type SpecialEffectDamageType int

const (
	EffectOnly SpecialEffectDamageType = iota
)

func (s SpecialEffectDamageType) String() string {
	return [...]string{"EffectOnly"}[s]
}

func (s SpecialEffectDamageType) Value() (driver.Value, error) {
	return int(s), nil
}

func (s *SpecialEffectDamageType) Scan(value interface{}) error {
	intValue, ok := value.(int64)
	if !ok {
		return errors.New("type assertion to int64 failed")
	}
	*s = SpecialEffectDamageType(intValue)
	return nil
}
