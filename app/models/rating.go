package models

import (
	"gorm.io/gorm"
)

type Rating struct {
	gorm.Model
	Rating float64
	JokeID uint
	Joke   Joke `gorm:"foreignKey:JokeID"`
	UserID uint
	User   User `gorm:"foreignKey:UserID"`
}

// TableName gives table name of model
func (m Rating) TableName() string {
	return "ratings"
}

type RatingStoreResponse struct {
	Rating float64 `json:"rating"`
}

type RatingStoreRequest struct {
	Rating float64 `json:"rating" validate:"min=0,max=1"`
}
