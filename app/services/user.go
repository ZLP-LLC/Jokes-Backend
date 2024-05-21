package services

import (
	"gorm.io/gorm"

	"jokes/domains"
	"jokes/lib"
	"jokes/models"
	"jokes/repository"
)

// UserService service layer
type UserService struct {
	logger     lib.Logger
	repository repository.UserRepository
}

// NewUserService creates a new userservice
func NewUserService(logger lib.Logger, repository repository.UserRepository) domains.UserService {
	return UserService{
		logger:     logger,
		repository: repository,
	}
}

// WithTrx delegates transaction to repository database
func (s UserService) WithTrx(trxHandle *gorm.DB) domains.UserService {
	s.repository = s.repository.WithTrx(trxHandle)
	return s
}

// Register call to register the user
func (s UserService) Register(q *models.RegisterRequest) (models.RegisterResponse, error) {
	user := models.User{
		Username: q.Username,
		Password: q.Password,
		Role:     models.UserRole,
	}

	if err := s.repository.Create(&user); err != nil {
		return models.RegisterResponse{}, err
	}

	newUser := models.RegisterResponse{
		Username: *user.Username,
		Role:     user.Role,
	}
	return newUser, nil
}

func (s UserService) GetByUsername(username *string) (*models.User, error) {
	return s.repository.GetByUsername(username)
}

func (s UserService) Get(id uint) (*models.User, error) {
	return s.repository.Get(id)
}
