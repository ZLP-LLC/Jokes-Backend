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

type AnnotationController struct {
	logger       lib.Logger
	service      domains.AnnotationService
	errorHandler lib.ErrorHandler
}

func NewAnnotationController(
	logger lib.Logger,
	service domains.AnnotationService,
	errorHandler lib.ErrorHandler,
) AnnotationController {
	return AnnotationController{
		logger:       logger,
		service:      service,
		errorHandler: errorHandler,
	}
}

// List
func (c AnnotationController) List(ctx *gin.Context) {
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

	annotations, err := c.service.List(uint(id))
	// TODO: Обработка ошибок
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	var resp []models.AnnotationGetResponse
	for _, annotation := range annotations {
		resp = append(resp, models.AnnotationGetResponse{
			ID:   annotation.ID,
			Text: annotation.Text,
			From: annotation.From,
			To:   annotation.To,
		})
	}

	ctx.JSON(http.StatusOK, resp)
}

// Store
func (c AnnotationController) Store(ctx *gin.Context) {
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

	var annotation models.AnnotationStoreRequest
	// TODO: Обработка ошибок
	if err := ctx.ShouldBindJSON(&annotation); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid request body",
		})
		return
	}
	if err := c.errorHandler.IsValid(annotation); err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"error": c.errorHandler.ParseValidationErrors(err),
		})
		return
	}

	err = c.service.Store(&models.Annotation{
		JokeID: uint(id),
		Text:   annotation.Text,
		From:   annotation.From,
		To:     annotation.To,
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	resp := models.AnnotationStoreResponse{
		Text: annotation.Text,
		From: annotation.From,
		To:   annotation.To,
	}

	ctx.JSON(http.StatusOK, resp)
}
