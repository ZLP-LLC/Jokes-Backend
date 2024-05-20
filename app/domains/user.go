package domains

import (
	"jokes/models"

	"gorm.io/gorm"
)

type UserService interface {
	WithTrx(trxHandle *gorm.DB) UserService
	Register(q *models.RegisterRequest) (models.RegisterResponse, error)
	GetByUsername(username *string) (*models.User, error)
	Get(id uint) (*models.User, error)
}
