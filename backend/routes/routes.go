package routes

import (
	"backend/controllers"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// SetupRouter configures and returns the main Gin router.
// It sets up middleware (CORS) and defines all API routes
// that map HTTP requests to controller handlers.
func SetupRouter() *gin.Engine {

	// Create a Gin router with default middleware:
	// logger and recovery (crash-free) middleware
	r := gin.Default()

	// Configure CORS to allow requests from the React frontend
	// running on http://localhost:5173
	r.Use(cors.New(cors.Config{

		AllowOrigins:  []string{"http://localhost:5173", "https://orange-desert-096e96800.1.azurestaticapps.net/"}, // Frontend development server
		AllowMethods:  []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},                                         // Allowed HTTP methods for cross-origin requests
		AllowHeaders:  []string{"Origin", "Content-Type", "Authorization"},                                         // Allowed request headers from the client
		ExposeHeaders: []string{"Content-Length"},                                                                  // Headers exposed to the client
	}))

	// Group all API routes under the /api prefix
	// Example: http://localhost:8080/api/login
	api := r.Group("/api")
	{
		// Simple test endpoint to confirm the backend is running
		api.GET("/hello", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "Hello from Go Backend (2026) User Login test!"})
		})

		// Authentication routes
		api.POST("/login", controllers.Login)
		api.POST("/signup", controllers.Signup)

		// User map pin routes
		// These endpoints manage CRUD operations for user-created map pins
		api.POST("/UserMapPins", controllers.UserMapPins)
		api.POST("/NewUserPin", controllers.NewUserPin)
		api.POST("/DeleteUserPin", controllers.DeleteUserPin)
		api.POST("/UpdateUserPin", controllers.UpdateUserPin)

	}

	return r
}
