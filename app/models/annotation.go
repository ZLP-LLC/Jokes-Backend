package models

import (
	"gorm.io/gorm"
)

// Model
type Annotation struct {
	gorm.Model
	JokeID   uint
	Joke     Joke `gorm:"foreignKey:JokeID"`
	Text     string
	From     uint
	To       uint
	Approved bool `gorm:"default:false"`
}

// TableName gives table name of model
func (m Annotation) TableName() string {
	return "annotations"
}

// Get
type AnnotationGetResponse struct {
	ID   uint   `json:"id"`
	Text string `json:"text"`
	From uint   `json:"from"`
	To   uint   `json:"to"`
}

// Store
type AnnotationStoreRequest struct {
	Text string `json:"text"`
	From uint   `json:"from" validate:"min=1"`
	To   uint   `json:"to" validate:"min=1"`
}

type AnnotationStoreResponse struct {
	Text string `json:"text"`
	From uint   `json:"from"`
	To   uint   `json:"to"`
}
