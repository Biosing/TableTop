package models

import (
	"database/sql/driver"
	"errors"
)

type WeaponComboType int

const (
	FlexSlot WeaponComboType = iota
	CombatSkillSlot
	AbilitySlot
)

func (ct WeaponComboType) String() string {
	return [...]string{"Flex slot", "Combat skill slot", "Ability slot"}[ct]
}

func (ct WeaponComboType) Value() (driver.Value, error) {
	return int(ct), nil
}

func (ct *WeaponComboType) Scan(value interface{}) error {
	intValue, ok := value.(int64)
	if !ok {
		return errors.New("type assertion to int64 failed")
	}
	*ct = WeaponComboType(intValue)
	return nil
}
