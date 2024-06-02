package game_session_players

import "github.com/google/uuid"

type RemovePlayerRequest struct {
	GameSessionID *uuid.UUID `json:"game_session_id,omitempty"`
	Nickname      string     `json:"nickname"`
}
