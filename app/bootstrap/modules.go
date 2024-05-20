package bootstrap

import (
	"go.uber.org/fx"

	"jokes/api/controllers"
	"jokes/api/middlewares"
	"jokes/api/routes"
	"jokes/lib"
	"jokes/repository"
	"jokes/services"
)

var CommonModules = fx.Options(
	controllers.Module,
	routes.Module,
	lib.Module,
	services.Module,
	middlewares.Module,
	repository.Module,
)
