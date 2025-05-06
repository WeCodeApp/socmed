package services

import (
	"bondly/config"
	"context"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/microsoft"
)

var oauth2Config oauth2.Config

// InitializeOAuth sets up the OAuth2 configuration
func InitializeOAuth() {
	oauth2Config = oauth2.Config{
		ClientID:     config.AppConfig.ClientID,
		ClientSecret: config.AppConfig.ClientSecret,
		RedirectURL:  config.AppConfig.RedirectURL, // Ensure this URL is correct
		Scopes:       []string{"openid", "profile", "email"},
		Endpoint:     microsoft.AzureADEndpoint(config.AppConfig.TenantID),
	}
}

// HandleMicrosoftLogin redirects the user to the Microsoft login page
func HandleMicrosoftLogin(c *gin.Context) {
	url := oauth2Config.AuthCodeURL("",
		oauth2.AccessTypeOffline,
		oauth2.SetAuthURLParam("response_mode", "query"),
	)
	// Make sure the redirect URL is pointing to the correct route
	c.Redirect(http.StatusFound, url)
}

// HandleMicrosoftCallback handles the callback from Microsoft after the user logs in
func HandleMicrosoftCallback(c *gin.Context) {
	code := c.DefaultQuery("code", "")
	if code == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No code in request"})
		return
	}

	// Exchange the code for an access token
	token, err := oauth2Config.Exchange(context.Background(), code)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to retrieve token from Microsoft"})
		return
	}

	// Get the user's information from Microsoft Graph
	client := oauth2Config.Client(context.Background(), token)
	resp, err := client.Get("https://graph.microsoft.com/v1.0/me")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to fetch user information"})
		return
	}
	defer resp.Body.Close()

	// Process user info
	var user MicrosoftUser
	if err := json.NewDecoder(resp.Body).Decode(&user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to decode user info"})
		return
	}

	// Create JWT Token for the user
	jwtToken, err := generateJWT(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to create JWT token"})
		return
	}

	// Redirect to your frontend with the token
	redirectURL := "http://localhost:5173/auth/callback?token=" + jwtToken
	c.Redirect(http.StatusFound, redirectURL)
}

// MicrosoftUser is a struct for the data returned from Microsoft Graph
type MicrosoftUser struct {
	ID          string `json:"id"`
	DisplayName string `json:"displayName"`
	Email       string `json:"mail"`
}

// generateJWT creates a JWT token using the user's info
func generateJWT(user MicrosoftUser) (string, error) {
	claims := jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(config.AppConfig.JWTSecret))
}
