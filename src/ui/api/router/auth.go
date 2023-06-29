package router

import (
	"backend_template/src/ui/api/dicontainer"
	"backend_template/src/ui/api/handlers"
	"backend_template/src/ui/api/middlewares"

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
	router.POST("/login", middlewares.EnhanceContext(r.handler.Login))
	router.POST("/logout", middlewares.EnhanceContext(r.handler.Logout))
	router.POST("/reset-password", middlewares.EnhanceContext(r.handler.AskPasswordResetMail))
	router.GET("/reset-password/:token", middlewares.EnhanceContext(r.handler.FindPasswordResetByToken))
	router.PUT("/reset-password/:token", middlewares.EnhanceContext(r.handler.UpdatePasswordByPasswordReset))
}
