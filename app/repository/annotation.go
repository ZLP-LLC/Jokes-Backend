package repository

import (
	"gorm.io/gorm"

	"jokes/lib"
	"jokes/models"
)

type AnnotationRepository struct {
	logger   lib.Logger
	Database lib.Database
}

func NewAnnotationRepository(logger lib.Logger, db lib.Database) AnnotationRepository {
	return AnnotationRepository{
		logger:   logger,
		Database: db,
	}
}

func (r AnnotationRepository) WithTrx(trxHandle *gorm.DB) AnnotationRepository {
	if trxHandle == nil {
		r.logger.Error("Transaction Database not found in gin context. ")
		return r
	}
	r.Database.DB = trxHandle
	return r
}

func (r AnnotationRepository) List(jokeID uint) ([]models.Annotation, error) {
	var annotations []models.Annotation
	err := r.Database.
		Where("joke_id = ?", jokeID).
		Where("approved = ?", true).
		Find(&annotations).Error
	return annotations, err
}

func (r AnnotationRepository) Store(annotation *models.Annotation) error {
	return r.Database.Create(&annotation).Error
}
