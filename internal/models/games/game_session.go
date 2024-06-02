package games

import (
	"time"

	"github.com/google/uuid"
)

type GameSession struct {
	ID             uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
	Name           string    `gorm:"type:varchar(32);not null" json:"name"`
	StartGameDate  time.Time `gorm:"type:timestamp with time zone;default:null" json:"start_game_date"`
	FinishGameDate time.Time `gorm:"type:timestamp with time zone;default:null" json:"finish_game_date"`
	MaxPlayers     int       `gorm:"type:integer;not null;default:1" json:"max_players"`
	Players        []Player  `gorm:"foreignKey:GameSessionID" json:"players"`
}
