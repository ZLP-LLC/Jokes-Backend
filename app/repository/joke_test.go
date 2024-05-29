package repository

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"jokes/lib"
	"jokes/models"
)

func TestJoke(t *testing.T) {
	logger := lib.GetLogger()
	env := lib.NewEnv()
	db := lib.NewDatabase(env, logger)
	jokeRepository := NewJokeRepository(logger, db)

	// Store
	jokeExp := models.Joke{Text: "test1"}
	err := jokeRepository.Store(&jokeExp)
	assert.NoError(t, err)

	// Get
	var jokeReal models.Joke
	jokeReal, err = jokeRepository.Get(jokeExp.ID)
	assert.NoError(t, err)

	assert.Equal(t, jokeExp.Text, jokeReal.Text)

	// List
	_, err = jokeRepository.List()
	assert.NoError(t, err)
}
