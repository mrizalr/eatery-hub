package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID          uuid.UUID `json:"id" gorm:"primary key"`
	Username    string    `json:"username" gorm:"type:varchar(30);index:idx_users_username,unique;not null"`
	Email       string    `json:"email" gorm:"type:varchar(50);index:idx_users_email,unique;not null"`
	Password    string    `json:"password" gorm:"not null"`
	PhoneNumber string    `json:"phone_number" gorm:"type:varchar(50);"`
	PhotoURL    string    `json:"photo_url"`
	IsAdmin     bool      `json:"is_admin" gorm:"default:false"`
}

func (u *User) BeforeCreate(tx *gorm.DB) error {
	u.ID = uuid.New()
	return nil
}
