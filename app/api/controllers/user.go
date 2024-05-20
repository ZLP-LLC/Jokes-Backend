package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"jokes/constants"
	"jokes/domains"
	"jokes/lib"
	"jokes/models"
)

// UserController struct
type UserController struct {
	logger  lib.Logger
	service domains.UserService
}

// NewUserController creates new controller
func NewUserController(
	logger lib.Logger,
	service domains.UserService,
) UserController {
	return UserController{
		logger:  logger,
		service: service,
	}
}

func (uc UserController) Get(c *gin.Context) {
	userID, ok := c.Get(constants.UserID)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to get user",
		})
		return
	}

	user, err := uc.service.Get(userID.(uint))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to get user",
		})
		return
	}

	response := models.UserGetResponse{
		Username: user.Username,
		Role:     user.Role,
	}

	c.JSON(http.StatusOK, response)
}
