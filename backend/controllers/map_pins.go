package controllers

import (
	"backend/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func UserMapPins(c *gin.Context) {
	var req struct {
		UserID string `json:"userId"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "BE: Invalid request"})
		return
	}

	pins, err := model.GetUserMapPins(req.UserID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "BE: Database error"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		//"token":    tokenString,
		"Pin ID": pins,
	})
}

func NewUserPin(c *gin.Context) {
	var req model.MapPin
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "BE: Invalid request"})
		return
	}

	pin, err := model.NewUserPinDB(req.UserID, req.Pintitle, req.Pindesc, req.Pincolor, req.Pinlat, req.Pinlong)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "BE: Database error"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"Pin ID": pin,
	})

}

func DeleteUserPin(c *gin.Context) {
	var req struct {
		UserID string `json:"user_id"`
		PinID  string `json:"pin_id"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "BE: Invalid request"})
		return
	}

	pin, err := model.DeletedUserMapPin(req.UserID, req.PinID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "BE: Database error"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"Pin ID": pin,
	})

}

func UpdateUserPin(c *gin.Context) {
	var req model.MapPin

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

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

	c.JSON(http.StatusOK, gin.H{
		"Pin ID": pin,
	})
}
