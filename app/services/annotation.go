package services

import (
	"errors"
	"fmt"
	"strings"

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

func (s AnnotationService) Store(annotation *models.Annotation) error {
	// Проверка на вшивость
	annotations, err := s.repository.List(annotation.JokeID)
	if err != nil {
		return err
	}

	var collides []string
	for _, model := range annotations {
		if ((annotation.From >= model.From) && (annotation.From <= model.To)) ||
			((annotation.To >= model.From) && (annotation.To <= model.To)) ||
			((annotation.From <= model.From) && (annotation.To >= model.To)) {
			collides = append(collides, fmt.Sprintf("%v", model.ID))
		}
	}
	if len(collides) != 0 {
		return errors.New(fmt.Sprintf("annotation collides with other annotations: %s", strings.Join(collides, ", ")))
	}

	return s.repository.Store(annotation)
}
