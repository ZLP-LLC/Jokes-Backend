package repository

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"jokes/lib"
	"jokes/models"
)

func TestRating(t *testing.T) {
	logger := lib.GetLogger()
	env := lib.NewEnv()
	db := lib.NewDatabase(env, logger)
	ratingRepository := NewRatingRepository(logger, db)

	// Store
	ratingExp := models.Rating{JokeID: 1, Rating: 1}
	err := ratingRepository.Store(&ratingExp)
	assert.NoError(t, err)

	// GetAverage
	_, err = ratingRepository.GetAverage(ratingExp.JokeID)
	assert.NoError(t, err)
}
