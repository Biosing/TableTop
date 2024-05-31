package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"table_top/internal/dtos/requests/game_sessions"
	services "table_top/internal/services/games"
)

type GameSessionHandler struct {
	service services.GameSessionService
}

func NewGameSessionHandler(service services.GameSessionService) *GameSessionHandler {
	return &GameSessionHandler{service: service}
}

// CreateGameSession godoc
// @Summary Create a new game session
// @Description Create a new game session
// @Tags GameSessions
// @Accept  json
// @Produce  json
// @Param gameSession body game_sessions.CreateRequest true "GameSession"
// @Success 200 {object} models.GameSession
// @Router /create_game [post]
func (h *GameSessionHandler) CreateGameSession(c *gin.Context) {
	ctx := c.Request.Context()

	var req game_sessions.CreateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	gameSession, err := h.service.CreateNewGame(ctx, c, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gameSession)
}

// GetGameSession godoc
// @Summary Get game session
// @Description Get game session
// @Tags GameSessions
// @Accept  json
// @Produce  json
// @Success 200 {object} models.GameSession
// @Router /game_session [get]
func (h *GameSessionHandler) GetGameSession(c *gin.Context) {
	ctx := c.Request.Context()

	gameSession, err := h.service.GameSession(ctx, c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gameSession)
}

// StartGameSession godoc
// @Summary Start a new game session
// @Description Start a new game session
// @Tags GameSessions
// @Accept  json
// @Produce  json
// @Success 200
// @Router /start_game [post]
func (h *GameSessionHandler) StartGameSession(c *gin.Context) {
	ctx := c.Request.Context()

	err := h.service.StartGame(ctx, c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Game session started"})

}

// FinishGameSession godoc
// @Summary Finish game session
// @Description Finish game session
// @Tags GameSessions
// @Accept  json
// @Produce  json
// @Success 200
// @Router /finish_game [post]
func (h *GameSessionHandler) FinishGameSession(c *gin.Context) {
	ctx := c.Request.Context()

	err := h.service.FinishGame(ctx, c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Game session finished"})
}

// ListGameSessions godoc
// @Summary List game sessions
// @Description List game sessions
// @Tags GameSessions
// @Accept  json
// @Produce  json
// @Param name query string false "Name"
// @Param page query int false "Page"
// @Param limit query int false "Limit"
// @Success 200 {object} []models.GameSession
// @Router /list_game_sessions [get]
func (h *GameSessionHandler) ListGameSessions(c *gin.Context) {
	ctx := c.Request.Context()

	var req game_sessions.ListRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	gameSessions, err := h.service.ListGameSessions(ctx, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gameSessions)
}
