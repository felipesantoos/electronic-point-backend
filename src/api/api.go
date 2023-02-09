package api

import (
	_ "dit_backend/src/api/docs"
	"dit_backend/src/api/middlewares"
	"dit_backend/src/api/router"
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

type API interface {
	Serve()
}

type api struct {
	host         string
	port         int
	echoInstance *echo.Echo
}

// @title DIT Backend API
// @version 1.0
// @description DIT Backend template for new backend projects
// @contact.name DIT - IFAL
// @contact.email wmrn1@aluno.ifal.edu.br
// @BasePath /api
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func NewAPI(host string, port int) API {
	echoInstance := echo.New()
	return &api{host, port, echoInstance}
}

func (instance *api) Serve() {
	instance.setupMiddlewares()
	instance.loadRoutes()
	instance.start()
}

func (instance *api) setupMiddlewares() {
	instance.echoInstance.Use(middleware.Logger())
	instance.echoInstance.Use(middleware.Recover())
	instance.echoInstance.Use(middlewares.CORSMiddleware())
	instance.echoInstance.Use(middlewares.GuardMiddleware)
}

func (instance *api) rootGroup() *echo.Group {
	return instance.echoInstance.Group("/api")
}

func (instance *api) loadRoutes() {
	router := router.New()
	router.Load(instance.rootGroup())
}

func (instance *api) start() {
	address := fmt.Sprintf("%s:%d", instance.host, instance.port)
	err := instance.echoInstance.Start(address)
	instance.echoInstance.Logger.Fatal(err)
}

func Logger() zerolog.Logger {
	return log.With().Str("layer", "api").Logger()
}
