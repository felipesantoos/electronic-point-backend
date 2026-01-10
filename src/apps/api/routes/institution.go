package routes

import (
	"eletronic_point/src/apps/api/dicontainer"
	"eletronic_point/src/apps/api/handlers"
	"eletronic_point/src/apps/api/middlewares"

	"github.com/labstack/echo/v4"
)

type institutionRouter struct {
	_handlers handlers.InstitutionHandlers
}

func NewInstitutionRouter() Router {
	services := dicontainer.InstitutionServices()
	_handlers := handlers.NewInstitutionHandlers(services)
	return &institutionRouter{_handlers}
}

func (this *institutionRouter) Load(rootEndpoint *echo.Group) {
	router := rootEndpoint.Group("/institutions")
	router.GET("", middlewares.EnhanceContext(this._handlers.List))
	router.GET("/:id", middlewares.EnhanceContext(this._handlers.Get))
	router.POST("", middlewares.EnhanceContext(this._handlers.Create))
	router.PUT("/:id", middlewares.EnhanceContext(this._handlers.Update))
	router.DELETE("/:id", middlewares.EnhanceContext(this._handlers.Delete))
}
