package routes

import (
	"eletronic_point/src/apps/api/dicontainer"
	"eletronic_point/src/apps/api/handlers"
	"eletronic_point/src/apps/api/middlewares"

	"github.com/labstack/echo/v4"
)

type fileRouter struct {
	_handlers handlers.FileHandlers
}

func NewFileRouter() Router {
	services := dicontainer.FileServices()
	_handlers := handlers.NewFileHandlers(services)
	return &fileRouter{_handlers}
}

func (this *fileRouter) Load(rootEndpoint *echo.Group) {
	router := rootEndpoint.Group("/files")
	router.GET("/:name", middlewares.EnhanceContext(this._handlers.Get))
}
