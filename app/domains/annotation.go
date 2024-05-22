package domains

import (
	"gorm.io/gorm"

	"jokes/models"
)

type AnnotationService interface {
	WithTrx(trxHandle *gorm.DB) AnnotationService
	List(jokeID uint) ([]models.Annotation, error)
	Store(annotation *models.Annotation) error
}
