package router

import (
	"backend_template/src/ui/api/dicontainer"
	"backend_template/src/ui/api/handlers"

	"github.com/labstack/echo/v4"
)

type accountRouter struct {
	handler handlers.AccountHandler
}

func NewAccountRouter() Router {
	service := dicontainer.AccountUseCase()
	handler := handlers.NewAccountHandler(service)
	return &accountRouter{handler}
}

func (instance *accountRouter) Load(group *echo.Group) {
	adminRouter := group.Group("/admin/accounts")
	adminRouter.GET("", instance.handler.List)
	adminRouter.POST("", instance.handler.Create)
	router := group.Group("/accounts")
	router.GET("/profile", instance.handler.FindProfile)
	router.PUT("/profile", instance.handler.UpdateProfile)
	router.PUT("/update-password", instance.handler.UpdatePassword)
}
