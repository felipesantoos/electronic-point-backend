package routes

import (
	"backend_template/src/apps/api/dicontainer"
	"backend_template/src/apps/api/handlers"
	"backend_template/src/apps/api/middlewares"

	"github.com/labstack/echo/v4"
)

type authRouter struct {
	handler handlers.AuthHandler
}

func NewAuthRouter() Router {
	usecase := dicontainer.AuthUseCase()
	handler := handlers.NewAuthHandler(usecase)
	return &authRouter{handler}
}

func (a *authRouter) Load(rootEndpoint *echo.Group) {
	router := rootEndpoint.Group("/auth")
	router.POST("/login", middlewares.EnhanceContext(a.handler.Login))
	router.POST("/logout", middlewares.EnhanceContext(a.handler.Logout))
	router.POST("/reset-password", middlewares.EnhanceContext(a.handler.AskPasswordResetMail))
	router.GET("/reset-password/:token", middlewares.EnhanceContext(a.handler.FindPasswordResetByToken))
	router.PUT("/reset-password/:token", middlewares.EnhanceContext(a.handler.UpdatePasswordByPasswordReset))
}
