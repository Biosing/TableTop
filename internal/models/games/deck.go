package games

import (
	"github.com/google/uuid"
)

type Deck struct {
	ID          uuid.UUID      `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
	PlayerID    uuid.UUID      `gorm:"type:uuid;not null" json:"player_id"`
	ComboCardID uuid.UUID      `gorm:"type:uuid;not null" json:"combo_card_id"`
	StatusType  DeckStatusType `gorm:"type:integer;default:1" json:"status_type"`
}
