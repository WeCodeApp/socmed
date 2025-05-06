package models

import "gorm.io/gorm"

// Post represents a post made by a user.
type Post struct {
	gorm.Model
	Content  string `json:"content"`
	ImageURL string `json:"image_url"`
	UserID   uint   `json:"user_id"`
	User     User   `json:"user"` // Association with the User model
}
