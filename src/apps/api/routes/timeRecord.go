package routes

import (
	"eletronic_point/src/apps/api/dicontainer"
	"eletronic_point/src/apps/api/handlers"
	"eletronic_point/src/apps/api/middlewares"

	"github.com/labstack/echo/v4"
)

type timeRecordRouter struct {
	_handlers handlers.TimeRecordHandlers
}

func NewTimeRecordRouter() Router {
	services := dicontainer.TimeRecordServices()
	_handlers := handlers.NewTimeRecordHandlers(services)
	return &timeRecordRouter{_handlers}
}

func (this *timeRecordRouter) Load(rootEndpoint *echo.Group) {
	router := rootEndpoint.Group("/time-records")
	router.GET("", middlewares.EnhanceContext(this._handlers.List))
	router.GET("/:id", middlewares.EnhanceContext(this._handlers.Get))
	router.POST("", middlewares.EnhanceContext(this._handlers.Create))
	router.PUT("/:id", middlewares.EnhanceContext(this._handlers.Update))
	router.DELETE("/:id", middlewares.EnhanceContext(this._handlers.Delete))
}
