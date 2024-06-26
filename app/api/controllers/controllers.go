package controllers

import "go.uber.org/fx"

// Module exported for initializing application
var Module = fx.Options(
	fx.Provide(NewJWTAuthController),
	fx.Provide(NewUserController),
	fx.Provide(NewJokeController),
	fx.Provide(NewRatingController),
	fx.Provide(NewAnnotationController),
)
