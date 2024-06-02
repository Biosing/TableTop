package special_effects

import (
	"github.com/google/uuid"

	"table_top/internal/models/combo_cards"
)

type UpdateRequest struct {
	ComboCardID uuid.UUID                           `json:"combo_card_id"`
	Type        combo_cards.SpecialEffectType       `json:"type"`
	DamageType  combo_cards.SpecialEffectDamageType `json:"damage_type"`
	Damage      int                                 `json:"damage"`
	Description string                              `json:"description"`
}
