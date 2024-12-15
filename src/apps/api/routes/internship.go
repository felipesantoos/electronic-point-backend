package routes

import (
	"eletronic_point/src/apps/api/dicontainer"
	"eletronic_point/src/apps/api/handlers"
	"eletronic_point/src/apps/api/middlewares"

	"github.com/labstack/echo/v4"
)

type internshipRouter struct {
	_handlers handlers.InternshipHandlers
}

func NewInternshipRouter() Router {
	services := dicontainer.InternshipServices()
	_handlers := handlers.NewInternshipHandlers(services)
	return &internshipRouter{_handlers}
}

func (this *internshipRouter) Load(rootEndpoint *echo.Group) {
	router := rootEndpoint.Group("/internships")
	router.GET("", middlewares.EnhanceContext(this._handlers.List))
	router.GET("/:id", middlewares.EnhanceContext(this._handlers.Get))
	router.POST("", middlewares.EnhanceContext(this._handlers.Create))
	router.PUT("/:id", middlewares.EnhanceContext(this._handlers.Update))
	router.DELETE("/:id", middlewares.EnhanceContext(this._handlers.Delete))
}
