package routes

import (
	"eletronic_point/src/apps/api/dicontainer"
	"eletronic_point/src/apps/api/handlers"
	"eletronic_point/src/apps/api/middlewares"

	"github.com/labstack/echo/v4"
)

type timeRecordStatusRouter struct {
	_handlers handlers.TimeRecordStatusHandlers
}

func NewTimeRecordStatusRouter() Router {
	services := dicontainer.TimeRecordStatusServices()
	_handlers := handlers.NewTimeRecordStatusHandlers(services)
	return &timeRecordStatusRouter{_handlers}
}

func (this *timeRecordStatusRouter) Load(rootEndpoint *echo.Group) {
	router := rootEndpoint.Group("/time-record-status")
	router.GET("", middlewares.EnhanceContext(this._handlers.List))
	router.GET("/:id", middlewares.EnhanceContext(this._handlers.Get))
}
