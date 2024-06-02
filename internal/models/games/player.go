package games

import (
	"github.com/google/uuid"
)

type Player struct {
	ID               uuid.UUID  `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
	UserID           uuid.UUID  `gorm:"type:uuid;default:null" json:"user_id"`
	GameSessionID    uuid.UUID  `gorm:"type:uuid;not null" json:"game_session_id"`
	Nickname         string     `gorm:"type:varchar(32);not null" json:"nickname"`
	CharacterID      uuid.UUID  `gorm:"type:uuid;default:null" json:"character_id"`
	EquippedWeaponID uuid.UUID  `gorm:"type:uuid;default:null" json:"equipped_weapon_id"`
	Experience       int        `gorm:"type:integer;default:0" json:"experience"`
	Level            int        `gorm:"type:integer;default:0" json:"level"`
	Health           int        `gorm:"type:integer;default:20" json:"health"`
	TurnOrder        int        `gorm:"type:integer;default:null" json:"turn_order"`
	Backpacks        []Backpack `gorm:"foreignKey:PlayerID" json:"backpacks"`
	Decks            []Deck     `gorm:"foreignKey:PlayerID" json:"decks"`
}
