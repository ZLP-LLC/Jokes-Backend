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
	ratingService domains.RatingService
}

func NewJokeController(
	logger lib.Logger,
	service domains.JokeService,
	ratingService domains.RatingService,
) JokeController {
	return JokeController{
		logger:        logger,
		service:       service,
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
