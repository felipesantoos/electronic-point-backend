package router

import (
	"backend_template/src/ui/api/dicontainer"
	"backend_template/src/ui/api/handlers"
	"backend_template/src/ui/api/middlewares"

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

func (r *accountRouter) Load(group *echo.Group) {
	adminRouter := group.Group("/admin/accounts")
	adminRouter.GET("", middlewares.EnhanceContext(r.handler.List))
	adminRouter.POST("", middlewares.EnhanceContext(r.handler.Create))
	router := group.Group("/accounts")
	router.GET("/profile", middlewares.EnhanceContext(r.handler.FindProfile))
	router.PUT("/profile", middlewares.EnhanceContext(r.handler.UpdateProfile))
	router.PUT("/update-password", middlewares.EnhanceContext(r.handler.UpdatePassword))
}
