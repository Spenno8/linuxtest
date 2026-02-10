package controllers

import (
	"backend/config"
	"backend/model"
	"backend/utils"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

// LoginRequest defines the expected JSON payload
// for a user login attempt.
type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// Login handles user authentication.
// It validates credentials, verifies the password hash,
// generates a JWT, and returns basic user information.
func Login(c *gin.Context) {

	var req LoginRequest

	// Parse and validate incoming JSON request
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	// Debug logging for login attempts (development only)
	fmt.Println("LOGIN ATTEMPT")
	fmt.Println("Email:", req.Email)
	fmt.Println("Password:", req.Password)

	// Attempt to find user by email first
	user, err := model.GetUserByCred("email", req.Email)

	// If not found by email, attempt lookup by username
	if err != nil {
		user, err = model.GetUserByCred("username", req.Email)
		if err != nil {
			// Do not reveal which field failed
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
			return
		}
	}

	// Debug: print hashed password stored in database
	fmt.Println("Auth Hash from DB:", user.Password)

	// Compare provided password with stored password hash
	if !utils.CheckPasswordHash(req.Password, user.Password) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		fmt.Println("User Password", user.Password)
		return
	}

	// Create a JWT containing the user's email
	// Token expires after 24 hours
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": req.Email,
		"exp":   time.Now().Add(24 * time.Hour).Unix(),
	})

	// Sign the token using the server's secret key
	tokenString, _ := token.SignedString(config.JwtSecret)

	// Debug: confirm correct user ID
	fmt.Println("USER ID CHECK:", user.ID)

	// Return authentication token and basic user info
	c.JSON(http.StatusOK, gin.H{
		"token":    tokenString,
		"username": user.Username,
		"email":    user.Email,
		"id":       user.ID,
	})

}
