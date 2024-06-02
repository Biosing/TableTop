package handlers

import (
	"net/http"

	"github.com/google/uuid"

	"github.com/gin-gonic/gin"

	dtos "table_top/internal/dtos/characters"
	services "table_top/internal/services/characters"
)

type CharacterHandler struct {
	service services.CharacterService
}

func NewCharacterHandler(service services.CharacterService) *CharacterHandler {
	return &CharacterHandler{service: service}
}

// CreateCharacter godoc
// @Summary Create a new character
// @Description Create a new character
// @Tags Characters
// @Accept  json
// @Produce  json
// @Param character body characters.CreateRequest true "Character"
// @Success 200 {object} characters.Character
// @Router /characters [post]
func (h *CharacterHandler) CreateCharacter(c *gin.Context) {
	ctx := c.Request.Context()
	var character dtos.CreateRequest
	if err := c.ShouldBindJSON(&character); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	characterRepo, err := h.service.CreateCharacter(ctx, &character)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, characterRepo)
}

// DeleteCharacter godoc
// @Summary Delete a character by ID
// @Description Delete a character by ID
// @Tags Characters
// @Param id path string true "Character ID"
// @Success 204
// @Router /characters/{id} [delete]
func (h *CharacterHandler) DeleteCharacter(c *gin.Context) {
	ctx := c.Request.Context()
	idStr := c.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	if err := h.service.DeleteCharacter(ctx, id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusNoContent)
}

// GetCharacter godoc
// @Summary Get a character by ID
// @Description Get a character by ID
// @Tags Characters
// @Produce  json
// @Param id path string true "Character ID"
// @Success 200 {object} characters.Character
// @Router /characters/{id} [get]
func (h *CharacterHandler) GetCharacter(c *gin.Context) {
	ctx := c.Request.Context()
	idStr := c.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	character, err := h.service.GetCharacterByID(ctx, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, character)
}

// ListCharacters godoc
// @Summary List all characters
// @Description List all characters
// @Tags Characters
// @Produce  json
// @Success 200 {array} characters.Character
// @Router /characters [get]
func (h *CharacterHandler) ListCharacters(c *gin.Context) {
	ctx := c.Request.Context()
	char, err := h.service.ListCharacters(ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, char)
}

// UpdateCharacter godoc
// @Summary Update a character
// @Description Update a character
// @Tags Characters
// @Accept  json
// @Produce  json
// @Param id path string true "Enemy ID"
// @Param character body characters.UpdateRequest true "Character"
// @Success 200 {object} characters.Character
// @Router /characters/{id} [put]
func (h *CharacterHandler) UpdateCharacter(c *gin.Context) {
	ctx := c.Request.Context()
	idStr := c.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var request dtos.UpdateRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	character, err := h.service.UpdateCharacter(ctx, id, &request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, character)
}
