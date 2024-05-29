package routes

import (
	"jokes/api/controllers"
	"jokes/api/middlewares"
	"jokes/lib"
	"jokes/models"
)

// JokeRoutes struct
type JokeRoutes struct {
	logger         lib.Logger
	handler        lib.RequestHandler
	controller     controllers.JokeController
	authMiddleware middlewares.JWTAuthMiddleware
	roleMiddleware middlewares.RoleMiddleware
}

// Setup joke routes
func (s JokeRoutes) Setup() {
	root := s.handler.Gin.Group("/api/v1")
	{
		root.
			GET("/joke/:id", s.controller.Get).
			GET("/joke", s.controller.List)
		root.Use(s.authMiddleware.Handler()).Use(s.roleMiddleware.Handler([]models.Role{models.AdminRole})).
			POST("/joke", s.controller.Store)
	}
}

// NewJokeRoutes creates new joke controller
func NewJokeRoutes(
	logger lib.Logger,
	handler lib.RequestHandler,
	controller controllers.JokeController,
	authMiddleware middlewares.JWTAuthMiddleware,
	roleMiddleware middlewares.RoleMiddleware,
) JokeRoutes {
	return JokeRoutes{
		logger:         logger,
		handler:        handler,
		controller:     controller,
		authMiddleware: authMiddleware,
		roleMiddleware: roleMiddleware,
	}
}
