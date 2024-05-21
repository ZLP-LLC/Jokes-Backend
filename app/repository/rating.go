package repository

import (
	"gorm.io/gorm"

	"jokes/lib"
	"jokes/models"
)

type RatingRepository struct {
	logger   lib.Logger
	Database lib.Database
}

func NewRatingRepository(logger lib.Logger, db lib.Database) RatingRepository {
	return RatingRepository{
		logger:   logger,
		Database: db,
	}
}

func (r RatingRepository) WithTrx(trxHandle *gorm.DB) RatingRepository {
	if trxHandle == nil {
		r.logger.Error("Transaction Database not found in gin context. ")
		return r
	}
	r.Database.DB = trxHandle
	return r
}

func (r RatingRepository) Store(rating *models.Rating) error {
	return r.Database.Create(&rating).Error
}

func (r RatingRepository) GetAverage(joke_id uint) (float64, error) {
	var avg float64
	err := r.Database.Model(&models.Rating{}).
		Select("AVG(rating)").
		Where("joke_id = ?", joke_id).
		Group("joke_id").
		Scan(&avg).Error
	return avg, err
}
