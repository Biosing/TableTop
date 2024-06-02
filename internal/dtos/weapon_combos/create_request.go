package weapon_combos

import (
	"github.com/google/uuid"

	"table_top/internal/models/items"
)

type CreateRequest struct {
	WeaponID uuid.UUID             `json:"weapon_id"`
	Type     items.WeaponComboType `json:"type"`
	Count    int                   `json:"count"`
	Order    int                   `json:"order"`
}
