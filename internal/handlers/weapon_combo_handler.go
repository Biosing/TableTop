package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"table_top/internal/dtos/requests/weapon_combos"
	"table_top/internal/services/items"
)

type WeaponComboHandler struct {
	service services.WeaponComboService
}

func NewWeaponComboHandler(service services.WeaponComboService) *WeaponComboHandler {
	return &WeaponComboHandler{service: service}
}

// CreateWeaponCombo godoc
// @Summary Create a new weapon combo
// @Description Create a new weapon combo
// @Tags WeaponCombos
// @Accept  json
// @Produce  json
// @Param weaponCombo body weapon_combos.CreateRequest true "WeaponCombo"
// @Success 200 {object} models.WeaponCombo
// @Router /weapon_combos [post]
func (h *WeaponComboHandler) CreateWeaponCombo(c *gin.Context) {
	var req weapon_combos.CreateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	weaponCombo, err := h.service.CreateWeaponCombo(c.Request.Context(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, weaponCombo)
}

// GetWeaponCombo godoc
// @Summary Get a weapon combo by ID
// @Description Get a weapon combo by ID
// @Tags WeaponCombos
// @Produce  json
// @Param id path string true "WeaponCombo ID"
// @Success 200 {object} models.WeaponCombo
// @Router /weapon_combos/{id} [get]
func (h *WeaponComboHandler) GetWeaponCombo(c *gin.Context) {
	id := c.Param("id")
	weaponComboID, err := uuid.Parse(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	weaponCombo, err := h.service.GetWeaponComboByID(c.Request.Context(), weaponComboID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, weaponCombo)
}

// UpdateWeaponCombo godoc
// @Summary Update a weapon combo
// @Description Update a weapon combo
// @Tags WeaponCombos
// @Accept  json
// @Produce  json
// @Param id path string true "WeaponCombo ID"
// @Param weaponCombo body weapon_combos.UpdateRequest true "WeaponCombo"
// @Success 200 {object} models.WeaponCombo
// @Router /weapon_combos/{id} [put]
func (h *WeaponComboHandler) UpdateWeaponCombo(c *gin.Context) {
	id := c.Param("id")
	weaponComboID, err := uuid.Parse(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var req weapon_combos.UpdateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	weaponCombo, err := h.service.UpdateWeaponCombo(c.Request.Context(), weaponComboID, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, weaponCombo)
}

// DeleteWeaponCombo godoc
// @Summary Delete a weapon combo by ID
// @Description Delete a weapon combo by ID
// @Tags WeaponCombos
// @Param id path string true "WeaponCombo ID"
// @Success 204
// @Router /weapon_combos/{id} [delete]
func (h *WeaponComboHandler) DeleteWeaponCombo(c *gin.Context) {
	id := c.Param("id")
	weaponComboID, err := uuid.Parse(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	if err := h.service.DeleteWeaponCombo(c.Request.Context(), weaponComboID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}

// ListWeaponCombos godoc
// @Summary List all weapon combos
// @Description List all weapon combos
// @Tags WeaponCombos
// @Produce  json
// @Success 200 {array} models.WeaponCombo
// @Router /weapon_combos [get]
func (h *WeaponComboHandler) ListWeaponCombos(c *gin.Context) {
	weaponCombos, err := h.service.ListWeaponCombos(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, weaponCombos)
}
