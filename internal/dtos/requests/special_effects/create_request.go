package special_effects

import (
	"github.com/google/uuid"

	models "table_top/internal/models/combo_cards"
)

type CreateRequest struct {
	ComboCardID uuid.UUID                      `json:"combo_card_id"`
	Type        models.SpecialEffectType       `json:"type"`
	DamageType  models.SpecialEffectDamageType `json:"damage_type"`
	Damage      int                            `json:"damage"`
	Description string                         `json:"description"`
}
