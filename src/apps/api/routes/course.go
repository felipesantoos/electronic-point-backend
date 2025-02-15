package routes

import (
	"eletronic_point/src/apps/api/dicontainer"
	"eletronic_point/src/apps/api/handlers"
	"eletronic_point/src/apps/api/middlewares"

	"github.com/labstack/echo/v4"
)

type courseRouter struct {
	_handlers handlers.CourseHandlers
}

func NewCourseRouter() Router {
	services := dicontainer.CourseServices()
	_handlers := handlers.NewCourseHandlers(services)
	return &courseRouter{_handlers}
}

func (this *courseRouter) Load(rootEndpoint *echo.Group) {
	router := rootEndpoint.Group("/courses")
	router.GET("", middlewares.EnhanceContext(this._handlers.List))
}
