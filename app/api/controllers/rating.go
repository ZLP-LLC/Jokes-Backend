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

type RatingController struct {
	logger       lib.Logger
	service      domains.RatingService
	errorHandler lib.ErrorHandler
}

func NewRatingController(
	logger lib.Logger,
	service domains.RatingService,
	errorHandler lib.ErrorHandler,
) RatingController {
	return RatingController{
		logger:       logger,
		service:      service,
		errorHandler: errorHandler,
	}
}

// Store
func (c RatingController) Store(ctx *gin.Context) {
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

	var rating models.RatingStoreRequest
	// TODO: Обработка ошибок
	if err := ctx.ShouldBindJSON(&rating); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid request body",
		})
		return
	}
	if err := c.errorHandler.IsValid(rating); err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"error": c.errorHandler.ParseValidationErrors(err),
		})
		return
	}

	err = c.service.Store(&models.Rating{
		JokeID: uint(id),
		Rating: rating.Rating,
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	resp := models.RatingStoreResponse{
		Rating: rating.Rating,
	}

	ctx.JSON(http.StatusOK, resp)
}
