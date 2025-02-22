package routes

import (
	"eletronic_point/src/apps/api/dicontainer"
	"eletronic_point/src/apps/api/handlers"
	"eletronic_point/src/apps/api/middlewares"

	"github.com/labstack/echo/v4"
)

type reportsRouter struct {
	_handlers handlers.ReportsHandlers
}

func NewReportsRouter() Router {
	services := dicontainer.ReportsServices()
	_handlers := handlers.NewReportsHandlers(services)
	return &reportsRouter{_handlers}
}

func (this *reportsRouter) Load(rootEndpoint *echo.Group) {
	router := rootEndpoint.Group("/reports")
	router.GET("/time-records-by-student", middlewares.EnhanceContext(this._handlers.GetTimeRecordsByStudent))
}
