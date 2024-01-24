package routes

import (
	"backend_template/src/apps/api/dicontainer"
	"backend_template/src/apps/api/handlers"
	"backend_template/src/apps/api/middlewares"

	"github.com/labstack/echo/v4"
)

type resourcesRouter struct {
	handler handlers.ResourcesHandler
}

func NewResourcesRouter() Router {
	usecase := dicontainer.ResourcesUseCase()
	handler := handlers.NewResourcesHandler(usecase)
	return &resourcesRouter{handler}
}

func (r *resourcesRouter) Load(rootEndpoint *echo.Group) {
	router := rootEndpoint.Group("/res")
	router.GET("/account-roles", middlewares.EnhanceContext(r.handler.ListAccountRoles))
}
