package users

import (
	"time"

	"github.com/google/uuid"

	"table_top/internal/models/games"
)

type User struct {
	ID        uuid.UUID      `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
	Login     string         `gorm:"type:varchar(128);not null" json:"login"`
	Email     string         `gorm:"type:varchar(128);default:null" json:"email"`
	Phone     string         `gorm:"type:varchar(64);default:null" json:"phone"`
	Password  string         `gorm:"type:varchar(255);not null" json:"password"`
	CreatedAt time.Time      `gorm:"type:timestamp with time zone;default:CURRENT_TIMESTAMP;not null" json:"created_at"`
	UpdatedAt time.Time      `gorm:"type:timestamp with time zone;default:null" json:"updated_at"`
	DeletedAt time.Time      `gorm:"type:timestamp with time zone;default:null" json:"deleted_at"`
	Players   []games.Player `gorm:"foreignKey:UserID" json:"players"`
}
