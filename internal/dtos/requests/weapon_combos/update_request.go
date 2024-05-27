package weapon_combos

import (
	"github.com/google/uuid"

	models "table_top/internal/models/items"
)

type UpdateRequest struct {
	WeaponID uuid.UUID              `json:"weapon_id"`
	Type     models.WeaponComboType `json:"type"`
	Count    int                    `json:"count"`
	Order    int                    `json:"order"`
}
