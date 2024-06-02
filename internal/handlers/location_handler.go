package handlers

import (
	"net/http"

	"github.com/google/uuid"

	"github.com/gin-gonic/gin"

	"table_top/internal/dtos/locations"
	services "table_top/internal/services/locations"
)

type LocationHandler struct {
	service services.LocationService
}

func NewLocationHandler(service services.LocationService) *LocationHandler {
	return &LocationHandler{service: service}
}

// CreateLocation godoc
// @Summary Create a new location
// @Description Create a new location
// @Tags Locations
// @Accept  json
// @Produce  json
// @Param location body locations.CreateRequest true "Location"
// @Success 200 {object} locations.Location
// @Router /locations [post]
func (h *LocationHandler) CreateLocation(c *gin.Context) {
	ctx := c.Request.Context()
	var location locations.CreateRequest
	if err := c.ShouldBindJSON(&location); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	createdLocation, err := h.service.CreateLocation(ctx, &location)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, createdLocation)
}

// GetLocation godoc
// @Summary Get a location by ID
// @Description Get a location by ID
// @Tags Locations
// @Produce  json
// @Param id path string true "Location ID"
// @Success 200 {object} locations.Location
// @Router /locations/{id} [get]
func (h *LocationHandler) GetLocation(c *gin.Context) {
	ctx := c.Request.Context()
	idStr := c.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	location, err := h.service.GetLocationByID(ctx, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, location)
}

// ListLocations godoc
// @Summary List all locations
// @Description List all locations
// @Tags Locations
// @Produce  json
// @Success 200 {object} []locations.Location
// @Router /locations [get]
func (h *LocationHandler) ListLocations(c *gin.Context) {
	ctx := c.Request.Context()
	loc, err := h.service.ListLocations(ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, loc)
}

// DeleteLocation godoc
// @Summary Delete a location by ID
// @Description Delete a location by ID
// @Tags Locations
// @Param id path string true "Location ID"
// @Success 204
// @Router /locations/{id} [delete]
func (h *LocationHandler) DeleteLocation(c *gin.Context) {
	ctx := c.Request.Context()
	idStr := c.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	if err := h.service.DeleteLocation(ctx, id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusNoContent)
}

// UpdateLocation godoc
// @Summary Update a location
// @Description Update a location
// @Tags Locations
// @Accept  json
// @Produce  json
// @Param id path string true "Location ID"
// @Param location body locations.UpdateRequest true "Location"
// @Success 200 {object} locations.Location
// @Router /locations/{id} [put]
func (h *LocationHandler) UpdateLocation(c *gin.Context) {
	ctx := c.Request.Context()
	idStr := c.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var req locations.UpdateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updatedLocation, err := h.service.UpdateLocation(ctx, id, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, updatedLocation)
}
