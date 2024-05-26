package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"table_top/internal/dtos/requests/enemy_moves"
	"table_top/internal/services"
)

type EnemyMoveHandler struct {
	service services.EnemyMoveService
}

func NewEnemyMoveHandler(service services.EnemyMoveService) *EnemyMoveHandler {
	return &EnemyMoveHandler{service: service}
}

// CreateEnemyMove godoc
// @Summary Create a new enemy move
// @Description Create a new enemy move
// @Tags EnemyMoves
// @Accept  json
// @Produce  json
// @Param enemyId path string true "Enemy ID"
// @Param enemyMove body enemy_moves.CreateRequest true "EnemyMove"
// @Success 200 {object} models.EnemyMove
// @Router /enemy_moves/{enemyId} [post]
func (h *EnemyMoveHandler) CreateEnemyMove(c *gin.Context) {
	ctx := c.Request.Context()
	enemyIdStr := c.Param("enemyId")
	enemyId, err := uuid.Parse(enemyIdStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Enemy ID"})
		return
	}

	var req enemy_moves.CreateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	enemyMove, err := h.service.CreateEnemyMove(ctx, enemyId, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, enemyMove)
}

// UpdateEnemyMove godoc
// @Summary Update an existing enemy move
// @Description Update an existing enemy move
// @Tags EnemyMoves
// @Accept  json
// @Produce  json
// @Param id path string true "EnemyMove ID"
// @Param enemyMove body enemy_moves.UpdateRequest true "EnemyMove"
// @Success 200 {object} models.EnemyMove
// @Router /enemy_moves/{id} [put]
func (h *EnemyMoveHandler) UpdateEnemyMove(c *gin.Context) {
	ctx := c.Request.Context()
	idStr := c.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var req enemy_moves.UpdateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	enemyMove, err := h.service.UpdateEnemyMove(ctx, id, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, enemyMove)
}

// GetEnemyMove godoc
// @Summary Get an enemy move by ID
// @Description Get an enemy move by ID
// @Tags EnemyMoves
// @Produce  json
// @Param id path string true "EnemyMove ID"
// @Success 200 {object} models.EnemyMove
// @Router /enemy_moves/{id} [get]
func (h *EnemyMoveHandler) GetEnemyMove(c *gin.Context) {
	ctx := c.Request.Context()
	idStr := c.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	enemyMove, err := h.service.GetEnemyMoveByID(ctx, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, enemyMove)
}

// DeleteEnemyMove godoc
// @Summary Delete an enemy move by ID
// @Description Delete an enemy move by ID
// @Tags EnemyMoves
// @Param id path string true "EnemyMove ID"
// @Success 204
// @Router /enemy_moves/{id} [delete]
func (h *EnemyMoveHandler) DeleteEnemyMove(c *gin.Context) {
	ctx := c.Request.Context()
	idStr := c.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	if err := h.service.DeleteEnemyMove(ctx, id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}

// ListEnemyMoves godoc
// @Summary List all enemy moves
// @Description List all enemy moves
// @Tags EnemyMoves
// @Produce  json
// @Success 200 {array} models.EnemyMove
// @Router /enemy_moves [get]
func (h *EnemyMoveHandler) ListEnemyMoves(c *gin.Context) {
	ctx := c.Request.Context()

	enemyMoves, err := h.service.ListEnemyMoves(ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, enemyMoves)
}
