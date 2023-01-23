package http

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mrizalr/eatery-hub/internal/models"
	"github.com/mrizalr/eatery-hub/internal/user"
	"github.com/mrizalr/eatery-hub/pkg/validator"
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

		errs := user.Validate(validator.Validator)
		if len(errs) > 0 {
			return c.Status(fiber.StatusBadRequest).
				JSON(models.ResponseBadRequest(errs))
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

func (h *userHandler) Login() fiber.Handler {
	type LoginRequest struct {
		Username string `json:"username" validate:"required,min=6"`
		Password string `json:"password" validate:"required,min=6"`
	}
	return func(c *fiber.Ctx) error {
		request := LoginRequest{}

		err := c.BodyParser(&request)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).
				JSON(models.ResponseBadRequest(err.Error()))
		}

		err = validator.Validator.Struct(request)
		if err != nil {
			errs := validator.TranslateErrors(err)
			return c.Status(fiber.StatusBadRequest).
				JSON(models.ResponseBadRequest(errs))
		}

		userWithToken, err := h.userUsecase.Login(models.User{
			Username: request.Username,
			Password: request.Password,
		})
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).
				JSON(models.ResponseUnauthorized(nil))
		}

		return c.Status(fiber.StatusOK).
			JSON(models.ResponseOK(userWithToken))
	}
}
