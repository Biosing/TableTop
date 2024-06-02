package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"table_top/internal/dtos/weapons"
	services "table_top/internal/services/items"
)

type WeaponHandler struct {
	service services.WeaponService
}

func NewWeaponHandler(service services.WeaponService) *WeaponHandler {
	return &WeaponHandler{service: service}
}

// CreateWeapon godoc
// @Summary Create a new weapon
// @Description Create a new weapon along with its combos
// @Tags Weapons
// @Accept  json
// @Produce  json
// @Param weapon body weapons.CreateRequest true "Weapon"
// @Success 200 {object} items.Weapon
// @Router /weapons [post]
func (h *WeaponHandler) CreateWeapon(c *gin.Context) {
	var req weapons.CreateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	weapon, err := h.service.CreateWeapon(c.Request.Context(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, weapon)
}

// GetWeapon godoc
// @Summary Get a weapon by ID
// @Description Get a weapon by ID
// @Tags Weapons
// @Produce  json
// @Param id path string true "Weapon ID"
// @Success 200 {object} items.Weapon
// @Router /weapons/{id} [get]
func (h *WeaponHandler) GetWeapon(c *gin.Context) {
	id := c.Param("id")
	weaponID, err := uuid.Parse(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	weapon, err := h.service.GetWeaponByID(c.Request.Context(), weaponID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, weapon)
}

// UpdateWeapon godoc
// @Summary Update a weapon
// @Description Update a weapon
// @Tags Weapons
// @Accept  json
// @Produce  json
// @Param id path string true "Weapon ID"
// @Param weapon body weapons.UpdateRequest true "Weapon"
// @Success 200 {object} items.Weapon
// @Router /weapons/{id} [put]
func (h *WeaponHandler) UpdateWeapon(c *gin.Context) {
	id := c.Param("id")
	weaponID, err := uuid.Parse(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var req weapons.UpdateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	weapon, err := h.service.UpdateWeapon(c.Request.Context(), weaponID, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, weapon)
}

// DeleteWeapon godoc
// @Summary Delete a weapon by ID
// @Description Delete a weapon by ID
// @Tags Weapons
// @Param id path string true "Weapon ID"
// @Success 204
// @Router /weapons/{id} [delete]
func (h *WeaponHandler) DeleteWeapon(c *gin.Context) {
	id := c.Param("id")
	weaponID, err := uuid.Parse(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	if err := h.service.DeleteWeapon(c.Request.Context(), weaponID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}

// ListWeapons godoc
// @Summary List all weapons
// @Description List all weapons
// @Tags Weapons
// @Produce  json
// @Success 200 {object} []items.Weapon
// @Router /weapons [get]
func (h *WeaponHandler) ListWeapons(c *gin.Context) {
	wpn, err := h.service.ListWeapons(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, wpn)
}
