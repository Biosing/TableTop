package models

import "github.com/google/uuid"

type Backpack struct {
	ID            uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"json:"id"`
	GameSessionID uuid.UUID `gorm:"type:uuid" json:"game_session_id"`
	ItemId        uuid.UUID `gorm:"type:uuid" json:"item_id"`
	ItemType      ItemType  `json:"item_type"`
}