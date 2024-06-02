package services

import (
	"context"
	"errors"
	"net/http"
	"time"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	dtos "table_top/internal/dtos/game_sessions"
	models "table_top/internal/models/games"
	repositories "table_top/internal/repositories/games"
)

type GameSessionService interface {
	CreateNewGame(ctx context.Context, c *gin.Context, req *dtos.CreateRequest) (*models.GameSession, error)
	GameSession(ctx context.Context, c *gin.Context) (*models.GameSession, error)
	StartGame(ctx context.Context, c *gin.Context) error
	FinishGame(ctx context.Context, c *gin.Context) error
	ListGameSessions(ctx context.Context, req *dtos.ListRequest) (*dtos.ListResponse, error)
}

type gameSessionService struct {
	repo repositories.GameSessionRepository
}

func NewGameSessionService(repo repositories.GameSessionRepository) GameSessionService {
	return &gameSessionService{repo: repo}
}

func (s *gameSessionService) CreateNewGame(ctx context.Context, c *gin.Context, req *dtos.CreateRequest) (*models.GameSession, error) {
	session := sessions.Default(c)

	gameSessionID := uuid.New()

	gameSession := &models.GameSession{
		ID:         gameSessionID,
		Name:       req.GameSessionName,
		MaxPlayers: req.MaxPlayers,
	}

	player := models.Player{
		Nickname:    req.Nickname,
		CharacterID: req.CharacterID,
	}
	gameSession.Players = append(gameSession.Players, player)

	if err := s.repo.Create(ctx, gameSession); err != nil {
		return nil, err
	}

	session.Set("game_session_id", gameSessionID.String())
	if err := session.Save(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save session"})
		return nil, errors.New("Failed to save session")
	}

	return gameSession, nil
}

func (s *gameSessionService) GameSession(ctx context.Context, c *gin.Context) (*models.GameSession, error) {
	session := sessions.Default(c)
	gameSessionID := session.Get("game_session_id")
	if gameSessionID == nil {
		return nil, errors.New("no game session found in the session")
	}

	gameSessionUUID, err := uuid.Parse(gameSessionID.(string))
	if err != nil {
		return nil, errors.New("invalid game session ID format")
	}

	gameSession, err := s.repo.GetByID(ctx, gameSessionUUID)
	if err != nil {
		return nil, err
	}

	return gameSession, nil
}

func (s *gameSessionService) StartGame(ctx context.Context, c *gin.Context) error {
	session := sessions.Default(c)
	gameSessionID := session.Get("game_session_id")
	if gameSessionID == nil {
		return errors.New("no game session found in the session")
	}

	gameSessionUUID, err := uuid.Parse(gameSessionID.(string))
	if err != nil {
		return errors.New("invalid game session ID format")
	}

	gameSession, err := s.repo.GetByID(ctx, gameSessionUUID)
	if err != nil {
		return err
	}

	now := time.Now()
	gameSession.StartGameDate = now

	if err := s.repo.Update(ctx, gameSession); err != nil {
		return err
	}

	return nil
}

func (s *gameSessionService) FinishGame(ctx context.Context, c *gin.Context) error {
	session := sessions.Default(c)
	gameSessionID := session.Get("game_session_id")
	if gameSessionID == nil {
		return errors.New("no game session found in the session")
	}

	gameSessionUUID, err := uuid.Parse(gameSessionID.(string))
	if err != nil {
		return errors.New("invalid game session ID format")
	}

	gameSession, err := s.repo.GetByID(ctx, gameSessionUUID)
	if err != nil {
		return err
	}

	now := time.Now()
	gameSession.FinishGameDate = now

	if err := s.repo.Update(ctx, gameSession); err != nil {
		return err
	}

	session.Delete("game_session_id")
	if err := session.Save(); err != nil {
		return errors.New("failed to save session")
	}

	return nil
}

func (s *gameSessionService) ListGameSessions(ctx context.Context, req *dtos.ListRequest) (*dtos.ListResponse, error) {
	offset := (req.Page - 1) * req.Limit

	gameSessions, totalCount, err := s.repo.List(ctx, req.Name, req.Limit, offset)
	if err != nil {
		return nil, err
	}

	var result []dtos.ListResult
	for _, gameSession := range gameSessions {
		item := dtos.ListResult{
			ID:   gameSession.ID,
			Name: gameSession.Name,
		}
		result = append(result, item)
	}

	response := &dtos.ListResponse{
		TotalCount: totalCount,
		Results:    result,
	}

	return response, nil
}
