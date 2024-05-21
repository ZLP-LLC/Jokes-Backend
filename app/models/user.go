package models

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type Role string

const (
	AdminRole Role = "admin"
	UserRole       = "user"
)

type User struct {
	gorm.Model
	Username *string `gorm:"unique"`
	Password string
	Role     Role
}

// TableName gives table name of model
func (u User) TableName() string {
	return "users"
}

func (user *User) BeforeSave(tx *gorm.DB) error {
	// Хэширование пароля
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hashPassword)

	return nil
}

// Структура ответа на GET запрос
type UserGetResponse struct {
	Username *string `json:"username"`
	Role     Role    `json:"role"`
}
