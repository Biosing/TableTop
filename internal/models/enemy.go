package models

import "github.com/google/uuid"

type Enemy struct {
	ID           uuid.UUID   `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
	Name         string      `gorm:"size:128" json:"name" example:"Goblin"`
	Level        int         `json:"level" example:"1"`
	Class        string      `gorm:"size:128" json:"class" example:"Warrior"`
	Description  string      `gorm:"size:2048" json:"description" example:"A small but fierce warrior"`
	Hp           int         `json:"hp" example:"100"`
	Experience   int         `json:"experience" example:"10"`
	QuantityDeck int         `json:"quantity_deck" example:"2"`
	Defense      int         `json:"defense" example:"2"`
	EnemyMoves   []EnemyMove `gorm:"foreignKey:EnemyID" json:"enemy_moves"`
}
