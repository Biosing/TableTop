package game_sessions

import "github.com/google/uuid"

type CreateRequest struct {
	Nickname    string    `json:"nickname"`
	CharacterID uuid.UUID `json:"character_id"`
}
