package models

import "github.com/google/uuid"

type EnemyMove struct {
	ID          uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
	EnemyID     uuid.UUID `gorm:"type:uuid" json:"enemy_id"`
	RangeFrom   int       `json:"range_from" example:"1"`
	RangeTo     int       `json:"range_to" example:"3"`
	Description string    `gorm:"size:2048" json:"description" example:"Slash attack"`
}
