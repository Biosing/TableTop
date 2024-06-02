package weapons

import "github.com/google/uuid"

type UpdateRequest struct {
	CharacterID uuid.UUID `json:"character_id"`
	Name        string    `json:"name"`
	CountCards  int       `json:"count_cards"`
	Defense     int       `json:"defense"`
}
