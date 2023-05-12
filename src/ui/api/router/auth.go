package router

import (
	"backend_template/src/ui/api/dicontainer"
	"backend_template/src/ui/api/handlers"

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

func (instance *authRouter) Load(apiGroup *echo.Group) {
	router := apiGroup.Group("/auth")
	router.POST("/login", instance.handler.Login)
	router.POST("/logout", instance.handler.Logout)
	router.POST("/reset-password", instance.handler.AskPasswordResetMail)
	router.GET("/reset-password/:token", instance.handler.FindPasswordResetByToken)
	router.PUT("/reset-password/:token", instance.handler.UpdatePasswordByPasswordReset)
}
