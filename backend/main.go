package main

import (
	"os"

	"backend/config"
	"backend/routes"
)

func main() {
	config.InitDB() // Initialize DB connection

	r := routes.SetupRouter() // Returns *gin.Engine

	// Get port from Azure (or fallback locally)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	r.Run(":" + port)

}
