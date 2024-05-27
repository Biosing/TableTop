package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"table_top/internal/dtos/requests/special_effects"
	"table_top/internal/services/combo_cards"
)

type SpecialEffectHandler struct {
	service services.SpecialEffectService
}

func NewSpecialEffectHandler(service services.SpecialEffectService) *SpecialEffectHandler {
	return &SpecialEffectHandler{service: service}
}

// CreateSpecialEffect godoc
// @Summary Create a new special effect
// @Description Create a new special effect
// @Tags SpecialEffects
// @Accept  json
// @Produce  json
// @Param specialEffect body special_effects.CreateRequest true "Special Effect"
// @Success 200 {object} models.SpecialEffect
// @Router /special_effects [post]
func (h *SpecialEffectHandler) CreateSpecialEffect(c *gin.Context) {
	var req special_effects.CreateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	specialEffect, err := h.service.CreateSpecialEffect(c.Request.Context(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, specialEffect)
}

// GetSpecialEffect godoc
// @Summary Get a special effect by ID
// @Description Get a special effect by ID
// @Tags SpecialEffects
// @Produce  json
// @Param id path string true "Special Effect ID"
// @Success 200 {object} models.SpecialEffect
// @Router /special_effects/{id} [get]
func (h *SpecialEffectHandler) GetSpecialEffect(c *gin.Context) {
	id := c.Param("id")
	uid, err := uuid.Parse(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	specialEffect, err := h.service.GetSpecialEffectByID(c.Request.Context(), uid)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, specialEffect)
}

// ListSpecialEffects godoc
// @Summary List all special effects
// @Description List all special effects
// @Tags SpecialEffects
// @Produce  json
// @Success 200 {array} models.SpecialEffect
// @Router /special_effects [get]
func (h *SpecialEffectHandler) ListSpecialEffects(c *gin.Context) {
	specialEffects, err := h.service.ListSpecialEffects(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, specialEffects)
}

// UpdateSpecialEffect godoc
// @Summary Update a special effect
// @Description Update a special effect
// @Tags SpecialEffects
// @Accept  json
// @Produce  json
// @Param id path string true "Special Effect ID"
// @Param specialEffect body special_effects.UpdateRequest true "Special Effect"
// @Success 200 {object} models.SpecialEffect
// @Router /special_effects/{id} [put]
func (h *SpecialEffectHandler) UpdateSpecialEffect(c *gin.Context) {
	id := c.Param("id")
	uid, err := uuid.Parse(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var req special_effects.UpdateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	specialEffect, err := h.service.UpdateSpecialEffect(c.Request.Context(), uid, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, specialEffect)
}

// DeleteSpecialEffect godoc
// @Summary Delete a special effect
// @Description Delete a special effect
// @Tags SpecialEffects
// @Param id path string true "Special Effect ID"
// @Success 204
// @Router /special_effects/{id} [delete]
func (h *SpecialEffectHandler) DeleteSpecialEffect(c *gin.Context) {
	id := c.Param("id")
	uid, err := uuid.Parse(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	if err := h.service.DeleteSpecialEffect(c.Request.Context(), uid); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}
