package repository

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"jokes/lib"
	"jokes/models"
)

func TestAnnotation(t *testing.T) {
	logger := lib.GetLogger()
	env := lib.NewEnv()
	db := lib.NewDatabase(env, logger)
	annotationRepository := NewAnnotationRepository(logger, db)

	// Store
	annotationExp := models.Annotation{Text: "test1", JokeID: 1, From: 1, To: 1}
	err := annotationRepository.Store(&annotationExp)
	assert.NoError(t, err)

	// List
	_, err = annotationRepository.List(annotationExp.JokeID)
	assert.NoError(t, err)
}
