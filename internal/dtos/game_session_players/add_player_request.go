package game_session_players

import "github.com/google/uuid"

type AddPlayerRequest struct {
	GameSessionID uuid.UUID `json:"game_session_id"`
	Nickname      string    `json:"nickname"`
	CharacterID   uuid.UUID `json:"character_id"`
}
