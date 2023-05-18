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

func (r *authRouter) Load(apiGroup *echo.Group) {
	router := apiGroup.Group("/auth")
	router.POST("/login", r.handler.Login)
	router.POST("/logout", r.handler.Logout)
	router.POST("/reset-password", r.handler.AskPasswordResetMail)
	router.GET("/reset-password/:token", r.handler.FindPasswordResetByToken)
	router.PUT("/reset-password/:token", r.handler.UpdatePasswordByPasswordReset)
}
