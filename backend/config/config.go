package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

type Config struct {
	DBUser       string
	DBPassword   string
	DBHost       string
	DBPort       string
	DBName       string
	ClientID     string
	ClientSecret string
	RedirectURL  string
	TenantID     string
	JWTSecret    string
}

var AppConfig Config

// LoadConfig loads the configuration from the .env file
func LoadConfig() Config {
	// Load .env file
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	// Get values from environment variables
	AppConfig = Config{
		DBUser:       os.Getenv("DB_USER"),
		DBPassword:   os.Getenv("DB_PASSWORD"),
		DBHost:       os.Getenv("DB_HOST"),
		DBPort:       os.Getenv("DB_PORT"),
		DBName:       os.Getenv("DB_NAME"),
		ClientID:     os.Getenv("CLIENT_ID"),
		ClientSecret: os.Getenv("CLIENT_SECRET"),
		RedirectURL:  os.Getenv("REDIRECT_URL"),
		TenantID:     os.Getenv("TENANT_ID"),
		JWTSecret:    os.Getenv("JWT_SECRET"),
	}

	return AppConfig
}
