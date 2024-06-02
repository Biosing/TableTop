package combo_cards

import (
	"github.com/google/uuid"
)

type DamageComboCard struct {
	ID          uuid.UUID           `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
	ComboCardID uuid.UUID           `gorm:"type:uuid;not null" json:"combo_card_id"`
	RangeFrom   int                 `gorm:"type:integer;not null" json:"range_from"`
	RangeTo     int                 `gorm:"type:integer;not null" json:"range_to"`
	Type        DamageComboCardType `gorm:"type:integer;not null" json:"type"`
	Damage      int                 `gorm:"type:integer;default:0" json:"damage"`
}
