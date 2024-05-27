package models

import "github.com/google/uuid"

type DamageComboCard struct {
	ID          uuid.UUID           `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
	ComboCardID uuid.UUID           `gorm:"type:uuid" json:"combo_card_id"`
	RangeFrom   int                 `json:"range_from"`
	RangeTo     int                 `json:"range_to"`
	Type        DamageComboCardType `json:"type"`
	Damage      int                 `json:"damage"`
}
