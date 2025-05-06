package main

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"log"
)

func generateJWTSecret() string {
	// Generate 32 random bytes
	secret := make([]byte, 32)
	_, err := rand.Read(secret)
	if err != nil {
		log.Fatalf("Error generating random bytes: %v", err)
	}

	// Encode the bytes to base64
	encodedSecret := base64.StdEncoding.EncodeToString(secret)
	return encodedSecret
}

func main() {
	jwtSecret := generateJWTSecret()
	fmt.Println("Generated JWT Secret:", jwtSecret)
}
