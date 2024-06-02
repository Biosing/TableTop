package services

import (
	"context"
	"errors"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	dtos "table_top/internal/dtos/game_session_players"
	models "table_top/internal/models/games"
	repositories "table_top/internal/repositories/games"
)

type GameSessionPlayerService interface {
	AddPlayer(ctx context.Context, req *dtos.AddPlayerRequest) error
	RemovePlayer(ctx context.Context, c *gin.Context, req *dtos.RemovePlayerRequest) error
}

type gameSessionPlayerService struct {
	repo repositories.GameSessionPlayerRepository
}

func NewGameSessionPlayerService(repo repositories.GameSessionPlayerRepository) GameSessionPlayerService {
	return &gameSessionPlayerService{repo: repo}
}

func (s *gameSessionPlayerService) AddPlayer(ctx context.Context, req *dtos.AddPlayerRequest) error {
	duplicateNickname, err := s.repo.CheckDuplicateNickname(ctx, req.GameSessionID, req.Nickname)
	if err != nil {
		return err
	}
	if duplicateNickname {
		return errors.New("nickname already exists in this session")
	}

	duplicateCharacterID, err := s.repo.CheckDuplicateCharacterID(ctx, req.GameSessionID, req.CharacterID)
	if err != nil {
		return err
	}
	if duplicateCharacterID {
		return errors.New("character already exists in this session")
	}

	player := &models.Player{
		GameSessionID: req.GameSessionID,
		CharacterID:   req.CharacterID,
		Nickname:      req.Nickname,
	}

	return s.repo.AddPlayer(ctx, player)
}

func (s *gameSessionPlayerService) RemovePlayer(ctx context.Context, c *gin.Context, req *dtos.RemovePlayerRequest) error {
	var gameSessionUUID uuid.UUID
	var err error

	if req.GameSessionID == nil {
		session := sessions.Default(c)
		gameSessionID := session.Get("game_session_id")
		if gameSessionID == nil {
			return errors.New("no game session found in the session")
		}

		gameSessionUUID, err = uuid.Parse(gameSessionID.(string))
		if err != nil {
			return errors.New("invalid game session ID format")
		}
	} else {
		gameSessionUUID = *req.GameSessionID
	}

	return s.repo.RemovePlayerBySessionIDAndNickname(ctx, gameSessionUUID, req.Nickname)
}
