package routes

import (
	"eletronic_point/src/apps/api/dicontainer"
	"eletronic_point/src/apps/api/handlers"
	"eletronic_point/src/apps/api/middlewares"

	"github.com/labstack/echo/v4"
)

type resourcesRouter struct {
	handler handlers.ResourcesHandler
}

func NewResourcesRouter() Router {
	usecase := dicontainer.ResourcesServices()
	handler := handlers.NewResourcesHandler(usecase)
	return &resourcesRouter{handler}
}

func (r *resourcesRouter) Load(rootEndpoint *echo.Group) {
	router := rootEndpoint.Group("/res")
	router.GET("/account-roles", middlewares.EnhanceContext(r.handler.ListAccountRoles))
}
