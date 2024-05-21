package services

import (
	"gorm.io/gorm"

	"jokes/domains"
	"jokes/lib"
	"jokes/models"
	"jokes/repository"
)

type RatingService struct {
	logger     lib.Logger
	repository repository.RatingRepository
}

func NewRatingService(
	logger lib.Logger,
	repository repository.RatingRepository,
) domains.RatingService {
	return RatingService{
		logger:     logger,
		repository: repository,
	}
}

func (s RatingService) WithTrx(trxHandle *gorm.DB) domains.RatingService {
	s.repository = s.repository.WithTrx(trxHandle)
	return s
}

func (s RatingService) GetAverage(jokeId uint) (float64, error) {
	rating, err := s.repository.GetAverage(jokeId)
	return rating, err
}

func (s RatingService) Store(rating *models.Rating) error {
	return s.repository.Store(rating)
}
