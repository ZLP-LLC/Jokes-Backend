package routes

import (
	"jokes/api/controllers"
	"jokes/api/middlewares"
	"jokes/lib"
)

// RatingRoutes struct
type RatingRoutes struct {
	logger         lib.Logger
	handler        lib.RequestHandler
	controller     controllers.RatingController
	authMiddleware middlewares.JWTAuthMiddleware
}

// Setup joke routes
func (s RatingRoutes) Setup() {
	root := s.handler.Gin.Group("/api/v1").Use(s.authMiddleware.Handler())
	{
		root.POST("/joke/:id/rating", s.controller.Store)
	}
}

// NewRatingRoutes creates new joke controller
func NewRatingRoutes(
	logger lib.Logger,
	handler lib.RequestHandler,
	controller controllers.RatingController,
	authMiddleware middlewares.JWTAuthMiddleware,
) RatingRoutes {
	return RatingRoutes{
		logger:         logger,
		handler:        handler,
		controller:     controller,
		authMiddleware: authMiddleware,
	}
}
