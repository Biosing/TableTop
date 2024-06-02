package combo_cards

import (
	"github.com/google/uuid"

	"table_top/internal/models/combo_cards"
)

type CreateRequest struct {
	CharacterID         uuid.UUID                 `json:"character_id"`
	Type                combo_cards.ComboCardType `json:"type"`
	Name                string                    `json:"name"`
	Description         string                    `json:"description"`
	TargetEnemyID       uuid.UUID                 `json:"target_enemy_id"`
	RequiredNumberCells int                       `json:"required_number_cells"`
	AddedNumberCells    int                       `json:"added_number_cells"`
	DamageComboCards    []DamageComboCardRequest  `json:"damage_combo_cards"`
	SpecialEffects      []SpecialEffectRequest    `json:"special_effects"`
}

type DamageComboCardRequest struct {
	RangeFrom int                             `json:"range_from"`
	RangeTo   int                             `json:"range_to"`
	Type      combo_cards.DamageComboCardType `json:"type"`
	Damage    int                             `json:"damage"`
}

type SpecialEffectRequest struct {
	Type        combo_cards.SpecialEffectType       `json:"type"`
	DamageType  combo_cards.SpecialEffectDamageType `json:"damage_type"`
	Damage      int                                 `json:"damage"`
	Description string                              `json:"description"`
}
