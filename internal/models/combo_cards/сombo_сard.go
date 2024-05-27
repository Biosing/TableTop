package models

import (
	"github.com/google/uuid"
)

type ComboCard struct {
	ID                  uuid.UUID         `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
	CharacterID         uuid.UUID         `gorm:"type:uuid" json:"character_id"`
	Type                ComboCardType     `json:"type"`
	Name                string            `gorm:"size:128" json:"name"`
	Description         string            `gorm:"size:2048" json:"description"`
	TargetEnemyID       uuid.UUID         `gorm:"type:uuid" json:"target_enemy_id"`
	RequiredNumberCells int               `json:"required_number_cells"`
	AddedNumberCells    int               `json:"added_number_cells"`
	DamageComboCards    []DamageComboCard `json:"damage_combo_cards"`
	SpecialEffects      []SpecialEffect   `json:"special_effects"`
}
