package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize the Gin router
	r := gin.Default()

	// Configure CORS middleware (allows your React app on port 5173 to access this server)
	r.Use(cors.Default())

	// Define an API endpoint that returns JSON
	r.GET("/api/hello", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello from Go Backend (2026)!",
		})
	})

	// Run the server on port 8080
	// Make sure this port is different from your React dev server (usually 5173)
	r.Run(":8080")
}
