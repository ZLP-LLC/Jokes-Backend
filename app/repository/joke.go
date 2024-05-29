package repository

import (
	"gorm.io/gorm"

	"jokes/lib"
	"jokes/models"
)

type JokeRepository struct {
	logger   lib.Logger
	Database lib.Database
}

func NewJokeRepository(logger lib.Logger, db lib.Database) JokeRepository {
	return JokeRepository{
		logger:   logger,
		Database: db,
	}
}

func (r JokeRepository) WithTrx(trxHandle *gorm.DB) JokeRepository {
	if trxHandle == nil {
		r.logger.Error("Transaction Database not found in gin context. ")
		return r
	}
	r.Database.DB = trxHandle
	return r
}

func (r JokeRepository) List() ([]models.Joke, error) {
	var jokes []models.Joke
	err := r.Database.Find(&jokes).Error
	return jokes, err
}

func (r JokeRepository) Get(id uint) (models.Joke, error) {
	var joke models.Joke
	err := r.Database.Where("id = ?", id).First(&joke).Error
	return joke, err
}

func (r JokeRepository) Store(joke *models.Joke) error {
	return r.Database.Create(&joke).Error
}
