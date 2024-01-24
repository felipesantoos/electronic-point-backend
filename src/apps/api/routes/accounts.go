package routes

import (
	"backend_template/src/apps/api/dicontainer"
	"backend_template/src/apps/api/handlers"
	"backend_template/src/apps/api/middlewares"

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

func (a *accountRouter) Load(rootEndpoint *echo.Group) {
	adminRouter := rootEndpoint.Group("/admin/accounts")
	adminRouter.GET("", middlewares.EnhanceContext(a.handler.List))
	adminRouter.POST("", middlewares.EnhanceContext(a.handler.Create))
	router := rootEndpoint.Group("/accounts")
	router.GET("/profile", middlewares.EnhanceContext(a.handler.FindProfile))
	router.PUT("/profile", middlewares.EnhanceContext(a.handler.UpdateProfile))
	router.PUT("/update-password", middlewares.EnhanceContext(a.handler.UpdatePassword))
}
