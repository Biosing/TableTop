package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"table_top/internal/dtos/requests/damage_combo_cards"
	services "table_top/internal/services/combo_cards"
)

type DamageComboCardHandler struct {
	service services.DamageComboCardService
}

func NewDamageComboCardHandler(service services.DamageComboCardService) *DamageComboCardHandler {
	return &DamageComboCardHandler{service: service}
}

// CreateDamageComboCard godoc
// @Summary Create a new damage combo card
// @Description Create a new damage combo card
// @Tags DamageComboCards
// @Accept  json
// @Produce  json
// @Param damageComboCard body damage_combo_cards.CreateRequest true "Damage Combo Card"
// @Success 200 {object} models.DamageComboCard
// @Router /damage_combo_cards [post]
func (h *DamageComboCardHandler) CreateDamageComboCard(c *gin.Context) {
	var req damage_combo_cards.CreateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	damageComboCard, err := h.service.CreateDamageComboCard(c.Request.Context(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, damageComboCard)
}

// GetDamageComboCard godoc
// @Summary Get a damage combo card by ID
// @Description Get a damage combo card by ID
// @Tags DamageComboCards
// @Produce  json
// @Param id path string true "Damage Combo Card ID"
// @Success 200 {object} models.DamageComboCard
// @Router /damage_combo_cards/{id} [get]
func (h *DamageComboCardHandler) GetDamageComboCard(c *gin.Context) {
	id := c.Param("id")
	uid, err := uuid.Parse(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	damageComboCard, err := h.service.GetDamageComboCardByID(c.Request.Context(), uid)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, damageComboCard)
}

// ListDamageComboCards godoc
// @Summary List all damage combo cards
// @Description List all damage combo cards
// @Tags DamageComboCards
// @Produce  json
// @Success 200 {array} models.DamageComboCard
// @Router /damage_combo_cards [get]
func (h *DamageComboCardHandler) ListDamageComboCards(c *gin.Context) {
	damageComboCards, err := h.service.ListDamageComboCards(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, damageComboCards)
}

// UpdateDamageComboCard godoc
// @Summary Update a damage combo card
// @Description Update a damage combo card
// @Tags DamageComboCards
// @Accept  json
// @Produce  json
// @Param id path string true "Damage Combo Card ID"
// @Param damageComboCard body damage_combo_cards.UpdateRequest true "Damage Combo Card"
// @Success 200 {object} models.DamageComboCard
// @Router /damage_combo_cards/{id} [put]
func (h *DamageComboCardHandler) UpdateDamageComboCard(c *gin.Context) {
	id := c.Param("id")
	uid, err := uuid.Parse(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var req damage_combo_cards.UpdateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	damageComboCard, err := h.service.UpdateDamageComboCard(c.Request.Context(), uid, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, damageComboCard)
}

// DeleteDamageComboCard godoc
// @Summary Delete a damage combo card
// @Description Delete a damage combo card
// @Tags DamageComboCards
// @Param id path string true "Damage Combo Card ID"
// @Success 204
// @Router /damage_combo_cards/{id} [delete]
func (h *DamageComboCardHandler) DeleteDamageComboCard(c *gin.Context) {
	id := c.Param("id")
	uid, err := uuid.Parse(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	if err := h.service.DeleteDamageComboCard(c.Request.Context(), uid); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}
