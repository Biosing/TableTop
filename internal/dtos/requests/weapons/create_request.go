package weapons

import (
	"github.com/google/uuid"

	models "table_top/internal/models/items"
)

type CreateRequest struct {
	CharacterID uuid.UUID     `json:"character_id"`
	Name        string        `json:"name"`
	CountCards  int           `json:"count_cards"`
	Defense     int           `json:"defense"`
	WeaponCombo []WeaponCombo `json:"weapon_combo"`
}

type WeaponCombo struct {
	Type  models.WeaponComboType `json:"type"`
	Count int                    `json:"count"`
	Order int                    `json:"order"`
}
