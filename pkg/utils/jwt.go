package utils

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/mrizalr/eatery-hub/config"
	"github.com/mrizalr/eatery-hub/internal/models"
)

func GenerateJWTToken(user *models.User, cfg *config.Config) (string, error) {
	claims := jwt.MapClaims{
		"id":       user.ID,
		"username": user.Username,
		"email":    user.Email,
		"is_admin": user.IsAdmin,
		"exp":      time.Now().Add(24 * time.Hour).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(cfg.Server.JwtSecretKey))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
