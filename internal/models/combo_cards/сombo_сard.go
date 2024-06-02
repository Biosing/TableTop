package combo_cards

import (
	"github.com/google/uuid"
)

type ComboCard struct {
	ID                  uuid.UUID         `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
	CharacterID         uuid.UUID         `gorm:"type:uuid;not null" json:"character_id"`
	Type                ComboCardType     `gorm:"type:integer;not null" json:"type"`
	Name                string            `gorm:"type:varchar(128);not null" json:"name"`
	Description         string            `gorm:"type:varchar(2048);not null" json:"description"`
	TargetEnemyID       uuid.UUID         `gorm:"type:uuid;not null" json:"target_enemy_id"`
	RequiredNumberCells int               `gorm:"type:integer;not null" json:"required_number_cells"`
	AddedNumberCells    int               `gorm:"type:integer;not null" json:"added_number_cells"`
	DamageComboCards    []DamageComboCard `gorm:"foreignKey:ComboCardID" json:"damage_combo_cards"`
	SpecialEffects      []SpecialEffect   `gorm:"foreignKey:ComboCardID" json:"special_effects"`
}
