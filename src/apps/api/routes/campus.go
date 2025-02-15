package routes

import (
	"eletronic_point/src/apps/api/dicontainer"
	"eletronic_point/src/apps/api/handlers"
	"eletronic_point/src/apps/api/middlewares"

	"github.com/labstack/echo/v4"
)

type campusRouter struct {
	_handlers handlers.CampusHandlers
}

func NewCampusRouter() Router {
	services := dicontainer.CampusServices()
	_handlers := handlers.NewCampusHandlers(services)
	return &campusRouter{_handlers}
}

func (this *campusRouter) Load(rootEndpoint *echo.Group) {
	router := rootEndpoint.Group("/campus")
	router.GET("", middlewares.EnhanceContext(this._handlers.List))
}
