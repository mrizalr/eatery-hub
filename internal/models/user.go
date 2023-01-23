package models

import (
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	_validator "github.com/mrizalr/eatery-hub/pkg/validator"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID          uuid.UUID `json:"id" gorm:"primary key"`
	Username    string    `json:"username" validate:"required,min=6" gorm:"type:varchar(30);index:idx_users_username,unique;not null"`
	Email       string    `json:"email" validate:"required,email" gorm:"type:varchar(50);index:idx_users_email,unique;not null"`
	Password    string    `json:"password,omitempty" validate:"required,min=6" gorm:"not null"`
	PhoneNumber string    `json:"phone_number" validate:"e164" gorm:"type:varchar(50);"`
	PhotoURL    string    `json:"photo_url"`
	IsAdmin     bool      `json:"is_admin,omitempty" gorm:"default:false"`
}

func (u *User) BeforeCreate(tx *gorm.DB) error {
	u.ID = uuid.New()

	return u.HashPassword()
}

func (u *User) Validate(v *validator.Validate) []string {
	errs := []string{}

	err := v.Struct(u)
	if err != nil {
		errs = _validator.TranslateErrors(err)
	}

	return errs
}

func (u *User) HashPassword() error {
	hashed, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	u.Password = string(hashed)
	return nil
}

func (u *User) CompareHashAndPassword(password string) error {
	return bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
}

type UserLoginResponse struct {
	ID          uuid.UUID `json:"id"`
	Username    string    `json:"username"`
	Email       string    `json:"email"`
	PhoneNumber string    `json:"phone_number"`
	PhotoURL    string    `json:"photo_url"`
	CreatedAt   time.Time `json:"created_at"`
}

type UserWithToken struct {
	User  UserLoginResponse `json:"user"`
	Token string            `json:"token"`
}
