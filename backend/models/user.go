package models

import "gorm.io/gorm"

// User represents a user in the system.
type User struct {
	gorm.Model
	DisplayName string `json:"display_name"`
	Email       string `json:"email"`
	MicrosoftID string `json:"microsoft_id"`
}
