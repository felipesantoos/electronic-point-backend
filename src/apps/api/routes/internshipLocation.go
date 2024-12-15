package routes

import (
	"eletronic_point/src/apps/api/dicontainer"
	"eletronic_point/src/apps/api/handlers"
	"eletronic_point/src/apps/api/middlewares"

	"github.com/labstack/echo/v4"
)

type internshipLocationRouter struct {
	_handlers handlers.InternshipLocationHandlers
}

func NewInternshipLocationRouter() Router {
	services := dicontainer.InternshipLocationServices()
	_handlers := handlers.NewInternshipLocationHandlers(services)
	return &internshipLocationRouter{_handlers}
}

func (this *internshipLocationRouter) Load(rootEndpoint *echo.Group) {
	router := rootEndpoint.Group("/internship-locations")
	router.GET("", middlewares.EnhanceContext(this._handlers.List))
	router.GET("/:id", middlewares.EnhanceContext(this._handlers.Get))
	router.POST("", middlewares.EnhanceContext(this._handlers.Create))
	router.PUT("/:id", middlewares.EnhanceContext(this._handlers.Update))
	router.DELETE("/:id", middlewares.EnhanceContext(this._handlers.Delete))
}
