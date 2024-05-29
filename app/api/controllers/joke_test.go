package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"

	"jokes/lib"
	"jokes/models"
	"jokes/repository"
	"jokes/services"
)

func setupRouter() *gin.Engine {
	router := gin.Default()
	return router
}

func TestGetJokes(t *testing.T) {
	router := setupRouter()
	logger := lib.GetLogger()
	errorHandler := lib.NewErrorHandler()
	env := lib.NewEnv()
	db := lib.NewDatabase(env, logger)
	jokeRepository := repository.NewJokeRepository(logger, db)
	jokeService := services.NewJokeService(logger, jokeRepository)
	ratingRepository := repository.NewRatingRepository(logger, db)
	ratingService := services.NewRatingService(logger, ratingRepository)
	jokeController := NewJokeController(logger, jokeService, errorHandler, ratingService)

	joke := &models.Joke{Text: "test1"}
	db.Create(joke)
	jokeRespExp := models.JokeGetResponse{ID: joke.ID, Text: "test1"}

	router.GET("/joke/:id", jokeController.Get)

	req, _ := http.NewRequest("GET", fmt.Sprintf("/joke/%v", joke.ID), nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	var jokeRespReal models.JokeGetResponse
	err := json.Unmarshal(w.Body.Bytes(), &jokeRespReal)
	assert.NoError(t, err)
	assert.Equal(t, jokeRespExp, jokeRespReal)
}
