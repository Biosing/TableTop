package models

import "github.com/google/uuid"

type Weapon struct {
	ID          uuid.UUID     `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
	CharacterID uuid.UUID     `gorm:"type:uuid" json:"character_id"`
	Name        string        `gorm:"size:128" json:"name"`
	CountCards  int           `json:"count_cards"`
	Defense     int           `json:"defense"`
	WeaponCombo []WeaponCombo `gorm:"foreignKey:WeaponID" json:"weapon_combo"`
}
