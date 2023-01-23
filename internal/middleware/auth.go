package middleware

import (
	"fmt"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"github.com/mrizalr/eatery-hub/config"
	"github.com/mrizalr/eatery-hub/internal/models"
)

func validateToken(tokenString string, cfg *config.Config) (*jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("token isn't valid : unexpected signing method %v: ", t.Method)
		}
		return []byte(cfg.Server.JwtSecretKey), nil
	})

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return nil, fmt.Errorf("token isn't valid")
	}

	return &claims, nil
}

func checkTokenExpire(claims *jwt.MapClaims) error {
	exp, ok := (*claims)["exp"].(float64)
	if !ok {
		return fmt.Errorf("token isn't valid")
	}

	if time.Now().Unix() > int64(exp) {
		return fmt.Errorf("Expired token")
	}

	return nil
}

func (mw *MiddlewareManager) Auth(c *fiber.Ctx) error {
	header := strings.Split(c.Get("Authorization"), " ")
	if len(header) != 2 || strings.ToLower(header[0]) != "bearer" {
		return c.Status(fiber.StatusUnauthorized).
			JSON(models.ResponseUnauthorized("invalid header format"))
	}

	tokenString := header[1]
	claims, err := validateToken(tokenString, mw.config)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).
			JSON(models.ResponseUnauthorized(err.Error()))
	}

	err = checkTokenExpire(claims)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).
			JSON(models.ResponseUnauthorized(err.Error()))
	}
	c.Locals("claims", claims)
	return c.Next()
}
