package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"table_top/internal/dtos/game_session_players"
	services "table_top/internal/services/games"
)

type GameSessionPlayerHandler struct {
	service services.GameSessionPlayerService
}

func NewGameSessionPlayerHandler(service services.GameSessionPlayerService) *GameSessionPlayerHandler {
	return &GameSessionPlayerHandler{service: service}
}

// AddPlayerToGameSession godoc
// @Summary Add player to game session
// @Description Add player to game session
// @Tags GameSessionsPlayers
// @Accept  json
// @Produce  json
// @Param gameSession body game_session_players.AddPlayerRequest true "GameSession"
// @Success 200
// @Router /add_player [post]
func (h *GameSessionPlayerHandler) AddPlayerToGameSession(c *gin.Context) {
	ctx := c.Request.Context()

	var req game_session_players.AddPlayerRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := h.service.AddPlayer(ctx, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Player added to game session"})
}

// RemovePlayerFromGameSession godoc
// @Summary Remove player from game session
// @Description Remove player from game session
// @Tags GameSessionsPlayers
// @Accept  json
// @Produce  json
// @Param gameSession body game_session_players.RemovePlayerRequest true "GameSession"
// @Success 200
// @Router /remove_player [post]
func (h *GameSessionPlayerHandler) RemovePlayerFromGameSession(c *gin.Context) {
	ctx := c.Request.Context()

	var req game_session_players.RemovePlayerRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := h.service.RemovePlayer(ctx, c, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Player removed from game session"})
}
