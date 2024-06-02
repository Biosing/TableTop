package game_sessions

import "github.com/google/uuid"

type ListResult struct {
	ID           uuid.UUID `json:"ID"`
	Name         string    `json:"name"`
	PlayerCounts int       `json:"player_counts"`
	MaxPlayers   int       `json:"max_players"`
}

type ListResponse struct {
	Results    []ListResult `json:"results"`
	TotalCount int          `json:"total_count"`
}
