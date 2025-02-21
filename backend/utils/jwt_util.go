package utils

import (
	"errors"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

// GenerateToken creates a JWT token with userID and role
func GenerateToken(userID uint, role string) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userID,
		"role":    role,
		"exp":     time.Now().Add(time.Hour * 24).Unix(), // Token expires in 24 hours
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(os.Getenv("SECRET_KEY")))
}

// ExtractToken retrieves the JWT token from the request header
func ExtractToken(c *gin.Context) string {
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		return ""
	}
	// Token format: "Bearer <TOKEN>"
	parts := strings.Split(authHeader, " ")
	if len(parts) == 2 {
		return parts[1]
	}
	return ""
}

// ExtractTokenData validates the token and extracts user ID and role
func ExtractTokenData(c *gin.Context) (uint, string, error) {
	tokenString := ExtractToken(c)
	if tokenString == "" {
		return 0, "", errors.New("no token provided")
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("SECRET_KEY")), nil
	})

	if err != nil {
		return 0, "", errors.New("invalid token")
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		userID := uint(claims["user_id"].(float64))
		role := claims["role"].(string)
		return userID, role, nil
	}

	return 0, "", errors.New("could not extract user data")
}
