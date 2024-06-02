package combo_cards

import (
	"github.com/google/uuid"

	"table_top/internal/models/combo_cards"
)

type UpdateRequest struct {
	CharacterID         uuid.UUID                 `json:"character_id"`
	Type                combo_cards.ComboCardType `json:"type"`
	Name                string                    `json:"name"`
	Description         string                    `json:"description"`
	TargetEnemyID       uuid.UUID                 `json:"target_enemy_id"`
	RequiredNumberCells int                       `json:"required_number_cells"`
	AddedNumberCells    int                       `json:"added_number_cells"`
}
