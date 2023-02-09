package router

import (
	"dit_backend/src/api/dicontainer"
	"dit_backend/src/api/handlers"

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

func (instance *resourcesRouter) Load(apiGroup *echo.Group) {
	router := apiGroup.Group("/res")
	router.GET("/account-roles", instance.handler.ListAccountRoles)
}
