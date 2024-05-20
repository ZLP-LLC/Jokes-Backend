package repository

import (
	"gorm.io/gorm"

	"jokes/lib"
	"jokes/models"
)

// UserRepository database structure
type UserRepository struct {
	logger   lib.Logger
	Database lib.Database
}

// NewUserRepository creates a new user repository
func NewUserRepository(logger lib.Logger, db lib.Database) UserRepository {
	return UserRepository{
		logger:   logger,
		Database: db,
	}
}

// WithTrx enables repository with transaction
func (r UserRepository) WithTrx(trxHandle *gorm.DB) UserRepository {
	if trxHandle == nil {
		r.logger.Error("Transaction Database not found in gin context. ")
		return r
	}
	r.Database.DB = trxHandle
	return r
}

func (r UserRepository) Create(model *models.User) error {
	return r.Database.Create(model).Error
}

func (r UserRepository) Get(id uint) (*models.User, error) {
	var user models.User
	err := r.Database.First(&user, id).Error
	return &user, err
}

func (r UserRepository) GetByUsername(username *string) (*models.User, error) {
	var user models.User
	err := r.Database.Where("username = ?", username).First(&user).Error
	return &user, err
}
