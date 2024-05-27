package models

import (
	"github.com/google/uuid"
)

type Location struct {
	ID            uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
	Name          string    `gorm:"size:128" json:"name"`
	Level         int       `json:"level"`
	DangerLevel   int       `json:"danger_level"`
	Description   string    `gorm:"size:2048" json:"description"`
	MonstersCount int       `json:"monsters_count"`
}
