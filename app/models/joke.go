package models

import (
	"gorm.io/gorm"
)

type Joke struct {
	gorm.Model
	Text   string
}

// TableName gives table name of model
func (m Joke) TableName() string {
	return "jokes"
}

// Структура ответа на GET запрос
type JokeGetResponse struct {
	ID     uint   `json:"id"`
	Rating uint   `json:"rating"`
	Text   string `json:"text"`
}
