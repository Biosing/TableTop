package models

import "github.com/google/uuid"

type Player struct {
	ID               uuid.UUID  `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
	GameSessionID    uuid.UUID  `gorm:"type:uuid; not null" json:"game_session_id"`
	Nickname         string     `gorm:"size:32; not null" json:"nickname"`
	CharacterID      uuid.UUID  `gorm:"type:uuid; not null" json:"character_id"`
	EquippedWeaponID *uuid.UUID `gorm:"type:uuid;default:null" json:"equipped_weapon_id"`
	Experience       int        `gorm:"not null; default:0; min:0;" json:"experience;"`
	Level            int        `gorm:"not null; default:0; min:0" json:"level"`
	Health           int        `gorm:"not null; default:20;max:30;" json:"health"`
	Deck             []Deck     `json:"deck"`
	Backpack         []Backpack `json:"backpack"`
	TurnOrder        int        `gorm:"not null; default:0; min:0" json:"turn_order"`
}
