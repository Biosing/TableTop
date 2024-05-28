package games

import (
	"context"
	"errors"
	"net/http"
	"time"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"table_top/internal/dtos/requests/game_sessions"
	models "table_top/internal/models/games"
	repositories "table_top/internal/repositories/games"
)

type GameSessionService interface {
	StartNewGame(ctx context.Context, c *gin.Context, req *game_sessions.CreateRequest) (*models.GameSession, error)
	GameSession(ctx context.Context, c *gin.Context) (*models.GameSession, error)
}

type gameSessionService struct {
	repo repositories.GameSessionRepository
}

func NewGameSessionService(repo repositories.GameSessionRepository) GameSessionService {
	return &gameSessionService{repo: repo}
}

func (s *gameSessionService) StartNewGame(ctx context.Context, c *gin.Context, req *game_sessions.CreateRequest) (*models.GameSession, error) {
	session := sessions.Default(c)

	now := time.Now()
	gameSessionID := uuid.New()

	gameSession := &models.GameSession{
		ID:            gameSessionID,
		Nickname:      req.Nickname,
		CharacterID:   req.CharacterID,
		StartGameDate: &now,
	}

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
