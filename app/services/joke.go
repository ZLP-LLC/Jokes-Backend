package services

import (
	"gorm.io/gorm"

	"jokes/domains"
	"jokes/lib"
	"jokes/models"
	"jokes/repository"
)

type JokeService struct {
	logger     lib.Logger
	repository repository.JokeRepository
}

func NewJokeService(
	logger lib.Logger,
	repository repository.JokeRepository,
) domains.JokeService {
	return JokeService{
		logger:     logger,
		repository: repository,
	}
}

func (s JokeService) WithTrx(trxHandle *gorm.DB) domains.JokeService {
	s.repository = s.repository.WithTrx(trxHandle)
	return s
}

func (s JokeService) Get(id uint) (models.Joke, error) {
	joke, err := s.repository.Get(id)
	if err != nil {
		return models.Joke{}, err
	}
	return joke, nil
}

func (s JokeService) List() ([]models.Joke, error) {
	jokes, err := s.repository.List()
	if err != nil {
		return []models.Joke{}, err
	}
	return jokes, nil
}
