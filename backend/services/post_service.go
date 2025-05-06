package services

import (
	"bondly/models"
	"bondly/utils"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"net/http"
)

// GetPosts fetches all posts from the database.
func GetPosts(c *gin.Context) {
	var posts []models.Post

	// Fetch all posts along with the user who created them
	if err := utils.DB.Preload("User").Find(&posts).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not fetch posts"})
		return
	}

	c.JSON(http.StatusOK, posts)
}

// CreatePost handles the creation of a new post.
func CreatePost(c *gin.Context) {
	var newPost models.Post
	var user models.User

	// Bind the incoming JSON payload to the newPost struct
	if err := c.ShouldBind(&newPost); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data"})
		return
	}

	// Get the user from the context (assuming you store user_id in context)
	userId, _ := c.Get("user_id") // Get user_id from context after login

	if err := utils.DB.Where("id = ?", userId).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching user"})
		}
		return
	}

	// Set the user ID for the new post
	newPost.UserID = user.ID

	// Save the post in the database
	if err := utils.DB.Create(&newPost).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not create post"})
		return
	}

	c.JSON(http.StatusOK, newPost)
}
