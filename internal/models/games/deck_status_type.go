package games

import (
	"database/sql/driver"
	"errors"
)

type DeckStatusType int

const (
	DeckStatusTypeInactive DeckStatusType = iota
	DeckStatusTypeActive1
)

func (dt DeckStatusType) String() string {
	return [...]string{"Inactive", "Active1"}[dt]
}

func (dt DeckStatusType) Value() (driver.Value, error) {
	return int(dt), nil
}

func (dt *DeckStatusType) Scan(value interface{}) error {
	intValue, ok := value.(int64)
	if !ok {
		return errors.New("type assertion to int64 failed")
	}
	*dt = DeckStatusType(intValue)
	return nil
}
