package characters

import (
	"github.com/google/uuid"

	"table_top/internal/models/combo_cards"
	"table_top/internal/models/items"
)

type Character struct {
	ID           uuid.UUID               `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
	Name         string                  `gorm:"type:varchar(128);not null" json:"name"`
	Class        string                  `gorm:"type:varchar(64);not null" json:"class"`
	Race         string                  `gorm:"type:varchar(128);not null" json:"race"`
	Age          int                     `gorm:"type:integer;not null" json:"age"`
	Sex          string                  `gorm:"type:varchar(64);not null" json:"sex"`
	Description  string                  `gorm:"type:varchar(2048);not null" json:"description"`
	BeginningWay string                  `gorm:"type:varchar(2048);not null" json:"beginning_way"`
	Weapons      []items.Weapon          `gorm:"foreignKey:CharacterID" json:"weapons"`
	ComboCards   []combo_cards.ComboCard `gorm:"foreignKey:CharacterID" json:"combo_cards"`
}
