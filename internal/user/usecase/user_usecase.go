package usecase

import (
	"github.com/google/uuid"
	"github.com/mrizalr/eatery-hub/config"
	"github.com/mrizalr/eatery-hub/internal/models"
	"github.com/mrizalr/eatery-hub/internal/user"
	"github.com/mrizalr/eatery-hub/pkg/utils"
)

type userUsecase struct {
	config              *config.Config
	mysqlUserRepository user.MysqlRepository
}

func NewUserUsecase(cfg *config.Config, mysqlUserRepository user.MysqlRepository) user.UserUsecase {
	return &userUsecase{cfg, mysqlUserRepository}
}

func (u *userUsecase) Register(user models.User) (uuid.UUID, error) {
	return u.mysqlUserRepository.Create(user)
}

func (u *userUsecase) Login(user models.User) (models.UserWithToken, error) {
	foundUser, err := u.mysqlUserRepository.FindByUsername(user.Username)
	if err != nil {
		return models.UserWithToken{}, err
	}

	err = foundUser.CompareHashAndPassword(user.Password)
	if err != nil {
		return models.UserWithToken{}, err
	}

	token, err := utils.GenerateJWTToken(&foundUser, u.config)
	if err != nil {
		return models.UserWithToken{}, err
	}

	response := models.UserLoginResponse{
		ID:          foundUser.ID,
		Username:    foundUser.Username,
		Email:       foundUser.Email,
		PhoneNumber: foundUser.PhoneNumber,
		PhotoURL:    foundUser.PhotoURL,
		CreatedAt:   foundUser.CreatedAt,
	}

	return models.UserWithToken{
		User:  response,
		Token: token,
	}, nil
}
