package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Character struct {
	ID           uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
	Name         string    `gorm:"size:128" json:"name"`
	Class        string    `gorm:"size:64" json:"class"`
	Race         string    `gorm:"size:128" json:"race"`
	Age          int       `json:"age"`
	Sex          string    `gorm:"size:64" json:"sex"`
	Description  string    `gorm:"size:2048" json:"description"`
	BeginningWay string    `gorm:"size:2048" json:"beginning_way"`
}

func (character *Character) BeforeCreate(tx *gorm.DB) (err error) {
	if character.ID == uuid.Nil {
		character.ID = uuid.New()
	}
	return
}
