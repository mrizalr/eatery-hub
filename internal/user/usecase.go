package user

import (
	"github.com/google/uuid"
	"github.com/mrizalr/eatery-hub/internal/models"
)

type UserUsecase interface {
	Register(user models.User) (uuid.UUID, error)
}
