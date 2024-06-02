package items

import (
	"github.com/google/uuid"
)

type WeaponCombo struct {
	ID       uuid.UUID       `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
	WeaponID uuid.UUID       `gorm:"type:uuid;not null" json:"weapon_id"`
	Type     WeaponComboType `gorm:"type:integer;not null" json:"type"`
	Count    int             `gorm:"type:integer;not null" json:"count"`
	Order    int             `gorm:"type:integer;not null" json:"order"`
}
