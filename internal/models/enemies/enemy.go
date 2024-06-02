package enemies

import (
	"github.com/google/uuid"
)

type Enemy struct {
	ID           uuid.UUID   `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
	Name         string      `gorm:"type:varchar(128);not null" json:"name"`
	Level        int         `gorm:"type:integer;default:1" json:"level"`
	Class        string      `gorm:"type:varchar(128);not null" json:"class"`
	Description  string      `gorm:"type:varchar(2048);default:null" json:"description"`
	HP           int         `gorm:"type:integer;default:1" json:"hp"`
	Experience   int         `gorm:"type:integer;default:0" json:"experience"`
	QuantityDeck int         `gorm:"type:integer;default:1" json:"quantity_deck"`
	Defense      int         `gorm:"type:integer;default:0" json:"defense"`
	EnemyMoves   []EnemyMove `gorm:"foreignKey:EnemyID" json:"enemy_moves"`
}
