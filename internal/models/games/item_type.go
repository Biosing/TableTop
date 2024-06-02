package games

import (
	"database/sql/driver"
	"errors"
)

type ItemType int

const (
	Weapon ItemType = iota
)

func (i ItemType) String() string {
	return [...]string{"Weapon"}[i]
}

func (i ItemType) Value() (driver.Value, error) {
	return int(i), nil
}

func (i *ItemType) Scan(value interface{}) error {
	intValue, ok := value.(int64)
	if !ok {
		return errors.New("type assertion to int64 failed")
	}
	*i = ItemType(intValue)
	return nil
}
