package combo_cards

import "github.com/google/uuid"

type CreateRequest struct {
	CharacterID         uuid.UUID         `json:"character_id"`
	ComboCardType       int               `json:"combo_card_type"`
	Name                string            `json:"name"`
	Description         string            `json:"description"`
	TargetEnemyID       uuid.UUID         `json:"target_enemy_id"`
	RequiredNumberCells int               `json:"required_number_cells"`
	AddedNumberCells    int               `json:"added_number_cells"`
	DamageComboCards    []DamageComboCard `json:"damage_combo_cards"`
	SpecialEffects      []SpecialEffect   `json:"special_effects"`
}

type DamageComboCard struct {
	RangeFrom  int `json:"range_from"`
	RangeTo    int `json:"range_to"`
	DamageType int `json:"damage_type"`
	Damage     int `json:"damage"`
}

type SpecialEffect struct {
	SpecialEffectType int    `json:"special_effect_type"`
	DamageType        int    `json:"damage_type"`
	Damage            int    `json:"damage"`
	Description       string `json:"description"`
}
