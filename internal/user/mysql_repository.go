package user

import (
	"github.com/google/uuid"
	"github.com/mrizalr/eatery-hub/internal/models"
)

type MysqlRepository interface {
	Create(user models.User) (uuid.UUID, error)
}
