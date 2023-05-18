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

func (r *accountRouter) Load(group *echo.Group) {
	adminRouter := group.Group("/admin/accounts")
	adminRouter.GET("", r.handler.List)
	adminRouter.POST("", r.handler.Create)
	router := group.Group("/accounts")
	router.GET("/profile", r.handler.FindProfile)
	router.PUT("/profile", r.handler.UpdateProfile)
	router.PUT("/update-password", r.handler.UpdatePassword)
}
