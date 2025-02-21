package routes

import (
	"eletronic_point/src/apps/api/dicontainer"
	"eletronic_point/src/apps/api/handlers"
	"eletronic_point/src/apps/api/middlewares"

	"github.com/labstack/echo/v4"
)

type studentRouter struct {
	_handlers handlers.StudentHandlers
}

func NewStudentRouter() Router {
	services := dicontainer.StudentServices()
	_handlers := handlers.NewStudentHandlers(services)
	return &studentRouter{_handlers}
}

func (this *studentRouter) Load(rootEndpoint *echo.Group) {
	router := rootEndpoint.Group("/students")
	router.GET("", middlewares.EnhanceContext(this._handlers.List))
	router.GET("/:id", middlewares.EnhanceContext(this._handlers.Get))
	router.POST("", middlewares.EnhanceContext(this._handlers.Create))
	router.PUT("/:id", middlewares.EnhanceContext(this._handlers.Update))
	router.DELETE("/:id", middlewares.EnhanceContext(this._handlers.Delete))
}
