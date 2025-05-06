package routes

import (
	"bondly/controllers" // Import the controller
	"github.com/gin-gonic/gin"
)

// RegisterPostRoutes registers the routes for posts
func RegisterPostRoutes(r *gin.Engine) {
	// Post Routes
	r.GET("/posts", controllers.GetPosts)          // Get all posts
	r.GET("/posts/:id", controllers.GetPost)       // Get a specific post
	r.POST("/posts", controllers.CreatePost)       // Create a new post
	r.PUT("/posts/:id", controllers.UpdatePost)    // Update a post
	r.DELETE("/posts/:id", controllers.DeletePost) // Delete a post
}
