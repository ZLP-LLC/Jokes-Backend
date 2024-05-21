package domains

import (
	"gorm.io/gorm"

	"jokes/models"
)

type RatingService interface {
	WithTrx(trxHandle *gorm.DB) RatingService
	GetAverage(jokeId uint) (float64, error)
	Store(rating *models.Rating) error
}
