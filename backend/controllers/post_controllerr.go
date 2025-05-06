package controllers

import (
	"bondly/models" // Import the Post model
	"bondly/utils"  // To access the DB connection
	"github.com/gin-gonic/gin"
	"net/http"
)

// GetPosts handles fetching all posts
func GetPosts(c *gin.Context) {
	var posts []models.Post
	if err := utils.DB.Find(&posts).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not fetch posts"})
		return
	}
	c.JSON(http.StatusOK, posts)
}

// GetPost handles fetching a specific post by ID
func GetPost(c *gin.Context) {
	id := c.Param("id")
	var post models.Post
	if err := utils.DB.Where("id = ?", id).First(&post).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Post not found"})
		return
	}
	c.JSON(http.StatusOK, post)
}

// CreatePost handles creating a new post
func CreatePost(c *gin.Context) {
	var post models.Post
	if err := c.ShouldBindJSON(&post); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	// Save the post to the database
	if err := utils.DB.Create(&post).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not create post"})
		return
	}

	c.JSON(http.StatusCreated, post)
}

// UpdatePost handles updating an existing post
func UpdatePost(c *gin.Context) {
	id := c.Param("id")
	var post models.Post
	if err := c.ShouldBindJSON(&post); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	// Update post in the database
	if err := utils.DB.Model(&models.Post{}).Where("id = ?", id).Updates(post).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not update post"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Post updated"})
}

// DeletePost handles deleting a post
func DeletePost(c *gin.Context) {
	id := c.Param("id")

	// Delete post from the database
	if err := utils.DB.Where("id = ?", id).Delete(&models.Post{}).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not delete post"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Post deleted"})
}
