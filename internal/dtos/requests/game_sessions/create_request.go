package game_sessions

import "github.com/google/uuid"

type CreateRequest struct {
	GameSessionName string    `json:"game_session_name"`
	Nickname        string    `json:"nickname"`
	CharacterID     uuid.UUID `json:"character_id"`
	MaxPlayers      int       `json:"max_players"`
}
