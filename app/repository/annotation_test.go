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

	// Test cases
	tests := map[string]struct {
		input  string
		result string
	}{
		"empty string": {
			input:  "",
			result: "",
		},
		"one character": {
			input:  "x",
			result: "x",
		},
		"long text": {
			input:  "joke text",
			result: "joke text",
		},
	}
	for name, test := range tests {
		test := test
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			// Store
			annotationExp := models.Annotation{Text: test.input, JokeID: 1, From: 1, To: 1}
			err := annotationRepository.Store(&annotationExp)
			assert.NoError(t, err)

			if got, expected := test.result, annotationExp.Text; got != expected {
				t.Fatalf("Store(%q) returned %q; expected %q", test.input, got, expected)
			}
		})
	}

	// Functions test
	// Store
	annotationExp := models.Annotation{Text: "test1", JokeID: 1, From: 1, To: 1}
	err := annotationRepository.Store(&annotationExp)
	assert.NoError(t, err)

	// List
	_, err = annotationRepository.List(annotationExp.JokeID)
	assert.NoError(t, err)
}
