package routes

import (
	"jokes/api/controllers"
	"jokes/api/middlewares"
	"jokes/lib"
)

// UserRoutes struct
type UserRoutes struct {
	logger         lib.Logger
	handler        lib.RequestHandler
	authMiddleware middlewares.JWTAuthMiddleware
	userController controllers.UserController
}

// Setup user routes
func (s UserRoutes) Setup() {
	root := s.handler.Gin.Group("/api/v1").Use(s.authMiddleware.Handler())
	{
		root.GET("/user", s.userController.Get)
	}
}

// NewUserRoutes creates new user controller
func NewUserRoutes(
	logger lib.Logger,
	handler lib.RequestHandler,
	authMiddleware middlewares.JWTAuthMiddleware,
	userController controllers.UserController,
) UserRoutes {
	return UserRoutes{
		logger:         logger,
		handler:        handler,
		authMiddleware: authMiddleware,
		userController: userController,
	}
}
