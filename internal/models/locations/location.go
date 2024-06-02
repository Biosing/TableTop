package locations

import (
	"github.com/google/uuid"
)

type Location struct {
	ID            uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
	Name          string    `gorm:"type:varchar(128);not null" json:"name"`
	Level         int       `gorm:"type:integer;default:1" json:"level"`
	DangerLevel   int       `gorm:"type:integer;default:0" json:"danger_level"`
	Description   string    `gorm:"type:varchar(2048);not null" json:"description"`
	MonstersCount int       `gorm:"type:integer;default:0" json:"monsters_count"`
}
