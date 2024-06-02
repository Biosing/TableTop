package combo_cards

import (
	"github.com/google/uuid"
)

type SpecialEffect struct {
	ID          uuid.UUID               `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
	ComboCardID uuid.UUID               `gorm:"type:uuid;not null" json:"combo_card_id"`
	Type        SpecialEffectType       `gorm:"type:integer;not null" json:"type"`
	DamageType  SpecialEffectDamageType `gorm:"type:integer;not null" json:"damage_type"`
	Damage      int                     `gorm:"type:integer;not null" json:"damage"`
	Description string                  `gorm:"type:varchar(2048);not null" json:"description"`
}
