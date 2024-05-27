package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"table_top/internal/dtos/requests/combo_cards"
	"table_top/internal/services"
)

type ComboCardHandler struct {
	service services.ComboCardService
}

func NewComboCardHandler(service services.ComboCardService) *ComboCardHandler {
	return &ComboCardHandler{service: service}
}

// CreateComboCard godoc
// @Summary Create a new combo card
// @Description Create a new combo card
// @Tags ComboCards
// @Accept  json
// @Produce  json
// @Param comboCard body combo_cards.CreateRequest true "Combo Card"
// @Success 200 {object} models.ComboCard
// @Router /combo_cards [post]
func (h *ComboCardHandler) CreateComboCard(c *gin.Context) {
	var req combo_cards.CreateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	comboCard, err := h.service.CreateComboCard(c.Request.Context(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, comboCard)
}

// GetComboCard godoc
// @Summary Get a combo card by ID
// @Description Get a combo card by ID
// @Tags ComboCards
// @Produce  json
// @Param id path string true "Combo Card ID"
// @Success 200 {object} models.ComboCard
// @Router /combo_cards/{id} [get]
func (h *ComboCardHandler) GetComboCard(c *gin.Context) {
	id := c.Param("id")
	uid, err := uuid.Parse(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	comboCard, err := h.service.GetComboCardByID(c.Request.Context(), uid)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, comboCard)
}

// ListComboCards godoc
// @Summary List all combo cards
// @Description List all combo cards
// @Tags ComboCards
// @Produce  json
// @Success 200 {array} models.ComboCard
// @Router /combo_cards [get]
func (h *ComboCardHandler) ListComboCards(c *gin.Context) {
	comboCards, err := h.service.ListComboCards(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, comboCards)
}

// UpdateComboCard godoc
// @Summary Update a combo card
// @Description Update a combo card
// @Tags ComboCards
// @Accept  json
// @Produce  json
// @Param id path string true "Combo Card ID"
// @Param comboCard body combo_cards.UpdateRequest true "Combo Card"
// @Success 200 {object} models.ComboCard
// @Router /combo_cards/{id} [put]
func (h *ComboCardHandler) UpdateComboCard(c *gin.Context) {
	id := c.Param("id")
	uid, err := uuid.Parse(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var req combo_cards.UpdateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	comboCard, err := h.service.UpdateComboCard(c.Request.Context(), uid, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, comboCard)
}

// DeleteComboCard godoc
// @Summary Delete a combo card
// @Description Delete a combo card
// @Tags ComboCards
// @Param id path string true "Combo Card ID"
// @Success 204
// @Router /combo_cards/{id} [delete]
func (h *ComboCardHandler) DeleteComboCard(c *gin.Context) {
	id := c.Param("id")
	uid, err := uuid.Parse(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	if err := h.service.DeleteComboCard(c.Request.Context(), uid); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}
