package routes

import (
	"jokes/api/controllers"
	"jokes/lib"
)

// JokeRoutes struct
type JokeRoutes struct {
	logger     lib.Logger
	handler    lib.RequestHandler
	controller controllers.JokeController
}

// Setup joke routes
func (s JokeRoutes) Setup() {
	root := s.handler.Gin.Group("/api/v1")
	{
		root.GET("/joke/:id", s.controller.Get)
		root.GET("/joke", s.controller.List)
	}
}

// NewJokeRoutes creates new joke controller
func NewJokeRoutes(
	logger lib.Logger,
	handler lib.RequestHandler,
	controller controllers.JokeController,
) JokeRoutes {
	return JokeRoutes{
		logger:     logger,
		handler:    handler,
		controller: controller,
	}
}
