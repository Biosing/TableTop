package models

import (
	"time"

	"github.com/google/uuid"
)

type GameSession struct {
	ID             uuid.UUID  `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
	Name           string     `gorm:"size:32; not null" json:"name"`
	StartGameDate  *time.Time `gorm:"default:null" json:"start_game_date"`
	FinishGameDate *time.Time `gorm:"default:null" json:"finish_game_date"`
	MaxPlayers     int        `gorm:"default:1; max:8; not null" json:"max_players"`
	Players        []Player   `json:"players"`
}
