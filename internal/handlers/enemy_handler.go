package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"table_top/internal/dtos/enemies"
	services "table_top/internal/services/enemies"
)

type EnemyHandler struct {
	service services.EnemyService
}

func NewEnemyHandler(service services.EnemyService) *EnemyHandler {
	return &EnemyHandler{service: service}
}

// CreateEnemy godoc
// @Summary Create a new enemy
// @Description Create a new enemy with enemy moves
// @Tags Enemies
// @Accept  json
// @Produce  json
// @Param enemy body enemies.CreateRequest true "Enemy"
// @Success 200 {object} enemies.Enemy
// @Router /enemies [post]
func (h *EnemyHandler) CreateEnemy(c *gin.Context) {
	ctx := c.Request.Context()
	var request enemies.CreateRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	enemy, err := h.service.CreateEnemy(ctx, &request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, enemy)
}

// UpdateEnemy godoc
// @Summary Update an existing enemy
// @Description Update an existing enemy
// @Tags Enemies
// @Accept  json
// @Produce  json
// @Param id path string true "Enemy ID"
// @Param enemy body enemies.UpdateRequest true "Enemy"
// @Success 200 {object} enemies.Enemy
// @Router /enemies/{id} [put]
func (h *EnemyHandler) UpdateEnemy(c *gin.Context) {
	ctx := c.Request.Context()
	idStr := c.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var req enemies.UpdateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	enemy, err := h.service.UpdateEnemy(ctx, id, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, enemy)
}

// GetEnemy godoc
// @Summary Get an enemy by ID
// @Description Get an enemy by ID
// @Tags Enemies
// @Produce  json
// @Param id path string true "Enemy ID"
// @Success 200 {object} enemies.Enemy
// @Router /enemies/{id} [get]
func (h *EnemyHandler) GetEnemy(c *gin.Context) {
	ctx := c.Request.Context()
	idStr := c.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	enemy, err := h.service.GetEnemyByID(ctx, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, enemy)
}

// DeleteEnemy godoc
// @Summary Delete an enemy by ID
// @Description Delete an enemy by ID
// @Tags Enemies
// @Param id path string true "Enemy ID"
// @Success 204
// @Router /enemies/{id} [delete]
func (h *EnemyHandler) DeleteEnemy(c *gin.Context) {
	ctx := c.Request.Context()
	idStr := c.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	if err := h.service.DeleteEnemy(ctx, id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}

// ListEnemies godoc
// @Summary List all enemies
// @Description List all enemies
// @Tags Enemies
// @Produce  json
// @Success 200 {array} enemies.Enemy
// @Router /enemies [get]
func (h *EnemyHandler) ListEnemies(c *gin.Context) {
	ctx := c.Request.Context()

	enm, err := h.service.ListEnemies(ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, enm)
}
