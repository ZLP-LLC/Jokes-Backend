package models

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username *string `gorm:"unique"`
	Password string
	Role     string
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
	Role     string  `json:"role"`
}
