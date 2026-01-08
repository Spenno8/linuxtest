package main

import (
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

var jwtSecret = []byte("super-secret-key")

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func main() {
	// Initialize the Gin router
	r := gin.Default()

	// Configure CORS middleware (allows your React app on port 5173 to access this server)
	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{"http://localhost:5173"},
		AllowMethods: []string{"POST"},
		AllowHeaders: []string{"Content-Type"},
	}))
	// Define an API endpoint that returns JSON
	r.GET("/api/hello", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello from Go Backend (2026) User Login test!",
		})
	})

	hashedPassword, _ := bcrypt.GenerateFromPassword(
		[]byte("password123"),
		bcrypt.DefaultCost,
	)

	r.POST("/login", func(c *gin.Context) {
		var req LoginRequest

		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
			return
		}

		// Fake user check
		if req.Email != "test@example.com" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
			return
		}

		if err := bcrypt.CompareHashAndPassword(hashedPassword, []byte(req.Password)); err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
			return
		}

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"email": req.Email,
			"exp":   time.Now().Add(24 * time.Hour).Unix(),
		})

		tokenString, _ := token.SignedString(jwtSecret)

		c.JSON(http.StatusOK, gin.H{
			"token": tokenString,
		})
	})

	// Run the server on port 8080
	// Make sure this port is different from your React dev server (usually 5173)
	r.Run(":8080")
}
