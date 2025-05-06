package main

import (
	"bondly/config"
	"bondly/services"
	"bondly/utils"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	// Load environment variables and config
	config.LoadConfig()
	utils.InitLogger()

	// ✅ Initialize OAuth2 config
	services.InitializeOAuth()

	// Initialize DB
	_, err := utils.InitDatabase()
	if err != nil {
		panic("Database connection failed: " + err.Error())
	}

	r := gin.Default()

	// Enable CORS for frontend access
	r.Use(cors.Default()) // This allows all origins, you can customize it if needed.

	// Microsoft OAuth routes
	r.GET("/auth/microsoft", services.HandleMicrosoftLogin)             // Correct route for Microsoft login
	r.GET("/auth/microsoft/callback", services.HandleMicrosoftCallback) // Correct callback route

	// Posts routes
	r.GET("/posts", services.GetPosts)    // Fetch all posts
	r.POST("/posts", services.CreatePost) // Create new post

	// Start server on port 8000
	r.Run(":8000")
}
