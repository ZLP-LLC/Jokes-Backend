package routes

import (
	"jokes/api/controllers"
	"jokes/api/middlewares"
	"jokes/lib"
)

// AnnotationRoutes struct
type AnnotationRoutes struct {
	logger         lib.Logger
	handler        lib.RequestHandler
	controller     controllers.AnnotationController
	authMiddleware middlewares.JWTAuthMiddleware
}

// Setup joke routes
func (s AnnotationRoutes) Setup() {
	root := s.handler.Gin.Group("/api/v1")
	{
		root.GET("/joke/:id/annotations", s.controller.List)
	}
}

// NewAnnotationRoutes creates new joke controller
func NewAnnotationRoutes(
	logger lib.Logger,
	handler lib.RequestHandler,
	controller controllers.AnnotationController,
	authMiddleware middlewares.JWTAuthMiddleware,
) AnnotationRoutes {
	return AnnotationRoutes{
		logger:         logger,
		handler:        handler,
		controller:     controller,
		authMiddleware: authMiddleware,
	}
}
