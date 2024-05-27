package weapons

import "github.com/google/uuid"

type CreateRequest struct {
	CharacterID uuid.UUID     `json:"character_id"`
	Name        string        `json:"name"`
	CountCards  int           `json:"count_cards"`
	Defense     int           `json:"defense"`
	WeaponCombo []WeaponCombo `json:"weapon_combo"`
}

type WeaponCombo struct {
	ComboType int `json:"combo_type"`
	Count     int `json:"count"`
	Order     int `json:"order"`
}
