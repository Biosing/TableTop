package enemies

import (
	"github.com/google/uuid"
)

type EnemyMove struct {
	ID          uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
	EnemyID     uuid.UUID `gorm:"type:uuid;not null" json:"enemy_id"`
	RangeFrom   int       `gorm:"type:integer;not null" json:"range_from"`
	RangeTo     int       `gorm:"type:integer;not null" json:"range_to"`
	Description string    `gorm:"type:varchar(2048);not null" json:"description"`
}
