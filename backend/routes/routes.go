package routes

import (
	"backend/controllers"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	// Initialize the Gin router
	r := gin.Default()

	// Configure CORS middleware (allows your React app on port 5173 to access this server)
	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{"http://localhost:5173"},
		AllowMethods: []string{"POST"},
		AllowHeaders: []string{"Content-Type"},
	}))

	// Define an API endpoint that returns JSON
	api := r.Group("/api")
	{
		api.GET("/hello", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "Hello from Go Backend (2026) User Login test!"})
		})
		api.POST("/login", controllers.Login)
		api.POST("/signup", controllers.Signup)
	}

	return r
}
