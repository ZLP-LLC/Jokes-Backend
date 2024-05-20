package controllers

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"net/http"

	"jokes/domains"
	"jokes/lib"
	"jokes/models"
)

// JWTAuthController struct
type JWTAuthController struct {
	logger      lib.Logger
	errors      lib.ErrorHandler
	service     domains.AuthService
	userService domains.UserService
}

// NewJWTAuthController creates new controller
func NewJWTAuthController(
	logger lib.Logger,
	errors lib.ErrorHandler,
	service domains.AuthService,
	userService domains.UserService,
) JWTAuthController {
	return JWTAuthController{
		logger:      logger,
		errors:      errors,
		service:     service,
		userService: userService,
	}
}

// Login
func (jwt JWTAuthController) Login(c *gin.Context) {
	// Парсинг запроса
	var q models.LoginRequest
	if err := c.ShouldBindJSON(&q); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid request body",
		})
		return
	}

	if err := jwt.errors.IsValid(q); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"error": jwt.errors.ParseValidationErrors(err),
		})
		return
	}

	// Нахождение пользователя по username пользователя
	user, err := jwt.userService.GetByUsername(q.Username)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Invalid username or password",
		})
		return
	}

	// Сравнение хэша и пароля
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(q.Password))
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Invalid username or password",
		})
		return
	}

	// Получение токена
	token, err := jwt.service.CreateToken(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to create token",
		})
		return
	}

	// Создание ответа
	response := models.LoginResponse{
		Token: token,
	}

	// Отправка токена
	c.JSON(http.StatusOK, response)
}

// Register
func (jwt JWTAuthController) Register(c *gin.Context) {
	// Парсинг запроса
	var q models.RegisterRequest
	if err := c.ShouldBindJSON(&q); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid request body",
		})
		return
	}

	if err := jwt.errors.IsValid(q); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"error": jwt.errors.ParseValidationErrors(err),
		})
		return
	}

	// Регистрация пользователя
	user, err := jwt.userService.Register(&q)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to register user",
		})
		return
	}

	// Отправка ответа
	c.JSON(http.StatusOK, user)
}
