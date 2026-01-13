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

type SignupRequest struct {
	Email     string `json:"email"`
	Username  string `json:"username"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Password  string `json:"password"`
}

func Signup(c *gin.Context) {
	var req SignupRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "BE: Invalid request"})
		return
	}

	user, err := model.GetUserByCred("email", req.Email)
	if err == nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Email already in use"})
		return
	}

	user, err = model.GetUserByCred("username", req.Username)
	if err == nil {
		c.JSON(http.StatusConflict, gin.H{"error": "Username already in use"})
		return
	}

	hashedPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "BE: Password hash error"})
		return
	}

	user, err = model.SignupNewUser(req.Email, req.Username, hashedPassword, req.Firstname, req.Lastname)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "BE: Database error"})
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": user.Email,
		"exp":   time.Now().Add(24 * time.Hour).Unix(),
	})

	tokenString, err := token.SignedString(config.JwtSecret)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "BE: Token generation failed"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"token":    tokenString,
		"username": user.Username,
		"email":    user.Email,
	})
}

// func Signup(c *gin.Context) {
// 	var req SignupRequest
// 	if err := c.ShouldBindJSON(&req); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
// 		return
// 	}

// 	fmt.Println("LOGIN ATTEMPT")
// 	fmt.Println("Email:", req.Email)
// 	fmt.Println("Username:", req.Username)
// 	fmt.Println("Firstname:", req.Firstname)
// 	fmt.Println("Lastname:", req.Lastname)
// 	fmt.Println("Password:", req.Password)

// 	user, err := model.GetUserByEmail(req.Email, req.Email)
// 	if err == nil {
// 		c.JSON(http.StatusUnauthorized, gin.H{"error": "Username or Email in use"})
// 		return
// 	}

// 	hashedpassword, err := utils.HashPassword(req.Password)
// 	if err != nil {
// 		c.JSON(http.StatusUnauthorized, gin.H{"error": "Password Hash Error"})
// 		return
// 	}

// 	user, err := model.SignupNewUser(req.Email, req.Username, hashedpassword, req.Firstname, req.Lastname)
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database Error"})
// 		return
// 	}

// 	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
// 		"email": req.Email,
// 		"exp":   time.Now().Add(24 * time.Hour).Unix(),
// 	})

// 	tokenString, err := token.SignedString(config.JwtSecret)
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "Token generation failed"})
// 		return
// 	}

// 	c.JSON(http.StatusOK, gin.H{
// 		"token":    tokenString,
// 		"username": req.Username,
// 		"email":    req.Email,
// 	})

// }
