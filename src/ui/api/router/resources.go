package router

import (
	"backend_template/src/ui/api/dicontainer"
	"backend_template/src/ui/api/handlers"

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

func (r *resourcesRouter) Load(apiGroup *echo.Group) {
	router := apiGroup.Group("/res")
	router.GET("/account-roles", r.handler.ListAccountRoles)
}
