package controllers

import (
	"backend/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

// UserMapPins retrieves all map pins belonging to a specific user.
// Expects a JSON body containing the user's ID.
func UserMapPins(c *gin.Context) {

	// Request payload containing the user ID
	var req struct {
		UserID string `json:"userId"`
	}

	// Parse and validate incoming JSON request
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "BE: Invalid request"})
		return
	}

	// Fetch all map pins for the given user
	pins, err := model.GetUserMapPins(req.UserID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "BE: Database error"})
		return
	}

	// Return the user's map pins
	c.JSON(http.StatusOK, gin.H{
		"Pin ID": pins,
	})
}

// NewUserPin creates a new map pin for a user.
// Expects a JSON body matching the MapPin model.
func NewUserPin(c *gin.Context) {
	var req model.MapPin

	// Parse and validate incoming JSON request
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "BE: Invalid request"})
		return
	}

	// Insert a new pin record into the database
	pin, err := model.NewUserPinDB(req.UserID, req.Pintitle, req.Pindesc, req.Pincolor, req.Pinlat, req.Pinlong)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "BE: Database error"})
		return
	}

	// Return the newly created pin ID
	c.JSON(http.StatusOK, gin.H{
		"Pin ID": pin,
	})

}

// DeleteUserPin removes a specific map pin owned by a user.
// Expects both the user ID and pin ID in the request body.
func DeleteUserPin(c *gin.Context) {

	// Request payload containing user ID and pin ID
	var req struct {
		UserID string `json:"user_id"`
		PinID  string `json:"pin_id"`
	}

	// Parse and validate incoming JSON request
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "BE: Invalid request"})
		return
	}

	// Delete the specified pin from the database
	pin, err := model.DeletedUserMapPin(req.UserID, req.PinID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "BE: Database error"})
		return
	}

	// Return the deleted pin ID (or confirmation)
	c.JSON(http.StatusOK, gin.H{
		"Pin ID": pin,
	})

}

// UpdateUserPin updates an existing map pin's details.
// Expects a JSON body matching the MapPin model.
func UpdateUserPin(c *gin.Context) {
	var req model.MapPin

	// Parse and validate incoming JSON request
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	// Update the pin record in the database
	pin, err := model.UpdateUserPinDB(
		req.ID,
		req.UserID,
		req.Pintitle,
		req.Pindesc,
		req.Pincolor,
		req.Pinlat,
		req.Pinlong,
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
		return
	}

	// Return the updated pin ID (or confirmation)
	c.JSON(http.StatusOK, gin.H{
		"Pin ID": pin,
	})
}
