package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"jokes/constants"
	"jokes/domains"
	"jokes/lib"
	"jokes/models"
)

type JokeController struct {
	logger        lib.Logger
	service       domains.JokeService
	errorHandler  lib.ErrorHandler
	ratingService domains.RatingService
}

func NewJokeController(
	logger lib.Logger,
	service domains.JokeService,
	errorHandler lib.ErrorHandler,
	ratingService domains.RatingService,
) JokeController {
	return JokeController{
		logger:        logger,
		service:       service,
		errorHandler:  errorHandler,
		ratingService: ratingService,
	}
}

// Get
func (c JokeController) Get(ctx *gin.Context) {
	param := ctx.Param("id")
	if param == "" {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": constants.FailedToGetUrlParam,
		})
		return
	}
	id, err := strconv.Atoi(param)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": constants.FailedToParseNum,
		})
		return
	}

	joke, err := c.service.Get(uint(id))
	// TODO: Обработка ошибок
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	rating, err := c.ratingService.GetAverage(uint(id))
	// TODO: Обработка ошибок
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	resp := models.JokeGetResponse{
		ID:     joke.ID,
		Rating: rating,
		Text:   joke.Text,
	}

	ctx.JSON(http.StatusOK, resp)
}

// List
func (c JokeController) List(ctx *gin.Context) {
	jokes, err := c.service.List()
	// TODO: Обработка ошибок
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	var resp []models.JokeGetResponse
	for _, joke := range jokes {
		rating, err := c.ratingService.GetAverage(joke.ID)
		// TODO: Обработка ошибок
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}

		resp = append(resp, models.JokeGetResponse{
			ID:     joke.ID,
			Rating: rating,
			Text:   joke.Text,
		})
	}

	ctx.JSON(http.StatusOK, resp)
}

// Store
func (c JokeController) Store(ctx *gin.Context) {
	var joke models.JokeStoreRequest
	// TODO: Обработка ошибок
	if err := ctx.ShouldBindJSON(&joke); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid request body",
		})
		return
	}
	if err := c.errorHandler.IsValid(joke); err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"error": c.errorHandler.ParseValidationErrors(err),
		})
		return
	}

	if err := c.service.Store(&models.Joke{
		Text: joke.Text,
	}); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	resp := models.JokeStoreResponse{
		Text: joke.Text,
	}

	ctx.JSON(http.StatusOK, resp)
}
