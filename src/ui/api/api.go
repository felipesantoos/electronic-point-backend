package api

import (
	"backend_template/src/core"
	_ "backend_template/src/ui/api/docs"
	"backend_template/src/ui/api/middlewares"
	"backend_template/src/ui/api/router"
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/rs/zerolog"
)

var logger = core.CoreLogger().With().Logger()

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
	instance.echoInstance.Use(middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogMethod:   true,
		LogError:    true,
		LogRemoteIP: true,
		LogURIPath:  true,
		LogURI:      true,
		LogStatus:   true,
		LogValuesFunc: func(_ echo.Context, v middleware.RequestLoggerValues) error {
			var event *zerolog.Event
			if v.Status < 400 {
				event = logger.Info()
			} else if v.Status >= 400 {
				event = logger.Error()
			}
			event.Str("PATH", v.URIPath).Str("REMOTEIP", v.RemoteIP).Int("STATUS", v.Status)
			if v.Error != nil {
				event.Str("ERROR", v.Error.Error())
			}
			event.Msg(fmt.Sprintf("[%-5s]", v.Method))
			return nil
		},
	}))
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
