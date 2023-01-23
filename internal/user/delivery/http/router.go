package http

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mrizalr/eatery-hub/internal/middleware"
)

func MapRoutes(r fiber.Router, h userHandler, mw *middleware.MiddlewareManager) {
	r.Post("register", h.Register())
	r.Post("login", h.Login())
	r.Get("/check", mw.Auth, func(c *fiber.Ctx) error {
		return c.JSON(c.Locals("claims"))
	})
}
