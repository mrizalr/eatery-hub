package http

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mrizalr/eatery-hub/internal/models"
	"github.com/mrizalr/eatery-hub/internal/user"
)

type userHandler struct {
	userUsecase user.UserUsecase
}

func NewUserHandler(userUsecase user.UserUsecase) *userHandler {
	return &userHandler{userUsecase}
}

func (h *userHandler) Register() fiber.Handler {
	return func(c *fiber.Ctx) error {
		user := models.User{}
		err := c.BodyParser(&user)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).
				JSON(models.ResponseBadRequest(err.Error()))
		}

		userID, err := h.userUsecase.Register(user)
		if err != nil {
			return c.Status(fiber.StatusBadGateway).
				JSON(models.ResponseBadGateway(err.Error()))
		}

		return c.Status(fiber.StatusCreated).
			JSON(models.ResponseCreated(fiber.Map{
				"user_id": userID,
			}))
	}
}
