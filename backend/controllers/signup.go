package controllers

import (
	"backend/config"
	"backend/model"
	"backend/utils"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

// Sign up data field expected to be recieved from the client
type SignupRequest struct {
	Email     string `json:"email"`
	Username  string `json:"username"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Password  string `json:"password"`
}

// Signup handles user registration requests.
// It validates input, checks for duplicate credentials,
// securely hashes the password, stores the user,
// and returns a JWT for authentication.
func Signup(c *gin.Context) {
	var req SignupRequest

	// Parse and validate incoming JSON request body
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "BE: Invalid request"})
		return
	}

	// Check if the email already exists in the database
	// GetUserByCred returns nil error if the record exists
	user, err := model.GetUserByCred("email", req.Email)
	if err == nil { //No error will return if it exists therefore email in use
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Email already in use"})
		return
	}

	// Same logic as email for username
	user, err = model.GetUserByCred("username", req.Username)
	if err == nil {
		c.JSON(http.StatusConflict, gin.H{"error": "Username already in use"})
		return
	}

	// Hash the user's password before storing it
	// This ensures the plain-text password is never saved
	hashedPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "BE: Password hash error"})
		return
	}

	// Create a new user record in the database
	user, err = model.SignupNewUser(req.Email, req.Username, hashedPassword, req.Firstname, req.Lastname)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "BE: Database error"})
		return
	}

	// Create a JWT token containing the user's email
	// Token expires after 24 hours
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": user.Email,
		"exp":   time.Now().Add(24 * time.Hour).Unix(),
	})

	// Sign the token using the server's secret key
	tokenString, err := token.SignedString(config.JwtSecret)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "BE: Token generation failed"})
		return
	}

	// Return the JWT and basic user details to the client
	c.JSON(http.StatusOK, gin.H{
		"token":    tokenString,
		"username": user.Username,
		"email":    user.Email,
	})
}
