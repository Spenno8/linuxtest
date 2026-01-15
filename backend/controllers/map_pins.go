package controllers

import (
	"backend/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

type MapPin struct {
	Pinuuid  string `json:"UUID"`
	Pinid    string `json:"id"`
	Pintitle string `json:"pintitle"`
	Pindesc  string `json:"pindesc"`
	Pincolor string `json:"pincolor"`
	Pinlat   string `json:"pinlat"`
	Pinlong  string `json:"pinlong"`
}

func UserMapPins(c *gin.Context) {
	var req MapPin
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "BE: Invalid request"})
		return
	}

	pins, err := model.GetUserMapPins(req.Pinuuid)
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
	var req MapPin
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "BE: Invalid request"})
		return
	}

	pin, err := model.NewUserPinDB(req.Pinuuid, req.Pintitle, req.Pindesc, req.Pincolor, req.Pinlat, req.Pinlong)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "BE: Database error"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"Pin ID": pin,
	})

}

func DeleteUserPin(c *gin.Context) {
	var req MapPin
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "BE: Invalid request"})
		return
	}

	pin, err := model.DeletedUserMapPin(req.Pinuuid, req.Pinid)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "BE: Database error"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"Pin ID": pin,
	})

}
