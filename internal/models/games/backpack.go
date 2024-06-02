package games

import (
	"github.com/google/uuid"
)

type Backpack struct {
	ID       uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
	PlayerID uuid.UUID `gorm:"type:uuid;not null" json:"player_id"`
	ItemID   uuid.UUID `gorm:"type:uuid;not null" json:"item_id"`
	ItemType int       `gorm:"type:integer;not null" json:"item_type"`
}
