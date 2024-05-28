package models

import (
	"time"

	"github.com/google/uuid"
)

type GameSession struct {
	ID               uuid.UUID  `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
	Nickname         string     `gorm:"size:32" json:"nickname"`
	CharacterID      uuid.UUID  `gorm:"type:uuid;" json:"character_id"`
	EquippedWeaponID *uuid.UUID `gorm:"type:uuid;default:null" json:"equipped_weapon_id"`
	Experience       int        `json:"experience"`
	Level            int        `json:"level"`
	Health           int        `gorm:"default:20;max:30;" json:"health"`
	Deck             []Deck     `json:"deck"`
	Backpack         []Backpack `json:"backpack"`
	StartGameDate    *time.Time `gorm:"default:null" json:"start_game_date"`
	FinishGameDate   *time.Time `gorm:"default:null" json:"finish_game_date"`
}
