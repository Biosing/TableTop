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

// StartGameSession godoc
// @Summary Start a new game session
// @Description Start a new game session
// @Tags GameSessions
// @Accept  json
// @Produce  json
// @Param gameSession body game_sessions.CreateRequest true "GameSession"
// @Success 200 {object} models.GameSession
// @Router /start_game [post]
func (h *GameSessionHandler) StartGameSession(c *gin.Context) {
	ctx := c.Request.Context()

	var req game_sessions.CreateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	gameSession, err := h.service.StartNewGame(ctx, c, &req)
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
