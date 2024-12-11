package routes

import (
	"eletronic_point/src/apps/api/dicontainer"
	"eletronic_point/src/apps/api/handlers"
	"eletronic_point/src/apps/api/middlewares"

	"github.com/labstack/echo/v4"
)

type studentRouter struct {
	handler handlers.StudentHandler
}

func NewStudentRouter() Router {
	usecase := dicontainer.StudentUseCase()
	handler := handlers.NewStudentHandler(usecase)
	return &studentRouter{handler}
}

func (r *studentRouter) Load(rootEndpoint *echo.Group) {
	router := rootEndpoint.Group("/students")
	router.GET("", middlewares.EnhanceContext(r.handler.List))
	router.GET("/:id", middlewares.EnhanceContext(r.handler.Get))
	router.POST("", middlewares.EnhanceContext(r.handler.Create))
	router.PUT("/:id", middlewares.EnhanceContext(r.handler.Update))
	router.DELETE("/:id", middlewares.EnhanceContext(r.handler.Delete))
}
