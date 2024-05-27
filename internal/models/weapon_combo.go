package models

import "github.com/google/uuid"

type WeaponCombo struct {
	ID        uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
	WeaponID  uuid.UUID `gorm:"type:uuid" json:"weapon_id"`
	ComboType int       `json:"combo_type"`
	Count     int       `json:"count"`
	Order     int       `json:"order"`
}
