package server

import (
	"github.com/gofiber/fiber/v2"
	userHttp "github.com/mrizalr/eatery-hub/internal/user/delivery/http"
	userRepository "github.com/mrizalr/eatery-hub/internal/user/repository"
	userUsecase "github.com/mrizalr/eatery-hub/internal/user/usecase"
)

func (s *Server) MapHandlers(app *fiber.App) error {
	// this will be set a repository, usecase, and handler
	mysqlUserRepository := userRepository.NewMysqlUserRepository(s.db)

	userUsecase := userUsecase.NewUserUsecase(mysqlUserRepository)

	userHandler := userHttp.NewUserHandler(userUsecase)

	v1 := app.Group("/api/v1")
	health := v1.Group("/health")
	user := v1.Group("/user")

	// this will set a routes
	userHttp.MapRoutes(user, *userHandler)

	health.Get("", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "ok"})
	})

	return nil
}
