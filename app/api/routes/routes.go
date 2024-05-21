package routes

import (
	"log"

	"go.uber.org/fx"
)

// Module exports dependency to container
var Module = fx.Options(
	fx.Provide(NewRoutes),
	fx.Provide(NewAuthRoutes),
	fx.Provide(NewUserRoutes),
	fx.Provide(NewJokeRoutes),
	fx.Provide(NewRatingRoutes),
)

// Routes contains multiple routes
type Routes []Route

// Route interface
type Route interface {
	Setup()
}

// NewRoutes sets up routes
func NewRoutes(
	authRoutes AuthRoutes,
	userRoutes UserRoutes,
	jokeRoutes JokeRoutes,
	ratingRoutes RatingRoutes,
) Routes {
	return Routes{
		authRoutes,
		userRoutes,
		jokeRoutes,
		ratingRoutes,
	}
}

// Setup all the route
func (r Routes) Setup() {
	log.Println("Setting up routes")
	for _, route := range r {
		route.Setup()
	}
}
