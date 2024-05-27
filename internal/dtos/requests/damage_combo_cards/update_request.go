package damage_combo_cards

import "github.com/google/uuid"

type UpdateRequest struct {
	ComboCardID uuid.UUID `json:"combo_card_id"`
	RangeFrom   int       `json:"range_from"`
	RangeTo     int       `json:"range_to"`
	DamageType  int       `json:"damage_type"`
	Damage      int       `json:"damage"`
}
