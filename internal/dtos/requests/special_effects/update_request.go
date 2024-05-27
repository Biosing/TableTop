package special_effects

import "github.com/google/uuid"

type UpdateRequest struct {
	ComboCardID       uuid.UUID `json:"combo_card_id"`
	SpecialEffectType int       `json:"special_effect_type"`
	DamageType        int       `json:"damage_type"`
	Damage            int       `json:"damage"`
	Description       string    `json:"description"`
}
