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
	logger  lib.Logger
	service domains.JokeService
}

func NewJokeController(
	logger lib.Logger,
	service domains.JokeService,
) JokeController {
	return JokeController{
		logger:  logger,
		service: service,
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

	resp := models.JokeGetResponse{
		ID:     joke.ID,
		Rating: 1, // TODO: Добавить рейтинг
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
		resp = append(resp, models.JokeGetResponse{
			ID:     joke.ID,
			Rating: 1, // TODO: Добавить рейтинг
			Text:   joke.Text,
		})
	}

	ctx.JSON(http.StatusOK, resp)
}
