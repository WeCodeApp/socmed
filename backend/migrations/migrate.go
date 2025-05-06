package migrations

import (
	"bondly/utils"
	"fmt"
	"log"
)

// MigrateSQL is the migration function to create the necessary tables.
func MigrateSQL() error {
	db := utils.DB

	// SQL statements to create the database and tables
	statements := []string{
		`CREATE DATABASE IF NOT EXISTS bondly;`,
		`USE bondly;`,
		`CREATE TABLE IF NOT EXISTS users (
			id INT AUTO_INCREMENT PRIMARY KEY,
			microsoft_id VARCHAR(255) NOT NULL UNIQUE,
			display_name VARCHAR(255),
			email VARCHAR(255),
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
		);`,
		`CREATE TABLE IF NOT EXISTS posts (
			id INT AUTO_INCREMENT PRIMARY KEY,
			user_id INT NOT NULL,
			content TEXT,
			image_url VARCHAR(255),
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			deleted_at TIMESTAMP DEFAULT NULL,  -- Add the deleted_at column
			FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
		);`,
		`CREATE TABLE IF NOT EXISTS likes (
			id INT AUTO_INCREMENT PRIMARY KEY,
			user_id INT NOT NULL,
			post_id INT NOT NULL,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
			FOREIGN KEY (post_id) REFERENCES posts(id) ON DELETE CASCADE,
			UNIQUE(user_id, post_id)
		);`,
	}

	// Execute each statement
	for _, stmt := range statements {
		if err := db.Exec(stmt).Error; err != nil {
			log.Printf("Error executing SQL statement: %s\n", stmt)
			return fmt.Errorf("error executing statement: %v", err)
		}
	}

	return nil
}
