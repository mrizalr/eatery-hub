package repository

import (
	"github.com/google/uuid"
	"github.com/mrizalr/eatery-hub/internal/models"
	"github.com/mrizalr/eatery-hub/internal/user"
	"gorm.io/gorm"
)

type mysqlUserRepository struct {
	db *gorm.DB
}

func NewMysqlUserRepository(db *gorm.DB) user.MysqlRepository {
	return &mysqlUserRepository{db}
}

func (r *mysqlUserRepository) Create(user models.User) (uuid.UUID, error) {
	tx := r.db.Create(&user)
	return user.ID, tx.Error
}
