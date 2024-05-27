package models

import "github.com/google/uuid"

type SpecialEffect struct {
	ID          uuid.UUID               `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
	ComboCardID uuid.UUID               `gorm:"type:uuid" json:"combo_card_id"`
	Type        SpecialEffectType       `json:"type"`
	DamageType  SpecialEffectDamageType `json:"damage_type"`
	Damage      int                     `json:"damage"`
	Description string                  `gorm:"size:2048" json:"description"`
}
