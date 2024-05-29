package domains

import (
	"gorm.io/gorm"

	"jokes/models"
)

type JokeService interface {
	WithTrx(trxHandle *gorm.DB) JokeService
	List() ([]models.Joke, error)
	Get(id uint) (models.Joke, error)
	Store(joke *models.Joke) error
}
