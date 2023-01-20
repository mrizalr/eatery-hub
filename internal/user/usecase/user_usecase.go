package usecase

import (
	"github.com/google/uuid"
	"github.com/mrizalr/eatery-hub/internal/models"
	"github.com/mrizalr/eatery-hub/internal/user"
)

type userUsecase struct {
	mysqlUserRepository user.MysqlRepository
}

func NewUserUsecase(mysqlUserRepository user.MysqlRepository) user.UserUsecase {
	return &userUsecase{mysqlUserRepository}
}

func (u *userUsecase) Register(user models.User) (uuid.UUID, error) {
	return u.mysqlUserRepository.Create(user)
}
