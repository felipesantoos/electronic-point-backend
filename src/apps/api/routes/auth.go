package routes

import (
	"eletronic_point/src/apps/api/dicontainer"
	"eletronic_point/src/apps/api/handlers"
	"eletronic_point/src/apps/api/middlewares"

	"github.com/labstack/echo/v4"
)

type authRouter struct {
	handler handlers.AuthHandler
}

func NewAuthRouter() Router {
	usecase := dicontainer.AuthServices()
	handler := handlers.NewAuthHandler(usecase)
	return &authRouter{handler}
}

func (a *authRouter) Load(rootEndpoint *echo.Group) {
	router := rootEndpoint.Group("/auth")
	router.POST("/login", middlewares.EnhanceContext(a.handler.Login))
	router.POST("/refresh", middlewares.EnhanceContext(a.handler.Refresh))
	router.POST("/logout", middlewares.EnhanceContext(a.handler.Logout))
	router.POST("/reset-password", middlewares.EnhanceContext(a.handler.AskPasswordResetMail))
	router.GET("/reset-password/:token", middlewares.EnhanceContext(a.handler.FindPasswordResetByToken))
	router.PUT("/reset-password/:token", middlewares.EnhanceContext(a.handler.UpdatePasswordByPasswordReset))
}
