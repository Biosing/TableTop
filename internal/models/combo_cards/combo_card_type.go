package combo_cards

import (
	"database/sql/driver"
	"errors"
)

type ComboCardType int

const (
	Flex ComboCardType = iota
	CombatSkill
	Ability
)

func (cct ComboCardType) String() string {
	return [...]string{"Flex", "Combat skill", "Ability"}[cct]
}

func (cct ComboCardType) Value() (driver.Value, error) {
	return int(cct), nil
}

func (cct *ComboCardType) Scan(value interface{}) error {
	intValue, ok := value.(int64)
	if !ok {
		return errors.New("type assertion to int64 failed")
	}
	*cct = ComboCardType(intValue)
	return nil
}
