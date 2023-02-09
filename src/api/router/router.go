package router

import (
	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
)

type Router interface {
	Load(*echo.Group)
}

type router struct {
}

func New() Router {
	return &router{}
}

func (r *router) Load(group *echo.Group) {
	group.GET("/docs/*", echoSwagger.WrapHandler)

	NewAuthRouter().Load(group)
	NewAccountRouter().Load(group)
	NewResourcesRouter().Load(group)
}
