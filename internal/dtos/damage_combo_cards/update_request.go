package damage_combo_cards

import (
	"github.com/google/uuid"

	"table_top/internal/models/combo_cards"
)

type UpdateRequest struct {
	ComboCardID uuid.UUID                       `json:"combo_card_id"`
	RangeFrom   int                             `json:"range_from"`
	RangeTo     int                             `json:"range_to"`
	Type        combo_cards.DamageComboCardType `json:"type"`
	Damage      int                             `json:"damage"`
}
