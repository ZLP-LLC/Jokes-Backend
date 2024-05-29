package repository

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"jokes/lib"
	"jokes/models"
)

func TestUser(t *testing.T) {
	logger := lib.GetLogger()
	env := lib.NewEnv()
	db := lib.NewDatabase(env, logger)
	userRepository := NewUserRepository(logger, db)

	// Create
	var username string = "test"
	userExp := models.User{
		Username: &username,
		Password: "testtest",
		Role:     "user",
	}
	err := userRepository.Create(&userExp)
	assert.NoError(t, err)

	// Get
	userReal, err := userRepository.Get(userExp.ID)
	assert.NoError(t, err)
	assert.Equal(t, userExp.Username, userReal.Username)

	// GetByUsername
	userReal, err = userRepository.GetByUsername(userExp.Username)
	assert.NoError(t, err)
	assert.Equal(t, userExp.Username, userReal.Username)
}
