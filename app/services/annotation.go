package services

import (
	"gorm.io/gorm"

	"jokes/domains"
	"jokes/lib"
	"jokes/models"
	"jokes/repository"
)

type AnnotationService struct {
	logger     lib.Logger
	repository repository.AnnotationRepository
}

func NewAnnotationService(
	logger lib.Logger,
	repository repository.AnnotationRepository,
) domains.AnnotationService {
	return AnnotationService{
		logger:     logger,
		repository: repository,
	}
}

func (s AnnotationService) WithTrx(trxHandle *gorm.DB) domains.AnnotationService {
	s.repository = s.repository.WithTrx(trxHandle)
	return s
}

func (s AnnotationService) List(jokeID uint) ([]models.Annotation, error) {
	rating, err := s.repository.List(jokeID)
	return rating, err
}

func (s AnnotationService) Store(rating *models.Annotation) error {
	return s.repository.Store(rating)
}
