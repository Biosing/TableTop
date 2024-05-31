package models

import "github.com/google/uuid"

type Deck struct {
	ID        uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
	PlayerID  uuid.UUID `gorm:"type:uuid" json:"player_id"`
	ComboCard uuid.UUID `gorm:"type:uuid" json:"combo_card"`
}
