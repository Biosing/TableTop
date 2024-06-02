package items

import (
	"github.com/google/uuid"
)

type Weapon struct {
	ID           uuid.UUID     `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
	CharacterID  uuid.UUID     `gorm:"type:uuid;not null" json:"character_id"`
	Name         string        `gorm:"type:varchar(128);not null" json:"name"`
	CountCards   int           `gorm:"type:integer;not null" json:"count_cards"`
	Defense      int           `gorm:"type:integer;default:0" json:"defense"`
	WeaponCombos []WeaponCombo `gorm:"foreignKey:WeaponID" json:"combos"`
}
