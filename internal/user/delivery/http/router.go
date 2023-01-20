package http

import "github.com/gofiber/fiber/v2"

func MapRoutes(r fiber.Router, h userHandler) {
	r.Post("", h.Register())
}
