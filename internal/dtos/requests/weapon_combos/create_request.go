package weapon_combos

import "github.com/google/uuid"

type CreateRequest struct {
	WeaponID  uuid.UUID `json:"weapon_id"`
	ComboType int       `json:"combo_type"`
	Count     int       `json:"count"`
	Order     int       `json:"order"`
}
